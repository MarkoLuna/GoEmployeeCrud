package controllers

import (
	"log"
	"net/http"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/dto"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/models"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/services"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/utils"
	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

var NewEmployee models.Employee

type EmployeeController struct {
	employeeService services.EmployeeService
}

func NewEmployeeController(employeeService services.EmployeeService) EmployeeController {
	return EmployeeController{employeeService}
}

func (eCtrl EmployeeController) CreateEmployee(c echo.Context) error {
	employee := dto.EmployeeRequest{}
	if err := c.Bind(&employee); err != nil {
		return err
	}

	v := utils.CreateValidator()
	err := v.Struct(employee)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			log.Println(e)
		}
		return c.String(http.StatusBadRequest, "")
	}

	e, err := eCtrl.employeeService.CreateEmployee(employee)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, e)
}

func (eCtrl EmployeeController) GetEmployees(c echo.Context) error {
	newEmployees, err := eCtrl.employeeService.GetEmployees()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, newEmployees)
}

func (eCtrl EmployeeController) GetEmployeeById(c echo.Context) error {
	employeeId := c.Param("employeeId")
	EmployeeDetails, err := eCtrl.employeeService.GetEmployeeById(employeeId)
	if err == nil {
		return c.JSON(http.StatusOK, EmployeeDetails)
	} else {
		return c.String(http.StatusNotFound, err.Error())
	}
}

func (eCtrl EmployeeController) UpdateEmployee(c echo.Context) error {
	var updateEmployee = dto.EmployeeRequest{}
	if err := c.Bind(&updateEmployee); err != nil {
		return err
	}

	log.Println("employee: " + updateEmployee.ToString())

	v := utils.CreateValidator()
	err := v.Struct(updateEmployee)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			log.Println(e)
		}
		return c.String(http.StatusNotFound, err.Error())
	}

	employeeId := c.Param("employeeId")
	employeeDetails, err := eCtrl.employeeService.UpdateEmployee(employeeId, updateEmployee)
	if err == nil {
		return c.JSON(http.StatusOK, employeeDetails)
	} else {
		return c.String(http.StatusNotFound, err.Error())
	}
}

func (eCtrl EmployeeController) DeleteEmployee(c echo.Context) error {
	employeeId := c.Param("employeeId")

	err := eCtrl.employeeService.DeleteEmployeeById(employeeId)
	if err == nil {
		return c.String(http.StatusOK, "")
	} else {
		return c.String(http.StatusNotFound, err.Error())
	}

}
