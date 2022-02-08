package main

import (
    "fmt"
    "time"
)

func func1() {
    for i:=0; i<10; i++ {
        time.Sleep(time.Second*1)
        fmt.Println("func1:",i)
    }
}

func main() {
    go func1()

    for i:=0; i<10; i++ {
        time.Sleep(time.Second/2)
        fmt.Println("main:",i)
    }

}