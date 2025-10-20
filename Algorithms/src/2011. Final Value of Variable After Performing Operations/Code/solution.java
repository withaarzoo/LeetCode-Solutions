class Solution {
    public int finalValueAfterOperations(String[] operations) {
        int X = 0; // initial value
        for (String op : operations) {
            // check if operation has '+' character
            if (op.indexOf('+') != -1) X++;
            else X--;
        }
        return X;
    }
}
