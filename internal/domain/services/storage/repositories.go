package storage

import (
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
)

type UsersRepository interface {
	AddUser(conn ConnDB, user interfaces.UserDTO) error
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
	GetPropertiesByUnit(conn ConnDB, layer string, unit interfaces.SearchUnitDTO) ([]entities.Property, error)
	SaveProperties(conn ConnDB, properties []string) ([]int, error)
}

type UnitsRepository interface {
	GetAllUnits(conn ConnDB, layer string) (interfaces.OutputUnitsDTO, error)
	GetUnitsByModels(conn ConnDB, layer string, modelsID []int) (interfaces.OutputUnitsDTO, error)
	GetUnitsByProperties(conn ConnDB, layer string, propertiesID []int) (interfaces.OutputUnitsDTO, error)
	SaveUnits(conn ConnDB, layer string, data interfaces.SaveUnitsDTO) error // stored procedure (SQL)
	RenameUnit(conn ConnDB, layer, lang, oldName, newName string) error
	SetUnitProperties(conn ConnDB, layer, lang, unitName string, propertiesID []int) error
}

type LayersRepository interface {
	LayerExist(conn ConnDB, layer string) (bool, error)
	GetAllLayers(conn ConnDB) ([]string, error)
	// GetLayerLanguages(conn ConnDB, layer string) ([]string, error)
	SaveLayer(conn ConnDB, name string) error
}

//type ContextsRepository interface {
//	GetByUnit(conn ConnDB, unit models.Unit) ([]models.Context, error)
//	// Save(conn ConnDB, contexts []string) ([]int, error)
//}

//type LinksRepository interface {
//	SaveModelElementLinks(conn ConnDB, layer string, links []models.Link) error
//	SaveModelUnitRuLinks(conn ConnDB, layer string, links []models.Link) error
//	SaveModelUnitEnLinks(conn ConnDB, layer string, links []models.Link) error
//	SaveUnitRuEnLinks(conn ConnDB, layer string, links []models.Link) error
//	SavePropertyUnitRuLinks(conn ConnDB, layer string, links []models.Link) error
//	SavePropertyUnitEnLinks(conn ConnDB, layer string, links []models.Link) error
//	SaveContextUnitRuLinks(conn ConnDB, layer string, links []models.Link) error
//	SaveContextUnitEnLinks(conn ConnDB, layer string, links []models.Link) error
//	SaveUserUnitRuLinks(conn ConnDB, layer string, links []models.Link) error
//	SaveUserUnitEnLinks(conn ConnDB, layer string, links []models.Link) error
//	// GetCount...() (int, err) ?
//}
