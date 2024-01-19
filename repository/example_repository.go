package repository

import (
	"context"
	"database/sql"
	"projectsistemkasir/model/domain"
)

type ExampleRepository interface {
	Save(ctx context.Context, tx *sql.Tx, example domain.Examples) domain.Examples
}