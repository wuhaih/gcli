package main

import (
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

	// version 1
	reader := bufio.NewReader(os.Stdin)
	log.Println("please input the project_folder_name:")
	project_folder_name, _ := reader.ReadBytes('\n')
	println(string(project_folder_name[0:len(project_folder_name)-1]))
	log.Println("please input the author name of github.com:")
	author_name, _ := reader.ReadBytes('\n')
	println(string(author_name[0:len(author_name)-1]))


}
