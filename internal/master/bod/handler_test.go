package bod_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/gofiber/fiber/v2"
	"github.com/jadahbakar/kawalrencanamu-backend/internal/master/bod"
	"github.com/jadahbakar/kawalrencanamu-backend/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// https://github.com/gofiber/recipes/blob/master/unit-test/main_test.go

// func TestGetAll(t *testing.T) {
// 	// app := fiber.New()
// 	// var mockBod bod.Bod
// 	// mockService := new(mocks.BodService)
// 	// mockListBod := make([]bod.Bod, 0)
// 	// mockListBod = append(mockListBod, mockBod)
// 	// mockService.On("Fetch", mock.Anything).Return(mockListBod)

// 	// req, _ := http.NewRequest(
// 	// 	"GET",
// 	// 	"http://localhost:8000/api/v1/bod/all",
// 	// 	nil,
// 	// )
// 	// res, err := app.Test(req, -1)
// 	// assert.Equalf(t, false, err != nil, "bod all route")
// 	// assert.Equalf(t, 200, res.StatusCode, "bod all route")
// 	// body, err := ioutil.ReadAll(res.Body)
// 	// assert.Nilf(t, err, "bod all route")
// 	// assert.Equalf(t, "OK", string(body), "bod all route")

// 	t.Run("Success", func(t *testing.T) {
// 		// setup
// 		testy := tester.New()
// 		c, rec := testy.RequestWithContext(http.MethodGet, "/todos", nil, nil)
// 		hen, ret := testy.SetupHandlerTest()
// 		ctx := c.Request().Context()

// 		// mocks
// 		ret.TodoUsecase.On("FetchTodos", ctx).Return([]*domain.Todo{}, nil)
// 		h := hen.FetchTodos(c)

// 		// Assertions
// 		assert.NoError(t, h)
// 		assert.Equal(t, http.StatusOK, rec.Code)
// 		assert.Equal(t, "[]\n", rec.Body.String())

// 		// end mock
// 		ret.TodoUsecase.AssertExpectations(t)
// 	})

// }

func TestGetAll(t *testing.T) {
	var mockBod bod.Bod
	err := faker.FakeData(&mockBod)
	assert.NoError(t, err)
	mockService := new(mocks.BodService)
	mockListBod := make([]bod.Bod, 0)
	mockListBod = append(mockListBod, mockBod)

	mockService.On("Getall", mock.Anything).Return(mockListBod, nil)

	f := fiber.New()
	req, err := http.NewRequest(fiber.MethodGet, "/bod/all", strings.NewReader(""))
	assert.NoError(t, err)
	c := f.Get()
	handler := bod.BodHandler{bodService: mockService}
	err = handler.GetAll(c)
	require.NoError(t, err)

	assert.Equal(t, "10", responseCursor)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}
