#include <vector>
using namespace std;

/*
 Simple Bank System - C++
*/
class Bank
{
private:
    vector<long long> bal; // store balances as 0-indexed

    // helper to check valid 1-indexed account
    inline bool valid(int account)
    {
        return account >= 1 && account <= (int)bal.size();
    }

public:
    Bank(vector<long long> &balance) : bal(balance) {}

    bool transfer(int account1, int account2, long long money)
    {
        if (!valid(account1) || !valid(account2))
            return false;
        int a = account1 - 1;
        int b = account2 - 1;
        if (bal[a] < money)
            return false; // insufficient funds
        bal[a] -= money;
        bal[b] += money;
        return true;
    }

    bool deposit(int account, long long money)
    {
        if (!valid(account))
            return false;
        bal[account - 1] += money;
        return true;
    }

    bool withdraw(int account, long long money)
    {
        if (!valid(account))
            return false;
        int a = account - 1;
        if (bal[a] < money)
            return false;
        bal[a] -= money;
        return true;
    }
};

/**
 * Your Bank object will be instantiated and called as such:
 * Bank* obj = new Bank(balance);
 * bool param_1 = obj->transfer(account1,account2,money);
 * bool param_2 = obj->deposit(account,money);
 * bool param_3 = obj->withdraw(account,money);
 */
