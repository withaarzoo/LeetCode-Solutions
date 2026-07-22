class Solution:
    def maxActiveSectionsAfterTrade(self, s: str, queries: List[List[int]]) -> List[int]:
        n = len(s) # Size of the string
        total_ones = s.count('1') # Base count of 1s without trades
        
        # Parallel lists to hold run-length encoded string segments
        type_arr = []
        start = []
        end_idx = []
        
        # Build segments of contiguous blocks of same character
        i = 0
        while i < n:
            j = i
            while j < n and s[j] == s[i]:
                j += 1
            # Append 0 or 1 identifier and its bounds
            type_arr.append(int(s[i]))
            start.append(i)
            end_idx.append(j - 1)
            i = j
            
        N = len(type_arr) # Count of blocks
        
        # Maps absolute string index to block ID so we don't have to binary search later
        pos_to_seg = [0] * n
        for idx in range(N):
            for j in range(start[idx], end_idx[idx] + 1):
                pos_to_seg[j] = idx
                
        # Calculate optimal full gain for converting any 1-block to 0s
        ans = [0] * N
        for idx in range(1, N - 1):
            if type_arr[idx] == 1:
                # The net gain is literally the sum of the adjacent 0-blocks' lengths
                ans[idx] = (end_idx[idx - 1] - start[idx - 1] + 1) + (end_idx[idx + 1] - start[idx + 1] + 1)
                
        # Setup table for Sparse Table initialization
        log_table = [0] * (N + 1)
        for idx in range(2, N + 1):
            log_table[idx] = log_table[idx // 2] + 1
            
        K = log_table[N] + 1
        st = [[0] * N for _ in range(K)] # Sparse Table grid
        
        # Load base layer
        for idx in range(N):
            st[0][idx] = ans[idx]
            
        # Compute RMQ tree layers
        for j in range(1, K):
            for idx in range(N - (1 << j) + 1):
                st[j][idx] = max(st[j - 1][idx], st[j - 1][idx + (1 << (j - 1))])
                
        # O(1) query function using max of overlapping intervals
        def query_rmq(L_q, R_q):
            if L_q > R_q:
                return 0
            j = log_table[R_q - L_q + 1]
            return max(st[j][L_q], st[j][R_q - (1 << j) + 1])
            
        # Function to test 1-segments abutting query edges where truncation occurs
        def eval_seg(idx, L, R, segL, segR):
            # Abort if not strictly nested or not a 1-block
            if idx <= segL or idx >= segR: return 0
            if type_arr[idx] == 0: return 0
            
            # Left 0-block might spill past the query start L, clip it if necessary
            if idx - 1 == segL:
                left_len = max(0, end_idx[idx - 1] - L + 1)
            else:
                left_len = end_idx[idx - 1] - start[idx - 1] + 1
                
            # Right 0-block might spill past the query end R, clip it if necessary
            if idx + 1 == segR:
                right_len = max(0, R - start[idx + 1] + 1)
            else:
                right_len = end_idx[idx + 1] - start[idx + 1] + 1
                
            return left_len + right_len

        res = []
        for L, R in queries:
            # Pinpoint the blocks L and R land in
            segL = pos_to_seg[L]
            segR = pos_to_seg[R]
            
            # If the span is small, there's no way to sandwich a 1 between two 0s
            if segR - segL < 2:
                res.append(total_ones)
                continue
                
            max_gain = 0
            # Test immediate boundaries manually since they're vulnerable to truncation
            max_gain = max(max_gain, eval_seg(segL + 1, L, R, segL, segR))
            max_gain = max(max_gain, eval_seg(segR - 1, L, R, segL, segR))
            
            # Fetch highest yield from safe middle blocks instantly
            if segL + 2 <= segR - 2:
                max_gain = max(max_gain, query_rmq(segL + 2, segR - 2))
                
            # Aggregate and append score
            res.append(total_ones + max_gain)
            
        return res