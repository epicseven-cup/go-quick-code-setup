package cmd

type Solution struct {
}

func Constructor(input []int) Solution {
	return Solution{}
}

func (this *Solution) method() int {
	return 0
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
