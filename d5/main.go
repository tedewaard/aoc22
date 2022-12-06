package main

import (
    "os"
    "fmt"
    "bufio"
)


func moveCrates(count int, from int, to int) {

}

func main() {
   
    //Open file
    //readfile,err := os.Open("d5_data")
    readfile,err := os.Open("test")
    if err != nil {
        fmt.Println(err)
    }

    filescanner := bufio.NewScanner(readfile)

    filescanner.Split(bufio.ScanLines)
    var blocks []string
    var moves []string

    for filescanner.Scan() {
        //s := strings.Split(filescanner.Text(), ",")
        //fmt.Println(filescanner.Text())
        data := filescanner.Text()
        if data != "" {
        if string(data[0]) != "m" {
            blocks = append(blocks, data)
        } else {
            moves = append(moves, data)
        }
    }

    }

    readfile.Close()

    var stacks [][]string

    n := len(blocks)-1
    //Index starts at 1 and jumps 4
    //1, 5, 9, 13

    //Grabbing the number of stacks and adding the stacks to an array
    var stack []int
    for j := 1; j <= len(blocks[n])-1; j+=4 {
       stack = append(stack, j) 
    }


    for i:=n-1; i >= 0; i-- {
        arr := []string{}
        for k := 1; k <= len(blocks[i])-1; k+=4 {
            a:=string(blocks[i][k])
            arr = append(arr, a)
        }
       stacks = append(stacks, arr) 
    }

    for _,i := range stacks {
        fmt.Println(i)
    }

    /*
    for _,i := range blocks {
        fmt.Println(i)
    }
    for _,i := range moves {
        fmt.Println(i)
    }
    */
}
