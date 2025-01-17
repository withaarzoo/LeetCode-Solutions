class Solution:
    def doesValidArrayExist(self, derived: List[int]) -> bool:
        n = len(derived)
        
        # Case 1: Assume original[0] = 0
        valid_case1 = True
        current = 0  # original[0]
        for i in range(n):
            current = derived[i] ^ current  # Compute original[i+1]
        valid_case1 = (current == 0)  # Wrap-around condition
        
        # Case 2: Assume original[0] = 1
        valid_case2 = True
        current = 1  # original[0]
        for i in range(n):
            current = derived[i] ^ current  # Compute original[i+1]
        valid_case2 = (current == 1)  # Wrap-around condition
        
        return valid_case1 or valid_case2
