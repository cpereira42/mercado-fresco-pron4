package handler

import (
	"net/http"

	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	sectionEntites "github.com/cpereira42/mercado-fresco-pron4/internal/section"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)

type SectionController struct {
	service sectionEntites.Service
}

func NewSectionController(route *gin.Engine, serviceSection section.Service) {
	controllerSection := &SectionController{service: serviceSection}
	NewRouteSection(route, controllerSection)
}

func NewRouteSection(route *gin.Engine, sectionController *SectionController) {
	section := route.Group("/api/v1/sections")
	section.GET("/", sectionController.ListarSectionAll())
	section.GET("/:id", sectionController.ListarSectionOne())
	section.POST("/", sectionController.CreateSection())
	section.PATCH("/:id", sectionController.UpdateSection())
	section.DELETE("/:id", sectionController.DeleteSection())
}

func (controller *SectionController) ListarSectionAll() gin.HandlerFunc {
	return func(context *gin.Context) {
		sections, err := controller.service.ListarSectionAll()
		if err != nil {
			context.JSON(http.StatusInternalServerError,
				web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}
		context.JSON(http.StatusOK, web.NewResponse(http.StatusOK, sections, ""))
	}
}

func (controller *SectionController) CreateSection() gin.HandlerFunc {
	return func(context *gin.Context) {
		var newSection sectionEntites.SectionRequestCreate

		if web.CheckIfErrorRequest(context, &newSection) {
			return
		}

		section, err := controller.service.CreateSection(newSection)
		if err != nil {
			context.JSON(http.StatusConflict,
				web.NewResponse(http.StatusConflict, nil, err.Error()))
			return
		}
		context.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, section, ""))
	}
}

func (controller *SectionController) ListarSectionOne() gin.HandlerFunc {
	return func(context *gin.Context) {
		param, err := util.IDChecker(context)
		if err != nil {
			context.JSON(http.StatusInternalServerError,
				web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}
		section, err := controller.service.ListarSectionOne(int64(param))
		if err != nil {
			context.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		context.JSON(http.StatusOK, web.NewResponse(http.StatusOK, section, ""))
	}
}

func (controller *SectionController) UpdateSection() gin.HandlerFunc {
	return func(context *gin.Context) {
		paramId, errconv := util.IDChecker(context)
		if errconv != nil {
			context.JSON(
				http.StatusInternalServerError,
				web.NewResponse(http.StatusInternalServerError, nil, errconv.Error()))
			return
		}
		sectionUp := sectionEntites.SectionRequestUpdate{}
		if web.CheckIfErrorRequest(context, &sectionUp) {
			return
		}
		updateSection, err := controller.service.UpdateSection(int64(paramId), sectionUp)
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
		paramId, err := util.IDChecker(context)
		if err != nil {
			context.JSON(
				http.StatusInternalServerError,
				web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}
		erro := controller.service.DeleteSection(int64(paramId))
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
