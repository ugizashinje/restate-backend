package tests

import (
	"net/http"
	"net/http/httptest"
	url "net/url"
	"os"
	"warrant-api/pkg/config"
	"warrant-api/pkg/ctx"
	"warrant-api/pkg/db"
	"warrant-api/pkg/wire"

	"github.com/gin-gonic/gin"
)

func SetupTest() *gin.Context {
	wire.Svc = wire.Init("testing")
	config.Db.DbLogging = true
	gormDB, _ := db.Init(config.Db)
	w := httptest.NewRecorder()
	g, _ := gin.CreateTestContext(w)
	g.Request = &http.Request{}
	var err error
	g.Request.URL, err = url.Parse("?page=1&pageSize=10")
	if err != nil {
		os.Exit(-1)
	}

	g.Set(ctx.Transaction, gormDB)

	return g
}
