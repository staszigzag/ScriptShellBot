package exec

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Run(command string) (answer string, err error) {
	parts := strings.Split(command, " ")

	//	The first part is the command, the rest are the args:
	head := parts[0]
	args := parts[1:len(parts)]

	//	Format the command
	cmd := exec.Command(head, args...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	//cmd.Stderr = &stderr
	cmd.Stderr = os.Stderr
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr

	//	Run the command
	cmd.Run()
	fmt.Println(stderr.String())
	return out.String(), nil
}
