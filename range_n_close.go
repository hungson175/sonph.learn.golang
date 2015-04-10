package learngo

func fibonacci(n int, c chan int) {
	pre, next := 0, 1
	for i := 0; i < n; i++ {
		c <- pre
		pre, next = next, pre+next
	}
	close(c)
}

func RunFiboChannel() {

}
