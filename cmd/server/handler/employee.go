package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cpereira42/mercado-fresco-pron4/internal/employee"
	"github.com/cpereira42/mercado-fresco-pron4/pkg/web"
	"github.com/gin-gonic/gin"
)

func checkID(ctx *gin.Context) (int, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, "invalid ID"))
		return id, err
	}
	return id, nil
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
		employees, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, fmt.Sprintf("%v", err)))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, employees, ""))
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
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, fmt.Sprintf("%v", err)))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, employee, ""))
	}
}

func (c *EmployeeController) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var request employee.RequestEmployeeCreate
		if web.CheckIfErrorRequest(ctx, &request) {
			return
		}

		employee, err := c.service.Create(request.CardNumberID, request.FirstName, request.LastName, request.WarehouseID)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, fmt.Sprintf("%v", err)))
			return
		}
		ctx.JSON(http.StatusCreated, web.NewResponse(http.StatusCreated, employee, ""))

	}
}

func (c *EmployeeController) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := checkID(ctx)
		if err != nil {
			return
		}

		var request employee.RequestEmployeeUpdate

		if web.CheckIfErrorRequest(ctx, &request) {
			return
		}

		employee, err := c.service.Update(id, request.CardNumberID, request.FirstName, request.LastName, request.WarehouseID)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, fmt.Sprintf("%v", err)))
			return
		}
		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, employee, ""))
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
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusNotFound, nil, err.Error()))
			return
		}
		ctx.JSON(http.StatusNoContent, web.NewResponse(http.StatusNoContent, nil, ""))
	}
}
