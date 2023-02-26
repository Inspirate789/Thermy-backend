package interfaces

import "github.com/Inspirate789/Thermy-backend/internal/domain/entities"

type UserRepository interface {
	AddUser(conn ConnDB, username string, role string) error
	GetUserPassword(conn ConnDB, username string) (string, error)
}

type ModelsRepository interface {
	GetAllModels(conn ConnDB, layer string) ([]entities.Model, error)
	SaveModels(conn ConnDB, layer string, models []string) ([]int, error)
}

type ModelElementsRepository interface {
	GetAllModelElements(conn ConnDB, layer string) ([]entities.ModelElement, error)
	SaveModelElements(conn ConnDB, layer string, modelElements []string) ([]int, error)
}

type PropertiesRepository interface {
	GetAllProperties(conn ConnDB) ([]entities.Property, error)
	GetPropertiesByUnit(conn ConnDB, layer string, unit SearchUnitDTO) ([]entities.Property, error)
	SaveProperties(conn ConnDB, properties []string) ([]int, error)
}

type UnitsRepository interface {
	GetAllUnits(conn ConnDB, layer string) (OutputUnitsDTO, error)
	GetUnitsByModels(conn ConnDB, layer string, modelsID []int) (OutputUnitsDTO, error)
	GetUnitsByProperties(conn ConnDB, layer string, propertiesID []int) (OutputUnitsDTO, error)
	SaveUnits(conn ConnDB, layer string, data SaveUnitsDTO) error // stored procedure (SQL)
	RenameUnit(conn ConnDB, layer string, oldName string, newName string) error
	SetUnitProperties(conn ConnDB, layer string, unitName string, propertiesID []int) error
}

type LayersRepository interface {
	LayerExist(conn ConnDB, layer string) (bool, error)
	GetAllLayers(conn ConnDB) ([]string, error)
	// GetLayerLanguages(conn ConnDB, layer string) ([]string, error)
	SaveLayer(conn ConnDB, name string) error
}

//type ContextsRepository interface {
//	GetByUnit(conn ConnDB, unit entities.Unit) ([]entities.Context, error)
//	// Save(conn ConnDB, contexts []string) ([]int, error)
//}

//type LinksRepository interface {
//	SaveModelElementLinks(conn ConnDB, layer string, links []entities.Link) error
//	SaveModelUnitRuLinks(conn ConnDB, layer string, links []entities.Link) error
//	SaveModelUnitEnLinks(conn ConnDB, layer string, links []entities.Link) error
//	SaveUnitRuEnLinks(conn ConnDB, layer string, links []entities.Link) error
//	SavePropertyUnitRuLinks(conn ConnDB, layer string, links []entities.Link) error
//	SavePropertyUnitEnLinks(conn ConnDB, layer string, links []entities.Link) error
//	SaveContextUnitRuLinks(conn ConnDB, layer string, links []entities.Link) error
//	SaveContextUnitEnLinks(conn ConnDB, layer string, links []entities.Link) error
//	SaveUserUnitRuLinks(conn ConnDB, layer string, links []entities.Link) error
//	SaveUserUnitEnLinks(conn ConnDB, layer string, links []entities.Link) error
//	// GetCount...() (int, err) ?
//}
