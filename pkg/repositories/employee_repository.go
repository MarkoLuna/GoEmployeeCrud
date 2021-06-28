package repositories

import (
	"github.com/MarkoLuna/GoEmployeeCrud/pkg/models"
)

type EmployeeRepository interface {
	Create(e models.Employee) (*models.Employee, error)

	FindAll() ([]models.Employee, error)

	FindById(ID int64) (models.Employee, error)

	DeleteById(ID int64) (int64, error)

	Update(e models.Employee) (int64, error)
}
