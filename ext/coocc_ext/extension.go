package coocc_ext

import (
	"embed"

	"github.com/meesooqa/tgtag/pkg/controllers"
	"github.com/meesooqa/tgtag/pkg/extensions"
	"github.com/meesooqa/tgtag/pkg/repositories"
)

//go:embed template/content/*.html
var fsContentTpl embed.FS

//go:embed template/static
var fsStaticDir embed.FS

type CooccExtension struct {
	extensions.BaseExtension
}

func NewCooccExtension(repo repositories.Repository) *CooccExtension {
	return &CooccExtension{extensions.BaseExtension{
		Name:         "coocc_ext",
		FsContentTpl: fsContentTpl,
		FsStaticDir:  fsStaticDir,
		Controllers: []controllers.Controller{
			NewCooccController(repo),
		},
	}}
}
