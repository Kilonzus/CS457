package main

import (
	"fmt"
	//"strings"
	//"bufio"
	//"os"
	"encoding/csv"
	"log"
	"os"
)

func main() {
	csvfile, err1 := os.Open("test.csv")
    if err1 != nil {
        log.Fatal(err1)
    }
	r := csv.NewReader(csvfile)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(records); i++ {
		for j := 0; j < len(records[i]); j++ {
			fmt.Printf("%s bb \t", records[i][j])
		}
		fmt.Printf("\n")
	}

	csvfile, err1 = os.Open("test1.csv")
    if err1 != nil {
        log.Fatal(err1)
    }
	r = csv.NewReader(csvfile)
	records1, err3 := r.ReadAll()
	if err3 != nil {
		log.Fatal(err3)
	}
	for i := 0; i < len(records1); i++ {
		for j := 0; j < len(records1[i]); j++ {
			fmt.Printf("%s bb \t", records1[i][j])
		}
		fmt.Printf("\n")
	}

	
	/*var cmds string
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