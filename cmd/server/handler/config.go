package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/cpereira42/mercado-fresco-pron4/internal/productbatch"
	mocks "github.com/cpereira42/mercado-fresco-pron4/internal/section/mock"
	"github.com/gin-gonic/gin"
)

// * Inicia uma request
// @param method string
// @param url string
// @param body string
func CreateRequestServer(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "123456")
	return req, httptest.NewRecorder()
}

/*
 * cria um server
 * @param mock mocks.SectionService
 * @param method string
 * @param url string
 * @param body string
 */
func CreateServerSection(serv *mocks.SectionService, method, url, body string) *httptest.ResponseRecorder {
	repoPB := productbatch.NewRepositoryProductBatches(nil)
	servicePB := productbatch.NewServiceProductBatches(repoPB)
	sectionController := NewSectionController(serv, servicePB)
	router := gin.Default()
	gp := router.Group("/api/v1/sections")
	gp.GET("/", sectionController.ListarSectionAll())
	gp.GET("/:id", sectionController.ListarSectionOne())
	gp.POST("/", sectionController.CreateSection())
	gp.PATCH("/:id", sectionController.UpdateSection())
	gp.DELETE("/:id", sectionController.DeleteSection())
	req, rr := CreateRequestServer(method, url, body)
	router.ServeHTTP(rr, req)
	return rr
}

var ObjetoResponse struct {
	Code  int   `json:"code"`
	Data  any   `json:"data"`
	Error error `json:"error"`
}
