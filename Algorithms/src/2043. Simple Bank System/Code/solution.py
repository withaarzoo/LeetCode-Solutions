# Simple Bank System - Python3
from typing import List

class Bank:

    def __init__(self, balance: List[int]):
        # store balances as 0-indexed list copy
        self.bal = balance[:]  

    def _valid(self, account: int) -> bool:
        return 1 <= account <= len(self.bal)

    def transfer(self, account1: int, account2: int, money: int) -> bool:
        if not (self._valid(account1) and self._valid(account2)):
            return False
        a, b = account1 - 1, account2 - 1
        if self.bal[a] < money:
            return False
        self.bal[a] -= money
        self.bal[b] += money
        return True

    def deposit(self, account: int, money: int) -> bool:
        if not self._valid(account):
            return False
        self.bal[account - 1] += money
        return True

    def withdraw(self, account: int, money: int) -> bool:
        if not self._valid(account):
            return False
        a = account - 1
        if self.bal[a] < money:
            return False
        self.bal[a] -= money
        return True

# Your Bank object will be instantiated and called as such:
# obj = Bank(balance)
# param_1 = obj.transfer(account1,account2,money)
# param_2 = obj.deposit(account,money)
# param_3 = obj.withdraw(account,money)
