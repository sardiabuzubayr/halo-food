package models

import (
	"halo_food/model"

	"gorm.io/gorm"
)

var db *gorm.DB

func (msl *MasterLevel) InitConnection(activeConnection *gorm.DB) {
	db = activeConnection
}

type MasterLevel struct {
	IdLevel   uint   `json:"id_level" gorm:"primaryKey"`
	NamaLevel string `json:"nama_level"`
	IsActive  bool   `json:"is_active"`
}

func (MasterLevel) TableName() string {
	return model.TBLevel
}
func GetLevelByID(IdLevel int) model.Response {
	var level MasterLevel
	result := db.Where("id_level = ?", IdLevel).First(&level)

	var data model.Response
	if result.Error != nil {
		data.ErrorCode = 1
		data.Message = "Gagal, data kosong"
	} else {
		data.ErrorCode = 0
		data.Message = "Berhasil"
		data.Data = level
	}

	return data
}

func GetAll(limit int, page int, keywords string) model.Response {
	var levels []MasterLevel
	offset := (limit * page) - limit
	if offset < 0 {
		offset = 0
	}
	result := db.Where("nama_level ILIKE ?", "%"+keywords+"%").Limit(limit).Offset(offset).Find(&levels)

	var data model.Response
	if result.Error != nil {
		data.ErrorCode = 1
		data.Message = "Gagal"
	} else {
		data.ErrorCode = 0
		data.Message = "Berhasil"
		data.Data = levels
	}

	return data
}
