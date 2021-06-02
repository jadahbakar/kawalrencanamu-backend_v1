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

type MasterRepository struct {
	db *pgx.Conn
}

func NewMasterRepository(db *pgx.Conn) entities.AssesmentEnvironmentRepository {
	return &MasterRepository{db: db}
}

func (r MasterRepository) GetAssEnvById() (*entities.AssesmentEnvironment, error) {
	return &entities.AssesmentEnvironment{}, nil
}
