package handler

import (
	"io"
	"net/http"
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/employee"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
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

func RequestValidator(req any) error {
	if err := validator.Validate(req); err != nil {
		return err
	}
	return nil
}

func CheckIfErrorRequest(ctx *gin.Context, req any) bool {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		if err == io.EOF {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "body must have all fields",
			})
			return true
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return true
	} else if err := RequestValidator(req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return true
	}
	return false
}

type EmployeeController struct {
	service employee.Service
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
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": employee,
		})
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

		var request employee.Request
		if CheckIfErrorRequest(ctx, &request) {
			return
		}

		employee, err := c.service.Create(request.CardNumberID, request.FirstName, request.LastName, request.WarehouseID)

		if err != nil {
			ctx.AbortWithStatus(http.StatusNotFound)
		}
		ctx.JSON(http.StatusCreated, gin.H{
			"data": employee,
		})

	}
}

func (c *EmployeeController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := checkID(ctx)
		if err != nil {
			return
		}

		var request employee.Employee

		if CheckIfErrorRequest(ctx, &request) {
			return
		}

		employee, err := c.service.Update(id, request.CardNumberID, request.FirstName, request.LastName, request.WarehouseID)

		if err != nil {
			ctx.JSON(http.StatusNotFound, err.Error())
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
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}
		ctx.AbortWithStatus(http.StatusNoContent)
	}
}
