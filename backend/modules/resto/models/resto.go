package models

import (
	model "halo_food/models"
	"halo_food/modules/users/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Resto struct {
	IdResto    uuid.UUID      `json:"id_resto" form:"id_resto" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	IdUser     uuid.UUID      `json:"id_user" form:"id_user"`
	UserOwn    models.Users   `gorm:"foreignKey:IdUser"`
	NoId       string         `json:"no_id" form:"no_id"`
	NamaUsaha  string         `json:"nama_usaha" form:"nama_usaha"`
	Alamat     string         `json:"alamat" form:"alamat"`
	Lokasi     model.Location `json:"lokasi"`
	Logo       string         `json:"logo" form:"logo"`
	IsActive   bool           `json:"is_active" form:"is_active"`
	IsCabang   bool           `json:"is_cabang" form:"is_cabang"`
	IsOpen     bool           `json:"is_open" form:"is_open"`
	StatusAkun string         `json:"status_akun" form:"status_akun" gorm:"default:'S'"`
	Ktp        string         `json:"ktp" form:"ktp"`
}

func (Resto) TableName() string {
	return model.TBResto
}

func (resto *Resto) RegisterNewResto(db *gorm.DB) error {
	result := db.Create(resto)
	return result.Error
}
