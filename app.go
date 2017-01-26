package main

import (
	"bufio"
	"os"
	"log"
	"os/exec"
	"fmt"
	"bytes"
	"github.com/zhangweilun/gcli/util"
)

/**
* 
* @author willian
* @created 2017-01-26 14:07
* @email 18702515157@163.com  
**/
func main() {

	// version 1
	reader := bufio.NewReader(os.Stdin)
	//log.Println("please input the project_folder_name:")
	//project_folder_name, _ := reader.ReadBytes('\n')
	log.Println("please input the author name of github.com:")
	author_name_line, _ := reader.ReadBytes('\n')
	author_name := bytes.TrimRight(author_name_line, "\r\n")
	log.Println("please input the project name of github.com:")
	project_name_line, _ := reader.ReadBytes('\n')
	project_name := bytes.TrimRight(project_name_line,"\r\n")
	//git@github.com:zhangweilun/goweb.git
	project_url := "git@github.com:"+util.ByteString(author_name)+"/"+util.ByteString(project_name)+".git"
	fmt.Println(project_url)
	args := "clone "+ project_url
	command := exec.Command("git", args)
	var out bytes.Buffer
	command.Stdout = &out
	err := command.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String()) //in all caps: "SOME INPUT"


}
