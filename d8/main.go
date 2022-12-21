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
    var upper bool
    for j:=0; j<i; j++  {
        if row[j] >= row[i] {
            upper = false 
            break
        }else {
            upper = true
        }
    }

    var lower bool
    for j:=len(row)-1; j > i; j-- {
        if row[j] >= row[i] {
            lower = false
            break
        } else {
            lower = true
        }
    }
    if upper || lower {
        return true
    }else{
        return false
    }
}

func checkUpDown(row int, col int) bool {
    var left bool
    for i:=0; i<row; i++ {
        if data[i][col] >= data[row][col] {
            left = false
            break
        } else {
            left = true
        }
    }

    var right bool
    for i:=len(data)-1; i>row; i-- {
        if data[i][col] >= data[row][col] {
            right = false
            break
        } else {
            right = true
        }
    }
    if left || right {
        return true
    } else {
        return false
    }
}


func main() {
    readData()

    visibleTrees += len(data[0])*2
    visibleTrees += (len(data)-2)*2

    for row := 1; row <= len(data)-2; row++ {
        for col := 1; col <= len(data[row])-2; col++ {
            fmt.Println(checkLeftRight(col, data[row]), checkUpDown(row, col))
            if checkLeftRight(col, data[row]) == true || checkUpDown(row, col) == true {
                visibleTrees++
            }
        }
    }



//Debugging
    for _, v := range data {
        fmt.Println(v)
    }

    fmt.Println("Total visible Trees = ", visibleTrees)
    //fmt.Println(checkLeftRight(2, data[3]))
    //fmt.Println(checkUpDown(1, 3))
}
