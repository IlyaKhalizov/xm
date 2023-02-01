package entities

type Company struct {
	Id              int    `db:"id"`
	Name            string `db:"name"`
	Description     string `db:"description"`
	EmployeesAmount int    `db:"employees_amount"`
	Registered      bool   `db:"registered"`
	Type            string `db:"type"`
}
