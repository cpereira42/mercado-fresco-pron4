package handler

import (
	//"fmt"
	//"os"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)


type SectionController struct {
	service section.Service
}
func NewSectionController(sectionService section.Service) *SectionController {
	return &SectionController{service: sectionService}
}
func (controller *SectionController)ListarSectionAll() gin.HandlerFunc {
	return func (context *gin.Context)  {
		/*token := context.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			erro := fmt.Errorf("error: você não tem permissão para fazer essa solicitação")
			context.JSON(http.StatusUnauthorized, 
				web.NewResponse(http.StatusUnauthorized, nil, erro.Error()))
			return
		}*/

		sections, err := controller.service.ListarSectionAll()
		if err != nil {
			context.JSON(http.StatusBadRequest, 
				web.NewResponse(http.StatusBadRequest, nil, err.Error() ))
			return
		}
		context.JSON(http.StatusOK, 
			web.NewResponse(http.StatusOK, sections, ""))
	}
}
func (controller *SectionController) CreateSection() gin.HandlerFunc {
	return func (context *gin.Context)  {
		/*token := context.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			erro := fmt.Errorf("error: você não tem permissão para fazer essa solicitação")
			context.JSON(http.StatusUnauthorized, 
				web.NewResponse(http.StatusUnauthorized, nil, erro.Error()))
			return
		}*/
		
		var newSection section.SectionRequest
		if err := context.ShouldBindJSON(&newSection); err != nil {
			context.JSON(http.StatusNotFound, 
				web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
			return
		}
		
		fields := []string{"SectionNumber", "CurrentTemperature", "MinimumTemperature", "CurrentCapacity", 
			"MinimumCapacity", "MaximumCapacity", "WareHouseId", "ProductTypeId"}
		
		mapSection := structs.Map(newSection)
		for _, value := range fields {
			if mapSection[value] == 0 {
				err := fmt.Errorf("field %s is required", value)
				context.JSON(
					http.StatusUnprocessableEntity, 
					web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error()))
				return 
			}
		}

		sections, err := controller.service.CreateSection(newSection)
		if err != nil {
			context.JSON(http.StatusConflict, 
				web.NewResponse(http.StatusConflict, nil, err.Error() ))
			return
		}
		context.JSON(http.StatusCreated, 
			web.NewResponse(http.StatusCreated, sections, ""))
	}
}
func (controller *SectionController) ListarSectionOne() gin.HandlerFunc {
	return func(context *gin.Context) {
		/*token := context.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			erro := fmt.Errorf("error: você não tem permissão para fazer essa solicitação")
			context.JSON(http.StatusUnauthorized, 
				web.NewResponse(http.StatusUnauthorized, nil, erro.Error()))
			return
		}*/

		paramId := context.Param("id")
		param, err := strconv.Atoi(paramId)
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
	return func (context *gin.Context)  {
		/*token := context.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			erro := fmt.Errorf("error: você não tem permissão para fazer essa solicitação")
			context.JSON(http.StatusUnauthorized, 
				web.NewResponse(http.StatusUnauthorized, nil, erro.Error()))
			return
		}*/

		id := context.Param("id")
		paramId, errconv := strconv.Atoi(id)
		if errconv != nil {
			context.JSON(
				http.StatusNotFound, 
				web.NewResponse(http.StatusNotFound, nil, errconv.Error()))
			return 
		} 
		var sectionUp section.Section
		err := context.ShouldBindJSON(&sectionUp)
		if err != nil {
			context.JSON(
				http.StatusNotFound, 
				web.NewResponse(http.StatusNotFound, nil, err.Error()))
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
		/*token := context.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			erro := fmt.Errorf("error: você não tem permissão para fazer essa solicitação")
			context.JSON(http.StatusUnauthorized, 
				web.NewResponse(http.StatusUnauthorized, nil, erro.Error()))
			return
		}*/

		id := context.Param("id")
		paramId, err := strconv.Atoi(id)
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
