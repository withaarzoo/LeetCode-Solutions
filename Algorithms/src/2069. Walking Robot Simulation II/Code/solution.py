class Robot:

    def __init__(self, width: int, height: int):
        self.width = width
        self.height = height
        self.perimeter = 2 * (width + height) - 4

        self.x = 0
        self.y = 0
        self.dir = 0

        # Directions: East, North, West, South
        self.dx = [1, 0, -1, 0]
        self.dy = [0, 1, 0, -1]
        self.dirs = ["East", "North", "West", "South"]

    def step(self, num: int) -> None:
        num %= self.perimeter

        # Full cycle case
        if num == 0:
            num = self.perimeter

        while num > 0:
            nx = self.x + self.dx[self.dir]
            ny = self.y + self.dy[self.dir]

            # Rotate if next move is outside grid
            if nx < 0 or nx >= self.width or ny < 0 or ny >= self.height:
                self.dir = (self.dir + 1) % 4
                continue

            self.x = nx
            self.y = ny
            num -= 1

    def getPos(self) -> List[int]:
        return [self.x, self.y]

    def getDir(self) -> str:
        return self.dirs[self.dir]