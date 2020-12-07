package cmd

import (
	"github.com/Shea11012/go_utils/internal/word"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

var str string
var mode int8

const (
	ModeUpper	= iota + 1	// 全部转大写
	ModeLower				// 全部转小写
	ModeUnderscoreToUpperCamelCase	// 下划线单词转为大写驼峰单词
	ModeUnderscoreToLowerCamelCase	// 下划线单词转为小写驼峰单词
	ModeCamelCaseToUnderscore		// 驼峰单词转下划线
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转换，模式如下：",
	"1: 全部单词转大写",
	"2: 全部单词转小写",
	"3: 下划线单词转大写驼峰单词",
	"4: 下划线单词转小写驼峰单词",
	"5: 驼峰单词转下划线单词",
},"\n")

var wordcmd = &cobra.Command{
	Use: "word",
	Short: "单词格式转换",
	Long: desc,
	Run: func(cmd *cobra.Command, args []string){
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("暂不支持该转换模式，请执行help word 查看帮助文档")
		}

		log.Printf("输出结果：%s",content)
	},
}

func init() {
	wordcmd.Flags().StringVarP(&str,"str","s","","请输入单词内容")
	wordcmd.Flags().Int8VarP(&mode,"mode","m",1,"请输入单词转换模式")
}