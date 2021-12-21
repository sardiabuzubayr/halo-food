package config

import "testing"

func TestKoneksi(t *testing.T) {
	ConnectDB()
	if DB == nil {
		t.Errorf("Gagal Koneksi")
	}
}
