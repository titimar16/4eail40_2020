package main

func fanOut(In <-chan int, OutA, OutB chan int) {
	for data := range In { // receives until closes
		select {
		case OutA <- data:
			//TODO
			break
		case OutB <- data:
			//TODO
			break
		}
	}
}
