class Solution {
public:
    double largestTriangleArea(vector<vector<int>>& points) {
        int n = points.size();
        double maxArea = 0.0;
        // iterate over all triples i < j < k
        for (int i = 0; i < n-2; ++i) {
            for (int j = i+1; j < n-1; ++j) {
                for (int k = j+1; k < n; ++k) {
                    // coordinates
                    int x1 = points[i][0], y1 = points[i][1];
                    int x2 = points[j][0], y2 = points[j][1];
                    int x3 = points[k][0], y3 = points[k][1];
                    // shoelace / cross-product for double area
                    double doubled = fabs( (double)x1*(y2 - y3)
                                         + (double)x2*(y3 - y1)
                                         + (double)x3*(y1 - y2) );
                    double area = doubled * 0.5;
                    if (area > maxArea) maxArea = area;
                }
            }
        }
        return maxArea;
    }
};
