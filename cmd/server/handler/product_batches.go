package handler

import (
	"net/http"
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/productbatch"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)

type ProductBatChesController struct {
	servicePB productbatch.ServicePB
}

func NewProductBatChesController(servicePB productbatch.ServicePB) *ProductBatChesController {
	return &ProductBatChesController{servicePB: servicePB}
}

func (controller *ProductBatChesController) ReadPB() gin.HandlerFunc {
	return func(context *gin.Context) {
		paramId := context.Query("id")
		if paramId != "" {
			sectionId, err := strconv.ParseInt(paramId, 10, 64)
			if err != nil {
				context.JSON(http.StatusInternalServerError,
					web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
				return
			}
			resultProductBatchesResponse, errPB := controller.servicePB.GetId(sectionId)
			if errPB != nil {
				context.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, errPB.Error()))
				return
			}
			context.JSON(http.StatusOK,
				web.NewResponse(http.StatusOK, resultProductBatchesResponse, ""))
			return
		}
		productBatchesResponse, err := controller.servicePB.GetAll()
		if err != nil {
			context.JSON(http.StatusInternalServerError,
				web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
			return
		}
		context.JSON(http.StatusOK, web.NewResponse(http.StatusOK, productBatchesResponse, ""))
	}
}

func (controller *ProductBatChesController) CreatePB() gin.HandlerFunc {
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
