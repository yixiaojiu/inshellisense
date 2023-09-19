// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["yomo"] = model.Subcommand{
		Name:        []string{"yomo"},
		Description: `CLI interface for YoMo`,
		Options: []model.Option{{
			Name:        []string{"--help", "-h"},
			Description: `Show help for yomo`,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"init"},
			Description: `Initial an example StreamFunction`,
			Args: []model.Arg{{
				Name:        "function name",
				Description: `StreamFunction name to initialize locally`,
			}},
			Options: []model.Option{{
				Name:        []string{"--rx"},
				Description: `Generate Rx code template`,
			}},
		}, {
			Name:        []string{"build"},
			Description: `Build a StreamFunction to WebAssembly`,
			Args: []model.Arg{{
				Name:        ".go file",
				Description: `The .go file to build`,
				Generator:   nil, // TODO: port over generator
			}},
			Options: []model.Option{{
				Name:        []string{"--target"},
				Description: `Build to wasm or binary`,
				Args: []model.Arg{{
					Name: "wasm",
				}, {
					Name: "binary",
				}},
			}, {
				Name:        []string{"-m", "--modfile"},
				Description: `Custom go.mod filepath`,
				Args: []model.Arg{{
					Name:      "module",
					Generator: nil, // TODO: port over generator
				}},
			}},
		}, {
			Name:        []string{"run"},
			Description: `Run a wasm stream function`,
			Args: []model.Arg{{
				Name:        ".wasm file",
				Description: `The .wasm file to run`,
				Generator:   nil, // TODO: port over generator
			}},
			Options: []model.Option{{
				Name:        []string{"-z", "--zipper"},
				Description: `Zipper endpoint this StreamFunction will connect to`,
			}, {
				Name:        []string{"-n", "--name"},
				Description: `Specify the name of this StreamFunction`,
			}},
		}},
	}
}