package main

import (
	"strconv"
	"strings"
)

// Spreadsheet holds number of rows and a map for explicitly set cells.
type Spreadsheet struct {
	rows int
	mp   map[int]int
}

// Constructor initializes the spreadsheet.
func Constructor(rows int) Spreadsheet {
	return Spreadsheet{
		rows: rows,
		mp:   make(map[int]int),
	}
}

// helper: convert "A1" -> integer key
func keyFromCell(cell string, rows int) int {
	col := int(cell[0] - 'A')
	r, _ := strconv.Atoi(cell[1:]) // safe: input guaranteed valid
	return col*rows + (r - 1)
}

func evalOperand(op string, rows int, mp map[int]int) int {
	// numeric literal?
	if op[0] >= '0' && op[0] <= '9' {
		v, _ := strconv.Atoi(op)
		return v
	}
	k := keyFromCell(op, rows)
	if val, ok := mp[k]; ok {
		return val
	}
	return 0
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	k := keyFromCell(cell, this.rows)
	this.mp[k] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	k := keyFromCell(cell, this.rows)
	delete(this.mp, k)
}

func (this *Spreadsheet) GetValue(formula string) int {
	expr := formula[1:] // drop '='
	parts := strings.Split(expr, "+")
	return evalOperand(parts[0], this.rows, this.mp) + evalOperand(parts[1], this.rows, this.mp)
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * obj := Constructor(rows);
 * obj.SetCell(cell,value);
 * obj.ResetCell(cell);
 * param_3 := obj.GetValue(formula);
 */
