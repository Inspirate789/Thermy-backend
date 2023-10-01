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

func (s *Service) checkLayer(layer string) error {
	exist, err := s.storage.LayerExist(layer)
	if err != nil {
		return errors.IdentifyStorageError(err)
	}
	if !exist {
		return errors.ErrUnknownLayerWrap(layer)
	}
	return nil
}

func (s *Service) GetAllUnits(layer string) (interfaces.OutputUnitsDTO, error) {
	err := s.checkLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, err
	}

	units, err := s.storage.GetAllUnits(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, errors.IdentifyStorageError(err)
	}

	return units, nil
}

func (s *Service) GetUnitsByModels(layer string, modelsDTO interfaces.ModelsIdDTO) (interfaces.OutputUnitsDTO, error) {
	err := s.checkLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, err
	}

	if len(modelsDTO.Models) == 0 {
		return interfaces.OutputUnitsDTO{
			Units:    make([]map[string]interfaces.OutputUnitDTO, 0),
			Contexts: make([]interfaces.ContextDTO, 0),
		}, nil
	}

	units, err := s.storage.GetUnitsByModels(layer, modelsDTO.Models)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, errors.IdentifyStorageError(err)
	}

	return units, nil
}

func (s *Service) GetUnitsByProperties(layer string, propertiesDTO interfaces.PropertiesIdDTO) (interfaces.OutputUnitsDTO, error) {
	err := s.checkLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, err
	}

	if len(propertiesDTO.Properties) == 0 {
		return interfaces.OutputUnitsDTO{
			Units:    make([]map[string]interfaces.OutputUnitDTO, 0),
			Contexts: make([]interfaces.ContextDTO, 0),
		}, nil
	}

	units, err := s.storage.GetUnitsByProperties(layer, propertiesDTO.Properties)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputUnitsDTO{}, errors.IdentifyStorageError(err)
	}

	return units, nil
}

func (s *Service) GetModels(layer string) (interfaces.OutputModelsDTO, error) {
	err := s.checkLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputModelsDTO{}, err
	}

	models, err := s.storage.GetAllModels(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputModelsDTO{}, errors.IdentifyStorageError(err)
	}

	modelsDTO := make([]interfaces.OutputModelDTO, 0, len(models))
	for i := range models {
		modelsDTO = append(modelsDTO, interfaces.OutputModelDTO(models[i]))
	}

	return interfaces.OutputModelsDTO{Models: modelsDTO}, nil
}

func (s *Service) GetModelElements(layer string) (interfaces.OutputModelElementsDTO, error) {
	err := s.checkLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputModelElementsDTO{}, err
	}

	modelElements, err := s.storage.GetAllModelElements(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputModelElementsDTO{}, errors.IdentifyStorageError(err)
	}

	modelElementsDTO := make([]interfaces.OutputModelElementDTO, 0, len(modelElements))
	for i := range modelElements {
		modelElementsDTO = append(modelElementsDTO, interfaces.OutputModelElementDTO(modelElements[i]))
	}

	return interfaces.OutputModelElementsDTO{Elements: modelElementsDTO}, nil
}

func (s *Service) GetProperties() (interfaces.OutputPropertiesDTO, error) {
	properties, err := s.storage.GetAllProperties()
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputPropertiesDTO{}, errors.IdentifyStorageError(err)
	}

	propertiesDTO := make([]interfaces.OutputPropertyDTO, 0, len(properties))
	for i := range properties {
		propertiesDTO = append(propertiesDTO, interfaces.OutputPropertyDTO(properties[i]))
	}

	return interfaces.OutputPropertiesDTO{Properties: propertiesDTO}, nil
}

func (s *Service) GetPropertiesByUnit(layer string, unit interfaces.SearchUnitDTO) (interfaces.OutputPropertiesDTO, error) {
	properties, err := s.storage.GetPropertiesByUnit(layer, unit)
	if err != nil {
		s.logger.Error(err)
		return interfaces.OutputPropertiesDTO{}, errors.IdentifyStorageError(err)
	}

	propertiesDTO := make([]interfaces.OutputPropertyDTO, 0, len(properties))
	for i := range properties {
		propertiesDTO = append(propertiesDTO, interfaces.OutputPropertyDTO(properties[i]))
	}

	return interfaces.OutputPropertiesDTO{Properties: propertiesDTO}, nil
}

