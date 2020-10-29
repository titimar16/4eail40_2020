package main

func fanOut(in <-chan int, outA, outB chan int) {
	for data := range in { // receives until closes
		select {
		case outA <- data:
			//TODO
			break
		case outB <- data:
			//TODO
			break
		}
	}
}
