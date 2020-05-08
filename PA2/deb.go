//AUTHOR: Peter Kilonzo Jr
//CS457 PRogramming Assignment 1
//5-8-2020

package main

import (
    "encoding/csv"
    "fmt"
    "bufio"
    "os"
    "strings"
    "log"
    "io"
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


func createTBL(dbname string, argss[]string) bool{
    if dbname == "none" {
        log.Fatalf("!Failed to make table because no directory was specified.")
    }
    
    filename := dbname + "/" + argss[2] + ".csv"
    if _, er3 := os.Stat(filename); er3 == nil {
        return false

    }
    param := argss[3:]

    file, err := os.Create(filename)
    checkError("Cannot create file", err)
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    er := writer.Write(param)
    checkError("Cannot write to file", er)

    if err == nil {
        return true
    } else {
        return false
    }
}


func deleteDB(nomme string) bool {

    if _, err := os.Stat(nomme); os.IsNotExist(err) {
    // path/to/whatever does not exist
        //fmt.Println("shit not here")
        return false

    } else {
        //fmt.Println("Guess it do")
        os.RemoveAll(nomme)
        return true
    }

}

func deleteTBL(dbname string, nomme string) bool {
    deltbl := dbname + "/" + nomme + ".csv"
    err := os.Remove(deltbl)

    if err == nil {
        return true
    } else {
        return false
    }
}

func useDB(dbname string) bool {
    if _, err := os.Stat(dbname); os.IsNotExist(err) {
        return false
    } else {
        return true
    }
}

func getTBL(dbname string, nomme string) bool {
    filename := dbname + "/" + nomme + ".csv"
    csvfile, err := os.Open(filename)
    if err != nil {
        return false
    }
    reader := csv.NewReader(csvfile)
    for {
        record, err1 := reader.Read()
        if err1 == io.EOF {
            break
        }
        if err1 != nil {
            log.Fatal(err)
        }
        for i := 0; i < len(record); i++ {
            fmt.Print(record[i])
            if i < len(record)-1 {
                fmt.Print(" | ")
            }
        }
        fmt.Println("")
    }

    return true
}

//new bug now appending columns will eat existing values
//almost as if there are only a set amount of values in the
//csvfile and it can not be expanded
func addCol(colName string, tblname string, dbname string) bool {
    if dbname == "none" {
        log.Fatalf("!Failed to alter table because no database was specified.")
        //return false
    }
    filename := dbname + "/" + tblname + ".csv"
    //fmt.Println(filename)
    f, err := os.Open(filename)
    //os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println("Error opening ", filename)
        return false
    }
    reader := csv.NewReader(bufio.NewReader(f))
    line, err1 := reader.Read()
    f.Close()
    if err1 == io.EOF {
            return false
        }
    if err1 != nil {
        fmt.Println("error happened err1 != nill")
        log.Fatal(err)
        }
    file, err2 := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
    if err2 != nil {
        fmt.Println("error happened err2 != nill")
        log.Fatal(err)
        }
    fmt.Println(line)
    //file.Truncate(0)
    file.Seek(0, 0)
    //col := []string{colName}
    line = append(line, colName)
    fmt.Println(line)
    writer := csv.NewWriter(file)
    writer.Write(line)
    writer.Flush()
    return true
}

func addData(args []string, dbname string) bool{
    filename := dbname + "/" + args[2] + ".csv"
    f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
    if err != nil {
        fmt.Printf("error opening file")
        return false
    }
    vals := args[4:]
    w := csv.NewWriter(f)
    w.Write(vals)
    w.Flush()
    //fmt.Println(vals)

    return true
}


func getArgs(line string) []string {
    var args []string
    if strings.Contains(line, "(") {
        linef := strings.TrimSuffix(line, ")")
        args = strings.Split(linef, "(")
        if len(args) > 2 {
            for i := 2; i < len(args); i++ {
                args[1] += ("("+args[i])
            }
            args = args[:2]
        }
        vals1 := strings.Split(args[0], " ")
        vals2 := strings.Split(args[1], ",")
        args = append(vals1, vals2...)
    } else {
        args= strings.Split(line, " ")
    }
    
    
    
    
    for i := 0; i < len(args); i++ {
        args[i] = strings.TrimSpace(args[i])
        args[i] = strings.TrimSuffix(args[i], "'")
        args[i] = strings.TrimPrefix(args[i], "'")
    }
    return args
}

func menu() {
    reader := bufio.NewReader(os.Stdin)
    var name string
    currDB := "none" // set current database to none

    cond := false
    for cond != true {
        name, _ = reader.ReadString(';') //read deach command until ; is reached

        name = strings.TrimRight(name, ";") //remove the ; from the string of commands

        argss := getArgs(name) //get the commands from the string
        for i := 0; i < len(argss); i++ {
            fmt.Printf("argss[%d] consist of %s\n", i, argss[i])
        }
        switch prod := strings.ToUpper(argss[0]); prod {
        case "CREATE":
        
            var success bool
            if strings.ToUpper(argss[1]) == "DATABASE" {
                success = createDB(argss)
            } else if strings.ToUpper(argss[1]) == "TABLE" {
                success = createTBL(currDB,argss)
            }
            
            if success == false {
                fmt.Println("!Failed to create", argss[1], argss[2], "because it already exists.")
            } else {
                fmt.Println("Creating", argss[1], argss[2])
            }
            break

        case "DROP":
            var success bool
            if strings.ToUpper(argss[1]) == "DATABASE" {
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

        case "SELECT":
            //var success bool
            success := getTBL(currDB,argss[3])
            if !success {
                fmt.Println("!Failed to query table", argss[3], "because it does not exist.")
            } 
        case "ALTER":
            var success bool
            if strings.ToUpper(argss[3]) == "ADD" {
                var colName string
                for i := 4; i < len(argss); i++ {
                    if i > 4 {
                        colName += " "
                    }
                    colName += argss[i]
                }
                success = addCol(colName, argss[2], currDB)
            } //else if strings.ToUpper(argss[2]) == "DROP" {

            //}
            if !success {
                fmt.Println("!Failed to alter table", argss[2], "because it does not exist.")
            } else {
                fmt.Println("Altering table", argss[2])
            }

        case "INSERT": 

            success := addData(argss, currDB)
            if !success {
                fmt.Printf("Failed to insert data into %s\n", argss[2])
            } else {
                fmt.Printf("1 new record added\n")
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