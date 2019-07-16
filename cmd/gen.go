package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var (
	genCmd = &cobra.Command{
		Use:   "gen",
		Short: "生成中间代码指令",
		Long:  `生成中间代码，包含所有代码`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("开始生成中间代码指令", files, searchDir, relathionPutAndPatch, outputDir, parseVendor, parseDependency, propertyStrategyFlag)

		},
	}
	files                string
	searchDir            string
	relathionPutAndPatch string
	outputDir            string
	parseVendor          bool
	parseDependency      bool
	propertyStrategyFlag string
)

func init() {
	// 是否通过具体文件来生成代码，如何有值，那么就searchDir就不能使用
	genCmd.Flags().StringVarP(&files, "files", "f", "", "Create data from files.\n Multi files Slipte from ','")
	genCmd.Flags().StringVarP(&searchDir, "search", "d", "./", "Directory you want to parse")
	// 文件输出位置
	genCmd.Flags().StringVarP(&outputDir, "outputDir", "o", "./models", "Output directory for al the generated files(*.go)")
	genCmd.Flags().StringVarP(&relathionPutAndPatch, "rpap", "r", "None", "Relation of put and patch method, Default None.values:C , B , None")
	genCmd.Flags().BoolVarP(&parseVendor, "parseVendor", "v", false, "Parse go files in 'vendor' folder, disabled by default")
	genCmd.Flags().BoolVarP(&parseDependency, "parseDependency", "p", false, "Parse go files in outside dependency folder, disabled by default")
	genCmd.Flags().StringVarP(&propertyStrategyFlag, "propertyStrategy", "s", "camelcase", "Property Naming Strategy like snakecase,camelcase,pascalcase")
	rootCmd.AddCommand(genCmd)

}
