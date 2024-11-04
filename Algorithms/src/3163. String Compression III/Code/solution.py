class Solution:
    def compressedString(self, word: str) -> str:
        comp = []
        count = 1
        
        for i in range(1, len(word) + 1):
            if i == len(word) or word[i] != word[i - 1] or count == 9:
                comp.append(f"{count}{word[i - 1]}")
                count = 1
            else:
                count += 1
        
        return ''.join(comp)
