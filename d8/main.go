package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

var data [][]int

func readData() {
    //Open file
    //readfile,err := os.Open("data")
    readfile,err := os.Open("test")
    if err != nil {
        fmt.Println(err)
    }

    filescanner := bufio.NewScanner(readfile)

    filescanner.Split(bufio.ScanLines)

    for filescanner.Scan() {
        line := strings.Split(filescanner.Text(), "")
        newline := []int{}
        for _, v := range line {
            j, err := strconv.Atoi(v)
            if err != nil {
                fmt.Println("Error converting string to int!")
            }
            newline = append(newline, j)
        }
        data = append(data, newline)
    }
    readfile.Close()  
}

var visibleTrees int


func checkLeftRight(i int, row []int) bool {
    var vis bool
    for j:=0; j<i; j++  {
        if row[j] >= row[i] {
            vis = false 
        }else {
            vis = true
        }
    }
    for j:=len(row)-1; j > i; j-- {
        if row[j] >= row[i] {
            vis = false
        } else {
            vis = true
        }
    }
    return vis
}

func checkUpDown(col int, row int) bool {
    var vis bool
    return vis
}


func main() {
    readData()

    visibleTrees += len(data[0])*2
    visibleTrees += (len(data)-2)*2



    for _, v := range data {
        fmt.Println(v)
    }

    fmt.Println("Total visible Trees = ", visibleTrees)
    fmt.Println(checkLeftRight(2, data[2]))
}
