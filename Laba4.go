package main

import ("fmt")



type token struct {
	data      string
	recipient int
	ttl       int
}



func worker(index int, in <-chan token, out chan<- token) {
	TokenNew := <- in
	fmt.Println("node", index)
	TokenNew.ttl--

	if index == TokenNew.recipient {
		fmt.Println("success")
		return
	}

	if TokenNew.ttl > 0 {
		out <- TokenNew
	} else {
		fmt.Println("ttl <= 0")
	}
}


func main() {
	var n, recipient, ttl int

	fmt.Print("n: ")
	fmt.Scanf("%d", &n)

	fmt.Print("recipient: ")
	fmt.Scanf("%d", &recipient)

	fmt.Print("ttl: ")
	fmt.Scanf("%d", &ttl)

	in := make(chan token)

	InTemp := in
	OutTemp := make(chan token)

	for i := 0; i < n; i++ {
		go worker(i, InTemp, OutTemp)
		InTemp = OutTemp
		OutTemp = make(chan token)
	}

	in <- token{"MY_SUPER_TOKEN", recipient, ttl}
}