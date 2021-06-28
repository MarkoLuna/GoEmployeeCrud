package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/MarkoLuna/GoEmployeeCrud/pkg/models"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/repositories"
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/utils"
	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
)

var NewEmployee models.Employee

type EmployeeController struct {
	employeeRepository repositories.EmployeeRepository
}

func NewEmployeeController(employeeRepository repositories.EmployeeRepository) EmployeeController {
	return EmployeeController{employeeRepository}
}

func (eCtrl EmployeeController) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	employee := &models.Employee{}
	utils.ParseBody(r.Body, employee)
	log.Println("employee: " + employee.ToString())

	v := utils.CreateValidator()
	err := v.Struct(employee)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			log.Println(e)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, err := eCtrl.employeeRepository.Create(*employee)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (eCtrl EmployeeController) GetEmployees(w http.ResponseWriter, r *http.Request) {
	newEmployees, err := eCtrl.employeeRepository.FindAll()
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
	ID, err := strconv.ParseInt(EmployeeId, 0, 0)
	if err != nil {
		log.Println("Error while parsing")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	EmployeeDetails, err := eCtrl.employeeRepository.FindById(ID)
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
	var updateEmployee = &models.Employee{}
	utils.ParseBody(r.Body, updateEmployee)

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
	EmployeeId := vars["employeeId"]
	ID, err := strconv.ParseInt(EmployeeId, 0, 0)
	if err != nil {
		log.Println("Error while parsing")
	}
	employeeDetails, err := eCtrl.employeeRepository.FindById(ID)
	if err == nil {
		if updateEmployee.FirstName != "" {
			employeeDetails.FirstName = updateEmployee.FirstName
		}
		if updateEmployee.LastName != "" {
			employeeDetails.LastName = updateEmployee.LastName
		}
		if updateEmployee.SecondLastName != "" {
			employeeDetails.SecondLastName = updateEmployee.SecondLastName
		}

		employeeDetails.DateOfBirth = updateEmployee.DateOfBirth
		employeeDetails.DateOfEmployment = updateEmployee.DateOfEmployment
		employeeDetails.Status = updateEmployee.Status

		log.Println("employee: " + employeeDetails.ToString())

		count, _ := eCtrl.employeeRepository.Update(employeeDetails)
		if count > 0 {
			res, _ := json.Marshal(employeeDetails)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(res)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func (eCtrl EmployeeController) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	EmployeeId := vars["employeeId"]
	ID, err := strconv.ParseInt(EmployeeId, 0, 0)
	if err != nil {
		log.Println("Error while parsing")
	}

	count, _ := eCtrl.employeeRepository.DeleteById(ID)
	if count > 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}