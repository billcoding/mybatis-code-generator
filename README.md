# mybatis-code-generator
A Mybatis & JPA code generator

```
Usage of mybatis-code-generator:

mybatis-code-generator -dsn DSN -db DATABASE -OPTIONS

Examples:

mybatis-code-generator -dsn "root:123@tcp(127.0.0.1:3306)/test" -db "Database"
mybatis-code-generator -dsn "root:123@tcp(127.0.0.1:3306)/test" -db "Database" -o "/to/path"
mybatis-code-generator -dsn "root:123@tcp(127.0.0.1:3306)/test" -db "Database" -au "bigboss" -o "/to/path"

Supports options:
  -au string
        The file copyright author (default "bill")
  -db string
        The Database name
  -dsn string
        The MySQL DSN (default "root:123@tcp(127.0.0.1:3306)/test")
  -e    The entity enable? (default true)
  -ea
        The entity @Entity generated? (default true)
  -ec
        The entity comment generated? (default true)
  -eca
        The entity @Column generated? (default true)
  -ecfs int
        The column to field name strategy[0: None,1: OnlyFirstLetterUpper,2: UnderlineToCamel,3: UnderlineToUpper] (default 2)
  -ecp string
        The entity class body prefix
  -ecs string
        The entity class body suffix
  -ee string
        The entity extends class
  -efc
        The entity field comment generated? (default true)
  -ei string
        The entity implements interfaces
  -eia
        The entity @Id generated? (default true)
  -el
        The entity field lombok plugin generated? (default true)
  -elaac
        The entity field lombok plugin @AllArgsConstructor generated? (default true)
  -eld
        The entity field lombok plugin @Date generated? (default true)
  -ellb
        The entity field lombok plugin @Builder generated? (default true)
  -elnac
        The entity field lombok plugin @NoArgsConstructor generated? (default true)
  -ep string
        The entity PKG (default "entity")
  -eta
        The entity @Table generated? (default true)
  -etes int
        The table to entity name strategy[0: None,1: OnlyFirstLetterUpper,2: UnderlineToCamel,3: UnderlineToUpper] (default 3)
  -get string
        The exclude table names[table_a,table_b]
  -git string
        The include table names[table_a,table_b]
  -h    The help info
  -m    The Mapper enable?
  -mc
        The Mapper comment? (default true)
  -mma
        The Mapper @Mapper generated? (default true)
  -mmb
        The Mapper supports Mybatis?
  -mnp string
        The Mapper name prefix
  -mns string
        The Mapper name suffix (default "Mapper")
  -mp string
        The Mapper PKG (default "mapper")
  -o string
        The output dir
  -r    The Repository enable?
  -rc
        The Repository comment? (default true)
  -rnp string
        The Repository name prefix
  -rns string
        The Repository name suffix (default "Repository")
  -rp string
        The Repository PKG (default "repository")
  -rra
        The Repository @Repository generated? (default true)
  -v    The version info
  -vb
        The verbose detail show?
  -x    The mapper xml enable?
  -xc
        The XML comment? (default true)
  -xd string
        The XML Dir generated (default "xml")
```