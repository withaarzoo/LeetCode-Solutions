class Solution:
    def canBeEqual(self, s1: str, s2: str) -> bool:
        # Get even indexed characters and sort them
        even1 = sorted([s1[0], s1[2]])
        even2 = sorted([s2[0], s2[2]])
        
        # Get odd indexed characters and sort them
        odd1 = sorted([s1[1], s1[3]])
        odd2 = sorted([s2[1], s2[3]])
        
        # Both even and odd groups must match
        return even1 == even2 and odd1 == odd2