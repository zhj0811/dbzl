package db

import (
	"time"

	"github.com/zhj0811/dbzl/common/define"
)

//保单表
type TPolicy struct {
	define.Policy
	TxID      string    `json:"tx_id" gorm:"column:tx_id; index"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//func InsertService() error {
//
//}
