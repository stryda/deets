package shell

import (
	"fmt"
	"os"
)
	//"os/exec"

//	//"os/exec"
//)
//
//func getCommand(command string, args []string) string {
//	cmd := exec.Command("bash", "-c", "ps", "-p", "$$", "-ocomm=")
//	//strings.Split(args, " ")...
//	//args...
//	output, err := cmd.CombinedOutput()
//
//	fmt.Println("Executing --->" + cmd.String())
//
//	if err != nil {
//		log.Fatal("Failed to derive shell: " + err.Error())
//	}
//
//	fmt.Println(string(output))
//
//	return string(output)
//}

func getEnvironment() []string {
	return os.Environ()
}

func printEnvironment() {
	for _, v := range getEnvironment() {
		fmt.Println(v)
	}
}

func getCurrentShell() string {
	// ps -p $$ -oargs=
	// ps -p $$ -ocomm=
	//cmd := exec.Command("ps", "-p", "$$", "ocomm=")
	fmt.Println("--------------------")
	printEnvironment()
	//sess := sh.Echo("ffs")

	//xx, err := sess.Output()

	// sh: count=$(echo "one two three" | wc -w)

	val := os.ExpandEnv("$0")

	os.Setenv("NAME", "gopher")
	os.Setenv("BURROW", "/usr/gopher")

	fmt.Println(os.ExpandEnv("$$ lives in ${BURROW}. "))




	fmt.Println("OUT: " + val + "|")

	return ""

	//return getCommand("ps", []string{})
}

func Print() {
	getCurrentShell()
}