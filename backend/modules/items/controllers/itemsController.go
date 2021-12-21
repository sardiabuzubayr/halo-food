package controllers

// func GetOneByID(c echo.Context) error {
// 	keywords := c.Param("kode_item")
// 	data := itemsModel.GetOneByID(config.DB, keywords)
// 	return c.JSON(http.StatusOK, data)
// }

// func GetAll(c echo.Context) error {
// 	limit, _ := strconv.Atoi(c.Param("limit"))
// 	page, _ := strconv.Atoi(c.Param("page"))
// 	keywords := c.QueryParam("keywords")
// 	data := itemsModel.GetAll(config.DB, limit, page, keywords)

// 	return c.JSON(http.StatusOK, data)
// }

// func AddNew(c echo.Context) {
// 	hargaModal, _ := strconv.ParseFloat(c.FormValue("harga_modal"), 32)
// 	hargaJual, _ := strconv.ParseFloat(c.FormValue("harga_jual"), 32)
// 	items := model.Items{
// 		KodeItem:   c.FormValue("kode_item"),
// 		NamaItem:   c.FormValue("nama_item"),
// 		HargaModal: float32(hargaModal),
// 		HargaJual:  float32(hargaJual),
// 	}

// 	result := model.Response{}
// 	_, err := itemsModel.SaveNew(config.DB, items)
// 	if err != nil {
// 		result.ErrorCode = 101
// 		result.Message = "Gagal insert " + string(err.Error())
// 	} else {
// 		result.ErrorCode = 0
// 		result.Message = "Berhasil insert"
// 	}

// 	c.JSON(http.StatusOK, result)
// }
