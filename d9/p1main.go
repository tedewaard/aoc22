package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "strconv"
)

type motion struct {
    dir string
    steps int
}
type position struct {
    x int
    y int
}
var data []motion

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
        line := strings.Split(filescanner.Text(), " ")
        j, err := strconv.Atoi(line[1])
        if err != nil {
            fmt.Println("Error converting string to int!")
        }
        m := motion{line[0], j}
        data  = append(data, m)
    }
    readfile.Close()  
}


func abs(x int) int {
    if x<0 {
        return -x
    }
    return x
}

/*
First we move the head and grab it's position
Next we move the tail accordingly and store it's position in an array
We only move the tail if it's no longer next to head or covered by head
If head goes up or down we move diagnally
Then we get size of array
*/
func checkTouching(head position, tail position) bool {
    //Check toucing on x or y access
    if head.y == tail.y && abs((head.x - tail.x)) <= 1 {
        return true
    } 
    if head.x == tail.x && abs((head.y - tail.y)) <= 1 {
        return true
    } 

    //Check diagnally touching
    if abs((head.y - tail.y)) == 1 && abs((head.x - tail.x)) == 1 {
        return true
    }

    return false
}

func movePosition(point position, direction string ) position {
    
        switch direction {
        case "U": 
            point.y++ 
        case "D":
            point.y--
        case "L":
            point.x--
        case "R":
            point.x++
        }

        return point
}



func main() {
    readData()
    positions := make(map[position]int)

    head := position{0, 0}
    tail := position{0, 0}

    for _, i := range data {
        j := 1
        for j<=i.steps {
            oldHead := head
            head = movePosition(head, i.dir)
            if checkTouching(head, tail) {
                //fmt.Println(head, " is touching ", tail)
            } else {
                //fmt.Println(head, " is not touching ", tail)
                positions[tail] = 1                //fmt.Println(tail, " is now moving to ", oldHead)
                tail = oldHead 
            }
            //fmt.Println(oldHead, " is now at ", head)
            j++
        }
    }
    positions[tail] = 1  


    fmt.Println(data)
    fmt.Println(positions)
    fmt.Println(len(positions))
}
