package repository

import (
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/jadahbakar/kawalrencanamu-backend/internal/master/entities"
)

var (
	ErrRepository = errors.New("Unable to Handle Request")
	ErrNotFound   = errors.New("Assesment Environment Not Found")
)

type postgresRepository struct {
	db *pgx.Conn
}

func (pgRepo *postgresRepository) Find(id string) (*entities.AssesmentEnvironment, error) {
	return nil, nil
}

func (pgRepo *postgresRepository) Strore(nama string) (*entities.AssesmentEnvironment, error) {
	return nil, nil

}

// func NewMasterRepository(db *pgx.Conn) entities.AssesmentEnvironmentRepository {
// 	return &MasterRepository{db: db}
// }

// func (r MasterRepository) GetAssEnvById() (*entities.AssesmentEnvironment, error) {
// 	return &entities.AssesmentEnvironment{}, nil
// }
