package cmd

import (
	"github.com/Shea11012/go_utils/internal/sql2struct"
	"github.com/spf13/cobra"
	"log"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string


var sqlCmd = &cobra.Command{
	Use: "sql",
	Short: "sql转换和处理",
	Long: "sql转换和处理",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var sql2StructCmd = &cobra.Command{
	Use: "struct",
	Short: "sql转换",
	Long: "sql转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType: dbType,
			Host: host,
			UserName: username,
			Password: password,
			Charset: charset,
		}

		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.connect err: %v",err)
		}

		columns,err := dbModel.GetColumns(dbName,tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumns err: %v",err)
		}

		template := sql2struct.NewStructTemplate()
		templateColumns := template.AssemblyColumns(columns)
		err = template.Generate(tableName,templateColumns)
		if err != nil {
			log.Fatalf("template.Generate err:%v",err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2StructCmd)

	sql2StructCmd.Flags().StringVarP(&username,"username","","zhuzi","请输入数据库账号")
	sql2StructCmd.Flags().StringVarP(&password,"password","","123456","请输入数据库密码")
	sql2StructCmd.Flags().StringVarP(&host,"host","","127.0.0.1","请输入数据库host")
	sql2StructCmd.Flags().StringVarP(&charset,"charset","","utf8mb4","请输入数据库编码")
	sql2StructCmd.Flags().StringVarP(&dbType,"type","","mysql","请输入数据类型")
	sql2StructCmd.Flags().StringVarP(&dbName,"db","","blog","请输入数据库名称")
	sql2StructCmd.Flags().StringVarP(&tableName,"table","","","请输入表名称")
}


