class Spreadsheet:

    def __init__(self, rows: int):
        # number of rows (columns are fixed A-Z)
        self.rows = rows
        # store only explicitly set cells: key -> value
        self._vals = {}

    def _key_from_cell(self, cell: str) -> int:
        # "A1" -> col index 0, row index 0 => key = col * rows + row
        col = ord(cell[0]) - ord('A')
        row = int(cell[1:]) - 1
        return col * self.rows + row

    def _eval_operand(self, op: str) -> int:
        # if operand starts with digit -> literal number; else cell reference
        if op[0].isdigit():
            return int(op)
        return self._vals.get(self._key_from_cell(op), 0)

    def setCell(self, cell: str, value: int) -> None:
        self._vals[self._key_from_cell(cell)] = value

    def resetCell(self, cell: str) -> None:
        self._vals.pop(self._key_from_cell(cell), None)

    def getValue(self, formula: str) -> int:
        # formula is like "=X+Y"
        expr = formula[1:]  # drop '='
        a, b = expr.split('+')
        return self._eval_operand(a) + self._eval_operand(b)


# Your Spreadsheet object will be instantiated and called as such:
# obj = Spreadsheet(rows)
# obj.setCell(cell,value)
# obj.resetCell(cell)
# param_3 = obj.getValue(formula)
