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
    readfile,err := os.Open("data")
    //readfile,err := os.Open("test")
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


func checkLeftRight(i int, row []int) int {
    left := 0
    for j:=i-1; j>=0; j--  {
        if row[j] >= row[i] {
            left++
            break
        }else {
            left++
        }
    }

    right := 0
    for j:=i+1; j <= len(row)-1; j++ {
        if row[j] >= row[i] {
            right++
            break
        } else {
            right++
        }
    }
    return left * right
}

func checkUpDown(row int, col int) int {
    up := 0
    for i:=row-1; i>=0; i-- {
        if data[i][col] >= data[row][col] {
            up++
            break
        } else {
            up++
        }
    }

    down := 0
    for i:=row+1; i<=len(data)-1; i++ {
        if data[i][col] >= data[row][col] {
            down++
            break
        } else {
            down++
        }
    }
    return up * down
}


func main() {
    readData()
    bestScore:=0
    for row := 0; row <= len(data)-1; row++ {
        for col := 0; col <= len(data[row])-1; col++ {
                //fmt.Println(checkLeftRight(col, data[row]) * checkUpDown(row, col))
                score := checkLeftRight(col, data[row]) * checkUpDown(row, col)
                if score > bestScore {
                    bestScore = score
                }
            }
        }



//Debugging
    for _, v := range data {
        fmt.Println(v)
    }

    //fmt.Println(checkLeftRight(2, data[3]))
    //fmt.Println(checkUpDown(2, 3))
    fmt.Println(bestScore)
   
}
