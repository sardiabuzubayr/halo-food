package models

import "halo_food/model"

type MasterRole struct {
	IdRole   uint   `json:"id_role"`
	NamaRole string `json:"nama_role"`
	IsActive bool   `json:"is_active"`
}

func (MasterRole) TableName() string {
	return model.TBRole
}
