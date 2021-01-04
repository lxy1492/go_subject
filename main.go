package main

import "fmt"
import "go_subject/Python"

func RunPython()  {
	err := Python.Main_python([]string{"abc"})
	if err!=nil{
		fmt.Println("执行错误")
	}
}

func Run()  {
	RunPython()
}

func main()  {
	fmt.Println("start Go ...")
	Run()
	fmt.Println("end!")
}
