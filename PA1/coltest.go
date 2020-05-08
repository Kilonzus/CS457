package main

import (
	"fmt"
	//"strings"
	"bufio"
	"os"
)

func main() {
	//var cmds string
	//test := "test string;to see whether;this idea;would work;i think it will;"
	reader := bufio.NewReader(os.Stdin)
	cmds, err := reader.ReadString(';')

	if err == nil {
		
		fmt.Printf("%T \n %s \n", cmds, cmds)
	}
	
	//testers := strings.Split(test, ";")
	/*for i:=0; i < len(testers)-1; i++ {
		fmt.Printf("%s\n", testers[i])
	}
	*/
}