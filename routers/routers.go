package routers

import (
	"github.com/capeskychung/game_slg/api/bind2group"
	"github.com/capeskychung/game_slg/api/closeclient"
	"github.com/capeskychung/game_slg/api/getonlinelist"
	"github.com/capeskychung/game_slg/api/register"
	"github.com/capeskychung/game_slg/api/send2client"
	"github.com/capeskychung/game_slg/api/send2clients"
	"github.com/capeskychung/game_slg/api/send2group"
	"github.com/capeskychung/game_slg/servers"
	"net/http"
	"github.com/capeskychung/game_slg/api/bilog"
)

func Init() {
	registerHandler := &register.Controller{}
	sendToClientHandler := &send2client.Controller{}
	sendToClientsHandler := &send2clients.Controller{}
	sendToGroupHandler := &send2group.Controller{}
	bindToGroupHandler := &bind2group.Controller{}
	getGroupListHandler := &getonlinelist.Controller{}
	closeClientHandler := &closeclient.Controller{}
	biLogHandler := &bilog.Controller{}

	http.HandleFunc("/api/register", registerHandler.Run)
	http.HandleFunc("/api/send_to_client", AccessTokenMiddleware(sendToClientHandler.Run))
	http.HandleFunc("/api/send_to_clients", AccessTokenMiddleware(sendToClientsHandler.Run))
	http.HandleFunc("/api/send_to_group", AccessTokenMiddleware(sendToGroupHandler.Run))
	http.HandleFunc("/api/bind_to_group", AccessTokenMiddleware(bindToGroupHandler.Run))
	http.HandleFunc("/api/get_online_list", AccessTokenMiddleware(getGroupListHandler.Run))
	http.HandleFunc("/api/close_client", AccessTokenMiddleware(closeClientHandler.Run))
	http.HandleFunc("/api/record_bi_log", AccessTokenMiddleware(biLogHandler.Run))

	servers.StartWebSocket()

	//db.InitDB()

	go servers.WriteMessage()
}
