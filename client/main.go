package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bregydoc/ergo/schema"

	"github.com/bregydoc/ergo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	melody "gopkg.in/olahol/melody.v1"
)

type ErgoUI struct {
	engine *gin.Engine
	ergo   ergo.Ergo
}

func NewErgoUI(engine *gin.Engine, ergo *ergo.Ergo) *ErgoUI {
	return &ErgoUI{
		engine: engine,
		ergo:   *ergo,
	}
}

// LaunchUIClientForDevelopers ...
func (ui *ErgoUI) LaunchUIClientForDevelopers(port string) error {
	r := ui.engine
	m := melody.New()
	m.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	r.Use(cors.Default())

	r.Static("/static", "./client/app/build/static")
	r.StaticFile("/asset-manifest.json", "./client/app/build/static/asset-manifest.json")
	r.StaticFile("/favicon.ico", "./client/app/build/static/favicon.ico")
	r.StaticFile("/idpasscircle.ico", "./client/app/build/static/idpasscircle.ico")
	r.StaticFile("/idpassidcon.ico", "./client/app/build/static/idpassidcon.ico")
	r.StaticFile("/manifest.json", "./client/app/build/static/manifest.json")
	r.StaticFile("/service-worker.js", "./client/app/build/static/service-worker.js")

	r.LoadHTMLFiles("./client/app/build/index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	ui.ergo.Repo.OnNewErrorHasBeenSaved(func(val *schema.ErrorInstance) {
		instances, err := ui.ergo.Repo.GetAllErrorsForDev()
		if err != nil {
			fmt.Println("error 0x01", err)
		}

		data, err := json.Marshal(instances)
		if err != nil {
			fmt.Println("error 0x02", err)
		}

		err = m.Broadcast(data)
		if err != nil {
			fmt.Println("error 0x03", err)
		}
	})

	m.HandleConnect(func(s *melody.Session) {
		fmt.Println("connected...", s.Keys)
		instances, err := ui.ergo.Repo.GetAllErrorsForDev()
		if err != nil {
			fmt.Println("error 0x01", err)
		}

		data, err := json.Marshal(instances)
		if err != nil {
			fmt.Println("error 0x02", err)
		}

		err = s.Write(data)
		if err != nil {
			fmt.Println("error 0x03", err)
		}
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		fmt.Println(string(msg))
	})

	return r.Run(port)
}
