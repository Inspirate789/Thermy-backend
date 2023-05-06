package storage

import (
	"context"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"io"
	"reflect"
	"testing"
)

func TestNewStorageService(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockStorage := new(MockStorage)

	type args struct {
		storage Storage
		logger  *log.Logger
	}
	tests := []struct {
		name string
		args args
		want *StorageService
	}{
		{
			name: "Simple positive test",
			args: args{
				storage: mockStorage,
				logger:  mockLogger,
			},
			want: &StorageService{
				storage: mockStorage,
				logger:  mockLogger,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewStorageService(tt.args.storage, tt.args.logger)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStorageService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorageService_AddUser(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockStorage := new(MockStorage)
	mockStorage.On("AddUser", mock.Anything, interfaces.UserDTO{
		Name:     "initial_admin",
		Password: "abcdefgh",
		Role:     "admin",
	}).Return(nil)

	type args struct {
		conn ConnDB
		user interfaces.UserDTO
	}
	tests := []struct {
		name    string
		ss      *StorageService
		args    args
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss: &StorageService{
				storage: mockStorage,
				logger:  mockLogger,
			},
			args: args{
				conn: nil,
				user: interfaces.UserDTO{
					Name:     "initial_admin",
					Password: "abcdefgh",
					Role:     "admin",
				},
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ss.AddUser(tt.args.conn, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		mockStorage.AssertNumberOfCalls(t, "AddUser", i+1)
	}
}

func TestStorageService_CloseConn(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockStorage := new(MockStorage)
	mockStorage.On("CloseConn", mock.Anything).Return(nil)

	type args struct {
		conn ConnDB
	}
	tests := []struct {
		name    string
		ss      *StorageService
		args    args
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss: &StorageService{
				storage: mockStorage,
				logger:  mockLogger,
			},
			args:    args{conn: nil},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ss.CloseConn(tt.args.conn); (err != nil) != tt.wantErr {
				t.Errorf("CloseConn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		mockStorage.AssertNumberOfCalls(t, "CloseConn", i+1)
	}
}

func TestStorageService_GetAllUnits(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockStorage := new(MockStorage)
	mockStorage.On("GetAllUnits", mock.Anything, "test_layer").Return(interfaces.OutputUnitsDTO{}, nil)
	mockStorage.On("LayerExist", mock.Anything, "test_layer").Return(true, nil)

	type args struct {
		conn  ConnDB
		layer string
	}
	tests := []struct {
		name    string
		ss      *StorageService
		args    args
		want    interfaces.OutputUnitsDTO
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss: &StorageService{
				storage: mockStorage,
				logger:  mockLogger,
			},
			args: args{
				conn:  nil,
				layer: "test_layer",
			},
			want:    interfaces.OutputUnitsDTO{},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.GetAllUnits(tt.args.conn, tt.args.layer)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUnits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllUnits() got = %v, want %v", got, tt.want)
			}
		})

		mockStorage.AssertNumberOfCalls(t, "GetAllUnits", i+1)
		mockStorage.AssertNumberOfCalls(t, "LayerExist", i+1)
	}
}

func TestStorageService_GetLayers(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockStorage := new(MockStorage)
	mockStorage.On("GetAllLayers", mock.Anything).Return([]string{}, nil)

	type args struct {
		conn ConnDB
	}
	tests := []struct {
		name    string
		ss      *StorageService
		args    args
		want    interfaces.LayersDTO
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss: &StorageService{
				storage: mockStorage,
				logger:  mockLogger,
			},
			args:    args{conn: nil},
			want:    interfaces.LayersDTO{Layers: []string{}},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.GetLayers(tt.args.conn)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLayers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLayers() got = %v, want %v", got, tt.want)
			}
		})

		mockStorage.AssertNumberOfCalls(t, "GetAllLayers", i+1)
	}
}

func TestStorageService_GetModelElements(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockStorage := new(MockStorage)
	mockStorage.On("GetAllModelElements", mock.Anything, "test_layer").Return([]entities.ModelElement{}, nil)
	mockStorage.On("LayerExist", mock.Anything, "test_layer").Return(true, nil)

	type args struct {
		conn  ConnDB
		layer string
	}
	tests := []struct {
		name    string
		ss      *StorageService
		args    args
		want    interfaces.OutputModelElementsDTO
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss: &StorageService{
				storage: mockStorage,
				logger:  mockLogger,
			},
			args: args{
				conn:  nil,
				layer: "test_layer",
			},
			want:    interfaces.OutputModelElementsDTO{Elements: []interfaces.OutputModelElementDTO{}},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.GetModelElements(tt.args.conn, tt.args.layer)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetModelElements() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetModelElements() got = %v, want %v", got, tt.want)
			}
		})

		mockStorage.AssertNumberOfCalls(t, "GetAllModelElements", i+1)
		mockStorage.AssertNumberOfCalls(t, "LayerExist", i+1)
	}
}

