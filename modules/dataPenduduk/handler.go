package datapenduduk

import (
	"example/data-penduduk-indonesia/dto"
	"example/data-penduduk-indonesia/entity"
	"example/data-penduduk-indonesia/modules/utils/pagination"
	"example/data-penduduk-indonesia/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DataPendudukHandler struct {
	data       []entity.DataPenduduk
	controller DataPendudukController
}

func NewRequestHandler(data []entity.DataPenduduk) DataPendudukHandler {
	return DataPendudukHandler{
		data: data,
	}
}

func (handler DataPendudukHandler) DataPendudukHandler(router *gin.Engine) {
	repoDataPenduduk := repository.NewDataPenduduk(handler.data)
	usecaseDataPenduduk := DataPendudukUseCase{
		dataPendudukRepository: repoDataPenduduk,
	}

	handler.controller = DataPendudukController{
		usecase: usecaseDataPenduduk,
	}

	routerGroup := router.Group("/data-penduduk")
	routerGroup.GET("/", handler.GetDataPenduduk)
	routerGroup.GET("/religion-chart", handler.GetChartByReligion)
	routerGroup.GET("/pernikahan-chart", handler.GetChartByPernikahan)
	routerGroup.GET("/pekerjaan-chart", handler.GetChartByPekerjaan)
	routerGroup.GET("/pendidikan-chart", handler.GetChartByPendidikan)
}

func (handler DataPendudukHandler) GetDataPenduduk(context *gin.Context) {

	paginate := pagination.GetPaginationQueryParam(context)

	response, err := handler.controller.GetDataPenduduk(paginate)

	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}

	context.JSON(http.StatusOK, response)
}

func (handler DataPendudukHandler) GetChartByReligion(context *gin.Context) {

	response, err := handler.controller.GetChartByReligion()

	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}

	context.JSON(http.StatusOK, response)
}

func (handler DataPendudukHandler) GetChartByPekerjaan(context *gin.Context) {

	response, err := handler.controller.GetChartByPekerjaan()

	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}

	context.JSON(http.StatusOK, response)
}

func (handler DataPendudukHandler) GetChartByPendidikan(context *gin.Context) {

	response, err := handler.controller.GetChartByPendidikan()

	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}

	context.JSON(http.StatusOK, response)
}

func (handler DataPendudukHandler) GetChartByPernikahan(context *gin.Context) {

	response, err := handler.controller.GetChartByPernikahan()

	if err != nil {
		context.JSON(http.StatusBadRequest, dto.DefaultErrorInvalidDataWithMessage(err.Error()))
		return
	}

	context.JSON(http.StatusOK, response)
}
