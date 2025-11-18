from typing import List

class Solution:
    def isOneBitCharacter(self, bits: List[int]) -> bool:
        n = len(bits)
        i = 0
        # process until the last bit (we don't need to parse beyond it)
        while i < n - 1:
            if bits[i] == 1:
                # 1 means a two-bit character -> skip two bits
                i += 2
            else:
                # 0 means one-bit character -> skip one bit
                i += 1
        # if pointer is exactly at last index -> last char is one-bit
        return i == n - 1
