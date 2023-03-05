package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/pkg/logger"
)

type StorageService struct {
	storage Storage
	log     logger.Logger
}

func NewStorageService(storage Storage, log logger.Logger) *StorageService {
	return &StorageService{
		storage: storage,
		log:     log,
	}
}

func (ss *StorageService) OpenConn(request *AuthRequest, ctx context.Context) (ConnDB, string, error) {
	return ss.storage.OpenConn(request, ctx)
}

func (ss *StorageService) CloseConn(conn ConnDB) error {
	return ss.storage.CloseConn(conn)
}

func (ss *StorageService) GetAllUnits(conn ConnDB, layer string) (interfaces.OutputUnitsDTO, error) {
	exist, err := ss.storage.LayerExist(conn, layer)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.OutputUnitsDTO{}, err
	}
	if !exist {
		err = errors.New(fmt.Sprintf("layer %s does not exist in database", layer))
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.OutputUnitsDTO{}, err
	}

	units, err := ss.storage.GetAllUnits(conn, layer)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.OutputUnitsDTO{}, err
	}

	return units, nil
}

func (ss *StorageService) GetUnitsByModels(conn ConnDB, layer string, modelsDTO interfaces.ModelsIdDTO) (interfaces.OutputUnitsDTO, error) {
	exist, err := ss.storage.LayerExist(conn, layer)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.OutputUnitsDTO{}, err
	}
	if !exist {
		err = errors.New(fmt.Sprintf("layer %s does not exist in database", layer))
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.OutputUnitsDTO{}, err
	}

	units, err := ss.storage.GetUnitsByModels(conn, layer, modelsDTO.Models)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.OutputUnitsDTO{}, err
	}

	return units, nil
}

func (ss *StorageService) GetUnitsByProperties(conn ConnDB, layer string, propertiesDTO interfaces.PropertiesIdDTO) (interfaces.OutputUnitsDTO, error) {
	exist, err := ss.storage.LayerExist(conn, layer)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.OutputUnitsDTO{}, err
	}
	if !exist {
		err = errors.New(fmt.Sprintf("layer %s does not exist in database", layer))
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.OutputUnitsDTO{}, err
	}

	units, err := ss.storage.GetUnitsByProperties(conn, layer, propertiesDTO.Properties)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.OutputUnitsDTO{}, err
	}

	return units, nil
}

func (ss *StorageService) GetModels(conn ConnDB, layer string) (interfaces.OutputModelsDTO, error) {
	exist, err := ss.storage.LayerExist(conn, layer)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.OutputModelsDTO{}, err
	}
	if !exist {
		err = errors.New(fmt.Sprintf("layer %s does not exist in database", layer))
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.OutputModelsDTO{}, err
	}

	models, err := ss.storage.GetAllModels(conn, layer)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.OutputModelsDTO{}, err
	}

	modelsDTO := make([]interfaces.OutputModelDTO, len(models))
	for i := range modelsDTO {
		modelsDTO[i] = interfaces.OutputModelDTO(models[i])
	}

	return interfaces.OutputModelsDTO{Models: modelsDTO}, nil
}

func (ss *StorageService) GetModelElements(conn ConnDB, layer string) (interfaces.OutputModelElementsDTO, error) {
	exist, err := ss.storage.LayerExist(conn, layer)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.OutputModelElementsDTO{}, err
	}
	if !exist {
		err = errors.New(fmt.Sprintf("layer %s does not exist in database", layer))
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.OutputModelElementsDTO{}, err
	}

	modelElements, err := ss.storage.GetAllModelElements(conn, layer)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.OutputModelElementsDTO{}, err
	}

	modelElementsDTO := make([]interfaces.OutputModelElementDTO, len(modelElements))
	for i := range modelElementsDTO {
		modelElementsDTO[i] = interfaces.OutputModelElementDTO(modelElements[i])
	}

	return interfaces.OutputModelElementsDTO{Elements: modelElementsDTO}, nil
}

func (ss *StorageService) GetProperties(conn ConnDB) (interfaces.OutputPropertiesDTO, error) {
	properties, err := ss.storage.GetAllProperties(conn)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.OutputPropertiesDTO{}, err
	}

	propertiesDTO := make([]interfaces.OutputPropertyDTO, len(properties))
	for i := range propertiesDTO {
		propertiesDTO[i] = interfaces.OutputPropertyDTO(properties[i])
	}

	return interfaces.OutputPropertiesDTO{Properties: propertiesDTO}, nil
}

