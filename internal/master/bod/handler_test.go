package bod_test

import (
	"testing"

	"github.com/jadahbakar/kawalrencanamu-backend/internal/master/bod"
	"github.com/jadahbakar/kawalrencanamu-backend/mocks"

	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	var mockBod bod.Bod
	mockService := new(mocks.BodService)
	mockListBod := make([]bod.Bod, 0)
	mockListBod = append(mockListBod, mockBod)
	mockService.On("Fetch", mock.Anything).Return(mockListBod)

	// app := New()
	// app.Get("/bod/all", func(c *Ctx) {
	// 	c.SendStatus(400)
	// })

	// resp, err := app.Test(httptest.NewRequest("GET", "/bod/all", nil))
	// utils.AssertEqual(t, nil, err, "app.Test")
	// utils.AssertEqual(t, 400, resp.StatusCode, "Status code")

	// e := echo.New()
	// req, err := http.NewRequest(echo.GET, "/article?num=1&cursor="+cursor, strings.NewReader(""))
	// assert.NoError(t, err)

	// rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)
	// handler := articleHttp.ArticleHandler{
	// 	AUsecase: mockUCase,
	// }
	// err = handler.FetchArticle(c)
	// require.NoError(t, err)

	// responseCursor := rec.Header().Get("X-Cursor")
	// assert.Equal(t, "10", responseCursor)
	// assert.Equal(t, http.StatusOK, rec.Code)
	// mockUCase.AssertExpectations(t)
}
