package constants

// Employee Status
//
// This model represents the employee status for this application.
//
// swagger:enum
type EmployeeStatus string

const (
	ACTIVE   EmployeeStatus = "ACTIVE"
	INACTIVE EmployeeStatus = "INACTIVE"
)
