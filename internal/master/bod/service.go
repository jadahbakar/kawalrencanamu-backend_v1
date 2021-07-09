package bod

type bodService struct {
	bodRepo BodRepository
}

func NewBodService(b BodRepository) BodService {
	return &bodService{bodRepo: b}
}

func (bs *bodService) FindAll() ([]*Bod, error) {
	return nil, nil
}

func (bs *bodService) FindById(id string) (*Bod, error) {
	return nil, nil
}