func TestStorageService_GetModels(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockStorage := new(MockStorage)
	mockStorage.On("GetAllModels", mock.Anything, "test_layer").Return([]entities.Model{}, nil)
	mockStorage.On("LayerExist", mock.Anything, "test_layer").Return(true, nil)

	type args struct {
		conn  ConnDB
		layer string
	}
	tests := []struct {
		name    string
		ss      *StorageService
		args    args
		want    interfaces.OutputModelsDTO
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss: &StorageService{
				storage: mockStorage,
				logger:  mockLogger,
			},
			args: args{
				conn:  nil,
				layer: "test_layer",
			},
			want:    interfaces.OutputModelsDTO{Models: []interfaces.OutputModelDTO{}},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.GetModels(tt.args.conn, tt.args.layer)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetModels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetModels() got = %v, want %v", got, tt.want)
			}
		})

		mockStorage.AssertNumberOfCalls(t, "GetAllModels", i+1)
		mockStorage.AssertNumberOfCalls(t, "LayerExist", i+1)
	}
}

func TestStorageService_GetProperties(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockStorage := new(MockStorage)
	mockStorage.On("GetAllProperties", mock.Anything).Return([]entities.Property{}, nil)

	type args struct {
		conn ConnDB
	}
	tests := []struct {
		name    string
		ss      *StorageService
		args    args
		want    interfaces.OutputPropertiesDTO
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss: &StorageService{
				storage: mockStorage,
				logger:  mockLogger,
			},
			args:    args{conn: nil},
			want:    interfaces.OutputPropertiesDTO{Properties: []interfaces.OutputPropertyDTO{}},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.GetProperties(tt.args.conn)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProperties() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProperties() got = %v, want %v", got, tt.want)
			}
		})

		mockStorage.AssertNumberOfCalls(t, "GetAllProperties", i+1)
	}
}

func TestStorageService_GetPropertiesByUnit(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockStorage := new(MockStorage)
	mockStorage.
		On("GetPropertiesByUnit", mock.Anything, "test_layer", interfaces.SearchUnitDTO{}).
		Return([]entities.Property{}, nil)

	type args struct {
		conn  ConnDB
		layer string
		unit  interfaces.SearchUnitDTO
	}
	tests := []struct {
		name    string
		ss      *StorageService
		args    args
		want    interfaces.OutputPropertiesDTO
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss: &StorageService{
				storage: mockStorage,
				logger:  mockLogger,
			},
			args: args{
				conn:  nil,
				layer: "test_layer",
				unit:  interfaces.SearchUnitDTO{},
			},
			want: interfaces.OutputPropertiesDTO{
				Properties: []interfaces.OutputPropertyDTO{},
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.GetPropertiesByUnit(tt.args.conn, tt.args.layer, tt.args.unit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPropertiesByUnit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPropertiesByUnit() got = %v, want %v", got, tt.want)
			}
		})

		mockStorage.AssertNumberOfCalls(t, "GetPropertiesByUnit", i+1)
	}
}

