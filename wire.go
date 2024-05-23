//go:build wireinject
// +build wireinject

package full_lab_cep

import (
	"full_cycle_cep/pkg/core/middleware"
	"full_cycle_cep/pkg/domain/contracts/viacep"
	"full_cycle_cep/pkg/domain/contracts/weatherapi"
	"full_cycle_cep/pkg/domain/use_cases/viacep/get_viacep"
	"full_cycle_cep/pkg/presentation/http"
	"full_cycle_cep/pkg/presentation/http/help"
	"full_cycle_cep/pkg/shared/log"
	"github.com/google/wire"
)

var superset = wire.NewSet(
	wire.Bind(new(log.LoggerManagerInterface), new(*log.LoggerManager)),
	log.NewLoggerManager,
	wire.Bind(new(viacep.ViaCepContractInterface), new(*viacep.ViaCepContract)),
	viacep.NewViaCepContract,
	wire.Bind(new(weatherapi.WeatherApiContractInterface), new(*weatherapi.WeatherApiContract)),
	weatherapi.NewWeatherApiContract,
	wire.Bind(new(get_viacep.GetViaCepUseCaseInterface), new(*get_viacep.GetViaCepUseCase)),
	get_viacep.NewGetViaCepUseCase,
	help.NewCreateHelpRoute,
	middleware.NewCepValidationMiddleware,

	http.ProvideHandlers,
)

func InitializeHandlers() *http.Handlers {
	wire.Build(
		superset,
	)
	return &http.Handlers{}
}
