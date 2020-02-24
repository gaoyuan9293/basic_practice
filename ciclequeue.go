package main

import (
	"errors"
	"fmt"
	"os"
)

type CicleQueue struct {
	MaxSize int
	Array   [5]int
	Front   int
	Rear    int
}

//队满
func (this *CicleQueue) IsFull() error {
	if (this.Rear+1)%this.MaxSize == this.Front {
		err := errors.New("quque full")
		return err
	}
	return nil
}

//队空
func (this *CicleQueue) IsEmpty() error {
	if (this.Rear) == this.Front {
		err := errors.New("queue empty")
		return err
	}
	return nil
}

func (this *CicleQueue) Size() int {
	size := (this.Rear + this.MaxSize - this.Front) % this.MaxSize
	return size
}

//入队
func (this *CicleQueue) Push(val int) error {
	err := this.IsFull()
	if err != nil {
		fmt.Println("push failed, err:", err)
		return err
	}
	this.Array[this.Rear] = val
	this.Rear = (this.Rear + 1) % this.MaxSize
	return nil
}

//出队

func (this *CicleQueue) Pop() (int, error) {
	err := this.IsEmpty()
	if err != nil {
		fmt.Println("pop failed, err:", err)
		return -1, err
	}
	val := this.Array[this.Front]
	this.Front = (this.Front + 1) % this.MaxSize
	return val, nil
}

//获取队列全部元素

func (this *CicleQueue) showQueue() {
	size := this.Size()
	if size == 0 {
		fmt.Println("quque empty")
	}

	tmp := this.Front
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d]=%d\n", tmp, this.Array[tmp])
		tmp = (tmp + 1) % this.MaxSize
	}
}

func main() {
	queue := &CicleQueue{
		MaxSize: 5,
		Front:   0,
		Rear:    0,
	}

	var key string
	var val int
	for {
		fmt.Println("1.add 标识添加元素到队列中")
		fmt.Println("2.show 获取队列中的全部元素")
		fmt.Println("3.get 从队列中取出元素")
		fmt.Println("4.exit 退出")

		fmt.Scan(&key)
		//fmt.Println("key:", &key)

		switch key {
		case "add":
			fmt.Scan(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println("add error is:", err)
			} else {
				fmt.Printf("add element [%d] to queue success", val)
			}
		case "show":
			queue.showQueue()
		case "get":
			i, err := queue.Pop()
			if err != nil {
				fmt.Println("get the element failed, err:", err)
			} else {
				fmt.Println("get the elelment from queue sueccess", i)
			}
		case "exit":
			os.Exit(0)
		}
	}
	return
}
