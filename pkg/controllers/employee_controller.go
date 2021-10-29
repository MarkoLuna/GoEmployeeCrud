package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/dto"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/models"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/services"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/utils"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

var NewEmployee models.Employee

type EmployeeController struct {
	employeeService services.EmployeeService
}

func NewEmployeeController(employeeService services.EmployeeService) EmployeeController {
	return EmployeeController{employeeService}
}

func (eCtrl EmployeeController) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	employee := dto.EmployeeRequest{}
	utils.ParseBody(r.Body, &employee)

	v := utils.CreateValidator()
	err := v.Struct(employee)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			log.Println(e)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	e, err := eCtrl.employeeService.CreateEmployee(employee)
	if err != nil {
		log.Fatalln(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(e)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (eCtrl EmployeeController) GetEmployees(w http.ResponseWriter, r *http.Request) {
	newEmployees, err := eCtrl.employeeService.GetEmployees()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(newEmployees)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (eCtrl EmployeeController) GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	EmployeeId := vars["employeeId"]
	EmployeeDetails, err := eCtrl.employeeService.GetEmployeeById(EmployeeId)
	if err == nil {
		res, _ := json.Marshal(EmployeeDetails)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (eCtrl EmployeeController) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	var updateEmployee = dto.EmployeeRequest{}
	utils.ParseBody(r.Body, &updateEmployee)

	log.Println("employee: " + updateEmployee.ToString())

	v := utils.CreateValidator()
	err := v.Struct(updateEmployee)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			log.Println(e)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	employeeId := vars["employeeId"]
	employeeDetails, err := eCtrl.employeeService.UpdateEmployee(employeeId, updateEmployee)
	if err == nil {
		res, _ := json.Marshal(employeeDetails)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (eCtrl EmployeeController) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employeeId := vars["employeeId"]

	err := eCtrl.employeeService.DeleteEmployeeById(employeeId)
	if err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}
