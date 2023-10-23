package bilog

import (
	"github.com/capeskychung/game_slg/define/retcode"
	"github.com/capeskychung/game_slg/api"
	"net/http"
	"encoding/json"
)

type Controller struct {
}

type inputData struct {
	ClientId string `json:"clientId" validate:"required"`
	UseId    string `json:"user_id" validate:"required"`
	BIKey    string `json:"bi_key" validate:"required"`
	Value    string `json:"value" validate:"value"`
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

	//systemId := r.Header.Get("SystemId")

	api.Render(w, retcode.SUCCESS, "success", map[string]string{})
	return
}
