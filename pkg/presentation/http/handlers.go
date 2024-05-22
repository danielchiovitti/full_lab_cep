package http

import (
	"full_cycle_cep/pkg/presentation/http/help"
	"full_cycle_cep/pkg/shared/log"
	"github.com/go-chi/chi/v5"
)

type Handlers struct {
	logger          log.LoggerManagerInterface
	createHelpRoute help.CreateHelpRoute
}

func ProvideHandlers(
	logger log.LoggerManagerInterface,

) *Handlers {
	return &Handlers{
		logger: logger,
	}
}

func (h *Handlers) GetRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/help", h.createHelpRoute.GetHelpRoute())
	return r
}
