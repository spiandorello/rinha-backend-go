package database

import (
	"fmt"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rinha-de-backend/src/structs"
)

type Postgres struct {
	DB *gorm.DB
}

func NewPostgres() *Postgres {
	dsn := getDSN()

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	//err = db.Use(
	//	dbresolver.Register(dbresolver.Config{
	//		//Replicas: []gorm.Dialector{
	//		//	postgres.Open(getDSNRead()),
	//		//},
	//	}).
	//		SetConnMaxIdleTime(time.Hour).
	//		SetConnMaxLifetime(24 * time.Hour).
	//		SetMaxIdleConns(10).
	//		SetMaxOpenConns(10),
	//)
	//if err != nil {
	//	panic(err)
	//}

	err = db.Use(otelgorm.NewPlugin())
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&structs.Person{}, &structs.Stack{})
	if err != nil {
		panic(err)
	}

	//sqlDB, err := db.DB()
	//sqlDB.SetMaxOpenConns(10000)
	//sqlDB.SetMaxIdleConns(1000)
	//sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return &Postgres{
		DB: db,
	}
}

func getDSN() string {
	//return fmt.Sprintf(
	//	"host=%s user=%s password=%s dbname=%s port=%v sslmode=disable",
	//	"localhost", "rinha", "rinha", "rinha", 5432,
	//)

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%v sslmode=disable",
		"rb-postgres", "rinha", "rinha", "rinha", 5432,
	)
}
