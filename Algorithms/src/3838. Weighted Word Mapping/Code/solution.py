class Solution:
    def mapWordWeights(self, words: List[str], weights: List[int]) -> str:
        result = []
        
        # Process each word
        for word in words:
            sum_weight = 0
            
            # Add weights of all characters
            for ch in word:
                sum_weight += weights[ord(ch) - ord('a')]
            
            # Take modulo 26
            value = sum_weight % 26
            
            # Reverse alphabetical mapping
            result.append(chr(ord('z') - value))
        
        return "".join(result)