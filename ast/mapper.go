package ast

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

func Scan(start *xml.StartElement) (Node, error) {
	var node Node
	switch start.Name.Local {
	case "mapper":
		node = NewMapper()
	case "sql", "update":
		node = NewSqlNode()
	default:
		return node, nil
		//return node, fmt.Errorf("unknow xml %s", start.Name.Local)
	}
	node.Scan(start)
	return node, nil
}

type Data struct {
}

type Param struct {
}

type IfNode struct {
}

type WhereNode struct {
}

type SqlNode struct {
	*ChildrenNode
	Id string
}

func NewSqlNode() *SqlNode {
	n := &SqlNode{}
	n.ChildrenNode = NewNode()
	return n
}

func (s *SqlNode) Scan(start *xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			s.Id = attr.Value
		}
	}
	return nil
}

func (s *SqlNode) String() string {
	return fmt.Sprintf("sql: %s", s.Id)
}

type Mapper struct {
	NameSpace string
	SqlNodes  map[string]*SqlNode
}

func NewMapper() *Mapper {
	return &Mapper{
		SqlNodes: map[string]*SqlNode{},
	}
}

func (m *Mapper) AddChildren(ns ...Node) error {
	for _, n := range ns {
		switch nt := n.(type) {
		case *SqlNode:
			if _, ok := m.SqlNodes[nt.Id]; ok {
				return fmt.Errorf("sql id %s is repeat", nt.Id)
			}
			m.SqlNodes[nt.Id] = nt
		}
	}
	return nil
}

func (m *Mapper) String() string {
	buff := &bytes.Buffer{}
	buff.WriteString("mapper: ")
	buff.WriteString(m.NameSpace)
	buff.WriteString("\n")
	for _, sql := range m.SqlNodes {
		buff.WriteString("\t")
		buff.WriteString(sql.String())
		buff.WriteString("\n")
	}
	return buff.String()
}

func (m *Mapper) Scan(start *xml.StartElement) error {
	for _, attr := range start.Attr {
		if attr.Name.Local == "namespace" {
			m.NameSpace = attr.Value
		}
	}
	return nil
}
