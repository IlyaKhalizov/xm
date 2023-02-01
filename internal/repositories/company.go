package repositories

import (
	"context"
	"xm/internal/handlers/dto"
	"xm/internal/repositories/entities"

	"github.com/jmoiron/sqlx"
)

type companyRepository struct {
	db  *sqlx.DB
	ctx context.Context
}

func NewCompanyRepository(ctx context.Context, db *sqlx.DB) companyRepository {
	return companyRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r companyRepository) GetCompany(ctx context.Context, id int) (entities.Company, error) {
	company := entities.Company{}

	err := r.db.GetContext(r.ctx, &company, `SELECT id, name, description, employees_amount, registered, type FROM company WHERE id = $1`, id)

	if err != nil {
		return entities.Company{}, err
	}

	return company, nil
}

func (r companyRepository) CreateCompany(ctx context.Context, c dto.Company) (entities.Company, error) {
	company := entities.Company{
		Name:            c.Name,
		Description:     c.Description,
		EmployeesAmount: c.EmployeesAmount,
		Registered:      c.Registered,
		Type:            c.Type,
	}

	stmt, err := r.db.PrepareNamedContext(r.ctx, `INSERT INTO company (name, description, employees_amount, registered, type) 
		VALUES (:name, :description, :employees_amount, :registered, :type) RETURNING ID`)

	if err != nil {
		return entities.Company{}, err
	}

	var id int
	err = stmt.Get(&id, company)

	company.Id = id
	if err != nil {
		return entities.Company{}, err
	}

	return company, err
}

func (r companyRepository) DeleteCompany(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(r.ctx, `DELETE FROM company WHERE id = $1`, id)

	return err
}

func (r companyRepository) UpdateCompany(ctx context.Context, id int, c dto.Company) (entities.Company, error) {
	var err error

	_, err = r.db.ExecContext(r.ctx, `UPDATE company 
		SET name = $1, description = $2, employees_amount = $3, registered = $4, type = $5
		WHERE id = $6`, c.Name, c.Description, c.EmployeesAmount, c.Registered, c.Type, id)

	if err != nil {
		return entities.Company{}, err
	}

	company := entities.Company{}

	err = r.db.GetContext(r.ctx, &company, `SELECT id, name, description, employees_amount, registered, type FROM company WHERE id = $1`, id)
	if err != nil {
		return entities.Company{}, err
	}

	return company, nil
}
