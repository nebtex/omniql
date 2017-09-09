package golang

import (
	"github.com/nebtex/omniql/pkg/io/omniql/corev1"
	"github.com/nebtex/omniql/pkg/utils"
	"strings"
)

func GenerateDocs(name string, d corev1.DocumentationReader) (value string) {
	if d != nil {
		if d.Short() == "" && d.Long() == "" {
			return "//" + strings.Title(name) + " ..."
		}

		if d.Short() != "" && d.Long() != "" {
			return strings.Title(name) + " " + d.Short() + "\n" + d.Long()
		}
		if d.Short() != "" {
			return utils.CommentGolang(strings.Title(name) + " " + d.Short())
		}
		if d.Long() != "" {
			return utils.CommentGolang(strings.Title(name) + " " + d.Long())
		}
	}
	return "//" + strings.Title(name) + " ..."
}
