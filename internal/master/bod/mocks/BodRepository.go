package mocks

import (
	"github.com/jadahbakar/kawalrencanamu-backend/internal/master/bod"
	"github.com/stretchr/testify/mock"
)

type BodRepository struct {
	mock.Mock
}

func (_m *BodRepository) SearchAll() ([]bod.Bod, error) {
	ret := _m.Called()
	var r0 []bod.Bod
	if rf, ok := ret.Get(0).(func() ([]bod.Bod, error)); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bod.Bod)
		}
	}
}
