package storage

import (
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/errors"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	storage         Storage
	logger          *log.Logger
	repeatOnFailure int
}

func NewStorageService(storage Storage, logger *log.Logger, repeatOnFailure int) *Service {
	return &Service{
		storage:         storage,
		logger:          logger,
		repeatOnFailure: repeatOnFailure,
	}
}

func (ss *Service) GetAllUnits(layer string) (interfaces.OutputUnitsDTO, error) {
	exist, err := ss.storage.LayerExist(layer)
	if err != nil {
		ss.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, errors.IdentifyStorageError(err)
	}
	if !exist {
		err = errors.ErrUnknownLayerWrap(layer)
		ss.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, err
	}

	units, err := ss.storage.GetAllUnits(layer)
	if err != nil {
		ss.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, errors.IdentifyStorageError(err)
	}

	return units, nil
}

func (ss *Service) GetUnitsByModels(layer string, modelsDTO interfaces.ModelsIdDTO) (interfaces.OutputUnitsDTO, error) {
	exist, err := ss.storage.LayerExist(layer)
	if err != nil {
		ss.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, errors.IdentifyStorageError(err)
	}
	if !exist {
		err = errors.ErrUnknownLayerWrap(layer)
		ss.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, err
	}

	if len(modelsDTO.Models) == 0 {
		return interfaces.OutputUnitsDTO{
			Units:    make([]map[string]interfaces.OutputUnitDTO, 0),
			Contexts: make([]interfaces.ContextDTO, 0),
		}, nil
	}

	units, err := ss.storage.GetUnitsByModels(layer, modelsDTO.Models)
	if err != nil {
		ss.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, errors.IdentifyStorageError(err)
	}

	return units, nil
}

func (ss *Service) GetUnitsByProperties(layer string, propertiesDTO interfaces.PropertiesIdDTO) (interfaces.OutputUnitsDTO, error) {
	exist, err := ss.storage.LayerExist(layer)
	if err != nil {
		ss.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, errors.IdentifyStorageError(err)
	}
	if !exist {
		err = errors.ErrUnknownLayerWrap(layer)
		ss.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, err
	}

	if len(propertiesDTO.Properties) == 0 {
		return interfaces.OutputUnitsDTO{
			Units:    make([]map[string]interfaces.OutputUnitDTO, 0),
			Contexts: make([]interfaces.ContextDTO, 0),
		}, nil
	}

	units, err := ss.storage.GetUnitsByProperties(layer, propertiesDTO.Properties)
	if err != nil {
		ss.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, errors.IdentifyStorageError(err)
	}

	return units, nil
}

func (ss *Service) GetModels(layer string) (interfaces.OutputModelsDTO, error) {
	exist, err := ss.storage.LayerExist(layer)
	if err != nil {
		ss.logger.Error(err)
		return interfaces.OutputModelsDTO{}, errors.IdentifyStorageError(err)
	}
	if !exist {
		err = errors.ErrUnknownLayerWrap(layer)
		ss.logger.Error(err)
		return interfaces.OutputModelsDTO{}, err
	}

	models, err := ss.storage.GetAllModels(layer)
	if err != nil {
		ss.logger.Error(err)
		return interfaces.OutputModelsDTO{}, errors.IdentifyStorageError(err)
	}

	modelsDTO := make([]interfaces.OutputModelDTO, 0, len(models))
	for i := range models {
		modelsDTO = append(modelsDTO, interfaces.OutputModelDTO(models[i]))
	}

	return interfaces.OutputModelsDTO{Models: modelsDTO}, nil
}

func (ss *Service) GetModelElements(layer string) (interfaces.OutputModelElementsDTO, error) {
	exist, err := ss.storage.LayerExist(layer)
	if err != nil {
		ss.logger.Error(err)
		return interfaces.OutputModelElementsDTO{}, errors.IdentifyStorageError(err)
	}
	if !exist {
		err = errors.ErrUnknownLayerWrap(layer)
		ss.logger.Error(err)
		return interfaces.OutputModelElementsDTO{}, err
	}

	modelElements, err := ss.storage.GetAllModelElements(layer)
	if err != nil {
		ss.logger.Error(err)
		return interfaces.OutputModelElementsDTO{}, errors.IdentifyStorageError(err)
	}

	modelElementsDTO := make([]interfaces.OutputModelElementDTO, 0, len(modelElements))
	for i := range modelElements {
		modelElementsDTO = append(modelElementsDTO, interfaces.OutputModelElementDTO(modelElements[i]))
	}

	return interfaces.OutputModelElementsDTO{Elements: modelElementsDTO}, nil
}

func (ss *Service) GetProperties() (interfaces.OutputPropertiesDTO, error) {
	properties, err := ss.storage.GetAllProperties()
	if err != nil {
		ss.logger.Error(err)
		return interfaces.OutputPropertiesDTO{}, errors.IdentifyStorageError(err)
	}

	propertiesDTO := make([]interfaces.OutputPropertyDTO, 0, len(properties))
	for i := range properties {
		propertiesDTO = append(propertiesDTO, interfaces.OutputPropertyDTO(properties[i]))
	}

	return interfaces.OutputPropertiesDTO{Properties: propertiesDTO}, nil
}

