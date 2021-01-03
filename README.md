# gosqltree

gosqltree基于[TiDB的语法解析器](https://github.com/pingcap/parser)，提供了对于MySQL SQL语句的离线解析，可生成易读的语法树、tidb/parser的原生语法树，或直接获取Schema、Table、Column列表。

### 安装说明


#### 二进制免安装


### 使用说明


```
# ./gosqltree --help
sqltree version: 1.0.0
Usage: sqltree [--sql sqltext] [--all] [--id] [--element] [--origin] [--pretty]

Options:  -all
        Print ALL.
  -element
        Print SQL element.
  -id
        Print SQL ID.
  -pretty
        Print pretty JSON.
  -sql string
        SQL text.
  -tree
        Print SQL tree.
```
