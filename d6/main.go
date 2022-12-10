package main

import (
    "os"
    "fmt"
    "bufio"
)


func main() {
   
    //Open file
    readfile,err := os.Open("d6_data")
    //readfile,err := os.Open("test")
    if err != nil {
        fmt.Println(err)
    }

    filescanner := bufio.NewScanner(readfile)

    filescanner.Split(bufio.ScanLines)
    var datastreams string

    for filescanner.Scan() {
        datastreams =filescanner.Text()
    }

    readfile.Close()  

    var marker int
    //var packet string 
    for j := 0; j<=len(datastreams)-4; j++ {
        check := 0
        substring := datastreams[j:j+4] 
        m := make(map[string]int)
        for _, v := range substring { //Looping through the string of 4
            if _, ok := m[string(v)]; ok { //Checking if key already exists
                check++
                break 
            }
            m[string(v)] = 1 
        }
        if check == 0 {
           marker = j+4 
           //packet = datastreams[j+5:]
           break
        }
    }

    fmt.Println(marker)

    var message int
    for j := marker; j<=len(datastreams)-14; j++ {
        check := 0
        substring := datastreams[j:j+14] 
        m := make(map[string]int)
        for _, v := range substring { //Looping through the string of 4
            if _, ok := m[string(v)]; ok { //Checking if key already exists
                check++
                //fmt.Println("Key exists, increase by 1")
                break 
            }
            m[string(v)] = 1 
        }
        if check == 0 {
           message = j+14 
           break
        }
    }
    

    fmt.Println(message)


/*
    for _, s := range datastreams {
        fmt.Println(s)
    }

*/

}
