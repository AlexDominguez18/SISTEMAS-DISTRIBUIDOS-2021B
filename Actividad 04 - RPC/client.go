package main

import (
	"fmt"
	"net/rpc"
	"bufio"
	"os"
	"./util"
)

const (
	SUBJECT_GRADES = 1
	STUDENT_GRADES = 2
	GENERAL_AVERAGE = 3
	SUBJECT_AVERAGE = 4
	EXIT = 5
)

type Args struct{
	Subject string 
	Student string
	Grade float64
}

func showMenu() {
	util.ClearScreen()
	fmt.Println("MENU:")
	fmt.Println("-----")
	fmt.Println("1. Add subject grades.")
	fmt.Println("2. Show student's average.")
	fmt.Println("3. Show general average.")
	fmt.Println("4. Show subject's average.")
	fmt.Println("5. Exit.")
	fmt.Print("\nOption: ")
}

func client() {
	var option int
	var studentName string
	var subjectName string
	var grade float64
	scanner := bufio.NewScanner(os.Stdin)

	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("err")
		return
	}

	for {
		showMenu()
		fmt.Scan(&option)
		util.ClearScreen()
		switch option {
		case SUBJECT_GRADES:
			var reply string
			fmt.Print("Student's name: ")
			scanner.Scan()
			studentName = scanner.Text()
			fmt.Print("Subject's name: ")
			scanner.Scan()
			subjectName = scanner.Text()
			fmt.Print("Grade: ")
			fmt.Scan(&grade)
			args := Args{subjectName, studentName, grade}
			err = c.Call("Server.AddSubjectGrades", args, &reply)
			if err != nil {
				fmt.Println("Err: ", err)
			} else {
				fmt.Println(reply)
			}
		case STUDENT_GRADES:
			var reply float64
			fmt.Print("Student's name: ")
			scanner.Scan()
			studentName = scanner.Text()
			err = c.Call("Server.ShowStudentAverage", studentName, &reply)
			if err != nil {
				fmt.Print("Err: ", err)
			} else {
				fmt.Println(reply)
				if reply > 0 {
					fmt.Printf("%s's average = %f\n", studentName, reply)
				} else {
					fmt.Println(studentName + " is not registered!")
				}
			}
		case GENERAL_AVERAGE:
			var reply float64
			err = c.Call("Server.ShowGeneralAverage", 0, &reply)
			if err != nil {
				fmt.Println("Err: ", err)
			} else {	
				fmt.Println("General average = ", reply)
			}
		case SUBJECT_AVERAGE:
			var reply float64
			fmt.Print("Subject's name: ")
			scanner.Scan()
			subjectName = scanner.Text()
			err = c.Call("Server.ShowSubjectAverage", subjectName, &reply)
			if err != nil {
				fmt.Println("Err: ", err)
			} else {
				fmt.Printf("%s's average: %.2f\n", subjectName, reply)
			}
		case EXIT:
			return
		}
		util.Pause()
	}
}

func main() {
	client()
}