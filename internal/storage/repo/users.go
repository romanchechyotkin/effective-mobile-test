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

func (repo *Users) Delete(ctx context.Context, id string) error {
	q := `delete from users where users.id = $1`

	res, err := repo.qb.Pool().Exec(ctx, q, id)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (repo *Users) GetUser(ctx context.Context, id string) (*dbo.User, error) {
	q := `select id, last_name, first_name, second_name, age, gender, nationality from users where id = $1`

	var user dbo.User
	err := repo.qb.Pool().QueryRow(ctx, q, id).Scan(&user.ID, &user.LastName, &user.FirstName, &user.SecondName, &user.Age, &user.Gender, &user.Nationality)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
