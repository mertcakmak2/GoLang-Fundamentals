package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	// out, err := exec.Command("cmd", "/c", "dir").Output()
	// out, err := exec.Command("cmd", "/c", "mkdir test-file").Output()
	// out, err := exec.Command("cmd", "/c", "npm --version").Output()
	// out, err := exec.Command("cmd", "/c", "docker version").Output()

	// cmd := exec.Command("cmd", "/c", "dir")
	// cmd.Dir = "C:/Users/P1608/Desktop/MERT/Projects/GoProjects/GoLang-Source"
	// out, err := cmd.Output()

	cmd := exec.Command("cmd", "/c", "git add * && git status && git commit -m 'cmd_execute' && git push origin master")
	cmd.Dir = "C:/Users/P1608/Desktop/MERT/Projects/GoProjects/GoLang-Source"
	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Output: ", string(out))

}
