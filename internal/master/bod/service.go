package bod

import "go.uber.org/zap"

type bodService struct {
	bodRepo BodRepository
}

func NewBodService(b BodRepository) BodService {
	return &bodService{bodRepo: b}
}

func (bs *bodService) FindAll() ([]Bod, error) {
	// return nil, nil
	res, err := bs.bodRepo.SearchAll()
	if err != nil {
		zap.Error(err)
		return nil, err
	}
	return res, nil
}

func (bs *bodService) FindById(Id int) (res Bod, err error) {
	res, err = bs.bodRepo.SearchById(Id)
	if err != nil {
		zap.Error(err)
		return
	}
	return
}
