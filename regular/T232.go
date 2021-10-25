package regular

type MyQueue struct {
	st1 []int
	st2 []int
}

func Constructor232() MyQueue {
	st1 := []int{}
	st2 := []int{}
	ret := MyQueue{st1, st2}
	return ret
}

func (this *MyQueue) Push(x int) {
	this.st1 = append(this.st1, x)
}

func (this *MyQueue) Pop() int {
	if len(this.st1) == 0 {
		return -1
	}
	this.moveToSt2()
	ret := this.st1[0]
	this.st1 = this.st1[:0]
	this.moveToSt1()
	return ret
}

func (this *MyQueue) moveToSt1() {
	for i := len(this.st2) - 1; i >= 0; i-- {
		this.st1 = append(this.st1, this.st2[i])
	}
	this.st2 = this.st2[0:0]
}

func (this *MyQueue) moveToSt2() {
	for {
		len1 := len(this.st1)
		if len1 != 1 {
			last := this.st1[len1-1]
			this.st2 = append(this.st2, last)
			this.st1 = this.st1[:len1-1]
		} else {
			break
		}
	}
}

func (this *MyQueue) Peek() int {
	if len(this.st1) == 0 {
		return -1
	}
	this.moveToSt2()
	ret := this.st1[0]
	this.moveToSt1()
	return ret
}

func (this *MyQueue) Empty() bool {
	return len(this.st1) == 0
}

func T232() {
	//queue := Constructor()
	//queue.Push(1)
	//queue.Pop()
	//fmt.Println(queue.Empty())
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
