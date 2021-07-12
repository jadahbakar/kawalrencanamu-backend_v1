package bod_test

import (
	"errors"
	"testing"

	"github.com/jadahbakar/kawalrencanamu-backend/internal/master/bod"
	"github.com/jadahbakar/kawalrencanamu-backend/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup() (mockBodRepo *mocks.BodRepository, srv bod.BodService) {
	mockBodRepo = new(mocks.BodRepository)
	srv = bod.NewBodService(mockBodRepo)
	return
}

func TestFindAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// setup
		repo, srv := setup()
		actual := []bod.Bod{}
		repo.On("SearchAll").Return(actual, nil)

		// testing
		data, err := srv.FindAll()

		// assertion
		assert.NoError(t, err)
		assert.Equal(t, actual, data)

		// end mock
		repo.AssertExpectations(t)
	})

	t.Run("Error", func(t *testing.T) {
		// setup
		repo, srv := setup()
		actual := []bod.Bod{}
		repo.On("SearchAll").Return(actual, errors.New("err"))

		// testing
		_, err := srv.FindAll()

		// assertion
		assert.Error(t, err)

		// end mock
		repo.AssertExpectations(t)
	})
}

func TestFindById(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		// setup
		repo, srv := setup()
		actual := bod.Bod{}
		repo.On("SearchById", 1).Return(actual, nil)

		// testing
		data, err := srv.FindById(1)

		// assertion
		assert.NoError(t, err)
		assert.Equal(t, actual, data)

		// end mock
		repo.AssertExpectations(t)
	})

	// t.Run("Error", func(t *testing.T) {
	// 	// setup
	// 	repo, srv := setup()
	// 	repo.On("SearchById", 1).Return(nil, errors.New("err"))

	// 	// testing
	// 	_, err := srv.FindById(1)

	// 	// assertion
	// 	assert.Error(t, err)

	// 	// end mock
	// 	repo.AssertExpectations(t)
	// })
	t.Run("Error", func(t *testing.T) {
		mockBodRepo := new(mocks.BodRepository)
		// mockBod := bod.Bod{
		// 	BodId:       1,
		// 	BodCallSign: "Call Sign",
		// 	BodNamaId:   "Nama",
		// 	BodNamaEn:   "Name",
		// }
		mockBodRepo.On("SearchById", mock.Anything, mock.AnythingOfType("string")).Return(bod.Bod{}, errors.New("Unexpected")).Once()

		srv := bod.NewBodService(mockBodRepo)
		paramId := 1

		a, err := srv.FindById(paramId)

		assert.Error(t, err)
		assert.Equal(t, bod.Bod{}, a)

		mockBodRepo.AssertExpectations(t)
		mockBodRepo.AssertExpectations(t)
	})
}
