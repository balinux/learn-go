package main

import (
	"fmt"
	"os/exec"
)

func main() {
	output1, _ := exec.Command("ls", "-la").Output()
	fmt.Printf("-ls\n%s\n", string(output1))
}
