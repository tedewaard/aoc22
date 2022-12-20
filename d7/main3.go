//DIRECTORY NAMES ARE NOT UNIQUE!!!!!! 
package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

//I think I want a map of directory as key and a struct that has two arrays:
//One for child directories and one for child files


//So what I really need is a map of folder to total filesize 
//And a map of folder to sub folders
//Which I kind of already have


func readData() []string {
    //Open file
    //readfile,err := os.Open("data")
    readfile,err := os.Open("test")
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

type file struct {
    size int
    dir string
}

var data = readData()
var directory = ""
var total = 0
var sum = 0
var stack []file




func getFolder(s string) string {
    a := strings.Split(s, " ") 
    return a[len(a)-1]
}
func getSize(s string) int {
    a := strings.Split(s, " ") 
    num, _ := strconv.Atoi(a[0])
    return num 
}

func addToPrev(n int) {
    for i:=0; i<=len(stack)-1; i++ {
       stack[i].size += n 
    }
}

/*
Grab the folder name and throw into the stack
Get the count and update the item in the stack
if the item is over size then we remove it
*/

func main() {

    for i:=0; i<=len(data)-1; i++ {
        if string(data[i][0:4]) == "$ cd" && data[i] != "$ cd .." {

            if i != 0 && total <= 100000 {
            fmt.Println(stack)
                n := len(stack)-1
                if n < 0 {
                    break
                }
                stack[n].size = total
                if n > 0 {
                    addToPrev(total)
                }
            fmt.Println(stack)
            } else if i != 0 && total > 100000 {
                stack = stack[:len(stack)-1]
            }

            total = 0

            directory = getFolder(data[i])
            f := file{total, directory}
            stack = append(stack, f)
            fmt.Println("Appended: ", f)

        }

        if string(data[i]) == "$ cd .." {
            n := len(stack)-1
            sum += stack[n].size
            stack = stack[:n]
        }
        if string(data[i][0]) != "$" && string(data[i][0]) != "d"{
            //fmt.Println(getSize(data[i]))
            total += getSize(data[i])
        }

    }



    for _, s := range stack {
        fmt.Println(s)
    }
    
    fmt.Println(sum)

}


