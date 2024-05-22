//go:build wireinject
// +build wireinject

package full_lab_cep

import (
	"full_cycle_cep/pkg/presentation/http"
	"full_cycle_cep/pkg/shared/log"
	"github.com/google/wire"
)

var superset = wire.NewSet(
	wire.Bind(new(log.LoggerManagerInterface), new(*log.LoggerManager)),
	log.NewLoggerManager,

	http.ProvideHandlers,
)

func InitializeHandlers() *http.Handlers {
	wire.Build(
		superset,
	)
	return &http.Handlers{}
}
