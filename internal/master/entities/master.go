package entities

type AssesmentEnvironment struct {
	id   int    `json:"assesment_environment_id"`
	nama string `json:"assesment_environment_nama"`
}

type AssesmentEnvironmentEntity interface {
	Fetch()
}

type AssesmentEnvironmentRepository interface {
	Fetch()
}
