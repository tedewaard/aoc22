package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Current issue is that if the number of moves is greater that 9 I'm not
//grabbing that number

func main() {
   
    //Open file
    readfile,err := os.Open("d5_data")
    //readfile,err := os.Open("test")
    if err != nil {
        fmt.Println(err)
    }

    filescanner := bufio.NewScanner(readfile)

    filescanner.Split(bufio.ScanLines)
    var cargo []string
    var moves []int

    for filescanner.Scan() {
        //s := strings.Split(filescanner.Text(), ",")
        //fmt.Println(filescanner.Text())
        data := filescanner.Text()
        if data != "" {
        if string(data[0]) != "m" {
            cargo = append(cargo, data)
        } else {
            split := strings.Split(data, " ")
            for _,s := range split {
                v, err := strconv.Atoi(s) 
                if err == nil {
                    moves = append(moves, v)
                }
            }
        }
    }

    }

    readfile.Close()

    n := len(cargo)-1

    //Make map
    stacks := make(map[int][]string)
    for i:=n-1; i>=0; i-- {
        for j:=1; j <= len(cargo[i]); j+=4 {
            val, _ := strconv.Atoi(string(cargo[n][j]))
            if string(cargo[i][j]) != " " {
                stacks[val] = append(stacks[val], string(cargo[i][j])) 
            }
        }
    }

    moveCrates := func(count int, from int, to int) {
        /*
        for i := 1; i<=count; i++ {
            f := stacks[from]
            if len(f) == 0 {
                break
            }
            stacks[to] = append(stacks[to], f[len(f)-1])
            stacks[from] = f[:len(f)-1]
        }
        */



        f := stacks[from]
        items := []string{}
        if count >= len(stacks[from]) {
            items = stacks[from]
            stacks[from] = []string{} 
        } else{
            items = f[len(f)-count:]
            stacks[from] = f[0:len(f)-count]
        }
        stacks[to] = append(stacks[to], items...) 
    }

    //Iterate through moves
    for i := 0; i <= len(moves)-1; i+=3 {
        //For each line of moves grab the character and 
        c := moves[i]
        f := moves[i+1]
        t := moves[i+2]
        moveCrates(c, f, t)
    }
//*/


    //Print out the current stacks
    for i:=1; i<=100; i++ {
        if val, ok := stacks[i]; ok {
            fmt.Println(i, val)
        }
    }

}
