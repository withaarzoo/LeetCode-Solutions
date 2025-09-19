import java.util.HashMap;

class Spreadsheet {
    private int rows;
    private HashMap<Integer, Integer> map;

    public Spreadsheet(int rows) {
        this.rows = rows;
        this.map = new HashMap<>();
    }

    // Convert "A1" to key
    private int keyFromCell(String cell) {
        int col = cell.charAt(0) - 'A';
        int row = Integer.parseInt(cell.substring(1)) - 1;
        return col * rows + row;
    }

    // Evaluate operand either numeric or cell
    private int evalOperand(String op) {
        if (Character.isDigit(op.charAt(0))) {
            return Integer.parseInt(op);
        } else {
            return map.getOrDefault(keyFromCell(op), 0);
        }
    }

    public void setCell(String cell, int value) {
        map.put(keyFromCell(cell), value);
    }

    public void resetCell(String cell) {
        map.remove(keyFromCell(cell));
    }

    public int getValue(String formula) {
        String expr = formula.substring(1); // remove '='
        int plus = expr.indexOf('+');
        String a = expr.substring(0, plus);
        String b = expr.substring(plus + 1);
        return evalOperand(a) + evalOperand(b);
    }
}

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * Spreadsheet obj = new Spreadsheet(rows);
 * obj.setCell(cell,value);
 * obj.resetCell(cell);
 * int param_3 = obj.getValue(formula);
 */
