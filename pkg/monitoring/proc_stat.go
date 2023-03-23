package monitoring

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/prometheus/procfs"
	"os"
	"time"
)

type ProcStatObserver struct {
	process procfs.Proc
	file    *os.File
}

type ProcStat struct {
	MeasuringTime string          `json:"measuring_time"`
	CPUTime       float64         `json:"cpu_time"`
	Vms           uint            `json:"virtual_memory_size"`
	Rss           int             `json:"resident_memory_size"`
	StartTime     float64         `json:"start_time"`
	ProcStat      procfs.ProcStat `json:"proc_stat"`
}

func NewProcStatObserver() (*ProcStatObserver, error) {
	p, err := procfs.Self()
	if err != nil {
		return nil, fmt.Errorf("could not get process: %w", err)
	}

	return &ProcStatObserver{process: p}, nil
}

func (o *ProcStatObserver) GetInfo() (*ProcStat, error) {
	stat, err := o.process.Stat()
	if err != nil {
		return nil, fmt.Errorf("could not get process stat: %w", err)
	}

	startTime, err := stat.StartTime()
	if err != nil {
		return nil, err
	}

	return &ProcStat{
		MeasuringTime: time.Now().Format(time.DateTime),
		CPUTime:       stat.CPUTime(),
		Vms:           stat.VirtualMemory(),
		Rss:           stat.ResidentMemory(),
		StartTime:     startTime,
		ProcStat:      stat,
	}, nil
}

func (o *ProcStatObserver) saveInfo(stat *ProcStat) error {
	data, err := json.MarshalIndent(stat, "", "	")
	if err != nil {
		return err
	}

	_, err = o.file.Write(data)

	return err
}

func (o *ProcStatObserver) Observe(ctx context.Context, d time.Duration, filename string) error {
	var err error
	o.file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer o.file.Close()

	var observeErr error
	ticker := time.NewTicker(d)
	go func() {
		for {
			select {
			case <-ticker.C:
				info, err := o.GetInfo()
				if err != nil {
					observeErr = err
					return
				}
				err = o.saveInfo(info)
				if err != nil {
					observeErr = err
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return observeErr
}
