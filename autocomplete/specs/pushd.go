// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["pushd"] = model.Subcommand{
		Name:        []string{"pushd"},
		Description: `Change the current directory, and push the old current directory onto the directory stack`,
		Args: []model.Arg{{
			Templates: []model.Template{model.TemplateFolders},
		}},
	}
}