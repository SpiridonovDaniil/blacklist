package postgres

import (
	"context"
	"fmt"
	"log"

	"blacklist/internal/config"
	"blacklist/internal/domain"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Db struct {
	db *sqlx.DB
}

func New(cfg config.Postgres) *Db {
	conn, err := sqlx.Connect("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.User,
			cfg.Pass,
			cfg.Address,
			cfg.Port,
			cfg.Db,
		))
	if err != nil {
		log.Fatal(err)
	}

	return &Db{db: conn}
}

func (d *Db) Create(ctx context.Context, data *domain.Person) error {
	query := `INSERT INTO blacklist (phone, name, reason, time, uploader) VALUES ($1, $2, $3, $4, $5)`
	_, err := d.db.ExecContext(ctx, query, data.Phone, data.Name, data.Reason, data.Time, data.Uploader)
	if err != nil {
		err = fmt.Errorf("incert person failed, error: %w", err)
		return err
	}

	return nil
}

func (d *Db) Delete(ctx context.Context, id int) error {
	_, err := d.db.ExecContext(ctx, "DELETE FROM blacklist WHERE id = $1", id)
	if err != nil {
		err = fmt.Errorf("delete person failed, error: %w", err)
		return err
	}

	return nil
}

func (d *Db) Get(ctx context.Context, name, phone string) ([]*domain.Person, error) {
	var result []*domain.Person
	query := `SELECT * FROM blacklist WHERE name = $1 OR phone = $2`
	err := d.db.SelectContext(ctx, &result, query, name, phone)
	if err != nil {
		return nil, fmt.Errorf("get person failed, error: %w", err)
	}

	return result, nil
}
