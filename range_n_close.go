package learngo

func fibonacci(n int, c chan int) {
	pre, next := 0, 1
	for i := 0; i < n; i++ {
		c <- pre
		pre, next = next, pre+next
	}
	close(c)
}

func RunFiboChannel(n int) []int {
	c := make(chan int, n)
	go fibonacci(n, c)
	var result []int
	for v := range c {
		result = append(result, v)
	}
	return result
}

func fiboSelect(c, quit chan int) {
	curr, next := 0, 1
	for {
		select {
		case c <- curr:
			curr, next = next, curr+next
		case <-quit:
			return
		}
	}
}
func RunFiboSelect(n int) []int {
	c, quit := make(chan int), make(chan int)
	var result []int
	go func() {
		for i := 0; i < n; i++ {
			result = append(result, <-c)
		}
		quit <- 1
	}()
	fiboSelect(c, quit)
	return result
}
