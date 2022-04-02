package models

import (
	util2 "kang/pkg/util"
)

type Model struct {
	Id        uint64           `gorm:"primaryKey;autoIncrement;column:id;type:bigint unsigned;NOT NULL;" json:"id"`
	CreatedAt util2.FormatTime `json:"created_at"`
	UpdatedAt util2.FormatTime `json:"updated_at"`
}
