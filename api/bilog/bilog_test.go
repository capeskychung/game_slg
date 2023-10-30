package bilog

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/capeskychung/game_slg/pkg/setting"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"fmt"
)

type testServer struct {
	*httptest.Server
	ClientURL string
}

type retMessage struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func newServer(t *testing.T) *testServer {
	var s testServer
	setting.Default()

	controller := &Controller{}
	s.Server = httptest.NewServer(http.HandlerFunc(controller.Run))
	s.ClientURL = s.Server.URL + "/api/bilog"

	return &s
}

func TestRun(t *testing.T) {
	fmt.Println("---------------0000")
	s := newServer(t)
	defer s.Close()
	fmt.Println("---------------111sss1")

	testContent := `{"client_id":"ade447d79f6489b5","user_id":"10112122","bi_key":"level", "value": "10"}`

	resp, err := http.Post(s.ClientURL, "application/json", strings.NewReader(testContent))
	Convey("测试bilog", t, func() {
		Convey("是否有报错", func() {
			So(err, ShouldBeNil)
		})
	})
	fmt.Println("---------------111111")
	defer resp.Body.Close()
}
