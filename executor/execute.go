package executor

import (
	"fmt"
	"os"
	"os/exec"
)

func Execute(commandString string) {
		execCmd := exec.Command("sh", "-c", commandString)
		result, err := execCmd.Output()
		if err != nil {
			fmt.Println("Error occured")
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(result))
	}
