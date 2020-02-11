package main

import "fmt"

import "time"

func main() {
	fmt.Println(mirroredQuery())
	time.Sleep((5 * time.Second))
}

func mirroredQuery() string {
	out := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		out <- request("us/east")
	}()
	go func() {
		time.Sleep(6 * time.Second)
		out <- request("us/west")
	}()
	go func() {
		time.Sleep(500 * time.Millisecond)
		out <- request("us/central")
	}()
	return <-out
}

func request(loc string) string {
	return fmt.Sprintf("Hello from %s", loc)
}
