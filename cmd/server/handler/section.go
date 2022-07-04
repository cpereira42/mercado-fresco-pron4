package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/productbatch"
	sectionEntites "github.com/cpereira42/mercado-fresco-pron4/internal/section"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)

type SectionController struct {
	service   sectionEntites.Service
	servicePB productbatch.ServicePB
}

func NewSectionController(sectionService sectionEntites.Service, servPB productbatch.ServicePB) *SectionController {
	return &SectionController{service: sectionService, servicePB: servPB}
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
		param, err := getRequestId(context, "Param")
		if err != nil {
			context.JSON(http.StatusInternalServerError, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
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

func (controller *SectionController) UpdateSection() gin.HandlerFunc {
	return func(context *gin.Context) {
		paramId, errconv := getRequestId(context, "Param")
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
		fmt.Println("Controller = ", sectionUp)
		updateSection, err := controller.service.UpdateSection(paramId, sectionUp)
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
		paramId, err := getRequestId(context, "Param")
		if err != nil {
			context.JSON(
				http.StatusInternalServerError,
				web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}
		erro := controller.service.DeleteSection(paramId)
		if erro != nil {
			code := http.StatusNotFound
			if erro.Error() == "falha ao remove o sections" {
				code = http.StatusInternalServerError
			}
			context.JSON(
				code,
				web.NewResponse(code, nil, erro.Error()))
			return
		}
		context.JSON(
			http.StatusNoContent,
			web.NewResponse(http.StatusNoContent, paramId, ""))
	}
}

func (controller *SectionController) ReadPB() gin.HandlerFunc {
	return func(context *gin.Context) {
		paramId := context.Query("id")

		if paramId != "" {
			sectionId, err := getRequestId(context, "Query")
			if err != nil {
				context.JSON(http.StatusInternalServerError,
					web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
				return
			}
			resultProductBatchesResponse, errPB := controller.servicePB.ReadPBSectionId(sectionId)
			if errPB != nil {
				context.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, errPB.Error()))
				return
			}
			context.JSON(http.StatusOK,
				web.NewResponse(http.StatusOK, resultProductBatchesResponse, ""))
			return
		}
		productBatchesResponse, err := controller.servicePB.ReadPBSectionTodo()
		if err != nil {
			context.JSON(http.StatusInternalServerError,
				web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}
		context.JSON(http.StatusOK, web.NewResponse(http.StatusOK, productBatchesResponse, ""))
	}
}

func (controller *SectionController) CreatePB() gin.HandlerFunc {
	return func(context *gin.Context) {
		var productBatches productbatch.ProductBatches

		if web.CheckIfErrorRequest(context, &productBatches) {
			return
		}

		prodPb, err := controller.servicePB.CreatePB(productBatches)
		if err != nil {
			context.JSON(http.StatusConflict,
				web.NewResponse(http.StatusConflict, nil, err.Error()))
			return
		}

		context.JSON(
			http.StatusCreated,
			web.NewResponse(http.StatusCreated, prodPb, ""))
	}
}

func getRequestId(context *gin.Context, typeParam string) (paramId int64, err error) {
	switch typeParam {
	case "Param":
		id := context.Param("id")
		paramId, err = strconv.ParseInt(id, 10, 64)
		return
	case "Query":
		id := context.Query("id")
		paramId, err = strconv.ParseInt(id, 10, 64)
		return
	}
	return
}
