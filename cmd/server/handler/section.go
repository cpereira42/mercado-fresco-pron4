package handler

import (
	"net/http" 
	"strconv"

	sectionEntites "github.com/cpereira42/mercado-fresco-pron4/internal/section"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)


type SectionController struct {
	service sectionEntites.Service
}
func NewSectionController(sectionService sectionEntites.Service) *SectionController {
	return &SectionController{service: sectionService}
}
func (controller *SectionController) ListarSectionAll() gin.HandlerFunc {
	return func (context *gin.Context)  {
		sections, err := controller.service.ListarSectionAll()
		if err != nil {
			context.JSON(http.StatusBadRequest, 
				web.NewResponse(http.StatusBadRequest, nil, err.Error() ))
			return
		}
		context.JSON(http.StatusOK,	web.NewResponse(http.StatusOK, sections, ""))
	}
}
func (controller *SectionController) CreateSection() gin.HandlerFunc {
	return func (context *gin.Context)  { 
		
		var newSection sectionEntites.SectionRequestCreate
		
		if web.CheckIfErrorRequest(context,  &newSection) {
			return
		} 
		
		section, err := controller.service.CreateSection(newSection)
		if err != nil {
			context.JSON(http.StatusConflict, 
				web.NewResponse(http.StatusConflict, nil, err.Error() ))
			return
		}
		context.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, section, ""))
	}
}

func (controller *SectionController) ListarSectionOne() gin.HandlerFunc {
	return func(context *gin.Context) { 
		param, err := getRequestId(context)
		if err != nil {
			context.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return 
		} 
		section, err := controller.service.ListarSectionOne(param)
		if err != nil {
			context.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		context.JSON(http.StatusOK, web.NewResponse(http.StatusOK, section, ""))
	}
}
func (controller *SectionController) UpdateSection() gin.HandlerFunc{
	return func (context *gin.Context) {
		paramId, errconv := getRequestId(context)
		if errconv != nil {
			context.JSON(
				http.StatusNotFound, 
				web.NewResponse(http.StatusNotFound, nil, errconv.Error()))
			return
		}
		var sectionUp sectionEntites.SectionRequestUpdate
		if web.CheckIfErrorRequest(context, &sectionUp) {
			return
		}
		updateSection,err := controller.service.UpdateSection(paramId, sectionUp)
		if err != nil {
			context.JSON(
				http.StatusNotFound, 
				web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return 
		}
		context.JSON(
			http.StatusOK, 
			web.NewResponse(http.StatusOK, updateSection, ""))
	}
}
func (controller *SectionController) DeleteSection() gin.HandlerFunc {
	return func(context *gin.Context) {
		paramId, err := getRequestId(context)
		if err != nil {
			context.JSON(
				http.StatusNotFound, 
				web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		erro := controller.service.DeleteSection(paramId)
		if erro != nil {
			context.JSON(
				http.StatusNotFound, 
				web.NewResponse(http.StatusNotFound, nil, erro.Error()))
			return 
		}
		context.JSON(
			http.StatusNoContent, 
			web.NewResponse(http.StatusNoContent, paramId, ""))
	}
} 


func getRequestId(context *gin.Context) (paramId int64, err error) {
	id := context.Param("id")
	paramId, err = strconv.ParseInt(id, 16,64)
	return
}

 