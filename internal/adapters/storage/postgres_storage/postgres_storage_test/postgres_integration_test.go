package postgres_storage_test

import (
	"context"
	"github.com/Inspirate789/Thermy-backend/internal/adapters/storage/postgres_storage"
	"github.com/Inspirate789/Thermy-backend/internal/domain/entities"
	"github.com/Inspirate789/Thermy-backend/internal/domain/interfaces"
	"github.com/Inspirate789/Thermy-backend/internal/domain/services/storage"
	"github.com/Inspirate789/go-randomdata"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"testing"
)

var (
	mockLogger    *log.Logger
	pgConn        storage.ConnDB
	pgStorage     = postgres_storage.NewPostgresStorage()
	testLayerName = randomdata.SillyName()
)

func TestPgStorageService_AddUser(t *testing.T) {
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
			ss:   storage.NewStorageService(pgStorage, mockLogger),
			args: args{
				conn: pgConn,
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

func TestPgStorageService_GetAllUnits(t *testing.T) {
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
			ss:   storage.NewStorageService(pgStorage, mockLogger),
			args: args{
				conn:  pgConn,
				layer: testLayerName,
			},
			want: interfaces.OutputUnitsDTO{
				Units:    make([]map[string]interfaces.OutputUnitDTO, 0),
				Contexts: nil,
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

func TestPgStorageService_GetLayers(t *testing.T) {
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
			ss:      storage.NewStorageService(pgStorage, mockLogger),
			args:    args{conn: pgConn},
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

func TestPgStorageService_GetModelElements(t *testing.T) {
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
			ss:   storage.NewStorageService(pgStorage, mockLogger),
			args: args{
				conn:  pgConn,
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

func TestPgStorageService_GetModels(t *testing.T) {
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
			ss:   storage.NewStorageService(pgStorage, mockLogger),
			args: args{
				conn:  pgConn,
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

func TestPgStorageService_GetProperties(t *testing.T) {
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
			ss:      storage.NewStorageService(pgStorage, mockLogger),
			args:    args{conn: pgConn},
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

func TestPgStorageService_GetPropertiesByUnit(t *testing.T) {
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
			name: "Simple positive test",
			ss:   storage.NewStorageService(pgStorage, mockLogger),
			args: args{
				conn:  pgConn,
				layer: testLayerName,
				unit:  interfaces.SearchUnitDTO{Lang: "ru"},
			},
			want: interfaces.OutputPropertiesDTO{
				Properties: []interfaces.OutputPropertyDTO{},
			},
			wantErr: false,
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

func TestPgStorageService_GetUnitsByModels(t *testing.T) {
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
			ss:   storage.NewStorageService(pgStorage, mockLogger),
			args: args{
				conn:  pgConn,
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

func TestPgStorageService_GetUnitsByProperties(t *testing.T) {
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
			ss:   storage.NewStorageService(pgStorage, mockLogger),
			args: args{
				conn:          pgConn,
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

//func TestPgStorageService_GetUserPassword(t *testing.T) {
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
//			ss:   storage.NewStorageService(pgStorage, mockLogger),
//			args: args{
//				conn:     pgConn,
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

func TestPgStorageService_SaveModelElements(t *testing.T) {
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
			ss:   storage.NewStorageService(pgStorage, mockLogger),
			args: args{
				conn:             pgConn,
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

func TestPgStorageService_SaveModels(t *testing.T) {
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
			ss:   storage.NewStorageService(pgStorage, mockLogger),
			args: args{
				conn:      pgConn,
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

func TestPgStorageService_SaveProperties(t *testing.T) {
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
			ss:   storage.NewStorageService(pgStorage, mockLogger),
			args: args{
				conn:          pgConn,
				propertiesDTO: interfaces.PropertyNamesDTO{Properties: []string{}},
			},
			want:    interfaces.PropertiesIdDTO{Properties: nil},
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

func TestPgStorageService_SaveUnits(t *testing.T) {
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
			ss:   storage.NewStorageService(pgStorage, mockLogger),
			args: args{
				conn:  pgConn,
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

func TestPgStorageService_UpdateUnits(t *testing.T) {
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
			ss:   storage.NewStorageService(pgStorage, mockLogger),
			args: args{
				conn:  pgConn,
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
	err := godotenv.Load("../../../../../.env")
	if err != nil {
		log.Fatal("No .env file found")
	}

	pgConn, _, err = pgStorage.OpenConn(&entities.AuthRequest{
		Username: os.Getenv("POSTGRES_ADMIN_USERNAME"),
		Password: os.Getenv("POSTGRES_ADMIN_PASSWORD"),
	}, context.Background())
	if err != nil {
		log.Fatal(err)
	}

	mockLogger = log.New()
	mockLogger.SetOutput(io.Discard)

	err = storage.NewStorageService(pgStorage, mockLogger).SaveLayer(pgConn, testLayerName)
	if err != nil {
		log.Fatal(err)
	}
}

func shutdown() {
	_ = pgStorage.CloseConn(pgConn)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}
