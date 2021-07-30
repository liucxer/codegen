package gen

import (
	"github.com/liucxer/codegen/internal/generate"
	"github.com/go-courier/packagesx"
	"github.com/go-courier/sqlx/v2/generator"
	"github.com/spf13/cobra"
)

var cmdGenModelFlagDatabase string
var cmdGenModelFlagTableName string
var cmdGenModelFlagWithTableName bool
var cmdGenModelFlagWithTableInterfaces bool
var cmdGenModelFlagWithTableMethods bool
var cmdGenModelFlagWithComments bool

func init() {
	cmdGenModel.Flags().
		StringVarP(&cmdGenModelFlagDatabase, "database", "", "", "(required) register model to database var")
	cmdGenModel.Flags().
		StringVarP(&cmdGenModelFlagTableName, "table-name", "t", "", "custom table name")
	cmdGenModel.Flags().
		BoolVarP(&cmdGenModelFlagWithTableName, "with-table-name", "", true, "with Register and interface TableName")
	cmdGenModel.Flags().
		BoolVarP(&cmdGenModelFlagWithTableInterfaces, "with-table-interfaces", "", true, "with table interfaces like Indexes Fields")
	cmdGenModel.Flags().
		BoolVarP(&cmdGenModelFlagWithTableMethods, "with-methods", "", true, "with table methods")
	cmdGenModel.Flags().
		BoolVarP(&cmdGenModelFlagWithComments, "with-comments", "", false, "use comments")

	CmdGen.AddCommand(cmdGenModel)
}

var cmdGenModel = &cobra.Command{
	Use:     "model",
	Aliases: []string{"model2"},
	Short:   "generate interfaces of db mode",
	Run: func(cmd *cobra.Command, args []string) {
		if cmdGenModelFlagDatabase == "" {
			panic("database must be defined")
		}

		for _, arg := range args {
			generate.RunGenerator(func(pkg *packagesx.Package) generate.Generator {
				g := generator.NewSqlFuncGenerator(pkg)
				g.WithComments = true
				g.WithTableInterfaces = true
				g.StructName = arg
				g.Database = cmdGenModelFlagDatabase
				g.TableName = cmdGenModelFlagTableName
				g.WithTableName = cmdGenModelFlagWithTableName
				g.WithTableInterfaces = cmdGenModelFlagWithTableInterfaces
				g.WithMethods = cmdGenModelFlagWithTableMethods
				g.WithComments = cmdGenModelFlagWithComments
				g.Scan()
				return g
			})
		}
	},
}
