class Solution:
    def isCircularSentence(self, sentence: str) -> bool:
        # Step 1: Split the sentence into words
        words = sentence.split()
        
        # Step 2: Check adjacent pairs and the circular condition
        for i in range(len(words)):
            last_char = words[i][-1]
            first_char = words[(i + 1) % len(words)][0]
            if last_char != first_char:
                return False
        
        return True
