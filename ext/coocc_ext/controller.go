package coocc_ext

import (
	"log/slog"
	"net/http"

	"github.com/meesooqa/tgtag/pkg/controllers"
	"github.com/meesooqa/tgtag/pkg/repositories"
)

type CooccController struct {
	controllers.BaseController
}

func NewCooccController(repo repositories.Repository) *CooccController {
	c := &CooccController{controllers.BaseController{
		Method:     http.MethodGet,
		Route:      "/coocc",
		Title:      "Coocc Extension Page",
		ContentTpl: "template/content/coocc.html",
	}}
	c.Self = c
	return c
}

func (c *CooccController) GetTplData(r *http.Request) map[string]any {
	data, err := c.Tpl.GetData(r, map[string]any{
		"Title": c.GetTitle(),
	})
	if err != nil {
		c.Log.Error("getting tpl data", slog.Any("err", err))
		return nil
	}
	return data
}
