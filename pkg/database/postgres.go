package postgres

import (
	"context"
	"fmt"
	"testtask/config"
	"github.com/golang-migrate/migrate/v4"
    _ "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
)




type Postgres struct {
	conn *pgxpool.Pool
	url string
}

func NewPostgres(config config.Config) *Postgres{
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	return &Postgres{
		url: url,
	}
}


func (p *Postgres) Connection() (conn *pgxpool.Pool, err error){
	conn, err = pgxpool.New(context.Background(), p.url)
	if err != nil{
		return nil, err
	}
	p.conn = conn
	return p.conn, nil
}

func (p *Postgres) MigrationUp() error{
	m, err := migrate.New("file://pkg/database/migrations", p.url)
	if err != nil{
		return err
	}
	if err = m.Up(); err != nil{
		if err.Error() == "no change"{
			return nil
		}
		return err
	}
	return nil
}