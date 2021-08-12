// A concurrent prime sieve

package main

// Send the sequence 2, 3, 4, ... to channel 'c0'.
func Step(ch chan<- int) {
    for i := 2; ; i++ {
        ch <- i // Send 'i' to channel 'c0'.
    }
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
    for {
        i := <-in // Receive value from 'in'.
        if i%prime != 0 {
            out <- i // Send 'i' to 'out'.
        }
    }
}
// The prime sieve: Daisy-chain Filter processes.
func Generate(rack chan int, count int) {
    c0 := make(chan int) // Create a new channel.
    c1 := make(chan int) // Create a new channel.
    go Step(c0)      // Launch Step goroutine.
    for i := 0; i != count; i++ {
        prime := <-c0
        rack <- prime
        c1 = make(chan int)
        go Filter(c0, c1, prime)
        c0 = c1
    }
    close(rack)
    close(c1)
}
func main() {
	count := 5
	rack := make(chan int)
	go Generate(rack, count)
	for i := 0; i < 10; i++ {
		ctr := 1
		for prime, open := <- rack; open; ctr++ {
			println(ctr, prime)
			prime, open = <- rack 
		}		
	}
}
