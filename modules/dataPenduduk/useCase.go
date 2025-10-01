package datapenduduk

import (
	"example/data-penduduk-indonesia/dto"
	"example/data-penduduk-indonesia/entity"
	"example/data-penduduk-indonesia/modules/utils/pagination"
	"example/data-penduduk-indonesia/repository"
)

type DataPendudukUseCase struct {
	dataPendudukRepository repository.DataPendudukRepository
}

func (dataPendudukUseCase *DataPendudukUseCase) GetDataPenduduk(paginate entity.Paginate) (dto.BaseResponseList, error) {

	count, data := dataPendudukUseCase.dataPendudukRepository.GetDataPenduduk(paginate)

	pageable := pagination.Paginate(paginate.CurrentPage, paginate.PerPage, count)

	return dto.BaseResponseList{
		PreviousPage: pageable.PreviousPage,
		CurrentPage:  pageable.CurrentPage,
		NextPage:     pageable.NextPage,
		Total:        int64(pageable.Total),
		PerPage:      pageable.PerPage,
		Success:      true,
		MessageTitle: "List Penduduk Indonesia",
		Message:      "Get list Penduduk Indonesia",
		Data:         data,
	}, nil
}

func (dataPendudukUseCase *DataPendudukUseCase) GetChartByReligion() (dto.ResponseMeta, error) {

	data := dataPendudukUseCase.dataPendudukRepository.GetChartByReligion()

	return dto.ResponseMeta{
		Success:      true,
		MessageTitle: "",
		Message:      "Success to get chart",
		ResponseTime: "",
		Data:         data,
	}, nil
}

func (dataPendudukUseCase *DataPendudukUseCase) GetChartByPekerjaan() (dto.ResponseMeta, error) {

	data := dataPendudukUseCase.dataPendudukRepository.GetChartByPekerjaan()

	return dto.ResponseMeta{
		Success:      true,
		MessageTitle: "",
		Message:      "Success to get chart",
		ResponseTime: "",
		Data:         data,
	}, nil
}

func (dataPendudukUseCase *DataPendudukUseCase) GetChartByPendidikan() (dto.ResponseMeta, error) {

	data := dataPendudukUseCase.dataPendudukRepository.GetChartByPendidikan()

	return dto.ResponseMeta{
		Success:      true,
		MessageTitle: "",
		Message:      "Success to get chart",
		ResponseTime: "",
		Data:         data,
	}, nil
}

func (dataPendudukUseCase *DataPendudukUseCase) GetChartByPernikahan() (dto.ResponseMeta, error) {

	data := dataPendudukUseCase.dataPendudukRepository.GetChartByPernikahan()

	return dto.ResponseMeta{
		Success:      true,
		MessageTitle: "",
		Message:      "Success to get chart",
		ResponseTime: "",
		Data:         data,
	}, nil
}
