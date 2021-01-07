# gosqltree

gosqltree基于[TiDB的语法解析器](https://github.com/pingcap/parser)，提供了对于MySQL SQL语句的离线解析，可生成易读的语法树、tidb/parser的原生语法树，或直接获取Schema、Table、Column列表。

&nbsp;
### 安装说明
[二进制免安装](https://github.com/Ben2277/gosqltree/releases)

&nbsp;
### 使用说明
#### 帮助信息
```
# ./gosqltree --help
gosqltree version: 1.0.0
Usage: gosqltree [--sql sqltext] [--all] [--id] [--element] [--origin] [--pretty]

Options:
  -all
        Print ALL.
  -element
        Print SQL element.
  -id
        Print SQL ID.
  -origin
        Print the original SQL syntax tree.
  -pretty
        Print pretty JSON.
  -sql string
        SQL text.
  -version
        Print gosqltree version.
```

&nbsp;
#### 简单示例
##### // 打印版本
```
# ./gosqltree --version
gosqltree Version: v1.0.1
```

##### // 打印SQL中的对象信息
```
# ./gosqltree --sql "select an1.c1,an2.c2 from s1.t1 an1 join s2.t2 an2 on an1.c3=an2.c3 where an1.c4='test'" --element
{"SchemaList":["s1","s2"],"TableList":["t1","t2"],"ColumnList":["c1","c2","c3","c4"],"AsNameList":["an1","an2"]}
```

##### // 打印create table语句的语法树
```
# ./gosqltree --sql "create table t1 (id int comment 'i am id', name varchar(2) not null) engine=innodb" --pretty
{
    "StmtType": "CreateTable",
    "StmtTree": {
        "IfNotExists": false,
        "IsTemporary": false,
        "Table": {
            "Schema": {
                "lower-case": "",
                "original": ""
            },
            "Name": {
                "lower-case": "t1",
                "original": "t1"
            },
            "IndexHints": null,
            "PartitionNames": []
        },
        "ReferTable": null,
        "Columns": [
            {
                "Name": {
                    "Schema": {
                        "lower-case": "",
                        "original": ""
                    },
                    "Table": {
                        "lower-case": "",
                        "original": ""
                    },
                    "Name": {
                        "lower-case": "id",
                        "original": "id"
                    }
                },
                "Type": {
                    "Datatype": "int",
                    "Flag": 0,
                    "Fieldlen": -1,
                    "Decimal": -1,
                    "Charset": "",
                    "Collate": "",
                    "Elems": null
                },
                "Options": [
                    {
                        "Type": "Comment",
                        "Expr": "i am id",
                        "Stored": false,
                        "Refer": null,
                        "StrValue": "",
                        "AutoRandomBitLength": 0,
                        "Enforced": false,
                        "ConstraintName": ""
                    }
                ]
            },
            {
                "Name": {
                    "Schema": {
                        "lower-case": "",
                        "original": ""
                    },
                    "Table": {
                        "lower-case": "",
                        "original": ""
                    },
                    "Name": {
                        "lower-case": "name",
                        "original": "name"
                    }
                },
                "Type": {
                    "Datatype": "varchar",
                    "Flag": 0,
                    "Fieldlen": 2,
                    "Decimal": -1,
                    "Charset": "",
                    "Collate": "",
                    "Elems": null
                },
                "Options": [
                    {
                        "Type": "NotNull",
                        "Expr": null,
                        "Stored": false,
                        "Refer": null,
                        "StrValue": "",
                        "AutoRandomBitLength": 0,
                        "Enforced": false,
                        "ConstraintName": ""
                    }
                ]
            }
        ],
        "Constraints": null,
        "Options": [
            {
                "Type": "Engine",
                "Default": false,
                "Value": "innodb",
                "TableNames": null
            }
        ],
        "Partition": null,
        "OnDuplicate": "ERROR",
        "Select": null
    }
}
```

&nbsp;
### 致谢
- [TiDB](https://github.com/pingcap/tidb)
