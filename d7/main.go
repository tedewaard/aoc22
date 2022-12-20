//DIRECTORY NAMES ARE NOT UNIQUE!!!!!! 
package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

/*
Grab the folder name and insert into array
Grab the size of files for that folder and update array
If size of files is > 100000 then we remove that item from the array
If it's less then we move on to the next one
If we hit a "cd .." then we add the size to Total and then pop off from stack
    and also add it to the new last folder in array

*/

func readData() []string {
    //Open file
    readfile,err := os.Open("data")
    //readfile,err := os.Open("test2")
    if err != nil {
        fmt.Println(err)
    }

    filescanner := bufio.NewScanner(readfile)

    filescanner.Split(bufio.ScanLines)
    var data []string

    for filescanner.Scan() {
        data = append(data, filescanner.Text())
    }
    readfile.Close()  
    return data
}

type folder struct {
    dir string
    size int
}

var data = readData()
var stack []folder
var total int

func getFolder(s string) string {
    a := strings.Split(s, " ") 
    return a[len(a)-1]
}

func getSize(s string) int {
    a := strings.Split(s, " ") 
    num, _ := strconv.Atoi(a[0])
    return num 
}

func addToAll(n int) {
    for i:=0; i<=len(stack)-1; i++ {
        stack[i].size += n
    }
}



func main() {

    for i:=0; i<=len(data)-1; i++ {
        if string(data[i][:4]) == "$ cd" && data[i] != "$ cd .." {
            record := getFolder(data[i])
            f := folder{record, 0}
            stack = append(stack, f)
        }

        if string(data[i][:3]) == "dir" {
            continue
        }

        if data[i] == "$ cd .." {
            //Add last size in stack to total and second from last item and then pop
            total += stack[len(stack)-1].size
            stack = stack[:len(stack)-1]

        }

        if string(data[i][0]) != "$" && string(data[i][0]) != "d"{
            size := getSize(data[i])
            addToAll(size)
        }


    }

    for _, d := range stack {
        fmt.Println(d)
    }

    fmt.Println(total)


}
