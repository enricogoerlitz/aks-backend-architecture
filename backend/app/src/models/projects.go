package models

import "time"

var PROJECTS_MODEL_NAME = "projects"

type ProjectModel struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"not null;unique"`
	Description string
	HourlyRate  float64       `gorm:"not null"`
	CustomerID  int           `gorm:"not null"`
	Customer    CustomerModel `gorm:"foreignKey:CustomerID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (ProjectModel) TableName() string {
	return PROJECTS_MODEL_NAME
}

type ProjectResponse struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"not null;unique"`
	Description string
	HourlyRate  float64 `gorm:"not null"`
	CustomerID  int     `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (ProjectResponse) TableName() string {
	return PROJECTS_MODEL_NAME
}

type ProjectPOST struct {
	Name        string `gorm:"not null;unique"`
	Description string
	HourlyRate  float64 `gorm:"not null"`
	CustomerID  int     `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (ProjectPOST) TableName() string {
	return PROJECTS_MODEL_NAME
}

type ProjectPATCH struct {
	Name        string `gorm:"not null;unique"`
	Description string
	HourlyRate  float64 `gorm:"not null"`
	CustomerID  int     `gorm:"not null"`
	UpdatedAt   time.Time
}

func (ProjectPATCH) TableName() string {
	return PROJECTS_MODEL_NAME
}