func (ss *StorageService) GetPropertiesByUnit(conn ConnDB, layer string, unit interfaces.SearchUnitDTO) (interfaces.OutputPropertiesDTO, error) {
	properties, err := ss.storage.GetPropertiesByUnit(conn, layer, unit)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.OutputPropertiesDTO{}, err
	}

	propertiesDTO := make([]interfaces.OutputPropertyDTO, len(properties))
	for i := range propertiesDTO {
		propertiesDTO[i] = interfaces.OutputPropertyDTO(properties[i])
	}

	return interfaces.OutputPropertiesDTO{Properties: propertiesDTO}, nil
}

func (ss *StorageService) GetLayers(conn ConnDB) (interfaces.LayersDTO, error) {
	layers, err := ss.storage.GetAllLayers(conn)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.LayersDTO{}, err
	}

	return interfaces.LayersDTO{Layers: layers}, nil
}

func (ss *StorageService) SaveUnits(conn ConnDB, layer string, unitsDTO interfaces.SaveUnitsDTO) error {
	exist, err := ss.storage.LayerExist(conn, layer)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return err
	}
	if !exist {
		err = errors.New(fmt.Sprintf("layer %s does not exist in database", layer))
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return err
	}

	err = ss.storage.SaveUnits(conn, layer, unitsDTO)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return err
	}

	return nil
}

func (ss *StorageService) UpdateUnits(conn ConnDB, layer string, unitsDTO interfaces.UpdateUnitsDTO) error {
	exist, err := ss.storage.LayerExist(conn, layer)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return err
	}
	if !exist {
		err = errors.New(fmt.Sprintf("layer %s does not exist in database", layer))
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return err
	}

	for _, unit := range unitsDTO.Units {
		name := unit.OldText

		if unit.NewText != "" {
			err = ss.storage.RenameUnit(conn, layer, unit.OldText, unit.NewText)
			if err != nil {
				ss.log.Print(logger.LogRecord{
					Name: "StorageService",
					Type: logger.Error,
					Msg:  err.Error(),
				})
				return err
			}
			name = unit.NewText
		}

		if len(unit.PropertiesID) != 0 {
			err = ss.storage.SetUnitProperties(conn, layer, name, unit.PropertiesID)
			if err != nil {
				ss.log.Print(logger.LogRecord{
					Name: "StorageService",
					Type: logger.Error,
					Msg:  err.Error(),
				})
				return err
			}
		}
	}

	return nil
}

func (ss *StorageService) SaveProperties(conn ConnDB, propertiesDTO interfaces.PropertyNamesDTO) (interfaces.PropertiesIdDTO, error) {
	propertiesID, err := ss.storage.SaveProperties(conn, propertiesDTO.Properties)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.PropertiesIdDTO{}, err
	}

	return interfaces.PropertiesIdDTO{Properties: propertiesID}, nil
}

func (ss *StorageService) SaveModels(conn ConnDB, layer string, modelsDTO interfaces.ModelNamesDTO) (interfaces.ModelsIdDTO, error) {
	exist, err := ss.storage.LayerExist(conn, layer)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.ModelsIdDTO{}, err
	}
	if !exist {
		err = errors.New(fmt.Sprintf("layer %s does not exist in database", layer))
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.ModelsIdDTO{}, err
	}

	modelsID, err := ss.storage.SaveModels(conn, layer, modelsDTO.Models)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.ModelsIdDTO{}, err
	}

	return interfaces.ModelsIdDTO{Models: modelsID}, nil
}

func (ss *StorageService) SaveModelElements(conn ConnDB, layer string, modelElementsDTO interfaces.ModelElementNamesDTO) (interfaces.ModelElementsIdDTO, error) {
	exist, err := ss.storage.LayerExist(conn, layer)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.ModelElementsIdDTO{}, err
	}
	if !exist {
		err = errors.New(fmt.Sprintf("layer %s does not exist in database", layer))
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.ModelElementsIdDTO{}, err
	}

	modelElementsID, err := ss.storage.SaveModelElements(conn, layer, modelElementsDTO.ModelElements)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return interfaces.ModelElementsIdDTO{}, err
	}

	return interfaces.ModelElementsIdDTO{ModelElements: modelElementsID}, nil
}

func (ss *StorageService) SaveLayer(conn ConnDB, layer string) error {
	err := ss.storage.SaveLayer(conn, layer)
	if err != nil {
		ss.log.Print(logger.LogRecord{
			Name: "StorageService",
			Type: logger.Error,
			Msg:  err.Error(),
		})
		return err
	}

	return nil
}

func (ss *StorageService) AddUser(conn ConnDB, username string, role string) error {
	return ss.storage.AddUser(conn, username, role)
}

func (ss *StorageService) GetUserPassword(conn ConnDB, username string) (string, error) {
	return ss.storage.GetUserPassword(conn, username)
}
