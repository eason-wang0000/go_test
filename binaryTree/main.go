package main

import (
	"fmt"
)

/**
 * 二叉树
 */
type Hero struct {
	No int
	Name string
	Left *Hero
	Right *Hero
}

// 前序遍历：先输出root节点，然后再输出左子树，然后再输出右子树
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

// 中序遍历: 先输出root的左子树，然后再输出root节点，然后再输出右子树
func InfixOrder (node *Hero) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		InfixOrder(node.Right)
	}
}

// 后序排序：先输出root的左子树，然后再输出右子树，然后再输出root节点
func PostOrder(node *Hero) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
	}
}

func main() {
	root := &Hero {
		No: 1,
		Name: "宋江",
	}

	left1 := &Hero {
		No: 2,
		Name: "吴用",
	}

	right1 := &Hero {
		No: 3,
		Name: "卢俊义",
	}

	right1.Left = &Hero {
		No: 666,
		Name: "xxxx",
	}

	root.Left = left1
	root.Right = right1 

	right2 := &Hero {
		No: 4,
		Name: "林冲",
	}

	left11 := &Hero {
		No: 5,
		Name: "xin",
	}
	right11 := &Hero {
		No: 6,
		Name: "xin1",
	}
	right22 := &Hero {
		No: 7,
		Name: "xin22",
	}
	left22 := &Hero {
		No: 8,
		Name: "xin~",
	}
	left1.Left = left11
	left1.Right = right11

	right1.Right = right2
	right2.Right = right22
	right2.Left = left22

	fmt.Println("原始树：")
	fmt.Println(root)
	fmt.Println("前序遍历：")
	PreOrder(root)
	fmt.Println("中序遍历：")
	InfixOrder(root)
	fmt.Println("后序遍历：")
	PostOrder(root)
}