package main

import (
    "encoding/csv"
    "fmt"
    "bufio"
    "os"
    "strings"
    "log"
)

func create(dbname string, argss []string, param [] string) bool{

    if argss[1] == "DATABASE" {
        err := os.Mkdir(argss[2], 0700)

        if err == nil {
            return true
        } else {
            return false
        } 
    } else if argss[1] == "TABLE" {
        if dbname == "none" {
            log.Fatalf("!Failed to make table because no directory was specified.")
        }
        filename := dbname + "/" + argss[2] + ".csv"
        csvfile, err := os.Create(filename)
        csvfile.Close()

        csvwriter := csv.NewWriter(csvfile)
        for i := 0; i < len(param); i++ {
            fmt.Println(param[i])
        }

        csvwriter.Flush()

        if err == nil {
            return true
        } else {
            return false
        }
    
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

func useDB(dbname string) bool {
    if _, err := os.Stat(dbname); os.IsNotExist(err) {
        return false;
    } else {
        return true;
    }
}

//to-do:
//split create into two functions one for db one for tables
//and create format string function.

func menu() {
    reader := bufio.NewReader(os.Stdin)
    var name string
    var currDB string
    currDB = "none"
    //fmt.Println(len(argss))
    cond := false
    for cond != true {
        name, _ = reader.ReadString('\n')
        name = strings.TrimRight(name, "\n")
        argss := strings.Split(name, " ")
        param := strings.Split(name, " (")
        param[len(param) - 1] = strings.TrimRight(param[len(param) - 1], ")")
        //fmt.Println(param[1])
        switch prod := argss[0]; prod {
        case "CREATE":
        
            success := create(currDB,argss, param)
            if success == false {
                fmt.Println("!Failed to create", argss[1], argss[2], "because it already exists.")
            } else {
                fmt.Println("Creating ", argss[1], argss[2])
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
        case "USE":
            success := useDB(argss[1])
            if !success {
                fmt.Println("!Failed to access database", argss[1], "because it does not exist.")
            } else {
                fmt.Println("Using database", argss[1])
                currDB = argss[1]
            }
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
    
    
}