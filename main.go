package main

import (
	"example/data-penduduk-indonesia/entity"
	datapenduduk "example/data-penduduk-indonesia/modules/dataPenduduk"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	router := gin.Default()
	var err error

	data := []entity.DataPenduduk{
		{
			ID:               "1",
			TanggalLahir:     "1990-05-12",
			JenisKelamin:     "Laki-laki",
			Nama:             "Budi Santoso",
			Alamat:           "Jl. Merdeka No. 45, Jakarta",
			NoHP:             "081234567890",
			Agama:            "Islam",
			StatusPernikahan: "Menikah",
			Pendidikan:       "S1",
			Pekerjaan:        "Pegawai Negeri",
		},
		{
			ID:               "2",
			TanggalLahir:     "1995-08-23",
			JenisKelamin:     "Perempuan",
			Nama:             "Siti Aminah",
			Alamat:           "Jl. Sudirman No. 12, Bandung",
			NoHP:             "082345678901",
			Agama:            "Islam",
			StatusPernikahan: "Belum Menikah",
			Pendidikan:       "SMA",
			Pekerjaan:        "Karyawan Swasta",
		},
		{
			ID:               "3",
			TanggalLahir:     "1987-01-15",
			JenisKelamin:     "Laki-laki",
			Nama:             "Andi Wijaya",
			Alamat:           "Jl. Diponegoro No. 88, Surabaya",
			NoHP:             "083456789012",
			Agama:            "Kristen",
			StatusPernikahan: "Menikah",
			Pendidikan:       "S2",
			Pekerjaan:        "Dosen",
		},
		{
			ID:               "4",
			TanggalLahir:     "2000-11-30",
			JenisKelamin:     "Perempuan",
			Nama:             "Maria Fransiska",
			Alamat:           "Jl. Gatot Subroto No. 5, Medan",
			NoHP:             "084567890123",
			Agama:            "Katolik",
			StatusPernikahan: "Belum Menikah",
			Pendidikan:       "Diploma",
			Pekerjaan:        "Mahasiswa",
		},
		{
			ID:               "5",
			TanggalLahir:     "1978-03-21",
			JenisKelamin:     "Laki-laki",
			Nama:             "Sugeng Raharjo",
			Alamat:           "Jl. Imam Bonjol No. 7, Yogyakarta",
			NoHP:             "085678901234",
			Agama:            "Hindu",
			StatusPernikahan: "Menikah",
			Pendidikan:       "SMA",
			Pekerjaan:        "Wiraswasta",
		},
	}

	penduduk := datapenduduk.NewRequestHandler(data)
	penduduk.DataPendudukHandler(router)

	err = router.Run(":" + os.Getenv("PORT"))

	if err != nil {
		log.Println("main router.Run:", err)
		return
	}
}
