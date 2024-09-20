package repository

import (
	"fmt"
	"order_service/internal/config"

	"github.com/DATA-DOG/go-sqlmock"
	gormMySQL "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func NewMockRepository() (*Repository, sqlmock.Sqlmock) {
	mockDb, mock, _ := sqlmock.New()
	mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("8.0.20"))

	dialector := gormMySQL.New(gormMySQL.Config{
		Conn:       mockDb,
		DriverName: "mysql",
	})
	db, _ := gorm.Open(dialector, &gorm.Config{})

	return &Repository{
		Db: db,
	}, mock
}

func NewRepository(config *config.ConfigService) *Repository {

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.DbUsername, config.DbPassword, config.DbHost, config.DbPort, config.DbName)

	db, err := gorm.Open(gormMySQL.New(gormMySQL.Config{
		DriverName: "mysql",
		DSN:        dbURI,
	}), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	println("Connected to database")

	return &Repository{
		Db: db,
	}
}
