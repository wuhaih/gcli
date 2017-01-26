package main

import (
	"fmt"
	"os/exec"
)

/**
* 
* @author willian
* @created 2017-01-26 14:07
* @email 18702515157@163.com  
**/
func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	//log.Println("the author name of the github.com")
	//scanner.Scan()
	//fmt.Println(scanner.Text())
	//log.Println("the prodject name of the github.com")
	//scanner.Scan()
	//fmt.Println(scanner.Text())

	//cmd := exec.Command("git", " clone git@github.com:zhangweilun/common.git")
	cmd := exec.Command("git", "init")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
