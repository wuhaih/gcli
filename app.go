package main

import (
	"fmt"
	"os/exec"
	"bufio"
	"os"
	"log"
)

/**
* 
* @author willian
* @created 2017-01-26 14:07
* @email 18702515157@163.com  
**/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	log.Println("the author name of the github.com")
	scanner.Scan()
	author := scanner.Text()
	log.Println("the prodject name of the github.com")
	scanner.Scan()
	project := scanner.Text()
	location, _ := os.Getwd()
	work_space := location+"/src/github.com/"+author
	err := os.MkdirAll(work_space, os.ModePerm)
	if err != nil {
		log.Fatal("创建目录出错~！！！！！")
	}
	project_url:= "git@github.com:"+author+"/"+project+".git"
	cmd := exec.Command("git", "clone",project_url)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(out))
	mv := exec.Command("mv", location+"/common", work_space)
	mv_out, err := mv.CombinedOutput()
	if err != nil {
		log.Fatalln(err)
	}else {
		fmt.Println(string(mv_out))
	}

}
