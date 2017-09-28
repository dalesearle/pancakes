package waiters

type waiter struct {
}

func New() *waiter {
	return new(waiter)
}

func (w *waiter) ArrangeStack(stack string) int {
	flips := 0
	r := []rune(stack)
	ref := r[0]
	for i, rn := range r {
		if ref != rn {
			ref = rn
			flipStack(r, i-1)
			flips++
		}
	}
	if ref == '-' {
		flipStack(r, len(r)-1)
		flips++
	}
	return flips
}

func flipStack(runes []rune, index int) {
	stack := make([]rune, len(runes))
	copy(stack, runes)
	for itr := 0; itr <= index; itr++ {
		runes[itr] = '-'
		if stack[index-itr] == '-' {
			runes[itr] = '+'
		}
	}
}
