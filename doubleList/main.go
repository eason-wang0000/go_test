package main

import (
	"fmt"
)

type HeroDList struct {
	no int
	name string
	nickname string
	prev *HeroDList
	next *HeroDList
}

func insertNode(head *HeroDList, node *HeroDList) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = node
	node.prev = temp
}

// 从小到大
func insertNode2(head *HeroDList, node *HeroDList) {
	temp := head
	for {
		if temp.next == nil {
			break
		}
		if temp.next.no > node.no {
			break
		}
		temp = temp.next
	}

	if temp.next == nil {
		node.next = temp.next
		node.prev = temp
	} else {
		node.next = temp.next
		node.prev = temp
		temp.next.prev = node
	}
	temp.next = node
}

func delNode(head *HeroDList, no int) {
	temp := head
	flag := false
	for {
		fmt.Println(temp.no, no)
		if temp.no == no {
			flag = true
			break
		}
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	// 不用flag时
	// // 空节点
	// if temp.next == nil && temp.prev == nil {
	// 	fmt.Println("空节点，没得删除")
	// 	return
	// }

	// // 证明到最后还是没有，也是不存在
	// if temp.next == nil && temp.no != no {
	// 	fmt.Println("该点不存在")
	// 	return
	// }

	// // 最后一位
	// if temp.next == nil {
	// 	temp.prev.next = temp.next
	// 	return
	// }

	// temp.next.prev = temp.prev
	// temp.prev.next = temp.next
	
	// 用flag时
	fmt.Println(temp)
	fmt.Println(flag)
	if flag == true {
		temp.prev.next = temp.next
		if temp.next != nil {
			temp.next.prev = temp.prev
		}
		return
	}
	fmt.Println("没找到")
}

func delNode2(head *HeroDList, no int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		}
		if temp.next.no == no {
			flag = true
			break
		}
		temp = temp.next
	}

	if flag == true {
		temp.next = temp.next.next
		if temp.next != nil {
			temp.next.prev = temp
		}
		return
	}
	fmt.Println("没找到")

}

// 显示链表所有的信息
func showHeroNode(head *HeroDList) {
	// 创建一个辅助节点
	temp := head
	// 先判断是否是空链表
	if temp.next == nil {
		fmt.Println("空空如也")
		return
	}
	// 遍历这个链表
	for {
		fmt.Printf("[%d , %s, %s]==>", temp.next.no, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
	fmt.Println()
}
func showHeroNode2(head *HeroDList) {
	// 创建一个辅助节点
	temp := head

	// 先判断是否是空链表
	if temp.next == nil {
		fmt.Println("空空如也")
		return
	}

	// 让temp定位到最后节点
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}


	// 遍历这个链表
	for {
		fmt.Printf("[%d , %s, %s]==>", temp.no, temp.name, temp.nickname)
		temp = temp.prev
		if temp.prev == nil {
			break
		}
	}
	fmt.Println()
}
func main() {
	head := &HeroDList {}

	hero := &HeroDList {
		no: 1,
		name: "宋江",
		nickname: "及时雨",
		prev: head,
	}
	hero2 := &HeroDList {
		no: 4,
		name: "吴用",
		nickname: "智多星",
	}
	hero3 := &HeroDList {
		no: 2,
		name: "卢俊义",
		nickname: "玉麒麟",
	}

	insertNode2(head, hero)
	insertNode2(head, hero2)
	insertNode2(head, hero3)
	
	delNode(head, 6)
	fmt.Println("顺序打印：")
	showHeroNode(head)
	fmt.Println("逆序打印：")
	showHeroNode2(head)
}