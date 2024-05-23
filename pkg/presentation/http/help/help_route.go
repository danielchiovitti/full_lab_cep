package help

import (
	"encoding/json"
	"full_cycle_cep/pkg/core/middleware"
	"full_cycle_cep/pkg/domain/use_cases/viacep/get_viacep"
	"full_cycle_cep/pkg/shared/business_error"
	"full_cycle_cep/pkg/shared/constants"
	"full_cycle_cep/pkg/shared/log"
	"github.com/go-chi/chi/v5"
	"net/http"
	"sync"
)

var lock sync.Mutex
var createHelpRouteInstance CreateHelpRoute

type CreateHelpRoute struct {
	logger                  log.LoggerManagerInterface
	getViaCepUseCase        get_viacep.GetViaCepUseCaseInterface
	cepValidationMiddleware middleware.CepValidationMiddleware
}

func NewCreateHelpRoute(
	logger log.LoggerManagerInterface,
	getViaCepUseCase get_viacep.GetViaCepUseCaseInterface,
) CreateHelpRoute {
	if createHelpRouteInstance == (CreateHelpRoute{}) {
		lock.Lock()
		defer lock.Unlock()
		if createHelpRouteInstance == (CreateHelpRoute{}) {
			createHelpRouteInstance = CreateHelpRoute{
				logger:           logger,
				getViaCepUseCase: getViaCepUseCase,
			}
		}
	}
	return createHelpRouteInstance
}

func (c *CreateHelpRoute) GetHelpRoute() *chi.Mux {
	r := chi.NewRouter()
	r.With(c.cepValidationMiddleware.Validate).Get("/{cep}", c.get)
	return r
}

func (c *CreateHelpRoute) get(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	res, err := c.getViaCepUseCase.Execute(cep)
	//er := &business_error.BusinessError{}

	if err != nil {
		if cepNotFoundErr, ok := err.(*business_error.BusinessError); ok {
			if cepNotFoundErr.Message == string(constants.Cep_not_found) {
				http.Error(w, "can not find zipcode", http.StatusNotFound)
				return
			}
		}
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "can not find zipcode", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
