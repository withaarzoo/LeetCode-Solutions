class Solution {
public:
    bool isRectangleOverlap(vector<int>& rec1, vector<int>& rec2) {
        // return true only when the rectangles share positive area on both axes
        // the four conditions below cover every way they can miss each other
        return !(rec1[2] <= rec2[0] ||   // rec1 completely left of rec2
                 rec1[0] >= rec2[2] ||   // rec1 completely right of rec2
                 rec1[3] <= rec2[1] ||   // rec1 completely below rec2
                 rec1[1] >= rec2[3]);    // rec1 completely above rec2
    }
};