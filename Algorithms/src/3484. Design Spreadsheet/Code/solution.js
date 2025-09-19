/**
 * @param {number} rows
 */
var Spreadsheet = function (rows) {
  this.rows = rows;
  this.map = new Map(); // key -> value
};

// helper: convert "A1" to integer key
Spreadsheet.prototype.keyFromCell = function (cell) {
  const col = cell.charCodeAt(0) - "A".charCodeAt(0);
  const row = parseInt(cell.slice(1), 10) - 1;
  return col * this.rows + row;
};

// helper: evaluate operand (number or cell)
Spreadsheet.prototype.evalOperand = function (op) {
  if (/^\d/.test(op)) return parseInt(op, 10);
  const key = this.keyFromCell(op);
  return this.map.has(key) ? this.map.get(key) : 0;
};

/**
 * @param {string} cell
 * @param {number} value
 * @return {void}
 */
Spreadsheet.prototype.setCell = function (cell, value) {
  this.map.set(this.keyFromCell(cell), value);
};

/**
 * @param {string} cell
 * @return {void}
 */
Spreadsheet.prototype.resetCell = function (cell) {
  this.map.delete(this.keyFromCell(cell));
};

/**
 * @param {string} formula
 * @return {number}
 */
Spreadsheet.prototype.getValue = function (formula) {
  const expr = formula.slice(1); // drop '='
  const parts = expr.split("+");
  return this.evalOperand(parts[0]) + this.evalOperand(parts[1]);
};

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * var obj = new Spreadsheet(rows)
 * obj.setCell(cell,value)
 * obj.resetCell(cell)
 * var param_3 = obj.getValue(formula)
 */
