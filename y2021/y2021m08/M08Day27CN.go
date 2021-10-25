package y2021m08

import "fmt"

type MedianFinder struct {
	counter map[int]int64
	elems   []int
	total   int64
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{counter: make(map[int]int64)}
}

func (this *MedianFinder) AddNum(num int) {
	this.total++
	if len(this.elems) == 0 {
		this.elems = append(this.elems, num)
	} else if this.elems[len(this.elems)-1] < num {
		this.elems = append(this.elems, num)
	} else if this.elems[0] > num {
		this.elems = append([]int{num}, this.elems...)
	} else {
		for i := 1; i < len(this.elems); i++ {
			if this.elems[i] == num {
				break
			} else if this.elems[i] > num {
				if this.elems[i-1] == num {
					break
				} else {
					this.elems = append(this.elems[:i+1], this.elems[i:]...)
					this.elems[i] = num
					break
				}
			}
		}
	}
	this.counter[num] += 1
}

func (this *MedianFinder) FindMedian() float64 {
	if this.total&0x1 == 1 {
		median := (this.total + 1) / 2
		acc := int64(0)
		for i := 0; i < len(this.elems); i++ {
			acc += this.counter[this.elems[i]]
			if acc >= median {
				return float64(this.elems[i])
			}
		}
	}
	median := (this.total + 1) / 2
	acc := int64(0)
	for i := 0; i < len(this.elems)-1; i++ {
		acc += this.counter[this.elems[i]]
		if acc >= median {
			if acc > median {
				return float64(this.elems[i])
			}
			return float64(this.elems[i]+this.elems[i+1]) / 2
		}
	}
	return 0
}

func Call0827Cn() {
	mf := Constructor()
	mf.AddNum(6)
	fmt.Println(mf.FindMedian())
	mf.AddNum(10)
	fmt.Println(mf.FindMedian())
	mf.AddNum(2)
	fmt.Println(mf.FindMedian())
	mf.AddNum(6)
	fmt.Println(mf.FindMedian())
	mf.AddNum(5)
	fmt.Println(mf.FindMedian())
	mf.AddNum(0)
	fmt.Println(mf.FindMedian())
	mf.AddNum(6)
	fmt.Println(mf.FindMedian())
	mf.AddNum(3)
	fmt.Println(mf.FindMedian())
	mf.AddNum(1)
	fmt.Println(mf.FindMedian())
	mf.AddNum(0)
	fmt.Println(mf.FindMedian())
	mf.AddNum(0)
	fmt.Println(mf.FindMedian())
}
