package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// 产生文档
var (
	docCmd = &cobra.Command{
		Use:   "doc",
		Short: "生成文档",
		Long:  `调用swag等生成文档`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("文档生成", files, searchDir, relathionPutAndPatch, outputDir, parseVendor, parseDependency, propertyStrategyFlag)

		},
	}
)

func init() {
	rootCmd.AddCommand(docCmd)

}
