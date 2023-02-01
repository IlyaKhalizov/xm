package dto

import "github.com/go-playground/validator/v10"

var Validator = validator.New()

type Company struct {
	Name            string `json:"name" validate:"required,max=15"`
	Description     string `json:"description" validate:"max=3000"`
	EmployeesAmount int    `json:"employees_amount" validate:"required"`
	Registered      bool   `json:"registered" validate:"required"`
	Type            string `json:"type" validate:"required"`
}
