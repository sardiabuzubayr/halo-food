package models

import (
	model "halo_food/models"
	rolemodel "halo_food/modules/master_role/models"
)

type LevelRole struct {
	IdLevelRole uint                 `json:"id_level_role"`
	Level       MasterLevel          `json:"id_level"`
	Role        rolemodel.MasterRole `json:"id_role"`
	IsActive    bool                 `json:"is_active"`
}

func (LevelRole) TableName() string {
	return model.TBLevelRole
}