func TestStorageService_GetUnitsByModels(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockStorage := new(MockStorage)
	mockStorage.On("GetUnitsByModels", mock.Anything, "test_layer", []int{}).Return(interfaces.OutputUnitsDTO{
		Units:    make(interfaces.UnitDtoMaps, 0),
		Contexts: []interfaces.ContextDTO{},
	}, nil)
	mockStorage.On("LayerExist", mock.Anything, "test_layer").Return(true, nil)

	type args struct {
		conn      ConnDB
		layer     string
		modelsDTO interfaces.ModelsIdDTO
	}
	tests := []struct {
		name    string
		ss      *StorageService
		args    args
		want    interfaces.OutputUnitsDTO
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss: &StorageService{
				storage: mockStorage,
				logger:  mockLogger,
			},
			args: args{
				conn:  nil,
				layer: "test_layer",
				modelsDTO: interfaces.ModelsIdDTO{
					Models: []int{},
				},
			},
			want: interfaces.OutputUnitsDTO{
				Units:    make(interfaces.UnitDtoMaps, 0),
				Contexts: []interfaces.ContextDTO{},
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.GetUnitsByModels(tt.args.conn, tt.args.layer, tt.args.modelsDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUnitsByModels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUnitsByModels() got = %v, want %v", got, tt.want)
			}
		})

		mockStorage.AssertNumberOfCalls(t, "GetUnitsByModels", 0)
		mockStorage.AssertNumberOfCalls(t, "LayerExist", i+1)
	}
}

func TestStorageService_GetUnitsByProperties(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockStorage := new(MockStorage)
	mockStorage.On("GetUnitsByProperties", mock.Anything, "test_layer", []int{}).Return(interfaces.OutputUnitsDTO{
		Units:    make(interfaces.UnitDtoMaps, 0),
		Contexts: []interfaces.ContextDTO{},
	}, nil)
	mockStorage.On("LayerExist", mock.Anything, "test_layer").Return(true, nil)

	type args struct {
		conn          ConnDB
		layer         string
		propertiesDTO interfaces.PropertiesIdDTO
	}
	tests := []struct {
		name    string
		ss      *StorageService
		args    args
		want    interfaces.OutputUnitsDTO
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss: &StorageService{
				storage: mockStorage,
				logger:  mockLogger,
			},
			args: args{
				conn:          nil,
				layer:         "test_layer",
				propertiesDTO: interfaces.PropertiesIdDTO{Properties: []int{}},
			},
			want: interfaces.OutputUnitsDTO{
				Units:    make(interfaces.UnitDtoMaps, 0),
				Contexts: []interfaces.ContextDTO{},
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.GetUnitsByProperties(tt.args.conn, tt.args.layer, tt.args.propertiesDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUnitsByProperties() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUnitsByProperties() got = %v, want %v", got, tt.want)
			}
		})

		mockStorage.AssertNumberOfCalls(t, "GetUnitsByProperties", 0)
		mockStorage.AssertNumberOfCalls(t, "LayerExist", i+1)
	}
}

//func TestStorageService_GetUserPassword(t *testing.T) {
//	mockLogger := log.New()
//	mockLogger.SetOutput(io.Discard)
//
//	mockStorage := new(MockStorage)
//	mockStorage.On("GetUserPassword", mock.Anything, "test_user").Return("test_password", nil)
//
//	type args struct {
//		conn     ConnDB
//		username string
//	}
//	tests := []struct {
//		name    string
//		ss      *StorageService
//		args    args
//		want    string
//		wantErr bool
//	}{
//		{
//			name: "Simple positive test",
//			ss: &StorageService{
//				storage: mockStorage,
//				logger:  mockLogger,
//			},
//			args: args{
//				conn:     nil,
//				username: "test_user",
//			},
//			want:    "test_password",
//			wantErr: false,
//		},
//	}
//	for i, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := tt.ss.GetUserPassword(tt.args.conn, tt.args.username)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetUserPassword() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if got != tt.want {
//				t.Errorf("GetUserPassword() got = %v, want %v", got, tt.want)
//			}
//		})
//
//		mockStorage.AssertNumberOfCalls(t, "GetUserPassword", i+1)
//	}
//}

