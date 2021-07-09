package propinsi

type PropinsiRepository interface {
	FindAll() ([]*Propinsi, error)
	FindById(Id int) (*Propinsi, error)
}

type PropinsiService interface {
	GetAll() ([]*Propinsi, error)
	GetById(Id int) (*Propinsi, error)
}
