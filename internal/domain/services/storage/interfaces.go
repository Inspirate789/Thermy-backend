package storage

import "context"

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ConnDB any

type Storage interface { // TODO: split?
	OpenConn(request *AuthRequest, ctx context.Context) (ConnDB, string, error) // Get conn, role in database and error
	UsersRepository
	ModelsRepository
	ModelElementsRepository
	PropertiesRepository
	UnitsRepository
	LayersRepository
	CloseConn(ConnDB) error
}