func TestStorageService_OpenConn(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockStorage := new(MockStorage)
	mockStorage.
		On("OpenConn", &entities.AuthRequest{
			Username: "initial_admin",
			Password: "12345",
		}, mock.Anything).
		Return(nil, "admin", nil)

	type args struct {
		request *entities.AuthRequest
		ctx     context.Context
	}
	tests := []struct {
		name    string
		ss      *StorageService
		args    args
		want    ConnDB
		want1   string
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss: &StorageService{
				storage: mockStorage,
				logger:  mockLogger,
			},
			args: args{
				request: &entities.AuthRequest{
					Username: "initial_admin",
					Password: "12345",
				},
				ctx: context.Background(),
			},
			want:    nil,
			want1:   "admin",
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.ss.OpenConn(tt.args.request, tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("OpenConn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenConn() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("OpenConn() got1 = %v, want %v", got1, tt.want1)
			}
		})

		mockStorage.AssertNumberOfCalls(t, "OpenConn", i+1)
	}
}

func TestStorageService_SaveLayer(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockStorage := new(MockStorage)
	mockStorage.On("SaveLayer", mock.Anything, "test_layer").Return(nil)

	type args struct {
		conn  ConnDB
		layer string
	}
	tests := []struct {
		name    string
		ss      *StorageService
		args    args
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss: &StorageService{
				storage: mockStorage,
				logger:  mockLogger,
			},
			args: args{
				conn:  nil,
				layer: "test_layer",
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ss.SaveLayer(tt.args.conn, tt.args.layer); (err != nil) != tt.wantErr {
				t.Errorf("SaveLayer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		mockStorage.AssertNumberOfCalls(t, "SaveLayer", i+1)
	}
}

func TestStorageService_SaveModelElements(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockStorage := new(MockStorage)
	mockStorage.On("SaveModelElements", mock.Anything, "test_layer", []string{}).Return([]int{}, nil)
	mockStorage.On("LayerExist", mock.Anything, "test_layer").Return(true, nil)

	type args struct {
		conn             ConnDB
		layer            string
		modelElementsDTO interfaces.ModelElementNamesDTO
	}
	tests := []struct {
		name    string
		ss      *StorageService
		args    args
		want    interfaces.ModelElementsIdDTO
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss: &StorageService{
				storage: mockStorage,
				logger:  mockLogger,
			},
			args: args{
				conn:             nil,
				layer:            "test_layer",
				modelElementsDTO: interfaces.ModelElementNamesDTO{ModelElements: []string{}},
			},
			want:    interfaces.ModelElementsIdDTO{ModelElements: []int{}},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.SaveModelElements(tt.args.conn, tt.args.layer, tt.args.modelElementsDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveModelElements() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SaveModelElements() got = %v, want %v", got, tt.want)
			}
		})

		mockStorage.AssertNumberOfCalls(t, "SaveModelElements", i+1)
		mockStorage.AssertNumberOfCalls(t, "LayerExist", i+1)
	}
}

func TestStorageService_SaveModels(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockStorage := new(MockStorage)
	mockStorage.On("SaveModels", mock.Anything, "test_layer", []string{}).Return([]int{}, nil)
	mockStorage.On("LayerExist", mock.Anything, "test_layer").Return(true, nil)

	type args struct {
		conn      ConnDB
		layer     string
		modelsDTO interfaces.ModelNamesDTO
	}
	tests := []struct {
		name    string
		ss      *StorageService
		args    args
		want    interfaces.ModelsIdDTO
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss: &StorageService{
				storage: mockStorage,
				logger:  mockLogger,
			},
			args: args{
				conn:      nil,
				layer:     "test_layer",
				modelsDTO: interfaces.ModelNamesDTO{Models: []string{}},
			},
			want:    interfaces.ModelsIdDTO{Models: []int{}},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.SaveModels(tt.args.conn, tt.args.layer, tt.args.modelsDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveModels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SaveModels() got = %v, want %v", got, tt.want)
			}
		})

		mockStorage.AssertNumberOfCalls(t, "SaveModels", i+1)
		mockStorage.AssertNumberOfCalls(t, "LayerExist", i+1)
	}
}

