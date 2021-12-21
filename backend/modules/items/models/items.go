package models

// func GetOneByID(con *gorm.DB, KodeItem string) model.Response {
// 	var item model.Items
// 	result := con.Where("kode_item = ?", KodeItem).First(&item)

// 	var data model.Response
// 	if result.Error != nil {
// 		data.ErrorCode = 1
// 		data.Message = "Gagal, data kosong"
// 	} else {
// 		data.ErrorCode = 0
// 		data.Message = "Berhasil"
// 		data.Data = item
// 	}

// 	return data
// }

// func GetAll(con *gorm.DB, limit int, page int, keywords string) model.Response {
// 	var items []model.Items
// 	offset := (limit * page) - limit
// 	if offset < 0 {
// 		offset = 0
// 	}
// 	result := con.Where("kode_item LIKE ?", "%"+keywords+"%").Limit(limit).Offset(offset).Find(&items)

// 	var data model.Response
// 	if result.Error != nil {
// 		data.ErrorCode = 1
// 		data.Message = "Gagal"
// 	} else {
// 		data.ErrorCode = 0
// 		data.Message = "Berhasil"
// 		data.Data = items
// 	}

// 	return data
// }

// func SaveNew(con *gorm.DB, item model.Items) (bool, error) {
// 	result := con.Create(&item)
// 	return result.RowsAffected > 0, result.Error
// }
