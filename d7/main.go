//DIRECTORY NAMES ARE NOT UNIQUE!!!!!! THATS WHY THIS ISN"T WORKING
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

type contents struct {
    dir []string
    files int
}


func readData() []string {
    //Open file
    //readfile,err := os.Open("data")
    readfile,err := os.Open("test2")
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


var directories map[string]contents 
var data = readData() 

func getFolder(s string) string {
    a := strings.Split(s, " ") 
    return a[len(a)-1]
}
func getDirectory(s string) string {
    a := strings.Split(s, " ") 
    return a[len(a)-1]
}
func fileSize(s string) int {
    a := strings.Split(s, " ") 
    num, _ := strconv.Atoi(a[0])
    return num 
}


func main() {
    
    directories = make(map[string]contents) //Initialize the map

    for i := 0; i <= len(data)-1; i++ {
        if data[i] == "$ ls" {
            f := getFolder(data[i-1])   //Grabs the parent folder
            //fmt.Println(f)
            //Need to find child folders and directories and assign them to childs
            var file int
            var dir []string
            j := i+1
            for string(data[j][0]) != "$" {
                if string(data[j][0:3]) == "dir" {
                    dir = append(dir, getDirectory(string(data[j])))
                    //fmt.Println(dir)
                } else {
                    file += fileSize(data[j])
                    //fmt.Println(file)
                }
                if j < len(data)-1 {
                    j++
                } else {
                    break
                }
                }
            childs := contents{dir, file}
            directories[f] = childs
            }
        }
    //fmt.Println(directories)

    total := make(map[string]int)
    for folder, file := range directories {
        //If file.folders isn't empty then we want to look at that map and grab the totals            
        t := file.files 
        if file.dir != nil {
            for _, d := range file.dir {
                t += directories[d].files
            }
        }
        total[folder] = t
    }

    fmt.Println(total)

    totalSize := 0
    for _, size := range total {
        if size <= 100000 {
            totalSize += size
        }
    }
    fmt.Println(totalSize)

}


