package main

import (
    "os"
    "fmt"
    "bufio"
    "strconv"
    "regexp"
)



func main() {
   
    //Open file
    //readfile,err := os.Open("d5_data")
    readfile,err := os.Open("d5_data")
    if err != nil {
        fmt.Println(err)
    }

    filescanner := bufio.NewScanner(readfile)

    filescanner.Split(bufio.ScanLines)
    var cargo []string
    var moves []string

    for filescanner.Scan() {
        //s := strings.Split(filescanner.Text(), ",")
        //fmt.Println(filescanner.Text())
        data := filescanner.Text()
        if data != "" {
        if string(data[0]) != "m" {
            cargo = append(cargo, data)
        } else {
            re := regexp.MustCompile("[0-9]+")
            moves = append(moves, re)
        }
    }

    }

    readfile.Close()



    //Index starts at 1 and jumps 4
    //1, 5, 9, 13
    //This will grab the number of stacks
    /*
    n := len(cargo)-1
    numStacks := []int{} 
    for i := 1; i <= len(cargo[n]); i+=4 {
        if string(cargo[n][i]) != "" {
            val, _ := strconv.Atoi(string(cargo[n][i]))
            numStacks = append(numStacks, val)
        }
    }
    */

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
///*
    moveCrates := func(count int, from int, to int) {
        for i := 1; i<=count; i++ {
            f := stacks[from]
            if len(f) == 0 {
                break
            }
            stacks[to] = append(stacks[to], f[len(f)-1])
            stacks[from] = f[:len(f)-1]
        }



        /*
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
        */
    }

    //Iterate through moves
    for i := 0; i <= len(moves)-1; i++ {

        for j := 0; j <=len(moves[i]); j++ {

        }



        //For each line of moves grab the character and 
        c, _ := strconv.Atoi(string(moves[i][5]))
        f, _ := strconv.Atoi(string(moves[i][12]))
        t, _ := strconv.Atoi(string(moves[i][17]))
        moveCrates(c, f, t)
    }
//*/


    //Print out the current stacks
    for i:=1; i<=100; i++ {
        if val, ok := stacks[i]; ok {
            fmt.Println(i, val)
        }
    }

    /*
    for _,i := range cargo {
        fmt.Println(i)
    }
    for _,i := range moves {
        fmt.Println(i)
    }
    */
}
