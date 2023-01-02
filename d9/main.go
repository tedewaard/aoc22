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
    //readfile,err := os.Open("data")
    readfile,err := os.Open("test2")
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

func checkDiagnal(knot1 position, knot2 position) bool {
    if knot1.x != knot2.x && knot1.y != knot2.y {
        //Start with up and right
        if knot1.x - knot2.x == 1 && knot1.y - knot2.y == 2 {
            return true
        }
        //Up and left
        if knot1.x - knot2.x == -1 && knot1.y - knot2.y == 2 {
            return true
        }
        //Down and right
        if knot1.x - knot2.x == 1 && knot1.y - knot2.y == -2 {
            return true
        }
        //Down and left
        if knot1.x - knot2.x == -1 && knot1.y - knot2.y == -2 {
            return true
        }
    }
    return false
}


func moveHead(point position, direction string ) position {
    
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

func moveDiagnal(knot1 position, knot2 position) position {
    // Diagnal
    if knot1.x != knot2.x && knot1.y != knot2.y {
        //Start with up and right
        if knot1.x - knot2.x == 1 && knot1.y - knot2.y == 2 {
            knot2.x++
            knot2.y++
        }
        //Up and left
        if knot1.x - knot2.x == -1 && knot1.y - knot2.y == 2 {
            knot2.x--
            knot2.y++
        }
        //Down and right
        if knot1.x - knot2.x == 1 && knot1.y - knot2.y == -2 {
            knot2.x++
            knot2.y--
        }
        //Down and left
        if knot1.x - knot2.x == -1 && knot1.y - knot2.y == -2 {
            knot2.x--
            knot2.y--
        }
    }


    return knot2

}

func moveKnots(knot1 position, knot2 position) position {
    /*
    //Same axes (y)
    if knot1.y == knot2.y {
        if knot1.x - knot2.x > 1 {
            knot2.x++
        } else if knot1.x - knot2.x < -1 {
            knot2.x--
        }
    } 
    
    //same axes (x)
    if knot1.x == knot2.x {
        if knot1.y - knot2.y > 1 {
            knot2.y++
        } else if knot1.y - knot2.y < -1 {
            knot2.y--
        }
    }
    */

    // Diagnal
    if knot1.x != knot2.x && knot1.y != knot2.y {
        //Start with up and right
        if knot1.x - knot2.x == 1 && knot1.y - knot2.y == 2 {
            knot2.x++
            knot2.y++
        }
        //Up and left
        if knot1.x - knot2.x == -1 && knot1.y - knot2.y == 2 {
            knot2.x--
            knot2.y++
        }
        //Down and right
        if knot1.x - knot2.x == 1 && knot1.y - knot2.y == -2 {
            knot2.x++
            knot2.y--
        }
        //Down and left
        if knot1.x - knot2.x == -1 && knot1.y - knot2.y == -2 {
            knot2.x--
            knot2.y--
        }
    }


    return knot2
}



func main() {
    knots := []position{
        {0, 0},
        {0, 0},
        {0, 0},
        {0, 0},
        {0, 0},
        {0, 0},
        {0, 0},
        {0, 0},
        {0, 0},
    }
    readData()
    tailPositions := make(map[position]int)
    //I want to move the head and then tell the next one to check if it's 
    //touching and if not then also move. Then the next should do the same.


    for _, i := range data {
        j:=1
        for j <= i.steps {
            //Move the head
            knots[0] = moveHead(knots[0], i.dir) 
            //Now loop through rest of knots
            for k:=1; k<=len(knots)-1; k++ {
                //Check if they are touching, if they aren't then check if diangal
                //and move move accordingly
                if checkTouching(knots[k-1], knots[k]) == false {
                    if checkDiagnal(knots[k-1], knots[k]) {
                        knots[k] = moveDiagnal(knots[k-1], knots[k]) 
                    } else {
                        knots[k] = moveHead(knots[k], i.dir)
                    }
                }
            }
            j++
        }
        tailPositions[knots[8]] = 1
    fmt.Println(knots)
}



    fmt.Println(tailPositions)
    fmt.Println(knots)
    fmt.Println(len(tailPositions))
}
