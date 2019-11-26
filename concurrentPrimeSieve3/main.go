package main

import (
    "flag";
    "fmt";
    "os";
)

var prime = flag.Int("p", 13, "test for primeness")

func f(left, right chan int, n int) {
    i := <-right;
    if i == 0 || n >= i { 
        left <- i;
        return;
    }   
    if i%n == 0 { 
        // too bad
        left <- 0;
        return;
    }   
    left <- i;
}

func main() {
    flag.Parse();

    if *prime < 2 { 
        fmt.Fprint(os.Stderr, "You know this already\n");
        os.Exit(1);
    }   

    leftmost := make(chan int);
    var left, right chan int = nil, leftmost;
    for n := 0; n < *prime-1; n++ {
        left, right = right, make(chan int);
        go f(left, right, *prime-n);
    }   
    right <- *prime;        // bang!
    x := <-leftmost;        // wait for completion
    fmt.Println(x);
}