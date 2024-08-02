package parser

import (
	"github.com/actiontech/mybatis-mapper-2-sql/ast"
	"github.com/pingcap/parser/format"
)

func SkipErrorQuery() func(*ast.Config) {
	return func(c *ast.Config) {
		c.SkipErrorQuery = true
	}
}

func PgRestoreSqlFlag() func(*ast.Config) {
	return func(c *ast.Config) {
		c.RestoreSqlFlag = format.RestoreNameDoubleQuotes | format.RestoreStringSingleQuotes
	}
}
