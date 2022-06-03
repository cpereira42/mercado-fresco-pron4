package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/employee"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

func checkID(ctx *gin.Context) int {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID",
		})
	}
	return id
}

type ErrorMsg struct {
	Message string `json:"message"`
}

func errorHandler(fe validator.FieldError) string {
	return fmt.Sprintf("field %s is required", fe.Field())
}

func CheckIfErrorRequestExists(ctx *gin.Context, err error) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{errorHandler(fe)}
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": out,
		})
	}

}

func CheckIfErrorBindJsonExists(ctx *gin.Context, req any) {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		CheckIfErrorRequestExists(ctx, err)

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
}

type EmployeeController struct {
	service employee.Service
}

type request struct {
	CardNumberID string `json:"cardNumberID" binding:"required"`
	FirstName    string `json:"firstName" binding:"required"`
	LastName     string `json:"lastName" binding:"required"`
	WarehouseID  int    `json:"warehouseID" binding:"required"`
}

func NewEmployee(employee employee.Service) *EmployeeController {
	return &EmployeeController{
		service: employee,
	}
}

func (c *EmployeeController) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		employee, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, employee)
	}
}

func (c *EmployeeController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := checkID(ctx)

		var request request
		CheckIfErrorBindJsonExists(ctx, &request)

		employee, err := c.service.GetByID(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": employee,
		})
	}
}

func (c *EmployeeController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request request
		CheckIfErrorBindJsonExists(ctx, &request)

		employee, err := c.service.Create(request.CardNumberID, request.FirstName, request.LastName, request.WarehouseID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(200, employee)

	}
}

func (c *EmployeeController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := checkID(ctx)

		var request request
		CheckIfErrorBindJsonExists(ctx, &request)

		employee, err := c.service.Update(id, request.CardNumberID, request.FirstName, request.LastName, request.WarehouseID)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, employee)
	}
}

func (c *EmployeeController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := checkID(ctx)

		err := c.service.Delete(id)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": fmt.Sprintf("The product with id %d was deleted", id),
		})
	}
}
