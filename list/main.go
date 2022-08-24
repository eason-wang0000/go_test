package main

import (
	"fmt"
)

type HeroNode struct {
	id int // 排名
	name string // 名称
	nickname string // 别名
	next *HeroNode // 下一个节点
}

// 给链表插入一个节点
// 1.在单链表的最后加入
func insertHeroNode(head *HeroNode, newNode *HeroNode) {
	// 思路
	// 先找到该链表的最后一个节点
	// 创建一个辅助节点
	temp := head
	for {
		if temp.next == nil { // 表示找到最后
			break
		}
		temp = head.next
	}
	temp.next = newNode
}

// 2.根据编号从小到大插入节点
func insertHeroNode2(head *HeroNode, newNode *HeroNode) {
	// 找到适当的节点
	// 创建一个辅助节点
	temp := head
	// 让插入的节点的id和temp节点的下一个节点的id作比较
	for {
		// 如果temp.next == nil，要么是头要么是尾
		if temp.next == nil {
			break
		}
		// 如果temp的next节点的id比当前插入的要大
		if temp.next.id < newNode.id {
			break
		}
		temp = temp.next
	}

	if temp.next == nil {
		temp.next = newNode
		return
	}
	newNode.next = temp.next
	temp.next = newNode
}

// 删除节点
func deleteHeroById(head *HeroNode, id int) {
	// 创建一个辅助节点
	temp := head
	for {
		if temp.next == nil {
			break
		}
		if temp.next.id == id {
			break
		}
		temp = temp.next
	}

	if temp.next == nil {
		fmt.Println("不存在")
		return
	}
	temp.next = temp.next.next
}

// 显示链表所有的信息
func showHeroNode(head *HeroNode) {
	// 创建一个辅助节点
	temp := head
	// 先判断是否是空链表
	if temp.next == nil {
		fmt.Println("空空如也")
		return
	}
	// 遍历这个链表
	for {
		fmt.Printf("[%d , %s, %s]==>", temp.next.id, temp.next.name, temp.next.nickname)
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func main() {
	// 创建一个头结点
	head := &HeroNode{}
	// 创建一个新的HeroNode节点
	newHero := &HeroNode {
		id: 1,
		name: "宋江",
		nickname: "及时雨",
	}
	newHero2 := &HeroNode {
		id: 4,
		name: "卢俊义",
		nickname: "玉麒麟",
	}
	newHero3 := &HeroNode {
		id: 2,
		name: "吴用",
		nickname: "智多星",
	}
	newHero4 := &HeroNode {
		id: 5,
		name: "李逵",
		nickname: "黑旋风",
	}
	insertHeroNode2(head, newHero)
	insertHeroNode2(head, newHero2)
	insertHeroNode2(head, newHero3)
	insertHeroNode2(head, newHero4)
	deleteHeroById(head, 2)
	showHeroNode(head)
}
