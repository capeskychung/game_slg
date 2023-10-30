package bilog

import (
	"github.com/capeskychung/game_slg/define/retcode"
	"github.com/capeskychung/game_slg/api"
	"net/http"
	"encoding/json"
	"github.com/capeskychung/game_slg/tools/db"
	"fmt"
)

type Controller struct {
}

type BILog struct {
	ClientId string `json:"client_id" validate:"required"`
	UseId    string `json:"user_id" validate:"required"`
	BIKey    string `json:"bi_key" validate:"required"`
	Value    string `json:"value" validate:"required"`
}

func (c *Controller) Run(w http.ResponseWriter, r *http.Request) {
	var inputData BILog
	if err := json.NewDecoder(r.Body).Decode(&inputData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("bi log decode json err:", err.Error())
		return
	}

	err := api.Validate(inputData)
	if err != nil {
		fmt.Println("bi log api validate err:", err.Error())
		api.Render(w, retcode.FAIL, err.Error(), []string{})
		return
	}
	//go func() {
	db := db.GetSqlDb()
	err1 := db.AutoMigrate(&BILog{})
	if err1 != nil {
		fmt.Println("bi log err: ", err1.Error())
		return
	}
	db.Create(&inputData)
	//}()
	api.Render(w, retcode.SUCCESS, "success", map[string]string{})
	return
}
