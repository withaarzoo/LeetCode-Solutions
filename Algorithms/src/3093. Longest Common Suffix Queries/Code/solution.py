class TrieNode:
    # Using slots prevents dynamic dict creation per instance, saving memory to avoid MLE
    __slots__ = ['children', 'bestLen', 'bestIdx']
    
    def __init__(self):
        self.children = {}
        self.bestLen = float('inf')
        self.bestIdx = float('inf')

class Solution:
    def stringIndices(self, wordsContainer: List[str], wordsQuery: List[str]) -> List[int]:
        root = TrieNode()
        
        # Process every string to build the suffix tree
        for i, word in enumerate(wordsContainer):
            n = len(word)
            curr = root
            
            # Root stores the best word globally (when suffix match length is 0)
            if n < curr.bestLen or (n == curr.bestLen and i < curr.bestIdx):
                curr.bestLen = n
                curr.bestIdx = i
                
            # Iterate backwards to treat suffixes as prefixes
            for char in reversed(word):
                if char not in curr.children:
                    curr.children[char] = TrieNode()
                
                curr = curr.children[char]
                
                # Continuously update the best match info down the branch
                if n < curr.bestLen or (n == curr.bestLen and i < curr.bestIdx):
                    curr.bestLen = n
                    curr.bestIdx = i
                    
        ans = []
        
        # Evaluate queries against the constructed tree
        for query in wordsQuery:
            curr = root
            
            # Traverse backwards until we hit a dead end
            for char in reversed(query):
                if char not in curr.children:
                    break
                curr = curr.children[char]
            
            # Append the stored best index at the deepest reachable node
            ans.append(curr.bestIdx)
            
        return ans