/*
依赖TiDB的SQL Parser进行SQL语法解析，功能如下：
1.打印MySQL语法树 - 原生
2.打印MySQL语法树 - 美化
*/
package main

import (
	"flag"
	"fmt"
	"gosqltree/session"
	"gosqltree/utils"
	"os"

	_ "github.com/pingcap/parser/test_driver"
)

type flags struct {
	all     bool
	pretty  bool
	id      bool
	element bool
	origin  bool
}

var (
	sql     string
	all     bool
	pretty  bool
	id      bool
	element bool
	origin  bool
)

// usage 自定义usage信息
func usage() {
	fmt.Fprintf(
		os.Stderr,
		`sqltree version: 1.0.0
Usage: sqltree [--sql sqltext] [--all] [--id] [--element] [--origin] [--pretty]
		
Options:`)
	flag.PrintDefaults()
}

func init() {
	flag.StringVar(&sql, "sql", "", "SQL text.")
	flag.BoolVar(&all, "all", false, "Print ALL.")
	flag.BoolVar(&pretty, "pretty", false, "Print pretty JSON.")
	flag.BoolVar(&id, "id", false, "Print SQL ID.")
	flag.BoolVar(&element, "element", false, "Print SQL element.")
	flag.BoolVar(&origin, "origin", false, "Print the original SQL syntax tree.")

	// 改变默认的 Usage
	flag.Usage = usage
}

// main
func main() {
	// 命令行参数解析
	// sql := flag.String("sql", "", "SQL text.")
	// all := flag.Bool("all", false, "Print ALL.")
	// pretty := flag.Bool("pretty", false, "Print pretty JSON.")
	// id := flag.Bool("id", false, "Print SQL ID.")
	// element := flag.Bool("element", false, "Print SQL element.")
	// origin := flag.Bool("origin", false, "Print Original SQL tree.")
	// flag.Usage = usage // 自定义Usage
	// flag.Parse()
	//
	// f := new(flags)
	// f.all = all
	// f.pretty = pretty
	// f.id = id
	// f.element = element
	// f.origin = origin

	flag.Usage = usage // 自定义Usage
	flag.Parse()

	f := new(flags)
	f.all = all
	f.pretty = pretty
	f.id = id
	f.element = element
	f.origin = origin

	// 初始化Session
	s := new(session.Session) // 等价于var s *Session = new(Session)

	// 解析SQL
	stmtNodes, err := utils.ParseSQL(sql, "", "")
	if err != nil {
		fmt.Printf("parse error: %v\n", err.Error())
		return
	}

	for _, stmtNode := range stmtNodes {
		/*
			tree, err := json.Marshal(stmtNode)
			tree, err := json.MarshalIndent(stmtNode, "", "    ")
			tree, err := jsonMarshalIndent(stmtNode, "", "    ")
			if err != nil {
			    fmt.Println(err.Error())
			}
			fmt.Println(string(tree))
			fmt.Println("=========================================================================================")
		*/

		// 根据flag解析stmtNode
		s.GetResult(stmtNode)

		// 根据flag输出结果
		switch {
		case f.all == false:
			if f.id == false && f.element == false && f.origin == false {
				fmt.Printf("%s", utils.PrintResult(s.SQLTree, f.pretty))
			} else {
				if f.id == true {
					fmt.Printf("%s", utils.PrintResult(s.SQLID, f.pretty))
				}

				if f.element == true {
					fmt.Printf("%s", utils.PrintResult(s.SQLElement, f.pretty))
				}

				if f.origin == true {
					fmt.Printf("%s", utils.PrintResult(stmtNode, f.pretty))
				}
			}
		case f.all == true:
			fmt.Printf("%s", utils.PrintResult(s, f.pretty))
		default:
		}
	}
}
