package config

import "testing"

func TestKoneksi(t *testing.T) {
	db := ConnectDB()
	if db == nil {
		t.Errorf("Gagal Koneksi")
	}
}
