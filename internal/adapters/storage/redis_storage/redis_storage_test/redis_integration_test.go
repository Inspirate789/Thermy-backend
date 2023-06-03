package redis_storage_test

import (
	"context"
	"github.com/Inspirate789/Thermy-backend/internal/adapters/storage/redis_storage"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	mockLogger   *log.Logger
	redisConn    storage.ConnDB
	redisStorage = redis_storage.NewRedisStorage(
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
		os.Getenv("REDIS_PASSWORD"),
	)
	testLayerName = randomdata.SillyName()
)

func TestRedisStorageService_AddUser(t *testing.T) {
	type args struct {
		conn storage.ConnDB
		user interfaces.UserDTO
	}
	tests := []struct {
		name    string
		ss      *storage.StorageService
		args    args
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss:   storage.NewStorageService(redisStorage, mockLogger),
			args: args{
				conn: redisConn,
				user: interfaces.UserDTO{
					Name:     randomdata.SillyName(),
					Password: randomdata.IpV6Address(),
					Role:     "admin",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ss.AddUser(tt.args.conn, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisStorageService_GetAllUnits(t *testing.T) {
	type args struct {
		conn  storage.ConnDB
		layer string
	}
	tests := []struct {
		name    string
		ss      *storage.StorageService
		args    args
		want    interfaces.OutputUnitsDTO
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss:   storage.NewStorageService(redisStorage, mockLogger),
			args: args{
				conn:  redisConn,
				layer: testLayerName,
			},
			want: interfaces.OutputUnitsDTO{
				Units:    make([]map[string]interfaces.OutputUnitDTO, 0),
				Contexts: make([]interfaces.ContextDTO, 0),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.GetAllUnits(tt.args.conn, tt.args.layer)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllUnits() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRedisStorageService_GetLayers(t *testing.T) {
	type args struct {
		conn storage.ConnDB
	}
	tests := []struct {
		name    string
		ss      *storage.StorageService
		args    args
		want    interfaces.LayersDTO
		wantErr bool
	}{
		{
			name:    "Simple positive test",
			ss:      storage.NewStorageService(redisStorage, mockLogger),
			args:    args{conn: redisConn},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.GetLayers(tt.args.conn)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLayers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Layers == nil {
				t.Errorf("GetLayers() got = %v, want not zero value", got)
				return
			}
		})
	}
}

func TestRedisStorageService_GetModelElements(t *testing.T) {
	type args struct {
		conn  storage.ConnDB
		layer string
	}
	tests := []struct {
		name    string
		ss      *storage.StorageService
		args    args
		want    interfaces.OutputModelElementsDTO
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss:   storage.NewStorageService(redisStorage, mockLogger),
			args: args{
				conn:  redisConn,
				layer: testLayerName,
			},
			want:    interfaces.OutputModelElementsDTO{Elements: []interfaces.OutputModelElementDTO{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.GetModelElements(tt.args.conn, tt.args.layer)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetModelElements() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRedisStorageService_GetModels(t *testing.T) {
	type args struct {
		conn  storage.ConnDB
		layer string
	}
	tests := []struct {
		name    string
		ss      *storage.StorageService
		args    args
		want    interfaces.OutputModelsDTO
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss:   storage.NewStorageService(redisStorage, mockLogger),
			args: args{
				conn:  redisConn,
				layer: testLayerName,
			},
			want:    interfaces.OutputModelsDTO{Models: []interfaces.OutputModelDTO{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.GetModels(tt.args.conn, tt.args.layer)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetModels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRedisStorageService_GetProperties(t *testing.T) {
	type args struct {
		conn storage.ConnDB
	}
	tests := []struct {
		name    string
		ss      *storage.StorageService
		args    args
		wantErr bool
	}{
		{
			name:    "Simple positive test",
			ss:      storage.NewStorageService(redisStorage, mockLogger),
			args:    args{conn: redisConn},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.GetProperties(tt.args.conn)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProperties() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotEqual(t, got.Properties, nil)
		})
	}
}

func TestRedisStorageService_GetPropertiesByUnit(t *testing.T) {
	type args struct {
		conn  storage.ConnDB
		layer string
		unit  interfaces.SearchUnitDTO
	}
	tests := []struct {
		name    string
		ss      *storage.StorageService
		args    args
		want    interfaces.OutputPropertiesDTO
		wantErr bool
	}{
		{
			name: "Simple negative test",
			ss:   storage.NewStorageService(redisStorage, mockLogger),
			args: args{
				conn:  redisConn,
				layer: testLayerName,
				unit:  interfaces.SearchUnitDTO{Lang: "ru", Text: "non-existing unit"},
			},
			want:    interfaces.OutputPropertiesDTO{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.GetPropertiesByUnit(tt.args.conn, tt.args.layer, tt.args.unit)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPropertiesByUnit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRedisStorageService_GetUnitsByModels(t *testing.T) {
	type args struct {
		conn      storage.ConnDB
		layer     string
		modelsDTO interfaces.ModelsIdDTO
	}
	tests := []struct {
		name    string
		ss      *storage.StorageService
		args    args
		want    interfaces.OutputUnitsDTO
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss:   storage.NewStorageService(redisStorage, mockLogger),
			args: args{
				conn:  redisConn,
				layer: testLayerName,
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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.GetUnitsByModels(tt.args.conn, tt.args.layer, tt.args.modelsDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUnitsByModels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRedisStorageService_GetUnitsByProperties(t *testing.T) {
	type args struct {
		conn          storage.ConnDB
		layer         string
		propertiesDTO interfaces.PropertiesIdDTO
	}
	tests := []struct {
		name    string
		ss      *storage.StorageService
		args    args
		want    interfaces.OutputUnitsDTO
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss:   storage.NewStorageService(redisStorage, mockLogger),
			args: args{
				conn:          redisConn,
				layer:         testLayerName,
				propertiesDTO: interfaces.PropertiesIdDTO{Properties: []int{}},
			},
			want: interfaces.OutputUnitsDTO{
				Units:    make(interfaces.UnitDtoMaps, 0),
				Contexts: []interfaces.ContextDTO{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.GetUnitsByProperties(tt.args.conn, tt.args.layer, tt.args.propertiesDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUnitsByProperties() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

//func TestRedisStorageService_GetUserPassword(t *testing.T) {
//	type args struct {
//		conn     storage.ConnDB
//		username string
//	}
//	tests := []struct {
//		name    string
//		ss      *storage.StorageService
//		args    args
//		want    string
//		wantErr bool
//	}{
//		{
//			name: "Simple positive test",
//			ss:   storage.NewStorageService(redisStorage, mockLogger),
//			args: args{
//				conn:     redisConn,
//				username: os.Getenv("POSTGRES_ADMIN_USERNAME"),
//			},
//			want:    os.Getenv("POSTGRES_ADMIN_PASSWORD"),
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
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
//	}
//}

func TestRedisStorageService_SaveModelElements(t *testing.T) {
	type args struct {
		conn             storage.ConnDB
		layer            string
		modelElementsDTO interfaces.ModelElementNamesDTO
	}
	tests := []struct {
		name    string
		ss      *storage.StorageService
		args    args
		want    interfaces.ModelElementsIdDTO
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss:   storage.NewStorageService(redisStorage, mockLogger),
			args: args{
				conn:             redisConn,
				layer:            testLayerName,
				modelElementsDTO: interfaces.ModelElementNamesDTO{ModelElements: []string{}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.SaveModelElements(tt.args.conn, tt.args.layer, tt.args.modelElementsDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveModelElements() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.ModelElements == nil {
				t.Errorf("SaveModelElements() got = %v, want not zero value", got)
				return
			}
		})
	}
}

func TestRedisStorageService_SaveModels(t *testing.T) {
	type args struct {
		conn      storage.ConnDB
		layer     string
		modelsDTO interfaces.ModelNamesDTO
	}
	tests := []struct {
		name    string
		ss      *storage.StorageService
		args    args
		want    interfaces.ModelsIdDTO
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss:   storage.NewStorageService(redisStorage, mockLogger),
			args: args{
				conn:      redisConn,
				layer:     testLayerName,
				modelsDTO: interfaces.ModelNamesDTO{Models: []string{}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.SaveModels(tt.args.conn, tt.args.layer, tt.args.modelsDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveModels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Models == nil {
				t.Errorf("SaveModels() got = %v, want not zero value", got)
				return
			}
		})
	}
}

func TestRedisStorageService_SaveProperties(t *testing.T) {
	type args struct {
		conn          storage.ConnDB
		propertiesDTO interfaces.PropertyNamesDTO
	}
	tests := []struct {
		name    string
		ss      *storage.StorageService
		args    args
		want    interfaces.PropertiesIdDTO
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss:   storage.NewStorageService(redisStorage, mockLogger),
			args: args{
				conn:          redisConn,
				propertiesDTO: interfaces.PropertyNamesDTO{Properties: []string{}},
			},
			want:    interfaces.PropertiesIdDTO{Properties: make([]int, 0)},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.ss.SaveProperties(tt.args.conn, tt.args.propertiesDTO)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveProperties() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRedisStorageService_SaveUnits(t *testing.T) {
	type args struct {
		conn     storage.ConnDB
		layer    string
		unitsDTO interfaces.SaveUnitsDTO
	}
	tests := []struct {
		name    string
		ss      *storage.StorageService
		args    args
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss:   storage.NewStorageService(redisStorage, mockLogger),
			args: args{
				conn:  redisConn,
				layer: testLayerName,
				unitsDTO: interfaces.SaveUnitsDTO{
					Contexts: make(map[string]string),
					Units:    make([]map[string]interfaces.SaveUnitDTO, 0),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ss.SaveUnits(tt.args.conn, tt.args.layer, tt.args.unitsDTO); (err != nil) != tt.wantErr {
				t.Errorf("SaveUnits() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisStorageService_UpdateUnits(t *testing.T) {
	type args struct {
		conn     storage.ConnDB
		layer    string
		unitsDTO interfaces.UpdateUnitsDTO
	}
	tests := []struct {
		name    string
		ss      *storage.StorageService
		args    args
		wantErr bool
	}{
		{
			name: "Simple positive test",
			ss:   storage.NewStorageService(redisStorage, mockLogger),
			args: args{
				conn:  redisConn,
				layer: testLayerName,
				unitsDTO: interfaces.UpdateUnitsDTO{
					Units: []interfaces.UpdateUnitDTO{},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ss.UpdateUnits(tt.args.conn, tt.args.layer, tt.args.unitsDTO); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUnits() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func setup() {
	err := redisStorage.AddUser(nil, interfaces.UserDTO{
		Name:     os.Getenv("POSTGRES_ADMIN_USERNAME"),
		Password: os.Getenv("POSTGRES_ADMIN_PASSWORD"),
		Role:     "admin",
	})
	if err != nil {
		log.Fatal(err)
	}
	redisConn, _, err = redisStorage.OpenConn(&entities.AuthRequest{
		Username: os.Getenv("POSTGRES_ADMIN_USERNAME"),
		Password: os.Getenv("POSTGRES_ADMIN_PASSWORD"),
	}, context.Background())
	if err != nil {
		log.Fatal(err)
	}

	mockLogger = log.New()
	mockLogger.SetOutput(os.Stdout)

	err = storage.NewStorageService(redisStorage, mockLogger).SaveLayer(redisConn, testLayerName)
	if err != nil {
		log.Fatal(err)
	}
}

func shutdown() {
	_ = redisStorage.CloseConn(redisConn)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
