package model

import (
	"time"

	"github.com/enzanumo/ky-theater-web/pkg/types"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// Model 公共Model
type Model struct {
	ID        int64 `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	IsDel     soft_delete.DeletedAt `gorm:"softDelete:flag" json:"is_del"`
}

type ConditionsT map[string]interface{}
type Predicates map[string]types.AnySlice

type Price int64

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	nowTime := time.Now().Unix()

	tx.Statement.SetColumn("created_on", nowTime)
	tx.Statement.SetColumn("modified_on", nowTime)
	return
}

func (m *Model) BeforeUpdate(tx *gorm.DB) (err error) {
	if !tx.Statement.Changed("modified_on") {
		tx.Statement.SetColumn("modified_on", time.Now().Unix())
	}

	return
}
