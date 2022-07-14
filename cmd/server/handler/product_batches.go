package handler

import (
	"net/http"

	"github.com/cpereira42/mercado-fresco-pron4/internal/productbatch"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/util"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)

type ProductBatChesController struct {
	servicePB productbatch.ServicePB
}

func NewProductBatChesController(route *gin.Engine, servicePB productbatch.ServicePB) {
	controllerPB := &ProductBatChesController{servicePB: servicePB}

	route.GET("/api/v1/sections/reportProducts", controllerPB.ReadPB())
	route.POST("/api/v1/productBatches", controllerPB.CreatePB())
}

func (controller *ProductBatChesController) ReadPB() gin.HandlerFunc {
	return func(context *gin.Context) {
		paramId, err := util.IDChecker(context)
		if err != nil {
			context.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "invalid ID"))
			return
		}
		if paramId != 0 {
			resultProductBatchesResponse, errPB := controller.servicePB.GetId(int64(paramId))
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
