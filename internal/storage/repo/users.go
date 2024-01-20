package repo

import (
	"context"

	"github.com/romanchechyotkin/effective-mobile-test-task/internal/storage/dbo"
	"github.com/romanchechyotkin/effective-mobile-test-task/pkg/db"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type QueryBuilder interface {
	Pool() *pgxpool.Pool
}

type Users struct {
	qb QueryBuilder
}

func NewUsers(qb *db.QBuilder) *Users {
	return &Users{qb: qb}
}

func (repo *Users) Create(ctx context.Context, model *dbo.User) (string, error) {
	args := pgx.NamedArgs{
		"last_name":   model.LastName,
		"first_name":  model.FirstName,
		"second_name": model.SecondName,
		"age":         model.Age,
		"gender":      model.Gender,
		"nationality": model.Nationality,
	}

	q := `insert into users(last_name, first_name, second_name, age, gender, nationality)
values (@last_name, @first_name, @second_name, @age, @gender, @nationality)
returning id`

	err := repo.qb.Pool().QueryRow(ctx, q, &args).Scan(&model.ID)
	if err != nil {
		return "", err
	}

	return model.ID, nil
}
