package models

import (
	"time"
)

var CUSTOMERS_MODEL_NAME = "customers"

type CustomerModel struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (CustomerModel) TableName() string {
	return CUSTOMERS_MODEL_NAME
}

type CustomerResponse struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (CustomerResponse) TableName() string {
	return CUSTOMERS_MODEL_NAME
}

type CustomerPOST struct {
	Name      string `gorm:"not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (CustomerPOST) TableName() string {
	return CUSTOMERS_MODEL_NAME
}

type CustomerPATCH struct {
	Name      string `gorm:"not null;unique"`
	UpdatedAt time.Time
}

func (CustomerPATCH) TableName() string {
	return CUSTOMERS_MODEL_NAME
}
