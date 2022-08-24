package main

import (
	"fmt"
)

/*
 * hash表（散列）
 */

// 定义一个hash table，含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

// 给hashtable编写insert雇员的方法
func (this *HashTable) Insert(employee *Employee) {
	// 确定要添加到哪个链表，使用散列函数
	linkNo := this.HashFunc(employee.Id)
	fmt.Println("linkNo:", linkNo)
	// 使用对应的链表添加
	this.LinkArr[linkNo].Insert(employee)
}

// 编写一个用于散列的方法:使用id进行散列
func (this *HashTable) HashFunc(id int) int {
	return id % 7; // 对应链表的下标
}
// 显示hashtabl所有的员工信息
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}
// 查找一个员工的信息
func (this *HashTable) Find(id int) {
	linkNo := this.HashFunc(id)
	this.LinkArr[linkNo].Find(linkNo, id)
}
// 编辑一个员工的信息
func (this *HashTable) Edit(emp *Employee) {
	linkNo := this.HashFunc(emp.Id)
	this.LinkArr[linkNo].Edit(linkNo, emp)
}
// 删除一个员工的信息
func (this *HashTable) Delete(id int) {
	linkNo := this.HashFunc(id)
	this.LinkArr[linkNo].Delete(linkNo, id)
}

// 定义一个员工结构体
type Employee struct {
	Id int
	Name string
	Next *Employee
}

// 定义一个员工头结构体
// 我们这里的empLink 不带表头，即第一个节点就存放雇员
type EmpLink struct {
	Head *Employee
}
// 添加员工的方法
// 如何保证添加时，id是从小到大的？
func (this *EmpLink) Insert(employee *Employee) {
	cur := this.Head // 这是个辅助指针
	fmt.Println("当前cur=", cur)
	var pre *Employee = nil // 辅助指针

	// 如果当前的EmpLink就是一个空链表
	if cur == nil {
		fmt.Println("当前是空链表，head直接指向员工")
		this.Head = employee
		return
	}

	// 如果不是一个空链表，给emp找到对应的位置并插入
	// 思路是：让cur跟emp比较，然后让pre保持在cur前面
	for {
		if cur != nil {
			// 比较，如果当前id大于传过来的emp的id，则传过来的emp应该插入在cur前面
			if cur.Id > employee.Id {
				// 找到位置
				break
			}
			pre = cur // 保证同步
			cur = cur.Next
		} else {
			break
		}
	}

	// 退出时，我们看下是否将emp添加到链表最后
	// if cur == nil {
	// 	pre.Next = employee
	// 	employee.Next = cur
	// } else {
	// 	pre.Next = employee
	// 	employee.Next = cur
	// }
	
	// 还要判断是否是添加到链表的首位
	if pre == nil {
		employee.Next = cur
		this.Head = employee
	} else {
		pre.Next = employee
		employee.Next = cur
	}
	
	

}

// 显示当前链表的信息
func (this *EmpLink) ShowLink(linkIdx int) {
	if this.Head == nil {
		fmt.Printf("链表%d 为空\n", linkIdx)
	}

	// 遍历当前链表并显示数据
	cur := this.Head // 辅助指针
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s ->", linkIdx, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}
func (this *EmpLink) Find(linkIdx int, id int) {
	cur := this.Head
	var emp *Employee = nil
	if cur == nil {
		fmt.Println("该员工不存在")
		return 
	}
	for {
		if cur == nil {
			break
		}
		if cur.Id == id {
			emp = cur
			break
		}
		cur = cur.Next
	}

	if emp == nil {
		fmt.Printf("id=%d的员工不存在", id)
		return
	}
	fmt.Printf("链表%d 雇员id=%d 名字=%s ->", linkIdx, cur.Id, cur.Name)
	fmt.Println()
}
func (this *EmpLink) Edit(linkIdx int, emp *Employee) {
	cur := this.Head
	if cur == nil {
		fmt.Println("该员工不存在")
		return 
	}
	for {
		if cur == nil {
			break
		}
		if cur.Id == emp.Id {
			break
		}
		cur = cur.Next
	}

	if cur == nil {
		fmt.Printf("id=%d的员工不存在", emp.Id)
		return
	}

	fmt.Printf("修改前的链表%d 雇员id=%d 名字=%s ->", linkIdx, cur.Id, cur.Name)
	fmt.Println()
	cur.Name = emp.Name
	fmt.Printf("修改后的链表%d 雇员id=%d 名字=%s ->", linkIdx, cur.Id, cur.Name)
	fmt.Println()
}
func (this *EmpLink) Delete(linkIdx int, id int) {
	cur := this.Head
	var pre *Employee = nil
	if cur == nil {
		fmt.Println("该员工不存在")
		return 
	}

	for {
		if cur.Id == id || cur == nil{
			break
		}
		pre = cur
		cur = cur.Next
	}

	if cur == nil {
		fmt.Printf("id=%d的员工不存在", id)
		return
	}

	// 当删除的是首位的时候
	if pre == nil {
		this.Head = cur.Next
		return
	}
	// 删除的是中间位置或者末位的时候
	pre.Next = cur.Next
	fmt.Println()
}


func main() {
	key := ""
	id := 0
	name := ""

	var hashTable HashTable
	for {
		fmt.Println("***********雇员系统菜单*********")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show 表示显示雇员")
		fmt.Println("find 表示查找雇员")
		fmt.Println("edit 表示编辑雇员")
		fmt.Println("delete 表示删除雇员")
		fmt.Println("exit 退出系统")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)

		switch key {
			case "input":
				fmt.Println("请输入雇员的id")
				fmt.Scanln(&id)
				fmt.Println("请输入姓名")
				fmt.Scanln(&name)

				emp := &Employee {
					Id: id,
					Name: name,
				}
				hashTable.Insert(emp)
			case "show":
				hashTable.ShowAll()
			case "find":
				fmt.Println("请输入查找雇员的id")
				fmt.Scanln(&id)
				hashTable.Find(id)
			case "edit":
				fmt.Println("请输入要编辑雇员的id")
				fmt.Scanln(&id)
				fmt.Println("请输入要修改的姓名")
				fmt.Scanln(&name)

				emp := &Employee {
					Id: id,
					Name: name,
				}
				hashTable.Edit(emp)
			case "delete":
				fmt.Println("请输入删除雇员的id")
				fmt.Scanln(&id)
				hashTable.Delete(id)
			case "exit":
			default:
				fmt.Println("输入错误")
		}

	}

}