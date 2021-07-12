package bod_test

import (
	"testing"

	"github.com/jadahbakar/kawalrencanamu-backend/internal/master/bod"
	"github.com/jadahbakar/kawalrencanamu-backend/mocks"
	"github.com/stretchr/testify/assert"
)

func TestFindById(t *testing.T) {
	mockBodRepo := new(mocks.BodRepository)
	mockBod := bod.Bod{
		BodId:       1,
		BodCallSign: "Call Sign 1",
		BodNamaId:   "Nama 1",
		BodNamaEn:   "Name 1",
	}
	mockListBod := make([]bod.Bod, 0)
	mockListBod = append(mockListBod, mockBod)

	t.Run("Success", func(t *testing.T) {

		// mockBodRepo.On("SearchAll", mock.Anything, mock.AnythingOfType("string"),
		// 	mock.AnythingOfType("int")).Return(mockListBod).Once()

		// repo.On("Fetch", ctx).Return(actual, errors.New("err"))
		mockBodRepo.On("SearchAll").Return(mockListBod).Once()
		srv := bod.NewBodService(mockBodRepo)
		_, err := srv.FindAll()
		// assert.Error(t, err)

		// assert.NoError(t, err)
		// assert.Len(t, list, len(mockListBod))
		// mockBodRepo.AssertExpectations(t)

		// assertion
		assert.Error(t, err)

		// end mock
		mockBodRepo.AssertExpectations(t)
	})
}
