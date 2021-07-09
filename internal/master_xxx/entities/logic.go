package entities

import "errors"

var (
	ErrRedirectNotFound = errors.New("Redirect Not Found")
	ErrRedirectInvalid  = errors.New("Redirect Invalid")
)

type masterService struct {
	masterRepo MasterRepository
}

func NewMasterService(masterRepo MasterRepository) MasterService {
	return &masterService{masterRepo}
}

func (m *masterService) Find(id string) (*AssesmentEnvironment, error) {
	return m.masterRepo.Find(id)
}

func (m *masterService) Store(nama string) (*AssesmentEnvironment, error) {
	return m.masterRepo.Store(nama)
}
