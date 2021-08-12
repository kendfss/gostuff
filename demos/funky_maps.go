package main

import (
    "fmt"

    "github.com/thoas/go-funk"
)

func main() {
    var map1 = map[string]int{"Mark": 10, "Sandy": 20,
        "Rocky": 30, "Rajiv": 40, "Kate": 50}

    // Map manipulates an iteratee and transforms it to another type.
    flip := funk.Map(map1, func(k string, v int) (int, string) {
        return v, k
    })
    fmt.Println(flip)

    // Keys creates an array of the own enumerable map keys.
    keys := funk.Keys(map1)
    fmt.Println(keys)

    // Values creates an array of the own enumerable map values.
    values := funk.Values(map1)
    fmt.Println(values)
}