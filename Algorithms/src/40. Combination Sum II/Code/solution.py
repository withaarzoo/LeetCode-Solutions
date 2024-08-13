class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        # Sort the candidates to make it easier to handle duplicates and to ensure combinations are in non-decreasing order.
        candidates.sort()
        result = []  # This will store all unique combinations that sum up to the target.
        
        def backtrack(start: int, target: int, current: List[int]):
            # Base case: If the target is exactly 0, it means the current combination adds up to the target.
            if target == 0:
                result.append(list(current))  # Add a copy of the current combination to the result list.
                return  # Return to explore other potential combinations.
            
            # Iterate over the candidates starting from the 'start' index to avoid reuse of the same elements.
            for i in range(start, len(candidates)):
                # If the current candidate is the same as the previous one, skip it to avoid duplicate combinations.
                if i > start and candidates[i] == candidates[i - 1]:
                    continue
                
                # If the current candidate exceeds the remaining target, there's no point in continuing, as all
                # subsequent candidates will also be greater (because the list is sorted).
                if candidates[i] > target:
                    break
                
                # Include the current candidate in the combination and move to the next candidate.
                current.append(candidates[i])
                
                # Recursively call backtrack with the next index (i + 1) and the updated target (target - candidates[i]).
                backtrack(i + 1, target - candidates[i], current)
                
                # Backtrack: Remove the last candidate added to try another possibility.
                current.pop()
        
        # Start the backtracking process from index 0 with the original target and an empty combination.
        backtrack(0, target, [])
        
        # Return the final list of unique combinations.
        return result
