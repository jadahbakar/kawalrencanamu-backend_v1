package entities

type MasterRepository interface {
	Find(id string) (*AssesmentEnvironment, error)
	Store(nama string) (*AssesmentEnvironment, error)
}
