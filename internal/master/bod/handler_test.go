package bod_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/kawalrencanamu-backend/internal/master/bod"
	"github.com/jadahbakar/kawalrencanamu-backend/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// https://github.com/gofiber/recipes/blob/master/unit-test/main_test.go

func TestGetAll(t *testing.T) {
	var mockBod bod.Bod
	err := faker.FakeData(&mockBod)
	assert.NoError(t, err)
	mockService := new(mocks.BodService)
	mockListBod := make([]bod.Bod, 0)
	mockListBod = append(mockListBod, mockBod)
	mockService.On("Fetch").Return(mockListBod, nil)

	f := fiber.New()
	req, err := http.NewRequest(fiber.MethodGet, "/bod/all", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	// c := fiber.Ctx(req, rec)
	handler := bod.BodHandler{
		bodService: mockService,
	}
	err = handler.GetAll(ctx)
	require.NoError(t, err)

}

// func TestGetAll(t *testing.T) {
// 	f := fiber.New()
// 	req := httptest.NewRequest(fiber.MethodGet, "/bod/all", nil)
// 	rec := httptest.NewRecorder()

// 	mockBodService := new(mocks.BodService)
// 	h := bod.NewBodHandler(f, mockBodService)
// }