func TestStorageService_SaveProperties(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockStorage := new(MockStorage)
	mockStorage.On("SaveProperties", mock.Anything, []string{}).Return([]int{}, nil)

	type args struct {
		conn          ConnDB
		propertiesDTO interfaces.PropertyNamesDTO
	}
	tests := []struct {
		name    string
		ss      *StorageService
		args    args
		want    interfaces.PropertiesIdDTO
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss: &StorageService{
				storage: mockStorage,
				logger:  mockLogger,
			},
			args: args{
				conn:          nil,
				propertiesDTO: interfaces.PropertyNamesDTO{Properties: []string{}},
			},
			want:    interfaces.PropertiesIdDTO{Properties: []int{}},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.SaveProperties(tt.args.conn, tt.args.propertiesDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveProperties() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SaveProperties() got = %v, want %v", got, tt.want)
			}
		})

		mockStorage.AssertNumberOfCalls(t, "SaveProperties", i+1)
	}
}

func TestStorageService_SaveUnits(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockStorage := new(MockStorage)
	mockStorage.On("SaveUnits", mock.Anything, "test_layer", interfaces.SaveUnitsDTO{
		Contexts: make(map[string]string),
		Units:    make([]map[string]interfaces.SaveUnitDTO, 0),
	}).Return(nil)
	mockStorage.On("LayerExist", mock.Anything, "test_layer").Return(true, nil)

	type args struct {
		conn     ConnDB
		layer    string
		unitsDTO interfaces.SaveUnitsDTO
	}
	tests := []struct {
		name    string
		ss      *StorageService
		args    args
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss: &StorageService{
				storage: mockStorage,
				logger:  mockLogger,
			},
			args: args{
				conn:  nil,
				layer: "test_layer",
				unitsDTO: interfaces.SaveUnitsDTO{
					Contexts: make(map[string]string),
					Units:    make([]map[string]interfaces.SaveUnitDTO, 0),
				},
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ss.SaveUnits(tt.args.conn, tt.args.layer, tt.args.unitsDTO); (err != nil) != tt.wantErr {
				t.Errorf("SaveUnits() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		mockStorage.AssertNumberOfCalls(t, "SaveUnits", i+1)
		mockStorage.AssertNumberOfCalls(t, "LayerExist", i+1)
	}
}

func TestStorageService_UpdateUnits(t *testing.T) {
	mockLogger := log.New()
	mockLogger.SetOutput(io.Discard)

	mockStorage := new(MockStorage)
	mockStorage.On("RenameUnit", mock.Anything, "test_layer", mock.Anything, mock.Anything).Return(nil)
	mockStorage.On("SetUnitProperties", mock.Anything, "test_layer", mock.Anything, mock.Anything).Return(nil)
	mockStorage.On("LayerExist", mock.Anything, "test_layer").Return(true, nil)

	type args struct {
		conn     ConnDB
		layer    string
		unitsDTO interfaces.UpdateUnitsDTO
	}
	tests := []struct {
		name    string
		ss      *StorageService
		args    args
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss: &StorageService{
				storage: mockStorage,
				logger:  mockLogger,
			},
			args: args{
				conn:  nil,
				layer: "test_layer",
				unitsDTO: interfaces.UpdateUnitsDTO{
					Units: []interfaces.UpdateUnitDTO{},
				},
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ss.UpdateUnits(tt.args.conn, tt.args.layer, tt.args.unitsDTO); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUnits() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		//mockStorage.AssertNumberOfCalls(t, "RenameUnit", i+1)
		//mockStorage.AssertNumberOfCalls(t, "SetUnitProperties", i+1)
		mockStorage.AssertNumberOfCalls(t, "LayerExist", i+1)
	}
}
