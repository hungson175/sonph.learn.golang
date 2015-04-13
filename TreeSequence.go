package learngo

type Tree struct {
	Left  *Tree
	value int
	Right *Tree
}

func (t *Tree) GetSequence() []int {
	var left, right []int
	if t.Left != nil {
		left = t.Left.GetSequence()
	}
	if t.Right != nil {
		right = t.Right.GetSequence()
	}
	return append(append(left, t.value), right...)
}

func Traverse(t *Tree, c chan int) {
	if t == nil {
		return
	}
	Traverse(t.Left, c)
	c <- t.value
	Traverse(t.Right, c)
}
func start(t *Tree, c chan int) {

	Traverse(t, c)
	close(c)
}
func (u *Tree) EqualSequence(v *Tree) bool {
	c1, c2 := make(chan int), make(chan int)

	go start(u, c1)
	go start(v, c2)
	for {
		x, openu := <-c1
		y, openv := <-c2
		if openv != openu {
			return false
		} else {
			if !openv {
				//both are closed, ok then !
				return true
			}
			if x != y {
				return false
			}
		}
	}

}
