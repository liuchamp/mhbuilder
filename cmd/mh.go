package cmd

import (
	"github.com/liuchamp/mhbuilder/log"
	"github.com/spf13/cobra"
)

// 产生文档
var (
	mhCmd = &cobra.Command{
		Use:   "mh",
		Short: "帮助文档生成",
		Long:  `帮助文档生成filter,post,add`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Info(args, cmd.Args)
		},
	}
	Dir  string
	file string
)

func init() {
	mhCmd.Flags().StringVarP(&Dir, "dir", "d", "./", "收集路径下的文件")
	mhCmd.Flags().StringVarP(&file, "file", "f", "", "目标文件")
	rootCmd.AddCommand(mhCmd)

}
