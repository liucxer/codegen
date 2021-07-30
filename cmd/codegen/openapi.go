package main

import (
	"github.com/liucxer/codegen/internal/generate"
	"github.com/go-courier/httptransport/openapi/generator"
	"github.com/go-courier/packagesx"
	_ "github.com/go-courier/validator/strfmt"
	"github.com/spf13/cobra"
)

var cmdSwagger = &cobra.Command{
	Use:     "openapi",
	Aliases: []string{"swagger"},
	Short:   "scan current project and generate openapi.json",
	Run: func(cmd *cobra.Command, args []string) {
		generate.RunGenerator(func(pkg *packagesx.Package) generate.Generator {
			g := generator.NewOpenAPIGenerator(pkg)
			g.Scan(nil)
			return g
		})
	},
}

func init() {
	cmdRoot.AddCommand(cmdSwagger)
}
