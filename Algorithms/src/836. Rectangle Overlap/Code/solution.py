class Solution:
    def isRectangleOverlap(self, rec1: List[int], rec2: List[int]) -> bool:
        # return true only when the rectangles share positive area on both axes
        # the four conditions below cover every way they can miss each other
        return not (rec1[2] <= rec2[0] or   # rec1 completely left of rec2
                    rec1[0] >= rec2[2] or   # rec1 completely right of rec2
                    rec1[3] <= rec2[1] or   # rec1 completely below rec2
                    rec1[1] >= rec2[3])     # rec1 completely above rec2