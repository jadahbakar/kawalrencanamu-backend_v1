package repository

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jadahbakar/kawalrencanamu-backend/internal/master/entities"
)

type MasterRepository struct {
	client *pgxpool.Pool
}

func NewMasterRepository(client *pgxpool.Pool) entities.AssesmentEnvironmentRepository {
	return &MasterRepository{client}
}

func (r *MasterRepository) Fetch() (result []entities.AssesmentEnvironment, err error) {

}