func (ss *Service) GetPropertiesByUnit(layer string, unit interfaces.SearchUnitDTO) (interfaces.OutputPropertiesDTO, error) {
	properties, err := ss.storage.GetPropertiesByUnit(layer, unit)
	if err != nil {
		ss.logger.Error(err)
		return interfaces.OutputPropertiesDTO{}, errors.IdentifyStorageError(err)
	}

	propertiesDTO := make([]interfaces.OutputPropertyDTO, 0, len(properties))
	for i := range properties {
		propertiesDTO = append(propertiesDTO, interfaces.OutputPropertyDTO(properties[i]))
	}

	return interfaces.OutputPropertiesDTO{Properties: propertiesDTO}, nil
}

func (ss *Service) GetLayers() (interfaces.LayersDTO, error) {
	layers, err := ss.storage.GetAllLayers()
	if err != nil {
		ss.logger.Error(err)
		return interfaces.LayersDTO{}, errors.IdentifyStorageError(err)
	}

	return interfaces.LayersDTO{Layers: layers}, nil
}

func (ss *Service) SaveUnits(layer string, unitsDTO interfaces.SaveUnitsDTO) error {
	exist, err := ss.storage.LayerExist(layer)
	if err != nil {
		ss.logger.Error(err)
		return errors.IdentifyStorageError(err)
	}
	if !exist {
		err = errors.ErrUnknownLayerWrap(layer)
		ss.logger.Error(err)
		return err
	}

	for i := 0; i < ss.repeatOnFailure; i++ {
		err = ss.storage.SaveUnits(layer, unitsDTO)
		if err == nil || errors.IdentifyStorageError(err) != errors.ErrConnectDatabase {
			break
		}
	}
	if err != nil {
		ss.logger.Error(err)
		return errors.IdentifyStorageError(err)
	}

	return nil
}

func (ss *Service) UpdateUnits(layer string, unitsDTO interfaces.UpdateUnitsDTO) error {
	exist, err := ss.storage.LayerExist(layer)
	if err != nil {
		ss.logger.Error(err)
		return errors.IdentifyStorageError(err)
	}
	if !exist {
		err = errors.ErrUnknownLayerWrap(layer)
		ss.logger.Error(err)
		return err
	}

	for _, unit := range unitsDTO.Units {
		name := unit.OldText

		if unit.NewText != nil {
			err = ss.storage.RenameUnit(layer, unit.Lang, unit.OldText, *unit.NewText)
			if err != nil {
				ss.logger.Error(err)
				return errors.IdentifyStorageError(err)
			}
			name = *unit.NewText
		}

		if len(unit.PropertiesID) != 0 {
			err = ss.storage.SetUnitProperties(layer, unit.Lang, name, unit.PropertiesID)
			if err != nil {
				ss.logger.Error(err)
				return errors.IdentifyStorageError(err)
			}
		}
	}

	return nil
}

func (ss *Service) SaveProperties(propertiesDTO interfaces.PropertyNamesDTO) (interfaces.PropertiesIdDTO, error) {
	propertiesID, err := ss.storage.SaveProperties(propertiesDTO.Properties)
	if err != nil {
		ss.logger.Error(err)
		return interfaces.PropertiesIdDTO{}, errors.IdentifyStorageError(err)
	}

	return interfaces.PropertiesIdDTO{Properties: propertiesID}, nil
}

func (ss *Service) SaveModels(layer string, modelsDTO interfaces.ModelNamesDTO) (interfaces.ModelsIdDTO, error) {
	exist, err := ss.storage.LayerExist(layer)
	if err != nil {
		ss.logger.Error(err)
		return interfaces.ModelsIdDTO{}, errors.IdentifyStorageError(err)
	}
	if !exist {
		err = errors.ErrUnknownLayerWrap(layer)
		ss.logger.Error(err)
		return interfaces.ModelsIdDTO{}, err
	}

	for _, modelName := range modelsDTO.Models {
		model := entities.Model{Name: modelName}
		err = model.IsValidName()
		if err != nil {
			ss.logger.Error(err)
			return interfaces.ModelsIdDTO{}, err
		}
	}

	modelsID, err := ss.storage.SaveModels(layer, modelsDTO.Models)
	if err != nil {
		ss.logger.Error(err)
		return interfaces.ModelsIdDTO{}, errors.IdentifyStorageError(err)
	}

	return interfaces.ModelsIdDTO{Models: modelsID}, nil
}

func (ss *Service) SaveModelElements(layer string, modelElementsDTO interfaces.ModelElementNamesDTO) (interfaces.ModelElementsIdDTO, error) {
	exist, err := ss.storage.LayerExist(layer)
	if err != nil {
		ss.logger.Error(err)
		return interfaces.ModelElementsIdDTO{}, errors.IdentifyStorageError(err)
	}
	if !exist {
		err = errors.ErrUnknownLayerWrap(layer)
		ss.logger.Error(err)
		return interfaces.ModelElementsIdDTO{}, err
	}

	modelElementsID, err := ss.storage.SaveModelElements(layer, modelElementsDTO.ModelElements)
	if err != nil {
		ss.logger.Error(err)
		return interfaces.ModelElementsIdDTO{}, errors.IdentifyStorageError(err)
	}

	return interfaces.ModelElementsIdDTO{ModelElements: modelElementsID}, nil
}

func (ss *Service) SaveLayer(layer string) error {
	err := ss.storage.SaveLayer(layer)
	if err != nil {
		ss.logger.Error(err)
		return errors.IdentifyStorageError(err)
	}

	return nil
}

func (ss *Service) AddUser(user interfaces.UserDTO) error {
	err := ss.storage.AddUser(user)
	if err != nil {
		ss.logger.Error(err)
		return errors.IdentifyStorageError(err)
	}
	ss.logger.Infof("storage user (name: %s, role: %s) inserted to database", user.Name, user.Role)

	return nil
}
