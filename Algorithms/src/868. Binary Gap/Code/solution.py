class Solution:
    def binaryGap(self, n: int) -> int:
        last_position = -1      # last index of 1
        max_distance = 0        # maximum gap
        current_position = 0    # bit index
        
        while n > 0:
            # Check if current bit is 1
            if n & 1:
                if last_position != -1:
                    max_distance = max(max_distance, current_position - last_position)
                last_position = current_position
            
            n >>= 1  # shift right
            current_position += 1
        
        return max_distance