package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *application) Home(ctx *gin.Context) {
	//ctx.IndentedJSON(http.StatusOK, "The home page")
	app.render(ctx, "home page", &TemplateData{})
}

func (app *application) render(ctx *gin.Context, t string, data *TemplateData) {
	ctx.HTML(http.StatusOK, "gohome.html", t)
}
