package help

import (
	"full_cycle_cep/pkg/shared/log"
	"github.com/go-chi/chi/v5"
)

type CreateHelpRoute struct {
	logger log.LoggerManagerInterface
}

func (c *CreateHelpRoute) GetHelpRoute() *chi.Mux {
	r := chi.NewRouter()
	return r
}
