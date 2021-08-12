package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func index(boo bool) int {
    var k int
    if (boo == false) {
        k = 0
    } else {
        k = 1
    }
    return k
}
func (ip IPAddr) String() string {
    var out string
    for i, v := range ip {
        out += fmt.Sprintf("%v", v)
        fmt.Println("%v", byte(v))
        // fmt.Println("%v", int(v))
        opts := [2]string {"", "."}
        out += opts[index(i < len(ip)-1)]
    }
    return out
}

func main() {
    hosts := map[string]IPAddr{
        "loopback":  {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}