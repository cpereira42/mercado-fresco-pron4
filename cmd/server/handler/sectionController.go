package handler

import ( 
	"net/http"
	"strconv"
	"github.com/cpereira42/mercado-fresco-pron4/internal/section"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)


type SectionController struct {
	service section.Service
}
func NewSectionController(sectionService section.Service) *SectionController {
	return &SectionController{
		service: sectionService,
	}
}
func (controller SectionController)ListarSectionAll() gin.HandlerFunc {
	return func (context *gin.Context)  {
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
func (controller SectionController)CreateSection() gin.HandlerFunc {
	return func (context *gin.Context)  {
		var newSection section.Section
		if err := context.ShouldBindJSON(&newSection); err != nil {
			context.JSON(http.StatusNotFound, 
				web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		sections, err := controller.service.CreateSection(newSection)
		if err != nil {
			context.JSON(http.StatusUnprocessableEntity, 
				web.NewResponse(http.StatusUnprocessableEntity, nil, err.Error() ))
			return
		}
		context.JSON(http.StatusCreated, 
			web.NewResponse(http.StatusCreated, sections, ""))
	}
}
func (controller SectionController) ListarSectionOne() gin.HandlerFunc {
	return func(context *gin.Context) {
		paramId := context.Param("id")
		param, err := strconv.Atoi(paramId)
		if err != nil {
			context.JSON(400, web.NewResponse(400, nil, err.Error()))
			return 
		} 
		section, err := controller.service.ListarSectionOne(param)
		if err != nil {
			context.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		context.JSON(200, web.NewResponse(200, section, ""))
	}
}
func (controller SectionController) UpdateSection() gin.HandlerFunc{
	return func (context *gin.Context)  {
		id := context.Param("id")
		paramId, errconv := strconv.Atoi(id)
		if errconv != nil {
			context.JSON(200, web.NewResponse(400, nil, errconv.Error()))
			return 
		} 
		var sectionUp section.Section
		err := context.ShouldBindJSON(&sectionUp)
		if err != nil {
			context.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}		
		updateSection,err := controller.service.UpdateSection(paramId, sectionUp)
		if err != nil {
			context.JSON(400, web.NewResponse(400, nil, err.Error()))
			return 
		}
		context.JSON(200, web.NewResponse(200, updateSection, ""))
	}
}
func (controller SectionController) DeleteSection() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")
		paramId, err := strconv.Atoi(id)
		if err != nil {
			context.JSON(400, web.NewResponse(400, nil, err.Error()))
			return
		}
		erro := controller.service.DeleteSection(paramId)
		if erro != nil {
			context.JSON(404, web.NewResponse(404, nil, erro.Error()))
			return 
		}
		context.JSON(204, web.NewResponse(204, paramId, ""))
	}
}
func (controller SectionController) ModifyParcial() gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")		
		paramId, err := strconv.Atoi(id)
		if err != nil {
			context.JSON(http.StatusBadRequest, 
				web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return
		}		
		var mps section.ModifyParcial
		if err := context.ShouldBindJSON(&mps); err != nil {
			context.JSON(http.StatusBadRequest, 
				web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return 
		}
		sectionModify, err := controller.service.ModifyParcial(paramId, mps)
		if err != nil {
			context.JSON(http.StatusBadRequest, 
				web.NewResponse(http.StatusBadRequest, nil, err.Error()))
			return 
		}
		context.JSON(http.StatusOK, 
			web.NewResponse(http.StatusOK, sectionModify, ""))
	}
}
