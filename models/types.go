package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int            `json:"id" gorm:"serial;primaryKey"`
	Firstname string         `json:"firstname" gorm:"not null"`
	Lastname  string         `json:"lastname" gorm:"not null"`
	Email     string         `json:"email" gorm:"not null;unique"`
	Password  string         `json:"-" gorm:"not null"`
	Photo     string         `json:"path,omitempty" gorm:"not null"`
	Role      string         `json:"role" gorm:"not null;default: user"`
	IsActive  bool           `json:"is_active" gorm:"not null;default:true"`
	CreatedAt time.Time      `json:"-" gorm:"autoCreate"`
	UpdatedAt time.Time      `json:"-" gorm:"autoUpdate"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type Project struct {
	ID          int       `json:"id" gorm:"serial;primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Department  string    `json:"department" gorm:"not null"`
	ManagerID   int       `json:"-" gorm:"manager_id"`
	ManagerName string    `json:"manager_name" gorm:"-"`
	Status      string    `json:"status" gorm:"not null;default:'Not started'"`
	StartDate   string    `json:"start_date,omitempty" gorm:"type:timestamp;not null;default: now()"`
	Deadline    string    `json:"deadline" gorm:"type:timestamp;not null"`
	IsActive    bool      `json:"-" gorm:"not null;default: true"`
	CreatedAt   time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"-" gorm:"autoUpdateTime"`
	DeletedAt   time.Time `json:"-" gorm:"index"`
	User        User      `json:"-" gorm:"foreignKey:ManagerID"`
}

type Projects []Project

type ProjectParticipant struct {
	ID            int     `json:"id" gorm:"serial;primaryKey"`
	ParticipantId int     `json:"participant_id" gorm:"participant_id"`
	Role          string  `json:"role" gorm:"not null;default:'participant'"`
	ProjectId     int     `json:"project_id" gorm:"project_id"`
	User          User    `json:"-" gorm:"foreignKey:ParticipantId"`
	Project       Project `json:"-" gorm:"foreignKey:ProjectId"`
}

type Task struct {
	ID           int       `json:"id" gorm:"serial;primaryKey"`
	Title        string    `json:"title" gorm:"not null"`
	Description  string    `json:"description" gorm:"not null"`
	ControllerId int       `json:"-" gorm:"controller_id"`
	ExecutorId   int       `json:"-" gorm:"executor_id"`
	ExecutorName string    `json:"executor_name" gorm:"-"`
	Status       string    `json:"status" gorm:"not null;default:'Not started'"`
	ProjectId    int       `json:"-" gorm:"project_id"`
	ProjectName  string    `json:"project_name" gorm:"-"`
	Deadline     string    `json:"deadline" gorm:"type:timestamp;not null"`
	IsActive     bool      `json:"-" gorm:"not null;default: true"`
	CreatedAt    time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"-" gorm:"autoUpdateTime"`
	DeletedAt    time.Time `json:"-" gorm:"index"`
	Controller   User      `json:"-" gorm:"foreignKey:ControllerId"`
	Executor     User      `json:"-" gorm:"foreignKey:ExecutorId"`
	Project      Project   `json:"-" gorm:"foreignKey:ProjectId"`
}

type Tasks []Task
