/**
 * Simple Bank System - Java
 */
class Bank {
    private long[] bal; // 0-indexed balances

    private boolean valid(int account) {
        return account >= 1 && account <= bal.length;
    }

    public Bank(long[] balance) {
        // copy for safety (optional), or assign directly
        this.bal = new long[balance.length];
        System.arraycopy(balance, 0, this.bal, 0, balance.length);
    }

    public boolean transfer(int account1, int account2, long money) {
        if (!valid(account1) || !valid(account2))
            return false;
        int a = account1 - 1;
        int b = account2 - 1;
        if (bal[a] < money)
            return false;
        bal[a] -= money;
        bal[b] += money;
        return true;
    }

    public boolean deposit(int account, long money) {
        if (!valid(account))
            return false;
        bal[account - 1] += money;
        return true;
    }

    public boolean withdraw(int account, long money) {
        if (!valid(account))
            return false;
        int a = account - 1;
        if (bal[a] < money)
            return false;
        bal[a] -= money;
        return true;
    }
}

/**
 * Your Bank object will be instantiated and called as such:
 * Bank obj = new Bank(balance);
 * boolean param_1 = obj.transfer(account1,account2,money);
 * boolean param_2 = obj.deposit(account,money);
 * boolean param_3 = obj.withdraw(account,money);
 */
