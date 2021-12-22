package models

import (
	model "halo_food/models"
	level "halo_food/modules/level/models"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Users struct {
	IdUser          uuid.UUID         `json:"id_user,omitempty" form:"id_user,omitempty" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Email           string            `json:"email" form:"email" validate:"required"`
	Alias           string            `json:"alias" form:"alias"`
	UserPassword    string            `json:"user_password" form:"user_password" gorm:"default:''"`
	Nama            string            `json:"nama" form:"nama"`
	Gender          string            `json:"gender" form:"gender" gorm:"default:NULL"`
	IsActive        bool              `json:"is_active" form:"is_active"`
	IsVerified      bool              `json:"is_verified" form:"is_verified"`
	VerifiedBy      uuid.UUID         `json:"verified_by" form:"verified_by" gorm:"type:uuid;default:NULL"`
	CreatedAt       time.Time         `json:"created_at,omitempty"`
	CreatedBy       uuid.UUID         `json:"created_by" gorm:"type:uuid;default:NULL"`
	IDLevel         int               `json:"id_level"`
	Level           level.MasterLevel `json:"level" form:"level" gorm:"foreignKey:IDLevel"`
	ConfirmKey      string            `json:"confirm_key"`
	ConfirmAt       time.Time         `json:"confirm_at,omitempty" gorm:"default:CURRENT_TIMESTAMP"`
	ConfirmKeyValid bool              `json:"confirm_key_valid"`
	Avatar          string            `json:"avatar" form:"avatar"`
}

func (Users) TableName() string {
	return model.TBUsers
}

func DoLoginUser(db *gorm.DB, username string, password string) bool {
	var user Users
	result := db.Where("email = ? OR alias = ?", username, username).Select("alias", "user_password").First(&user)

	if result.Error == nil {
		err := bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(password))
		return err == nil
	}

	return false
}

func GetOneByUsername(db *gorm.DB, username string) (Users, error) {
	var user Users
	result := db.Where("email = ? OR alias = ?", username, username).Joins("Level").First(&user)
	if result.Error == nil {
		return user, nil
	}
	return user, result.Error
}

func (user *Users) RegisterNewCustomer(db *gorm.DB) error {
	result := db.Create(user)
	return result.Error
}

// func GetAll(con *gorm.db, limit int, page int, keywords string) model.Response {
// 	var film []model.Film
// 	offset := (limit * page) - limit
// 	if offset < 0 {
// 		offset = 0
// 	}
// 	result := con.Where("title LIKE ?", "%"+keywords+"%").Limit(limit).Offset(offset).Table("film").Find(&film)

// 	var data model.Response
// 	if result.Error != nil {
// 		data.ErrorCode = 1
// 		data.Message = "Gagal"
// 	} else {
// 		data.ErrorCode = 0
// 		data.Message = "Berhasil"
// 		data.Data = film
// 	}

// 	return data
// }