func (s *Service) GetLayers() (interfaces.LayersDTO, error) {
	layers, err := s.storage.GetAllLayers()
	if err != nil {
		s.logger.Error(err)
		return interfaces.LayersDTO{}, errors.IdentifyStorageError(err)
	}

	return interfaces.LayersDTO{Layers: layers}, nil
}

func (s *Service) SaveUnits(layer string, unitsDTO interfaces.SaveUnitsDTO) error {
	err := s.checkLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	for i := 0; i < s.repeatOnFailure; i++ {
		err = s.storage.SaveUnits(layer, unitsDTO)
		if err == nil || errors.IdentifyStorageError(err) != errors.ErrConnectDatabase {
			break
		}
	}
	if err != nil {
		s.logger.Error(err)
		return errors.IdentifyStorageError(err)
	}

	return nil
}

func (s *Service) UpdateUnits(layer string, unitsDTO interfaces.UpdateUnitsDTO) error {
	err := s.checkLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return err
	}

	for _, unit := range unitsDTO.Units {
		name := unit.OldText

		if unit.NewText != nil {
			err = s.storage.RenameUnit(layer, unit.Lang, unit.OldText, *unit.NewText)
			if err != nil {
				s.logger.Error(err)
				return errors.IdentifyStorageError(err)
			}
			name = *unit.NewText
		}

		if len(unit.PropertiesID) != 0 {
			err = s.storage.SetUnitProperties(layer, unit.Lang, name, unit.PropertiesID)
			if err != nil {
				s.logger.Error(err)
				return errors.IdentifyStorageError(err)
			}
		}
	}

	return nil
}

func (s *Service) DeleteUnits(layer string, unitDTO interfaces.SearchUnitDTO) error {
	err := s.storage.DeleteUnits(layer, unitDTO.Lang, unitDTO.Text)
	if err != nil {
		s.logger.Error(err)
		return errors.IdentifyStorageError(err)
	}
	return nil
}

func (s *Service) SaveProperties(propertiesDTO interfaces.PropertyNamesDTO) (interfaces.PropertiesIdDTO, error) {
	propertiesID, err := s.storage.SaveProperties(propertiesDTO.Properties)
	if err != nil {
		s.logger.Error(err)
		return interfaces.PropertiesIdDTO{}, errors.IdentifyStorageError(err)
	}

	return interfaces.PropertiesIdDTO{Properties: propertiesID}, nil
}

func (s *Service) SaveModels(layer string, modelsDTO interfaces.ModelNamesDTO) (interfaces.ModelsIdDTO, error) {
	err := s.checkLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.ModelsIdDTO{}, err
	}

	for _, modelName := range modelsDTO.Models {
		model := entities.Model{Name: modelName}
		err = model.IsValidName()
		if err != nil {
			s.logger.Error(err)
			return interfaces.ModelsIdDTO{}, err
		}
	}

	modelsID, err := s.storage.SaveModels(layer, modelsDTO.Models)
	if err != nil {
		s.logger.Error(err)
		return interfaces.ModelsIdDTO{}, errors.IdentifyStorageError(err)
	}

	return interfaces.ModelsIdDTO{Models: modelsID}, nil
}

func (s *Service) SaveModelElements(layer string, modelElementsDTO interfaces.ModelElementNamesDTO) (interfaces.ModelElementsIdDTO, error) {
	err := s.checkLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return interfaces.ModelElementsIdDTO{}, err
	}

	modelElementsID, err := s.storage.SaveModelElements(layer, modelElementsDTO.ModelElements)
	if err != nil {
		s.logger.Error(err)
		return interfaces.ModelElementsIdDTO{}, errors.IdentifyStorageError(err)
	}

	return interfaces.ModelElementsIdDTO{ModelElements: modelElementsID}, nil
}

func (s *Service) SaveLayer(layer string) error {
	err := s.storage.SaveLayer(layer)
	if err != nil {
		s.logger.Error(err)
		return errors.IdentifyStorageError(err)
	}

	return nil
}

func (s *Service) AddUser(user interfaces.UserDTO) error {
	err := s.storage.AddUser(user)
	if err != nil {
		s.logger.Error(err)
		return errors.IdentifyStorageError(err)
	}
	s.logger.Infof("storage user (name: %s, role: %s) inserted to database", user.Name, user.Role)

	return nil
}

func (s *Service) GetUser(request entities.AuthRequest) (entities.User, error) {
	user, err := s.storage.GetUser(request)
	if err != nil {
		s.logger.Error(err)
		return entities.User{}, errors.IdentifyStorageError(err)
	}

	return user, nil
}
