#include <vector>
#include <string>
using namespace std;

class Solution {
public:
    int finalValueAfterOperations(vector<string>& operations) {
        int X = 0; // start from 0
        for (const string &op : operations) {
            // if operation contains '+', it's an increment, else it's a decrement
            if (op.find('+') != string::npos) X++;
            else X--;
        }
        return X;
    }
};
