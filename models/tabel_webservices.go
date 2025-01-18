package models

type PesertaDidikAPI struct {
	RegistrasiID          string
	JenisPendaftaranID    string
	JenisPendaftaranIDStr string
	NIPD                  string
	TanggalMasukSekolah   string
	SekolahAsal           string
	PesertaDidikID        string
	Nama                  string
	Nisn                  string
	JenisKelamin          string
	Nik                   string
	TempatLahir           string
	TanggalLahir          string
	AgamaID               string
	AgamaIDStr            string
	AlamatJalan           string
	NomorTeleponRumah     string
	NomorTeleponSeluler   string
	NamaAyah              string
	PekerjaanAyahID       string
	PekerjaanAyahIDStr    string
	NamaIbu               string
	PekerjaanIbuID        string
	PekerjaanIbuIDStr     string
	NamaWali              string
	PekerjaanWaliID       string
	PekerjaanWaliIDStr    string
	AnakKeberapa          string
	Email                 string
	SemesterID            string
	AnggotaRombelID       string
	KebutuhanKhusus       string
}

type GTKAPI struct {
	TahunAjaranID           string
	PtkTerdaftarID          string
	PtkID                   string
	PtkInduk                int16
	TanggalSuratTugas       string
	Nama                    string
	JenisKelamin            string
	TempatLahir             string
	TanggalLahir            string
	AgamaID                 int16
	AgamaIDStr              string
	Nuptk                   string
	Nik                     string
	JenisPtkID              string
	JenisPtkIDStr           string
	StatusKepegawaianID     string
	StatusKepegawaianIDStr  string
	Nip                     string
	PendidikanTerakhir      string
	BidangStudiTerakhir     string
	PangkatGolonganTerakhir string
}

type SekolahAPI struct {
	SekolahID             string
	Nama                  string
	Nss                   string
	Npsn                  string
	BentukPendidikanID    string
	BentukPendidikanIDStr string
	StatusSekolah         int16
	StatusSekolahStr      string
	AlamatJalan           string
	RT                    string
	RW                    string
	KodeWilayah           string
	KodePos               string
	NomorTelepon          string
	NomorFax              string
	Email                 string
	Website               string
	IsSks                 bool
	Lintang               float32
	Bujur                 float32
	Dusun                 string
	DesaKelurahan         string
	Kecamatan             string
	KabupatenKota         string
	Provinsi              string
}

type RombelApi struct {
	RombonganBelajarID     string
	Nama                   string
	TingkatPendidikanID    int8
	TingkatPendidikanIDStr string
	SemesterID             string
	JenisRombel            int8
	JenisRombelStr         string
	KurikulumID            string
	KurikulumIDStr         string
	IDRuang                string
	IDRuangStr             string
	MovingClass            bool
	PtkID                  string
	PtkIDStr               string
	JurusanID              string
	JurusanIDStr           string
}

type AnggotaRombelAPI struct {
	AnggotaRombelID       string
	PesertaDidikID        string
	JenisPendaftaranID    string
	JenisPendaftaranIDStr string
}
