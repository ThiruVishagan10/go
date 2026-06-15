package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func name() {
	fmt.Print("Enter your name:")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello %s", strings.TrimSpace(name))
}

func dream_project() {
	fmt.Print("Enter your dream project:")
	reader := bufio.NewReader(os.Stdin)
	project, _ := reader.ReadString('\n')
	fmt.Printf("Your dream project is: %s", strings.TrimSpace(project))
}

func Goal() {
	fmt.Print("Enter your goal:")
	reader := bufio.NewReader(os.Stdin)
	goal, _ := reader.ReadString('\n')
	fmt.Printf("Your goal is: %s", strings.TrimSpace(goal))
}

func main() {
	name()
	fmt.Println()
	dream_project()
	fmt.Println()
	Goal()
}
