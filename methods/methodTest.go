package main

import (
	"fmt"
)

import "../intf"

type QE struct {
	name string
	priLan string
	dept string
}

func (qe QE) progHello() {
	fmt.Println("My name is " + qe.name)
	fmt.Println("My favarite language is " + qe.priLan)
	fmt.Println()
	switch qe.priLan {
	case "Java":
		fmt.Println("public class Demo{")
		fmt.Println("\tpublic static void main(String[] args){")
		fmt.Println("\t\tSystem.out.println(\"Hello World\");")
		fmt.Println("\t}")
		fmt.Println("}")
	case "C":
		fmt.Println("import <stdio.h>")
		fmt.Println("int main(int argc, char** argv)")
		fmt.Println("{")
		fmt.Println("\tprintf(\"Hello World\\n\");")
		fmt.Println("}")
	case "Python":
		fmt.Println("print(\"Hello world!\")")
	case "Go":
		fmt.Println("import \"fmt\"")
		fmt.Println("")
		fmt.Println("func main(){")
		fmt.Println("\tfmt.Println(\"Hello world!\")")
		fmt.Println("}")
	}
}

func (qe QE) changeDept(newDept string) {
	qe.dept = newDept
}


func main() {
    tristan := QE{
		"Tristan Yan",
		"Go",
		"finance",
	}

	tristan.progHello()
	tristan.changeDept("IDEA")

	fmt.Println("I work for " + tristan.dept)
}