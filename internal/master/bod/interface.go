package bod

type BodRepository interface {
	SearchAll() ([]Bod, error)
	SearchById(Id int) (Bod, error)
}

type BodService interface {
	FindAll() ([]Bod, error)
	FindById(id string) (Bod, error)
}
