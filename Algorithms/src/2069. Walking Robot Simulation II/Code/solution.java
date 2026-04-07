class Robot {
    private int width, height, perimeter;
    private int x, y, dir;

    // Directions: East, North, West, South
    private int[] dx = { 1, 0, -1, 0 };
    private int[] dy = { 0, 1, 0, -1 };
    private String[] dirs = { "East", "North", "West", "South" };

    public Robot(int width, int height) {
        this.width = width;
        this.height = height;
        this.perimeter = 2 * (width + height) - 4;

        this.x = 0;
        this.y = 0;
        this.dir = 0; // Facing East
    }

    public void step(int num) {
        num %= perimeter;

        // Full cycle case
        if (num == 0) {
            num = perimeter;
        }

        while (num > 0) {
            int nx = x + dx[dir];
            int ny = y + dy[dir];

            // Rotate if out of boundary
            if (nx < 0 || nx >= width || ny < 0 || ny >= height) {
                dir = (dir + 1) % 4;
                continue;
            }

            x = nx;
            y = ny;
            num--;
        }
    }

    public int[] getPos() {
        return new int[] { x, y };
    }

    public String getDir() {
        return dirs[dir];
    }
}