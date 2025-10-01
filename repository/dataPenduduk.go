package repository

import (
	"example/data-penduduk-indonesia/entity"
	"strings"
)

type DataPendudukRepository struct {
	data []entity.DataPenduduk
}

func NewDataPenduduk(data []entity.DataPenduduk) DataPendudukRepository {
	return DataPendudukRepository{
		data: data,
	}
}

func (dataPenduduk *DataPendudukRepository) GetDataPenduduk(paginate entity.Paginate) (int, []entity.DataPenduduk) {
	data := dataPenduduk.data

	// Filtering by keyword (case-insensitive) across relevant fields
	if strings.TrimSpace(paginate.Keyword) != "" {
		keyword := strings.ToLower(strings.TrimSpace(paginate.Keyword))
		filtered := make([]entity.DataPenduduk, 0, len(data))
		for _, d := range data {
			if strings.Contains(strings.ToLower(d.ID), keyword) ||
				strings.Contains(strings.ToLower(d.TanggalLahir), keyword) ||
				strings.Contains(strings.ToLower(d.JenisKelamin), keyword) ||
				strings.Contains(strings.ToLower(d.Nama), keyword) ||
				strings.Contains(strings.ToLower(d.Alamat), keyword) ||
				strings.Contains(strings.ToLower(d.NoHP), keyword) ||
				strings.Contains(strings.ToLower(d.Agama), keyword) ||
				strings.Contains(strings.ToLower(d.StatusPernikahan), keyword) ||
				strings.Contains(strings.ToLower(d.Pendidikan), keyword) ||
				strings.Contains(strings.ToLower(d.Pekerjaan), keyword) {
				filtered = append(filtered, d)
			}
		}
		data = filtered
	}

	// Pagination
	if paginate.PerPage <= 0 || paginate.CurrentPage <= 0 {
		return len(data), data
	}

	start := (paginate.CurrentPage - 1) * paginate.PerPage
	if start >= len(data) {
		return 0, []entity.DataPenduduk{}
	}

	end := start + paginate.PerPage

	if end > len(data) {
		end = len(data)
	}

	return len(data), data[start:end]
}

func (datapenduduk *DataPendudukRepository) GetChartByReligion() interface{} {
	data := datapenduduk.data

	grouped := make(map[string]int)

	for _, d := range data {
		grouped[d.Agama] += 1
	}

	return grouped

}

func (datapenduduk *DataPendudukRepository) GetChartByPekerjaan() interface{} {
	data := datapenduduk.data

	grouped := make(map[string]int)

	for _, d := range data {
		grouped[d.Pekerjaan] += 1
	}

	return grouped
}

func (datapenduduk *DataPendudukRepository) GetChartByPendidikan() interface{} {
	data := datapenduduk.data

	grouped := make(map[string]int)

	for _, d := range data {
		grouped[d.Pendidikan] += 1
	}

	return grouped
}

func (datapenduduk *DataPendudukRepository) GetChartByPernikahan() interface{} {
	data := datapenduduk.data

	grouped := make(map[string]int)

	for _, d := range data {
		grouped[d.StatusPernikahan] += 1
	}

	return grouped
}
