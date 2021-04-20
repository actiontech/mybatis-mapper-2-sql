package ast

import (
	"bytes"
	"encoding/xml"
)

type QueryNode struct {
	*ChildrenNode
	Id   string
	Type string
}

func NewQueryNode() *QueryNode {
	n := &QueryNode{}
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

//func (s *QueryNode) String() string {
//	buff := bytes.Buffer{}
//	for _, child := range s.Children {
//		buff.WriteString(child.String())
//	}
//	return buff.String()
//}

func (s *QueryNode) GetStmt(ctx *Context) (string, error){
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