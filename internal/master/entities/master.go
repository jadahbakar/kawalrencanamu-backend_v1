package entities

type AssesmentEnvironment struct {
	Id   int    `json:"assesment_environment_id"`
	Nama string `json:"assesment_environment_nama"`
}

type AssesmentEnvironmentRepository interface {
	GetAssEnvById() (*AssesmentEnvironment, error)
}
type AssesmentEnvironmentService interface {
	GetAssEnvById() (*AssesmentEnvironment, error)
}
