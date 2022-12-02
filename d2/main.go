package main

import (
    "os"
    "fmt"
    "bufio"
)


func main() {
   
    //Open file
    readfile,err := os.Open("d2_data")
    if err != nil {
        fmt.Println(err)
    }

    filescanner := bufio.NewScanner(readfile)

    filescanner.Split(bufio.ScanLines)
    var data []string

    for filescanner.Scan() {
        data = append(data, filescanner.Text())
   //     fmt.Println(filescanner.Text())
    }

    readfile.Close()
    var points []int

    for i :=0; i<len(data); i++ {
//        fmt.Println(string(data[i][0]))
/*
Win = 6
Tie = 3
Loose = 0
Rock = 1
Paper = 2
Scissors = 3
x = Loose
y = tie
z = win
*/
        sum := 0
        switch string(data[i][0]) {
        case "A":
            switch string(data[i][2]) {
            case "X":
                sum = 3
            case "Y":
                sum = 4
            case "Z":
                sum = 8
            }
        case "B":
            switch string(data[i][2]) {
            case "X":
                sum = 1
            case "Y":
                sum = 5
            case "Z":
                sum = 9 
            }
        case "C":
            switch string(data[i][2]) {
            case "X":
                sum = 2
            case "Y":
                sum = 6
            case "Z":
                sum = 7
            }
        }
        points = append(points, sum)
    }

    var totalpoints int

    for _, j := range points {
        //fmt.Println(j)
        totalpoints += j
    }

    fmt.Println(totalpoints)

}
