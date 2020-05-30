//AUTHOR: Peter Kilonzo Jr
//CS457 PRogramming Assignment 3
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

//creates a database methods explained more in documentation pdf
func createDB(argss []string) bool{

    err := os.Mkdir(argss[2], 0700)

    if err == nil {
        return true
    } else {
        return false
    } 
    
    
}

//creates a csv(table)
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

//deletes a database
func deleteDB(nomme string) bool {

    if _, err := os.Stat(nomme); os.IsNotExist(err) {

        return false

    } else {
        os.RemoveAll(nomme)
        return true
    }

}

//delete a table(csv)
func deleteTBL(dbname string, nomme string) bool {
    deltbl := dbname + "/" + nomme + ".csv"
    err := os.Remove(deltbl)

    if err == nil {
        return true
    } else {
        return false
    }
}

//this function just checks whether a database exist if it does it returns true and the menu function
//sets it current database to dbname 
func useDB(dbname string) bool {
    if _, err := os.Stat(dbname); os.IsNotExist(err) {
        return false
    } else {
        return true
    }
}


//This is the select statement. This statement is highly optimized to work for the example it considers
//where, innerjoin, and leftouterjoin. This function will determing the correct output depending on 
//what the condition is
func getTBL(dbname string, args []string) bool {
    filename := dbname + "/" + args[3] + ".csv" //filename for first table
    csvfile, err := os.Open(filename)
    if err != nil { // if there is an error opening the file then return false menu function
        return false// states that there was a problem oprning it
    }
    reader := csv.NewReader(csvfile)

    file1record, err1 := reader.ReadAll() //this reads all data from the csv into a 2D array
    if err1 != nil { // if there is an error reading the output for whatever reason display error
        log.Fatal(err1)
    }
    csvfile.Close()

    var filename2 string
    if args[5] == "inner" || args[5] == "outer" { // These statements check for whether the 
        filename2 = dbname + "/" + args[7] + ".csv"//select command has specific join commands
    } else if args[5] == "left" || args[5] == "right" {//such as inner join and by knowing that
        filename2 = dbname + "/" + args[8] + ".csv"// we get the correct string for the second
    } else {                                        //filename
        filename2 = dbname + "/" + args[5] + ".csv"
    }
    csvfile2, err2 := os.Open(filename2)
    if err2 != nil {
        return false
    }
    reader = csv.NewReader(csvfile2)
    file2record, err3 := reader.ReadAll() //read all data into another 2D array
    if err3 != nil {
        log.Fatal(err3)
    }
    /*
    if len(file2record) > 1 {
        fmt.Println("")
    } 
    if len(file1record) > 1 {
        fmt.Println("")
    } 
    */
    if args[7] == "where" { //if using a where condition the selection statement evalues this
        name1 := strings.Split(args[8], ".")[1] //sidentifies the column identifier
        name2 := strings.Split(args[10], ".")[1]//i.e. this gets "id" from "E.id"

        x, y:=-1, -1
        for i := 0; i < len(file1record[0]); i++ {
            if name1 == strings.Split(file1record[0][i], " ")[0] {
                x = i 
            }
        }
        for i := 0; i < len(file2record[0]); i++ {
            if name2 == strings.Split(file2record[0][i], " ")[0] {
                y = i 
            }
        }

        if x == -1 || y == -1 {
            fmt.Printf("Error one of the requested fields does not exist in one of the tables \n")
            return false
        }

        title := false
        for i := 0; i < len(file1record); i++ {
            for j := 0; j < len(file2record); j++ {
                if file1record[i][x] == file2record[j][y] {
                    if !title {
                        fmt.Printf("%s|%s|%s|%s\n", file1record[0][0],file1record[0][1],
                                file2record[0][0], file2record[0][1])
                        title = true
                    }
                    fmt.Printf("%s|%s|%s|%s\n", file1record[i][0],file1record[i][1],
                                file2record[j][0], file2record[j][1])
                }
            }
        }
    } else if args[5] + args[6] == "innerjoin" { //if inner join is the condition
        mat := innerJoin(file1record, file2record, args[10], args[12])//calls innerJoin function which returns the innerjoin table (2d array)
        for i := 0; i < len(mat); i++ {//this is to print the data from our innerjoin table
            for j := 0; j < len(mat[i]); j++ {
                fmt.Printf("%s", mat[i][j])
                if j < len(mat[i])-1 { //makes sure we dont print a "|" at the end
                    fmt.Printf("|")
                }
            }
            fmt.Printf("\n")
        }
    } else if args[5] + args[6] == "leftouter" { //if leftouter this both innerjoin and leftouter
        mat := innerJoin(file1record, file2record, args[11], args[13])//leftouterjoin = innerjoin + left table values
        mat = leftOuterJoin(file1record, mat)//Use our innerjoin as the parameter and also result
        for i := 0; i < len(mat); i++ {//similar print as above
            for j := 0; j < len(mat[i]); j++ {
                fmt.Printf("%s", mat[i][j])
                if j < len(mat[i])-1 {
                    fmt.Printf("|")
                }
            }
            fmt.Printf("\n")
        }
    }
    return true
}


//this function is for Appending rows this is alter from project 1
func addCol(colName string, tblname string, dbname string) bool {
    if dbname == "none" {
        log.Fatalf("!Failed to alter table because no database was specified.")
    }
    filename := dbname + "/" + tblname + ".csv"
    f, err := os.Open(filename)
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
    file.Seek(0, 0)
    line = append(line, colName)
    fmt.Println(line)
    writer := csv.NewWriter(file)
    writer.Write(line)
    writer.Flush()
    return true
}

