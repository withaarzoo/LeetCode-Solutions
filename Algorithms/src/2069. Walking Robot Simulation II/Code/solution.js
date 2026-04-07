var Robot = function (width, height) {
  this.width = width;
  this.height = height;
  this.perimeter = 2 * (width + height) - 4;

  this.x = 0;
  this.y = 0;
  this.dir = 0;

  this.dx = [1, 0, -1, 0];
  this.dy = [0, 1, 0, -1];
  this.dirs = ["East", "North", "West", "South"];
};

Robot.prototype.step = function (num) {
  num %= this.perimeter;

  // Full cycle case
  if (num === 0) {
    num = this.perimeter;
  }

  while (num > 0) {
    let nx = this.x + this.dx[this.dir];
    let ny = this.y + this.dy[this.dir];

    // Rotate if next move is invalid
    if (nx < 0 || nx >= this.width || ny < 0 || ny >= this.height) {
      this.dir = (this.dir + 1) % 4;
      continue;
    }

    this.x = nx;
    this.y = ny;
    num--;
  }
};

Robot.prototype.getPos = function () {
  return [this.x, this.y];
};

Robot.prototype.getDir = function () {
  return this.dirs[this.dir];
};
