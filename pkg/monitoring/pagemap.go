package monitoring

import (
	"github.com/prometheus/procfs"
	"periph.io/x/host/v3/pmem"
)

const pageSize = 4096

type Page struct {
	Addr                 uintptr `json:"addr"`
	PhysicalAddr         uint64  `json:"physical_addr"`
	Present              bool    `json:"present"`
	Swapped              bool    `json:"swapped"`
	PteSoftDirty         bool    `json:"pte_soft_dirty"`
	ExclusivelyMapped    bool    `json:"exclusively_mapped"`
	FilePageOrSharedAnon bool    `json:"file_page_or_shared_anon"`
}

func getBit(addr uint64, bitIndex uint) bool {
	return ((addr >> bitIndex) & 0x1) != 0
}

func NewPage(virtualAddr uintptr) (*Page, error) {
	physicalAddr, err := pmem.ReadPageMap(virtualAddr)
	if err != nil {
		return nil, err
	}

	return &Page{
		Addr:                 virtualAddr,
		PhysicalAddr:         physicalAddr,
		Present:              getBit(physicalAddr, 63),
		Swapped:              getBit(physicalAddr, 62),
		PteSoftDirty:         getBit(physicalAddr, 55),
		ExclusivelyMapped:    getBit(physicalAddr, 56),
		FilePageOrSharedAnon: getBit(physicalAddr, 61),
	}, nil
}

type Region struct {
	StartAddr            uintptr `json:"start_addr"`
	EndAddr              uintptr `json:"end_addr"`
	PagesCount           uint64  `json:"pages_count"`
	Present              uint64  `json:"present"`
	Swapped              uint64  `json:"swapped"`
	PteSoftDirty         uint64  `json:"pte_soft_dirty"`
	ExclusivelyMapped    uint64  `json:"exclusively_mapped"`
	FilePageOrSharedAnon uint64  `json:"file_page_or_shared_anon"`
	Undefined            uint64  `json:"undefined"`
	Pages                []*Page `json:"-"`
}

func NewRegion(startAddr, endAddr uintptr) (*Region, error) {
	pagesCount := uint64((endAddr - startAddr) / pageSize)
	pages := make([]*Page, 0, pagesCount)
	var (
		present              uint64 = 0
		swapped              uint64 = 0
		pteSoftDirty         uint64 = 0
		exclusivelyMapped    uint64 = 0
		filePageOrSharedAnon uint64 = 0
		undefined            uint64 = 0
	)
	for addr := startAddr; addr < endAddr; addr += pageSize {
		page, err := NewPage(addr)
		if err != nil {
			undefined++
			continue
		}
		pages = append(pages, page)
		if page.Present {
			present++
		}
		if page.Swapped {
			swapped++
		}
		if page.PteSoftDirty {
			pteSoftDirty++
		}
		if page.ExclusivelyMapped {
			exclusivelyMapped++
		}
		if page.FilePageOrSharedAnon {
			filePageOrSharedAnon++
		}
	}

	return &Region{StartAddr: startAddr,
		EndAddr:              endAddr,
		PagesCount:           pagesCount,
		Present:              present,
		Swapped:              swapped,
		PteSoftDirty:         pteSoftDirty,
		ExclusivelyMapped:    exclusivelyMapped,
		FilePageOrSharedAnon: filePageOrSharedAnon,
		Undefined:            undefined,
		Pages:                pages,
	}, nil
}

type PageMap struct {
	Regions []*Region `json:"regions"`
}

func NewPageMaps(maps []*procfs.ProcMap) (*PageMap, error) {
	regions := make([]*Region, 0, len(maps))
	for _, m := range maps {
		region, err := NewRegion(m.StartAddr, m.EndAddr)
		if err != nil {
			return nil, err
		}
		regions = append(regions, region)
	}

	return &PageMap{Regions: regions}, nil
}
