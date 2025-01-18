package models

// import "gorm.io/gorm"

type APIDapodikWebservices struct {
	// gorm.Model
	SekolahID   string
	AplikasiID  string
	Nama        string
	Token       string
	Password    string
	IPAddress   string
	Port        string
	MacAddress  string
	AsalData    string
	Aktif       string
	ExpiredDate string
	// CreateDate  string
	// LastUpdate  string
	// UpdaterID   string
	// LastSync    string
	// SoftDelete  string
}
