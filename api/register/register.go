package register

import (
	"encoding/json"
	"github.com/capeskychung/game_slg/api"
	"github.com/capeskychung/game_slg/define/retcode"
	"github.com/capeskychung/game_slg/servers"
	"net/http"
)

type Controller struct {
}

type inputData struct {
	SystemId string `json:"systemId" validate:"required"`
}

func (c *Controller) Run(w http.ResponseWriter, r *http.Request) {
	var inputData inputData
	if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := api.Validate(inputData)
	if err != nil {
		api.Render(w, retcode.FAIL, err.Error(), []string{})
		return
	}

	err = servers.Register(inputData.SystemId)
	if err != nil {
		api.Render(w, retcode.FAIL, err.Error(), []string{})
		return
	}

	api.Render(w, retcode.SUCCESS, "success", []string{})
	return
}
