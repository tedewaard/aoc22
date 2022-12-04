package main

import (
    "os"
    "strconv"
    "fmt"
    "strings"
    "bufio"
)


func main() {
   
    //Open file
    readfile,err := os.Open("d4_data")
    //readfile,err := os.Open("test")
    if err != nil {
        fmt.Println(err)
    }

    filescanner := bufio.NewScanner(readfile)

    filescanner.Split(bufio.ScanLines)
    var data1 []string
    var data2 []string

    for filescanner.Scan() {
        s := strings.Split(filescanner.Text(), ",")


        data1 = append(data1, s[0])
        data2 = append(data2, s[1])

        //fmt.Println(s[0])
        //fmt.Println(s[1])

    }

    readfile.Close()


    sum := 0
    for i := 0; i <= len(data1)-1; i++ {
        d1 := strings.Split(data1[i], "-")
        d2 := strings.Split(data2[i], "-")


        l1,err := strconv.Atoi(string(d1[0]))
        if err != nil {
            fmt.Println("Error converting string to int")
        }
        l2,err := strconv.Atoi(string(d2[0]))
        if err != nil {
            fmt.Println("Error converting string to int")
        }
        u1,err := strconv.Atoi(string(d1[1]))
        if err != nil {
            fmt.Println("Error converting string to int")
        }
        u2,err := strconv.Atoi(string(d2[1]))
        if err != nil {
            fmt.Println("Error converting string to int")
        }

        fmt.Println(l1,"-",u1)
        fmt.Println(l2,"-",u2)
        
        //3-4 2-3
        //2-3 3-4
        //1-10
        //2-11

        if l1 <= l2 && u1 >= l2 {
            sum += 1
        } else if l1 >= l2 && l1 <= u2 {
            sum += 1
        }


    }


fmt.Println(sum)


}
