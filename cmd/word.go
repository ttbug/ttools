package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ttbug/ttools/internal/word"
)

var str string
var mode int8

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeCamelCaseToUnderScore:
			content = word.CamelCaseToUnderScore(str)
		case ModeUnderScoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeUnderScoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行hrlp word查看帮助文档")
		}

		log.Printf("输出结果: %s\n", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换模式")
}

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderScoreToUpperCamelCase
	ModeUnderScoreToLowerCamelCase
	ModeCamelCaseToUnderScore
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1: 全部单词转换为大写",
	"2: 全部单词转换为小写",
	"3: 下划线单词转换为大写驼峰单词",
	"4: 下划线单词转换为小写驼峰单词",
	"5: 驼峰单词转换为下划线单词",
}, "\n")
