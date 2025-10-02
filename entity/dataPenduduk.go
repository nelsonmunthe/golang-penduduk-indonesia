package entity

type DataPenduduk struct {
	ID               string `json:"id"`
	TanggalLahir     string `json:"tanggal_lahir"`
	JenisKelamin     string `json:"jenis_kelamin"`
	Nama             string `json:"nama"`
	Alamat           string `json:"alamat"`
	NoHP             string `json:"no_hp"`
	Agama            string `json:"agama"`
	StatusPernikahan string `json:"status_pernikahan"`
	Pendidikan       string `json:"pendidikan"`
	Pekerjaan        string `json:"pekerjaan"`
}

type Chart struct {
	Name       string  `json:"name"`
	Count      float32 `json:"count"`
	Percentage float32 `json:"percentage"`
}
