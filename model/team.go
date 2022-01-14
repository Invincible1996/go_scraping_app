package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

//
type Team struct {
	CNName   string    `json:"cn_name"`
	TeamIcon string    `json:"team_icon"`
	ENName   string    `json:"en_name"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func (t *Team) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreateAt", time.Now())
	if err != nil {
		return err
	}
	err2 := scope.SetColumn("UpdateAt", time.Now())
	if err2 != nil {
		return err2
	}
	return nil
}

func (t *Team) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedAt", time.Now())
	if err != nil {
		return err
	}
	err2 := scope.SetColumn("UpdatedAt", time.Now())
	if err2 != nil {
		return err2
	}
	return nil
}
