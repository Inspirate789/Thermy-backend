package storage

import (
	"context"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
)

type ConnDB any

type Storage interface { // TODO: split?
	OpenConn(request *entities.AuthRequest, ctx context.Context) (ConnDB, string, error) // Get conn, role in database and error
	UsersRepository
	ModelsRepository
	ModelElementsRepository
	PropertiesRepository
	UnitsRepository
	LayersRepository
	CloseConn(ConnDB) error
}

type ConnManager interface {
	OpenConn(request *entities.AuthRequest, ctx context.Context) (ConnDB, string, error)
	CloseConn(conn ConnDB) error
}

type UnitsManager interface {
	GetAllUnits(conn ConnDB, layer string) (interfaces.OutputUnitsDTO, error)
	GetUnitsByModels(conn ConnDB, layer string, modelsDTO interfaces.ModelsIdDTO) (interfaces.OutputUnitsDTO, error)
	GetUnitsByProperties(conn ConnDB, layer string, propertiesDTO interfaces.PropertiesIdDTO) (interfaces.OutputUnitsDTO, error)
	SaveUnits(conn ConnDB, layer string, unitsDTO interfaces.SaveUnitsDTO) error
	UpdateUnits(conn ConnDB, layer string, unitsDTO interfaces.UpdateUnitsDTO) error
}

type ModelsManager interface {
	GetModels(conn ConnDB, layer string) (interfaces.OutputModelsDTO, error)
	SaveModels(conn ConnDB, layer string, modelsDTO interfaces.ModelNamesDTO) (interfaces.ModelsIdDTO, error)
}

type ModelElementsManager interface {
	GetModelElements(conn ConnDB, layer string) (interfaces.OutputModelElementsDTO, error)
	SaveModelElements(conn ConnDB, layer string, modelElementsDTO interfaces.ModelElementNamesDTO) (interfaces.ModelElementsIdDTO, error)
}

type PropertiesManager interface {
	GetProperties(conn ConnDB) (interfaces.OutputPropertiesDTO, error)
	GetPropertiesByUnit(conn ConnDB, layer string, unit interfaces.SearchUnitDTO) (interfaces.OutputPropertiesDTO, error)
	SaveProperties(conn ConnDB, propertiesDTO interfaces.PropertyNamesDTO) (interfaces.PropertiesIdDTO, error)
}

type LayersManager interface {
	GetLayers(conn ConnDB) (interfaces.LayersDTO, error)
	SaveLayer(conn ConnDB, layer string) error
}

type UsersManager interface {
	AddUser(conn ConnDB, user interfaces.UserDTO) error
	//GetUserPassword(conn ConnDB, username string) (string, error)
}

type StorageManager interface {
	ConnManager
	UnitsManager
	ModelsManager
	ModelElementsManager
	PropertiesManager
	LayersManager
	UsersManager
}
