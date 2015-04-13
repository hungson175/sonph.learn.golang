package learngo

func Swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func Sort(a []int, compare func(int, int) int) []int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if compare(a[i], a[j]) == 1 {
				Swap(&a[i], &a[j])
			}
		}
	}
	return a
}
