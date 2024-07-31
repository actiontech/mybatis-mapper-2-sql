package ast

import (
	"bytes"
	"encoding/xml"

	"github.com/pingcap/parser/format"
	"github.com/actiontech/mybatis-mapper-2-sql/sqlfmt"
)

type QueryNode struct {
	*ChildrenNode
	Id        string
	Type      string
	StartLine uint64
}

func NewQueryNode(startLine uint64) *QueryNode {
	n := &QueryNode{
		StartLine: startLine,
	}
	n.ChildrenNode = NewNode()
	return n
}

func (s *QueryNode) Scan(start *xml.StartElement) error {
	s.Type = start.Name.Local
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			s.Id = attr.Value
		}
	}
	return nil
}

func (s *QueryNode) GetStmt(ctx *Context) (string, error) {
	buff := bytes.Buffer{}
	ctx.QueryType = s.Type
	for _, a := range s.Children {
		data, err := a.GetStmt(ctx)
		if err != nil {
			return "", err
		}
		buff.WriteString(data)
	}
	flag := format.RestoreNameBackQuotes|format.RestoreStringDoubleQuotes
	if ctx.Config != nil && ctx.Config.RestoreSqlFlag != 0 {
		flag = format.RestoreFlags(ctx.Config.RestoreSqlFlag)
	}
	return sqlfmt.FormatSQL(buff.String(), flag), nil
}
