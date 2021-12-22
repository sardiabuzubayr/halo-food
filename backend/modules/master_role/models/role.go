package models

import model "halo_food/models"

type MasterRole struct {
	IdRole   uint   `json:"id_role"`
	NamaRole string `json:"nama_role"`
	IsActive bool   `json:"is_active"`
}

func (MasterRole) TableName() string {
	return model.TBRole
}
