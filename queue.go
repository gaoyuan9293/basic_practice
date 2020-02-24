package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	MaxSize int
	Array   [5]int
	Front   int
	Rear    int
}

func (this *Queue) addQueue(val int) error {
	if this.Rear == this.MaxSize-1 {
		err := errors.New("queue full")
		fmt.Println(err)
		return err
	}
	this.Rear++
	this.Array[this.Rear] = val
	return nil
}

func (this *Queue) showQueue() {
	if this.Rear == this.Front {
		err := errors.New("queue empty")
		fmt.Println(err)
	}
	for i := this.Front + 1; i <= this.Rear; i++ {
		fmt.Println(this.Array[i])
	}

}

func (this *Queue) getQueue() (int, error) {
	if this.Front == this.Rear {
		err := errors.New("quque empty")
		fmt.Println(err)
		return -1, err
	}
	this.Front++
	val := this.Array[this.Front]
	return val, nil
}

func main() {
	queue := &Queue{
		MaxSize: 5,
		Front:   -1,
		Rear:    -1,
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
			err := queue.addQueue(val)
			if err != nil {
				fmt.Println("add error is:", err)
			} else {
				fmt.Println("add element to queue success")
			}
		case "show":
			queue.showQueue()
		case "get":
			i, err := queue.getQueue()
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
