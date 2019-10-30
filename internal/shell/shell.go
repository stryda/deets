package shell

import (
	"fmt"
	"os"
	"os/exec"
)

func getCommand(cmd string, args []string, newShell bool) string {
	if newShell {
		args = append([]string{"-c", cmd}, args[0:]...)
		cmd = "bash"
	}

	out, err := exec.Command(cmd, args...).Output()
	if err != nil {
		return ""
	}

	return string(out)
}

func getEnv() []string {
	return os.Environ()
}

func getEnvVariable(v string) string {
	return os.Getenv(v)
}

func printEnv() {
	for _, v := range getEnv() {
		fmt.Println(v)
	}
}

func getShell() string {
	return getCommand("ps", []string{"-p", "$$", "-ocomm="}, false)
}

func Print() string {
	return getShell()
}
