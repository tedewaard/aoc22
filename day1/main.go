package main

import (
    "os"
    "fmt"
    "strconv"
    "bufio"
    "sort"
)


func main() {
   
    //Open file
    readfile,err := os.Open("cals")
    if err != nil {
        fmt.Println(err)
    }

    filescanner := bufio.NewScanner(readfile)

    filescanner.Split(bufio.ScanLines)
    var data []int

    for filescanner.Scan() {
        num,_ := strconv.Atoi(filescanner.Text())
        data = append(data, num)
       // fmt.Println(num)
    }

    readfile.Close()
    var total []int

    for i:=0; i < len(data); i++ {
        sum := 0
        for data[i] != 0 {
            value := data[i]
            sum = sum + value
            if i == len(data)-1 {
               break 
           }
            i++
        }
        total = append(total, sum) 
    }

    sort.Ints(total)
    

    /*
    for _, j := range total {
        fmt.Println(j)
    }
    */
    sum := 0
    for i := len(total)-1; i >= len(total)-3; i-- {
        sum = sum + total[i] 
    }

    fmt.Println(sum)

}
