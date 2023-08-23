package ast

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

type SqlNode struct {
	*ChildrenNode
	Id        string
	namespace string
}

func NewSqlNode(ctx *Context) *SqlNode {
	if ctx == nil {
		ctx = NewContext()
	}
	return &SqlNode{
		namespace:    ctx.Namespace,
		ChildrenNode: NewNode(),
	}
}

func (s *SqlNode) Scan(start *xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			s.Id = fmt.Sprintf("%v.%v", s.namespace, attr.Value)
		}
	}
	return nil
}

func (s *SqlNode) GetStmt(ctx *Context) (string, error) {
	buff := bytes.Buffer{}
	for _, a := range s.Children {
		data, err := a.GetStmt(ctx)
		if err != nil {
			return "", err
		}
		buff.WriteString(data)
	}
	return buff.String(), nil
}
