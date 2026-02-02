import heapq

class Solution:
    def minimumCost(self, nums: List[int], k: int, dist: int) -> int:
        target = k - 2
        
        # We need two heaps with lazy removal
        # L: Max-heap (simulated with negative values) storing the smallest 'target' numbers
        # R: Min-heap storing the rest
        L = [] 
        R = []
        L_sum = 0
        
        # To handle lazy removal
        to_remove = {}
        
        def add_lazy(val):
            nonlocal L_sum
            # Naively push to L, balance later
            heapq.heappush(L, -val)
            L_sum += val
            
        def remove_lazy(val):
            # Mark for removal
            to_remove[val] = to_remove.get(val, 0) + 1
            
        def clean(heap, is_max_heap):
            # Remove invalid top elements
            while heap:
                val = -heap[0] if is_max_heap else heap[0]
                if to_remove.get(val, 0) > 0:
                    to_remove[val] -= 1
                    if to_remove[val] == 0:
                        del to_remove[val]
                    heapq.heappop(heap)
                else:
                    break

        def balance():
            nonlocal L_sum
            # Clean heaps before checking tops
            clean(L, True)
            clean(R, False)
            
            # Ensure L has size 'target' (conceptually)
            # Since we can't easily track size with lazy removal, we rely on shifting logic
            # However, simpler logic: Always keep L "full" and R "valid"
            
            # 1. Move largest from L to R if we have excess
            # We track size manually? 
            # Given Python's limitations, we can just maintain:
            # - valid elements in L = target
            pass 

        # RE-APPROACH FOR PYTHON:
        # Standard Two-Heaps is complex due to Lazy Removal size tracking.
        # Instead, we will simulate the logic carefully.
        
        # Re-initialize for a clean structure
        L, R = [], [] # L is max-heap (-val), R is min-heap (val)
        L_sum = 0
        L_size = 0 # tracks valid elements in L
        
        # Helper to push to proper heap
        def push(val):
            nonlocal L_sum, L_size
            if not L or val < -L[0]:
                heapq.heappush(L, -val)
                L_sum += val
                L_size += 1
            else:
                heapq.heappush(R, val)
                
        # Helper to balance
        def rebalance():
            nonlocal L_sum, L_size
            clean(L, True)
            clean(R, False)
            
            # If L has too many
            while L_size > target:
                clean(L, True)
                val = -heapq.heappop(L)
                L_sum -= val
                L_size -= 1
                heapq.heappush(R, val)
                clean(L, True) # Clean again to check next true top
            
            # If L has too few
            while L_size < target and R:
                clean(R, False)
                val = heapq.heappop(R)
                heapq.heappush(L, -val)
                L_sum += val
                L_size += 1
                clean(R, False)
                
            # Note: The swapping logic (L.max > R.min) is handled naturally 
            # because we only push to L if val < L.max, otherwise R.
            # Then we rebalance sizes.
                
        # Initial Window
        for j in range(2, min(len(nums), dist + 2)):
            push(nums[j])
        
        rebalance()
        
        min_cost = nums[0] + nums[1] + L_sum
        
        for i in range(2, len(nums) - (k - 1) + 1):
            out_val = nums[i]
            
            # We don't know immediately if out_val is in L or R because of duplicates.
            # We check if it could be in L (<= max(L))
            # But with lazy removal, we just decrement size if we think it's in L
            
            clean(L, True)
            in_L = L and out_val <= -L[0]
            
            remove_lazy(out_val)
            
            if in_L:
                L_sum -= out_val
                L_size -= 1
            
            if i + dist < len(nums):
                push(nums[i + dist])
                
            rebalance()
            min_cost = min(min_cost, nums[0] + nums[i] + L_sum)
            
        return min_cost