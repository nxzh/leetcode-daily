package regular

import "fmt"

type Robot struct {
	width     int
	height    int
	x         int
	y         int
	direction int // 0. East 1. West, 2. South 3. North
	round     int
}

func Constructor(width int, height int) Robot {
	rob := Robot{}
	rob.width = width
	rob.height = height
	rob.round = 2*width + 2*height - 4
	rob.direction = 2
	return rob
}
func Constructor2() Robot {
	rob := Robot{}
	rob.width = 18
	rob.height = 11
	rob.direction = 2
	rob.x = 0
	rob.y = 1
	return rob
}
func (this *Robot) Move(num int) {
	if num > this.round {
		num = num % this.round
	}
	this.moveInternal(num)
}

func (this *Robot) moveInternal(num int) {
	if this.direction == 0 {
		if this.x+num > this.width-1 {
			this.direction = 3
			oldX := this.x
			this.x = this.width - 1
			this.Move(oldX + num - this.width + 1)
		} else {
			this.x += num
		}
	} else if this.direction == 1 {
		if this.x-num < 0 {
			this.direction = 2
			oldX := this.x
			this.x = 0
			this.Move(num - oldX)
		} else {
			this.x -= num
		}
	} else if this.direction == 2 {
		if this.y-num < 0 {
			this.direction = 0
			oldY := this.y
			this.y = 0
			this.Move(num - oldY)
		} else {
			this.y -= num
		}
	} else if this.direction == 3 {
		if this.y+num > this.height-1 {
			this.direction = 1
			oldY := this.y
			this.y = this.height - 1
			this.Move(oldY + num - this.height + 1)
		} else {
			this.y += num
		}
	}
}

func (this *Robot) GetPos() []int {
	return []int{this.x, this.y}
}

func (this *Robot) GetDir() string {
	name := []string{"East", "West", "South", "North"}
	return name[this.direction]
}

func T5911() {
	robot := Constructor(6, 3)
	robot.Move(2)
	robot.Move(2)
	fmt.Println(robot.GetPos())
	fmt.Println(robot.GetDir())
	robot.Move(2)
	robot.Move(1)
	robot.Move(4)
	fmt.Println(robot.GetPos())
	fmt.Println(robot.GetDir())
}

/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Move(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */
