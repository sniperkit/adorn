package cmd

import (
	"io"
	"os"

	"github.com/mmcloughlin/adorn"
	"github.com/spf13/cobra"
)

var (
	cfg      adorn.Config
	filename string
)

func init() {
	generateCmd.Flags().StringVarP(&cfg.Package, "package", "p", "", "package name")
	generateCmd.Flags().StringVarP(&cfg.TypeName, "type", "t", "", "type name")
	generateCmd.Flags().StringVarP(&cfg.MethodName, "method", "m", "", "method name")
	generateCmd.Flags().StringSliceVarP(&cfg.ArgumentTypes, "args", "a", nil, "argument types")
	generateCmd.Flags().StringSliceVarP(&cfg.ReturnTypes, "return", "r", nil, "return types")
	generateCmd.Flags().StringVarP(&cfg.OutputFilename, "output", "o", "", "output filename (defaults to stdout)")
	generateCmd.Flags().StringVarP(&filename, "config", "c", "", "config filename")

	RootCmd.AddCommand(generateCmd)
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate type and adornments",
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		if filename != "" {
			cfg, err = adorn.LoadConfigFromFile(filename)
			if err != nil {
				return err
			}
		}

		var w io.Writer = os.Stdout
		if cfg.OutputFilename != "" {
			f, err := os.Create(cfg.OutputFilename)
			if err != nil {
				return err
			}
			defer f.Close()
			w = f
		}

		return adorn.Generate(cfg, w)
	},
}
