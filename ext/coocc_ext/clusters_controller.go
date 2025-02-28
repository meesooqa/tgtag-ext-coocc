package coocc_ext

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/meesooqa/tgtag/pkg/controllers"
	"github.com/meesooqa/tgtag/pkg/data"
	"github.com/meesooqa/tgtag/pkg/repositories"
)

type ClustersController struct {
	controllers.BaseController
	provider data.Provider
}

func NewClustersController(repo repositories.Repository) *ClustersController {
	c := &ClustersController{
		BaseController: controllers.BaseController{
			MethodApi:  http.MethodGet,
			RouteApi:   "/api/coocc/clusters",
			Method:     http.MethodGet,
			Route:      "/coocc/clusters",
			Title:      "Tag Clustering",
			ContentTpl: "template/content/clusters.html",
		},
		provider: NewClustersDataProvider(repo),
	}
	c.Self = c
	return c
}

func (c *ClustersController) GetApiData(r *http.Request) map[string]any {
	c.provider.SetLogger(c.Log)
	apiData, err := c.provider.GetData(context.TODO(), r.URL.Query().Get("group"))
	if err != nil {
		c.Log.Error("getting api data", slog.Any("err", err))
		return nil
	}
	return map[string]any{"data": apiData}
}

func (c *ClustersController) GetTplData(r *http.Request) map[string]any {
	tplData, err := c.Tpl.GetData(r, map[string]any{
		"Title": c.GetTitle(),
	})
	if err != nil {
		c.Log.Error("getting template data", slog.Any("err", err))
		return nil
	}
	return tplData
}
