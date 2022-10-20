package main

func main() {

	k := make(chan bool, 2)

	go func() {
		k <- true
		k <- false
	}()

	// fmt.Println(<-k, <-k)
	//output: true false

	// fmt.Println(<-k, <-k, <-k) // dead lock

}
