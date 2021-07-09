package entities

type AssesmentEnvironment struct {
	Id   int    `json:"assesment_environment_id" validate:"empty=false"`
	Nama string `json:"assesment_environment_nama"`
}