//Inner join works by creating a table where both give tables have equal values
//This function does that by first getting the comparison column and then
//creating a new table of combined values that are the same in both 
func innerJoin(arr [][]string, arr1 [][]string, comp string, comp1 string) [][]string {
    res:= [][]string{}                  //return 2D array
    name1 := strings.Split(comp, ".")[1]// Split springs by '.' to get column name
    name2 := strings.Split(comp1, ".")[1]

    x, y:=-1, -1 // these variable get the 2nd index of the columns specified
    for i := 0; i < len(arr[0]); i++ {
        if name1 == strings.Split(arr[0][i], " ")[0] {
            x = i 
        }
    }
    for i := 0; i < len(arr1[0]); i++ {
        if name2 == strings.Split(arr1[0][i], " ")[0] {
            y = i 
        }
    }

    if x == -1 || y == -1 { // check to make sure thise columns were found
        fmt.Printf("either column %s or %s was not found\n", name1, name2)
        return res
    }

    title := false //I use this variable to print the row of columns names
    for i := 0; i < len(arr); i++ {//iterate through both arrays completely
        for j := 0; j < len(arr1); j++ {
            if arr[i][x] == arr1[j][y] {//if we have a match in values add them to array
                if !title { //if we have not printed the title, print them
                    colNames := []string{arr[0][0], arr[0][1], arr1[0][0], arr1[0][1]}
                    res = append(res, colNames)
                    title = true
                }
                //we create an array of data from both tables and add it to our result array
                vals := []string{arr[i][0], arr[i][1], arr1[j][0], arr1[j][1]} 
                res = append(res, vals)
            }
        }
    }

    return res
} 

//This function preforms part of the left outer join. We know the left outer join is
//a combination of the inner join of the two tables as well as the data in the left table
//therefore for left outer join we just run innerJoin then this function which add all the
//parts of the left array not in innerjoin to the innerjoni matrix
func leftOuterJoin(arr [][]string, arr1[][]string) [][]string{// parameters are the left table
                                                            //and the innerjoin table
    res := arr1
    var valm map[string]bool //map to keep track of all values from the left already in the
                            //matrix
    valm = make(map[string]bool)
    for i := 1; i < len(arr1); i++ { //gets all IDs in the innerJoin table
        valm[arr1[i][0]] = true
    }

    for i := 1; i < len(arr); i++ {//this check the left table if the value is aleady in our
                                    //matrix we skip it if not we add it to that matrix
        if valm[arr[i][0]] == true {
            continue
        } else {
            vals := []string{arr[i][0], arr[i][1], "", ""}
            res = append(res, vals)
        }
    }

    return res
}

//This function is used for insert and it makes sure to add the data to the next row in the
//CSV
func addData(args []string, dbname string) bool{
    filename := dbname + "/" + args[2] + ".csv" //filename to open
    f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644) //open the
                                            //file with write and append permissions
    if err != nil {
        fmt.Printf("error opening file")
        return false
    }
    vals := args[4:] // takes the 4th argument of the command to the end and makes an array
    w := csv.NewWriter(f) 
    w.Write(vals) //writes that previous array to the file
    w.Flush() //closes file

    return true
}

//This function takes the unformatted string from input and breaks it into
//a string array where each element was a word in the unformatted string
//it then returns that array of strings. Essentially it returns an array of the arguments
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

//"menu" function to control the flow of the program
func menu() {
    reader := bufio.NewReader(os.Stdin) //reads information from terminal input
    var name string //this string holds the unformated terminal input
    currDB := "none" // variable to indicate current DB set current database to none

    cond := false // break condition for the for loop
    for cond != true {

        char, _:= reader.ReadByte() //this read the first byte of input to check whether it is
        if char == '.' {    //is a period then it breaks because a period -> .exit
            cond = true
            break
        } 

        err := reader.UnreadByte() //Go will not read the input correctly if i do not unread
        if err != nil { //the byte that i read in the previous step, this unreads it
            fmt.Println(err)
        }
        
        name, _ = reader.ReadString(';') //read each command until ; is reached
        reader.Reset(os.Stdin)// this flushes the buffer otherwise subsequent commands do not work
        name = strings.TrimRight(name, ";") //remove the ; from the string of commands

        argss := getArgs(name) //get the commands from the string
        


        switch prod := strings.ToUpper(argss[0]); prod {
        case "CREATE": // case statements to do appropriate action
        
            var success bool
            if strings.ToUpper(argss[1]) == "DATABASE" { // two different functions
                success = createDB(argss)               // for create database and 
                                                        // create table
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
        case "USE": // simply sets current database to argument passed if it exist
            success := useDB(argss[1])
            if !success {
                fmt.Println("!Failed to access database", argss[1], "because it does not exist.")
            } else {
                fmt.Println("Using database", argss[1])
                currDB = argss[1]
            }

        case "SELECT": //select case to return information from tables
            success := getTBL(currDB,argss)
            if !success {
                fmt.Println("!Failed to query table", argss[3], "because it does not exist.")
            } 
        case "ALTER": //alter statement to add or subtract columns
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
            } 
            if !success {
                fmt.Println("!Failed to alter table", argss[2], "because it does not exist.")
            } else {
                fmt.Println("Altering table", argss[2])
            }

        case "INSERT": //insert statement to insert data into table

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
            fmt.Printf("\n")
            break
        }
    }
    
}

func main() {
    
    menu()
    
}