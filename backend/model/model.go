package model

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var DefaultSchema string = "halo_food."
var TBLevel = DefaultSchema + "master_level"
var TBRole = DefaultSchema + "master_role"
var TBLevelRole = DefaultSchema + "level_role"
var TBUsers = DefaultSchema + "users"
var TBDriver = DefaultSchema + "driver"
var TBResto = DefaultSchema + "resto"

type Location struct {
	X, Y float64
}

func (loc Location) GormDataType() string {
	return "geometry"
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "POINT(?,?)",
		Vars: []interface{}{loc.X, loc.Y},
	}
}

type Response struct {
	ErrorCode    uint8       `json:"error_code"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
	Errors       interface{} `json:"errors,omitempty"`
	Token        string      `json:"token,omitempty"`
	RefreshToken string      `json:"refresh_token,omitempty"`
}

type ErrorValidator struct {
	FieldName    string `json:"field_name"`
	ErrorMessage string `json:"error_message"`
}
