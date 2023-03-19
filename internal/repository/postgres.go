package repository

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose"
	"github.com/qustavo/dotsql"
	"github.com/sirupsen/logrus"
	_ "smartway-test-task/internal/migrations"
)

const (
	createUser                = "create-user"
	createDepartment          = "create-department"
	createPassport            = "create-passport"
	findEmployeesByCompanyId  = "find-employees-by-company-id"
	findEmployeesByDepartment = "find-employees-by-department"
	isEmployeePresent         = "is-employee-present"
	updateEmployee            = "update-employee"
	updateDepartmentByUserId  = "update-department-by-user-id"
	updatePassportByUserId    = "update-passport-by-user-id"
	addDepartmentToUser       = "add-department-to-user"
	addPassportToUser         = "add-passport-to-user"
	deleteUser                = "delete-user"
)

type PostgresDb struct {
	db  *sql.DB
	dot *dotsql.DotSql
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
	Reload   bool
}

func NewPostgresDB(cfg Config) (*PostgresDb, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	if cfg.Reload {
		logrus.Printf("Start reloading database")
		err := goose.DownTo(db, "./internal/migrations", 0)
		if err != nil {
			return nil, err
		}
	}

	logrus.Printf("Start migrating database \n")
	err = goose.Up(db, "./internal/migrations")
	if err != nil {
		return nil, err
	}
	dot, err := dotsql.LoadFromFile("./sql/employee.sql")
	if err != nil {
		return nil, err
	}

	return &PostgresDb{
		db:  db,
		dot: dot,
	}, nil
}
