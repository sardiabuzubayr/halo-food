package driver

import (
	"halo_food/model"
	usermodel "halo_food/modules/users/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Driver struct {
	IDUser         uuid.UUID       `json:"id_user" form:"id_user" gorm:"primaryKey"`
	Users          usermodel.Users `gorm:"foreignKey:IDUser"`
	NoId           string          `json:"no_id" form:"no_id"`
	NoPlat         string          `json:"no_plat" form:"no_plat"`
	NamaKendaraan  string          `json:"nama_kendaraan" form:"nama_kendaraan"`
	WarnaKendaraan string          `json:"warna_kendaraan" form:"warna_kendaraan"`
	Status         string          `json:"status" form:"status" gorm:"default:NULL"`
	StatusAkun     string          `json:"status_akun" form:"status_akun" gorm:"default:NULL"`
	Ktp            string          `json:"ktp" form:"ktp"`
}

func (Driver) TableName() string {
	return model.TBDriver
}

func (driver *Driver) RegisterNewDriver(db *gorm.DB) error {
	result := db.Create(driver)
	return result.Error
}
