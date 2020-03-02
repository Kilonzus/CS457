package main

import (
    "encoding/csv"
    "fmt"
    "bufio"
    "os"
    "strings"
    "log"
)

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}

func createDB(argss []string) bool{

    err := os.Mkdir(argss[2], 0700)

    if err == nil {
        return true
    } else {
        return false
    } 
    
    
}

func createTBL(dbname string, argss[]string, param[] string) bool{
    if dbname == "none" {
        log.Fatalf("!Failed to make table because no directory was specified.")
    }

    for i := 0; i < len(param); i++ {
        fmt.Println(param[i])
    }

    //var data = [][]string{{"Line1", "Hello Readers of"}, {"Line2", "golangcode.com"}}
    
    filename := dbname + "/" + argss[2] + ".csv"

    file, err := os.Create(filename)
    checkError("Cannot create file", err)
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    //for _, value := range data {
    //    err := writer.Write(value)
    //    checkError("Cannot write to file", err)
    //}
        er := writer.Write(param)
        checkError("Cannot write to file", er)

    if err == nil {
        return true
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

func deleteTBL(dbname string, nomme string) bool {
    deltbl := dbname + "/" + nomme + ".csv"
    err := os.Remove(deltbl)

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

func fmtstr(val string) []string {
    values := strings.Split(val, " (")
    params := strings.Split(values[1], ",")
    params[len(params)-1] = strings.TrimSuffix(params[len(params)-1], ")")
    fmt.Println("Parameters are", params, "length of ", len(params))
    for i := 0; i < len(params); i++ {
        params[i] = strings.TrimLeft(params[i], " ")
    }

    return params
}

func menu() {
    reader := bufio.NewReader(os.Stdin)
    var name string
    var currDB string
    currDB = "none"

    cond := false
    for cond != true {
        name, _ = reader.ReadString('\n')
        name = strings.TrimRight(name, "\n")
        argss := strings.Split(name, " ")
        switch prod := argss[0]; prod {
        case "CREATE":
        
            var success bool
            if argss[1] == "DATABASE" {
                success = createDB(argss)
            } else if argss[1] == "TABLE" {
                param := fmtstr(name)
                success = createTBL(currDB,argss,param)
            }
            
            if success == false {
                fmt.Println("!Failed to create", argss[1], argss[2], "because it already exists.")
            } else {
                fmt.Println("Creating ", argss[1], argss[2])
            }
            break

        case "DROP":
            var success bool
            if argss[1] == "DATABASE" {
                success = deleteDB(argss[2])
            } else if argss[1] == "TABLE" {
                success = deleteTBL(currDB,argss[2])
            }
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