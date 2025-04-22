package models

import "time"

var PROJECTTIMES_MODEL_NAME = "projecttimes"

type ProjecttimeModel struct {
	ID            int          `gorm:"primaryKey"`
	UserID        int          `gorm:"not null"`
	User          UserModel    `gorm:"foreignKey:UserID"`
	ProjectID     int          `gorm:"not null"`
	Project       ProjectModel `gorm:"foreignKey:ProjectID"`
	Hours         float64      `gorm:"not null"`
	BillableHours float64      `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (ProjecttimeModel) TableName() string {
	return PROJECTTIMES_MODEL_NAME
}

type ProjecttimeResponse struct {
	ID            int     `gorm:"primaryKey"`
	UserID        int     `gorm:"not null"`
	ProjectID     int     `gorm:"not null"`
	Hours         float64 `gorm:"not null"`
	BillableHours float64 `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (ProjecttimeResponse) TableName() string {
	return PROJECTTIMES_MODEL_NAME
}

type ProjecttimeResponseDetail struct {
	ID            int          `gorm:"primaryKey"`
	UserID        int          `gorm:"not null"`
	User          UserModel    `gorm:"foreignKey:UserID"`
	ProjectID     int          `gorm:"not null"`
	Project       ProjectModel `gorm:"foreignKey:ProjectID"`
	Hours         float64      `gorm:"not null"`
	BillableHours float64      `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (ProjecttimeResponseDetail) TableName() string {
	return PROJECTTIMES_MODEL_NAME
}

type ProjecttimePOST struct {
	ID            int     `gorm:"primaryKey"`
	UserID        int     `gorm:"not null"`
	ProjectID     int     `gorm:"not null"`
	Hours         float64 `gorm:"not null"`
	BillableHours float64 `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (ProjecttimePOST) TableName() string {
	return PROJECTTIMES_MODEL_NAME
}

type ProjecttimePATCH struct {
	ID            int     `gorm:"primaryKey"`
	UserID        int     `gorm:"not null"`
	ProjectID     int     `gorm:"not null"`
	Hours         float64 `gorm:"not null"`
	BillableHours float64 `gorm:"not null"`
	UpdatedAt     time.Time
}

func (ProjecttimePATCH) TableName() string {
	return PROJECTTIMES_MODEL_NAME
}
