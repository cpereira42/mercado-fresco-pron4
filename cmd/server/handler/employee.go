package handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/employee"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
)

func checkID(ctx *gin.Context) (int, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID",
		})
		return id, err
	}
	return id, nil
}

type ErrorMsg struct {
	Message string `json:"message"`
}

func errorHandler(fe validator.FieldError) string {
	return fmt.Sprintf("field %s is required", fe.Field())
}

func CheckIfErrorRequestExists(ctx *gin.Context, err error) bool {

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{errorHandler(fe)}
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": out,
		})
		return true
	}
	return false

}

func CheckIfErrorBindJsonExists(ctx *gin.Context, req *request) bool {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		if CheckIfErrorRequestExists(ctx, err) {
			return true
		}

		if err == io.EOF {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "body could not be empty",
			})
			return true
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return true
	}
	return false
}

type EmployeeController struct {
	service employee.Service
}

type request struct {
	CardNumberID string `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name" binding:"required"`
	LastName     string `json:"last_name" binding:"required"`
	WarehouseID  int    `json:"warehouse_id" binding:"required"`
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
			return
		}
		ctx.JSON(http.StatusOK, employee)
	}
}

func (c *EmployeeController) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := checkID(ctx)
		if err != nil {
			return
		}

		employee, err := c.service.GetByID(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": employee,
		})
	}
}

func (c *EmployeeController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request request
		if CheckIfErrorBindJsonExists(ctx, &request) {
			return
		}

		employee, err := c.service.Create(request.CardNumberID, request.FirstName, request.LastName, request.WarehouseID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, employee)

	}
}

func (c *EmployeeController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := checkID(ctx)
		if err != nil {
			return
		}

		var request request
		if CheckIfErrorBindJsonExists(ctx, &request) {
			return
		}

		fmt.Println("id param", id)

		employee, err := c.service.Update(id, request.CardNumberID, request.FirstName, request.LastName, request.WarehouseID)

		fmt.Println("id employee", employee.ID)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, employee)
	}
}

func (c *EmployeeController) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := checkID(ctx)
		if err != nil {
			return
		}

		err = c.service.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": fmt.Sprintf("The product with id %d was deleted", id),
		})
	}
}
