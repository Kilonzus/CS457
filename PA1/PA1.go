package main

import (
"fmt"
"bufio"
"os"
"strings"
)

func createDB(nomme string) bool {
    err := os.Mkdir(nomme, 0700)

    if err == nil {
        return true;
    } else {
        return false
    }
}

func deleteDB(nomme string) bool {
    err := os.RemoveAll(nomme)

    if err == nil {
        return true;
    } else {
        return false
    }
}

func menu() {
    reader := bufio.NewReader(os.Stdin)
    var name string
    
    //fmt.Println(len(argss))
    cond := false
    for cond != true {
        name, _ = reader.ReadString('\n')
        name = strings.TrimRight(name, "\n")
        argss := strings.Split(name, " ")
        switch prod := argss[0]; prod {
        case "CREATE":
        
            success := createDB(argss[2])
            if !success {
                fmt.Println("!Failed to create database", argss[2], "because it already exists.")
            } else {
                fmt.Println("Creating database", argss[2])
            }
            break

        case "DROP":
            success := deleteDB(argss[2])
            if !success {
                fmt.Println("!Failed to delete database", argss[2], "because it does not exist.")
            } else {
                fmt.Println("Deleting database", argss[2])
            }
            break

        case ".EXIT" :
            cond = true
            break

        default:
            fmt.Println("Invalid Selection!")
            break
        }
    }
    

}

func main() {

    menu()
    
    //reading a string

/*
    var teststring string
    teststring = "CREATE"
    res1 := name == (teststring + "\n")
    
    if res1 {
    	fmt.Println("You wish to create a database.")
    }
    fmt.Println("1 is", name)
    fmt.Println("2 is", teststring)
    fmt.Println(res1)
    */
}