// An implementation of the quick sort algorithm
package main

import (
    "time"
    "math/rand"
    "fmt"
)

var arr = []int{6,109,8,2,4,1,6,9,18,4,2,-9,14,62,-45,8,12}


func main() {
    fmt.Println(quickSort(arr))
}

func quickSort(a []int) []int {

    if len(a) < 2 {
        return a
    }

    // Generate a random pivot
    rand.Seed(time.Now().UnixNano())
    p := rand.Intn(len(a))

    var a1 []int
    a2 := a[p]
    var a3 []int

    // Split the values into 2 slices one for values smaller than the pivot and one for values greater
    for i, v := range a {
        if i == p {
            continue
        }
        if v <= a2 {
            a1 = append(a1, v)
        } else {
            a3 = append(a3, v)
        }
    }

    // Recursively call quickSort on the two slices
    r1 := quickSort(a1)
    r2 := a2
    r3 := quickSort(a3)
   
    r1 = append(r1, r2)
    r1 = append(r1, r3...)

    return r1
}
