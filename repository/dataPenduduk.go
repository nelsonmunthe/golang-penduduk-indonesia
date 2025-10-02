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

func (datapenduduk *DataPendudukRepository) GetChartByReligion() []entity.Chart {
	data := datapenduduk.data

	grouped := make(map[string]entity.Chart)

	for _, d := range data {
		grouped[d.Agama] = entity.Chart{
			Name:       d.Agama,
			Count:      (grouped[d.Agama].Count) + 1,
			Percentage: (grouped[d.Agama].Count + 1) / float32(len(data)),
		}
	}

	charts := []entity.Chart{}

	for _, v := range grouped {
		charts = append(charts, v)
	}

	return charts

}

func (datapenduduk *DataPendudukRepository) GetChartByPekerjaan() []entity.Chart {
	data := datapenduduk.data

	grouped := make(map[string]entity.Chart)

	for _, d := range data {
		grouped[d.Pekerjaan] = entity.Chart{
			Name:       d.Pekerjaan,
			Count:      (grouped[d.Pekerjaan].Count) + 1,
			Percentage: (grouped[d.Pekerjaan].Count + 1) / float32(len(data)),
		}
	}

	charts := []entity.Chart{}

	for _, v := range grouped {
		charts = append(charts, v)
	}

	return charts
}

func (datapenduduk *DataPendudukRepository) GetChartByPendidikan() []entity.Chart {
	data := datapenduduk.data

	grouped := make(map[string]entity.Chart)

	for _, d := range data {
		grouped[d.Pendidikan] = entity.Chart{
			Name:       d.Pendidikan,
			Count:      (grouped[d.Pendidikan].Count) + 1,
			Percentage: (grouped[d.Pendidikan].Count + 1) / float32(len(data)),
		}
	}

	charts := []entity.Chart{}

	for _, v := range grouped {
		charts = append(charts, v)
	}

	return charts
}

func (datapenduduk *DataPendudukRepository) GetChartByPernikahan() []entity.Chart {
	data := datapenduduk.data

	grouped := make(map[string]entity.Chart)

	for _, d := range data {
		grouped[d.StatusPernikahan] = entity.Chart{
			Name:       d.StatusPernikahan,
			Count:      (grouped[d.StatusPernikahan].Count) + 1,
			Percentage: (grouped[d.StatusPernikahan].Count + 1) / float32(len(data)),
		}
	}

	charts := []entity.Chart{}

	for _, v := range grouped {
		charts = append(charts, v)
	}

	return charts
}
