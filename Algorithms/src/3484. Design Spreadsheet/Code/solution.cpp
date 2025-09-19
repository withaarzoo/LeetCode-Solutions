#include <unordered_map>
#include <string>

using namespace std;

class Spreadsheet {
private:
    int rows;
    unordered_map<int,int> vals; // key -> value

    // Convert a cell like "A1" to integer key
    int keyFromCell(const string &cell) {
        int col = cell[0] - 'A';
        int row = stoi(cell.substr(1)) - 1; // input is 1-indexed
        return col * rows + row;
    }

    // Evaluate operand: either integer literal or cell reference
    int evalOperand(const string &op) {
        if (isdigit(op[0])) return stoi(op); // numeric literal
        auto it = vals.find(keyFromCell(op));
        return (it == vals.end()) ? 0 : it->second;
    }

public:
    Spreadsheet(int rows) : rows(rows) {}

    void setCell(string cell, int value) {
        vals[keyFromCell(cell)] = value;
    }

    void resetCell(string cell) {
        vals.erase(keyFromCell(cell));
    }

    int getValue(string formula) {
        // formula is like "=X+Y"
        string expr = formula.substr(1); // drop '='
        size_t plus = expr.find('+');
        string a = expr.substr(0, plus);
        string b = expr.substr(plus + 1);
        return evalOperand(a) + evalOperand(b);
    }
};

/**
 * Your Spreadsheet object will be instantiated and called as such:
 * Spreadsheet* obj = new Spreadsheet(rows);
 * obj->setCell(cell,value);
 * obj->resetCell(cell);
 * int param_3 = obj->getValue(formula);
 */
