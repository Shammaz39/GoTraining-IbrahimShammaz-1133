package models

import (
	"time"

	"gorm.io/datatypes"
)

type OrderEvent struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID   string         `gorm:"not null;index" json:"order_Id"`
	Event     string         `gorm:"type:varchar(50);not null" json:"even"`
	Timestamp time.Time      `gorm:"autoCreateTime" json:"timeStamp"`
	Meta      datatypes.JSON `gorm:"type:jsonb" json:"meta"`
}
