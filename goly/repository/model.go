package repository

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Goly struct {
	ID       uint64 `json:"id" gorm:"column:id;primaryKey"`
	Redirect string `json:"redirect" gorm:"column:redirect;not null"`
	Goly     string `json:"goly" gorm:"column:goly;unique;not null"`
	Clicked  bool   `json:"clicked" gorm:"column:clicked"`
	Random   bool   `json:"random" gorm:"column:random"`
}

func Setup() {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Goly{})
	if err != nil {
		fmt.Println(err)
	}
}
