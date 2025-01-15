class Solution:
    def minimizeXor(self, num1: int, num2: int) -> int:
        count2 = bin(num2).count('1') # Number of 1s in num2
        count1 = bin(num1).count('1') # Number of 1s in num1

        if count1 == count2:
            return num1

        result = num1
        if count1 > count2:
            for i in range(32):
                if count1 == count2:
                    break
                if result & (1 << i):
                    result &= ~(1 << i) # Clear bit
                    count1 -= 1
        else:
            for i in range(32):
                if count1 == count2:
                    break
                if not (result & (1 << i)):
                    result |= (1 << i) # Set bit
                    count1 += 1

        return result
