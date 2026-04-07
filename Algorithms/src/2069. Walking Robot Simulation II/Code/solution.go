type Robot struct {
    width, height, perimeter int
    x, y, dir int
    dx []int
    dy []int
    dirs []string
}

func Constructor(width int, height int) Robot {
    return Robot{
        width: width,
        height: height,
        perimeter: 2 * (width + height) - 4,
        x: 0,
        y: 0,
        dir: 0,
        dx: []int{1, 0, -1, 0},
        dy: []int{0, 1, 0, -1},
        dirs: []string{"East", "North", "West", "South"},
    }
}

func (this *Robot) Step(num int) {
    num %= this.perimeter

    // Full cycle case
    if num == 0 {
        num = this.perimeter
    }

    for num > 0 {
        nx := this.x + this.dx[this.dir]
        ny := this.y + this.dy[this.dir]

        // Rotate if next move is invalid
        if nx < 0 || nx >= this.width || ny < 0 || ny >= this.height {
            this.dir = (this.dir + 1) % 4
            continue
        }

        this.x = nx
        this.y = ny
        num--
    }
}

func (this *Robot) GetPos() []int {
    return []int{this.x, this.y}
}

func (this *Robot) GetDir() string {
    return this.dirs[this.dir]
}