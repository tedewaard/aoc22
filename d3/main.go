package main

import (
    "os"
    "fmt"
    "bufio"
)


func main() {
   
    //Open file
    readfile,err := os.Open("d3_data")
    //readfile,err := os.Open("test")
    if err != nil {
        fmt.Println(err)
    }

    filescanner := bufio.NewScanner(readfile)

    filescanner.Split(bufio.ScanLines)
    var data []string

    for filescanner.Scan() {
        data = append(data, filescanner.Text())
        //fmt.Println(filescanner.Text())
    }

    readfile.Close()

    var item []string

/*  PART 1
    for i := 0; i<=len(data)-1; i++ {
        n := len(data[i])/2
        for j := 0; j <= n-1; j++ {
            a := string(data[i][j])
            b := ""
            for k := n; k <= len(data[i])-1; k++ {
                b = string(data[i][k])
                if a == b {
                    item = append(item, b)
                    break
                    //fmt.Println(b)
            }
            }
            if a == b {
                break
        }
    }
    }
*/

// PART 2
    for i := 0; i<len(data)-1; i+=3 {
        a := ""
        b := ""
        for j := 0; j <= len(data[i])-1; j++ {
            a = string(data[i][j])
            for k := 0; k <= len(data[i+1])-1; k++ {
                if a == string(data[i+1][k]) {
                    for l := 0; l <= len(data[i+2])-1; l++ {
                        b = string(data[i+2][l])
                        if a == b {
                            item = append(item, a)
                            break
                        }
                    }
                }
            if a == b {
                break
            }
            }
            if a == b {
                break
            }
        }
    }





    m := make(map[string]int)
    m["a"] = 1
    m["b"] = 2
    m["c"] = 3
    m["d"] = 4
    m["e"] = 5
    m["f"] = 6
    m["g"] = 7
    m["h"] = 8
    m["i"] = 9
    m["j"] = 10
    m["k"] = 11
    m["l"] = 12
    m["m"] = 13
    m["n"] = 14
    m["o"] = 15
    m["p"] = 16
    m["q"] = 17
    m["r"] = 18
    m["s"] = 19
    m["t"] = 20
    m["u"] = 21 
    m["v"] = 22
    m["w"] = 23
    m["x"] = 24
    m["y"] = 25
    m["z"] = 26
    m["A"] = 27 
    m["B"] = 28 
    m["C"] = 29 
    m["D"] = 30
    m["E"] = 31 
    m["F"] = 32
    m["G"] = 33
    m["H"] = 34 
    m["I"] = 35 
    m["J"] = 36 
    m["K"] = 37 
    m["L"] = 38 
    m["M"] = 39 
    m["N"] = 40 
    m["O"] = 41 
    m["P"] = 42 
    m["Q"] = 43 
    m["R"] = 44 
    m["S"] = 45 
    m["T"] = 46 
    m["U"] = 47 
    m["V"] = 48 
    m["W"] = 49 
    m["X"] = 50 
    m["Y"] = 51 
    m["Z"] = 52 

    sum := 0
    for _, v := range item {
        fmt.Println(v)
        fmt.Println(m[v])
        sum += m[v]
        //fmt.Println(sum)
    }

    fmt.Println(sum)
}
