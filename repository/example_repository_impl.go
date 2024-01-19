package repository

import (
	"context"
	"database/sql"
	"projectsistemkasir/helper"
	"projectsistemkasir/model/domain"
)

type ExampleRepositoryImpl struct {
}

func NewExampleRepository() ExampleRepository {
	return &ExampleRepositoryImpl{}
}

func (repository *ExampleRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, example domain.Examples) domain.Examples {
	SQL := "insert into examples(name,jurusan) values (?,?)"
	result, err := tx.ExecContext(ctx, SQL, example.Name, example.Jurusan)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	example.Id = int(id)
	return example
}