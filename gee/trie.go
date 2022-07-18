package gee

import "strings"

type node struct {
	Pattern  string  //待匹配路由
	Part     string  //当前路由
	Children []*node //子节点
	IsWild   bool    //
}

func (n *node) matchChild(part string) *node {
	if len(n.Children) > 0 {
		for _, child := range n.Children {
			if child.Part == part || child.IsWild {
				return child
			}
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	if len(n.Children) > 0 {
		for _, child := range n.Children {
			if child.Part == part || child.IsWild {
				nodes = append(nodes, child)
			}
		}
	}
	return nodes
}

//Insert 新增路由 /p/blog
func (n *node) Insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.Pattern = pattern
		return
	}
	// parts: ["p", "blog"]
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{Part: part, IsWild: part[0] == '*' || part[0] == ':'}
		n.Children = append(n.Children, child)
	}
	child.Insert(pattern, parts, height+1)
}

//Search 匹配路由
func (n *node) Search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.Part, "*") {
		if n.Pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		if child == nil {
			return nil
		}
		result := child.Search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
