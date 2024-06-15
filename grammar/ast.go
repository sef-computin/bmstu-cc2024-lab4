package grammar

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Node struct {
	Id       int
	Value    string
	Children []*Node
}

var idCounter = 0

func NewNode(val string) *Node {
	idCounter++
	return &Node{Id: idCounter, Value: val, Children: []*Node{}}
}

func (n *Node) AddChild(token string) (ret *Node) {
	ret = NewNode(token)
	n.Children = append(n.Children, ret)
	return
}

func (n *Node) AddChildrenRef(children ...*Node) {
	for _, c := range children {
		n.Children = append(n.Children, c)
	}
	return
}

func (n *Node) GetName() string {
	return fmt.Sprintf("\"%s-%d\"", n.Value, n.Id)
}

func DrawTree(filename string, root *Node) {
	var buf []byte = []byte("digraph AST{\nforcelabels=true;\n")

	buf = drawNodesAST(buf, root)

	buf = append(buf, []byte("\n}")...)

	name := fmt.Sprintf("./%s.gv", filename)

	os.WriteFile(name, buf, 0666)

	cmd := exec.Command("dot", "-Tsvg", name)
	stdout, err := cmd.Output()
	if err == nil {
		file, _ := os.Create(name + ".svg")
		file.Write(stdout)
	}
	os.Remove(name)

	cmd = exec.Command("firefox", name+".svg")
	_, _ = cmd.Output()

	// os.Remove(name+"")

	return
}

func drawNodesAST(buf []byte, node *Node) []byte {
	if node.Children == nil || len(node.Children) == 0 {
		buf = append(buf, []byte(fmt.Sprintf("%s [ shape=ellipse, label=\"%s\" ];\n", node.GetName(), node.Value))...)
	} else {
		buf = append(buf, []byte(fmt.Sprintf("%s [ shape=box, label=\"%s\" ];\n", node.GetName(), node.Value))...)
		for _, child := range node.Children {
			buf = drawNodesAST(buf, child)
			buf = append(buf, []byte(fmt.Sprintf("%s -> %s;\n", node.GetName(), child.GetName()))...)
		}
	}

	return buf
}

func (node *Node) isLeaf() bool {
	return node.Children == nil || len(node.Children) == 0
}

func (node *Node) getPolishString() (ret string) {
	if node.isLeaf() {
		if node.Value != "begin" && node.Value != "end" && node.Value != ";" {
			return node.Value
		}
		return ""
	}

	for _, child := range node.Children {
		if child.isLeaf() && matchTokenType(child.Value) != IDENTIFIER {
			// term = append(term, node.getPolishString())
			ret = fmt.Sprintf("%s %s", ret, child.getPolishString())
			ret = strings.TrimSpace(ret)
		} else {
			// nonterm = append(nonterm, node.getPolishString())
			ret = fmt.Sprintf("%s %s", child.getPolishString(), ret)
			ret = strings.TrimSpace(ret)
		}
	}

	return
}

// func GetAstFromPTree(root *Node) *Node {
//
// 	root = root.Children[0]
// 	// root = root.Children[1]
//
// 	ast := NewNode("Программа")
// 	ret := ast
// 	// ast = ast.AddChild("Список операторов").Children[0]
//
// 	ast.AddChildrenRef(getAstFromNode(root)...)
//
// 	return ret
// }
//
// func getAstFromNode(node *Node) (ret []*Node) {
// 	ret = []*Node{}
//   if node.isLeaf() && {
//
//   }
// 	switch node.Value {
// 	case "Блок":
// 		ret = []*Node{}
// 		nod := NewNode("Блок")
// 		nod.AddChildrenRef(getAstFromNode(node.Children[1])...)
// 		ret = append(ret, nod)
// 		break
// 	case "Список операторов":
// 		ret = []*Node{}
// 		for _, child := range node.Children {
// 			if child.Value != ";" {
// 				ret = append(ret, getAstFromNode(child)...)
// 			}
// 		}
//   case "Оператор""Выражение":
//
// 	}
// 	return
// }
