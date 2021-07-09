package entities

type MasterService interface {
	Find(id string) (*AssesmentEnvironment, error)
	Store(nama string) (*AssesmentEnvironment, error)
}
