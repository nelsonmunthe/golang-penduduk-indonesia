package datapenduduk

import (
	"example/data-penduduk-indonesia/dto"
	"example/data-penduduk-indonesia/entity"
)

type DataPendudukController struct {
	usecase DataPendudukUseCase
}

func (controller *DataPendudukController) GetDataPenduduk(paginate entity.Paginate) (dto.BaseResponseList, error) {
	return controller.usecase.GetDataPenduduk(paginate)
}

func (controller *DataPendudukController) GetChartByReligion() (dto.ResponseMeta, error) {
	return controller.usecase.GetChartByReligion()
}

func (controller *DataPendudukController) GetChartByPekerjaan() (dto.ResponseMeta, error) {
	return controller.usecase.GetChartByPekerjaan()
}

func (controller *DataPendudukController) GetChartByPendidikan() (dto.ResponseMeta, error) {
	return controller.usecase.GetChartByPendidikan()
}

func (controller *DataPendudukController) GetChartByPernikahan() (dto.ResponseMeta, error) {
	return controller.usecase.GetChartByPernikahan()
}
