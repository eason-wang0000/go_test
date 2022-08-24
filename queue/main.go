package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用数组实现队列的思路
// 1.创建一个array，作为队列的一个字段
// 2.front初始化为-1
// 3.rear，表示队列尾，初始化为-1
// 4.完成队列的基本查找
// AddQueue加入数据到队列
// GetQueue从队列取出数据
// ShowQueue显示队列
// 简单的队列，单向队列
type NormalQueue struct {
	maxSize int // 队列长度？
	array [5]int // 队列数组
	front int // 队首
	rear int // 队尾
}


func (this *NormalQueue) AddQueue(val int) error {
	// 先判断队列是否满
	if this.rear == this.maxSize - 1 {
		return errors.New("queue full")
	}

	this.rear++
	this.array[this.rear] = val
	fmt.Println("添加数据成功")
	return nil
}

func (this *NormalQueue) ShowQueue() {
	// 找到队首，然后遍历到队尾
	// this.front不包含队首的元素
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t\n", i, this.array[i])
	}
}

func (this *NormalQueue) GetQueue() (val int, err error) {
	// 判断队列是否为空
	if this.front == this.rear {
		return -1, errors.New("queue empty")
	}
	this.front++
	val = this.array[this.front]
	return val, nil
}

// 环形队列
// 尾索引的下一个为头索引时表示队列满：队列需要空出一个位置作为约定，实际长度为数组长度-1
//（rear + 1）% maxSize == head 【满】 
// head==tail	
// 统计有几个元素：(tail+maxSize-head)%maxSize
type CircleQueue struct {
	maxSize int
	array [4]int
	head int // 0
	tail int // 0
}

func (this *CircleQueue) Push(val int) error {
	// 判断是否满
	if (this.tail + 1) % this.maxSize == this.head {
		return errors.New("queue full")
	}

	// this.tail在队列尾部，但是不包含最后的元素
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize
	return nil
}

func (this *CircleQueue) Pop() (val int, err error) {
	if this.head == this.tail {
		return -1, errors.New("queue empty")
	}
	// 取,head指向队首，并且包含队首的元素
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return val, nil
}

func (this *CircleQueue) List() {
	// 取出当前队列有多少个元素
	size := this.Size()
	if size == 0 {
		fmt.Println("queue empty")
		return
	}
	fmt.Println(this.array)
	// 设计一个辅助变量，指向head
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

func (this *CircleQueue) Size() int {
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func main() {
	queue := &NormalQueue {
		maxSize: 5,
		front: -1,
		rear: -1,
	}

	cirleQueue := &CircleQueue{
		maxSize: 4,
		head: 0,
		tail: 0,
	}

	var key string
	var val int

	for {
		fmt.Println("1. add1 添加数据到普通队列")
		fmt.Println("2. get1 获取普通队列数据")
		fmt.Println("3. show1 显示队列")
		fmt.Println("4. add2 添加数据到环形队列")
		fmt.Println("5. get2 获取环形队列数据")
		fmt.Println("6. show2 显示环形队列")
		fmt.Println("7. exit 退出")

		fmt.Scanln(&key)
		switch key {
			case "add1":
				fmt.Println("输入入队列数：")
				fmt.Scanln(&val)
				err := queue.AddQueue(val)
				if err != nil {
					fmt.Println(err)
				}
			case "get1":
				fmt.Println("开始获取队列数据")
				val, err := queue.GetQueue()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("获取到的数据为：%d\n", val)
				}
			case "show1":
				fmt.Println("显示普通队列")
				queue.ShowQueue()
			case "add2":
				fmt.Println("输入入队列数：")
				fmt.Scanln(&val)
				err := cirleQueue.Push(val)
				if err != nil {
					fmt.Println(err)
				}
			case "get2":
				fmt.Println("获取环形队列数据")
				val, err := cirleQueue.Pop()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Printf("获取到的数据为：%d\n", val)
				}
			case "show2":
				fmt.Println("显示环形队列")
				cirleQueue.List()
			case "exit":
				fmt.Println("退出")
				os.Exit(0)
			default:
				fmt.Println("输入错误，请重新输入")
		}
	}
}