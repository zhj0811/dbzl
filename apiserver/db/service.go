package db

import (
	"time"

	"github.com/zhj0811/dbzl/common/define"
)

//服务表
type TService struct {
	ID uint64 `json:"id" gorm:"column:id;primary_key;autoIncrement"`
	define.Service
	TxId        string    `json:"tx_id" gorm:"column:tx_id"`
	BlockHeight string    `json:"block_height"`
	UploadAt    time.Time `json:"upload_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
