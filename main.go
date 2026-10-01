package main

import "fmt"

// Task 表示一条学习任务。
type Task struct {
	// TODO 1：添加 ID（int）、Title（string）、Done（bool）三个字段。
		ID int
		Title string
		Done bool
}

func main() {
	// TODO 2：创建任务，编号为 1，标题为“学习 Go”，尚未完成。
	task := Task{
		ID: 1
		Title: 
	}
	fmt.Printf("任务：%+v\n", task)

	// TODO 3：把任务改为已完成，再打印一次。
}
