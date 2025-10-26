/**
 * Simple Bank System - JavaScript
 */

/**
 * @param {number[]} balance
 */
var Bank = function (balance) {
  // keep balances as numbers (0-indexed)
  this.bal = balance.slice();
};

/**
 * @param {number} account1
 * @param {number} account2
 * @param {number} money
 * @return {boolean}
 */
Bank.prototype.transfer = function (account1, account2, money) {
  const n = this.bal.length;
  if (account1 < 1 || account1 > n || account2 < 1 || account2 > n)
    return false;
  const a = account1 - 1,
    b = account2 - 1;
  if (this.bal[a] < money) return false;
  this.bal[a] -= money;
  this.bal[b] += money;
  return true;
};

/**
 * @param {number} account
 * @param {number} money
 * @return {boolean}
 */
Bank.prototype.deposit = function (account, money) {
  const n = this.bal.length;
  if (account < 1 || account > n) return false;
  this.bal[account - 1] += money;
  return true;
};

/**
 * @param {number} account
 * @param {number} money
 * @return {boolean}
 */
Bank.prototype.withdraw = function (account, money) {
  const n = this.bal.length;
  if (account < 1 || account > n) return false;
  const a = account - 1;
  if (this.bal[a] < money) return false;
  this.bal[a] -= money;
  return true;
};

/**
 * Your Bank object will be instantiated and called as such:
 * var obj = new Bank(balance)
 * var param_1 = obj.transfer(account1,account2,money)
 * var param_2 = obj.deposit(account,money)
 * var param_3 = obj.withdraw(account,money)
 */
