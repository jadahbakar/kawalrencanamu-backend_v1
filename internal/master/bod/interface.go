package bod

type BodRepository interface {
	SearchAll() ([]*Bod, error)
	SearchById(id string) (*Bod, error)
}

type BodService interface {
	FindAll() ([]*Bod, error)
	FindById(id string) (*Bod, error)
}
