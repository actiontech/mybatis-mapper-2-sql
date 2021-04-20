package ast

import "encoding/xml"

type Node interface {
	Scan(start *xml.StartElement) error
	AddChildren(ns ...Node) error
	String() string
	GetStmt(ctx *Context) (string, error)
}

type ChildrenNode struct {
	Children []Node
}

func NewNode() *ChildrenNode {
	return &ChildrenNode{
		Children: []Node{},
	}
}

func (n *ChildrenNode) AddChildren(ns ...Node) error {
	n.Children = append(n.Children, ns...)
	return nil
}

type emptyPrint struct{}

func (ep *emptyPrint) String() string {
	return ""
}
