import "math"

func largestTriangleArea(points [][]int) float64 {
    n := len(points)
    maxArea := 0.0
    // iterate all triples i < j < k
    for i := 0; i < n-2; i++ {
        for j := i+1; j < n-1; j++ {
            for k := j+1; k < n; k++ {
                x1, y1 := points[i][0], points[i][1]
                x2, y2 := points[j][0], points[j][1]
                x3, y3 := points[k][0], points[k][1]
                // doubled area using shoelace / cross product
                doubled := math.Abs(float64(x1*(y2 - y3) + x2*(y3 - y1) + x3*(y1 - y2)))
                area := doubled * 0.5
                if area > maxArea {
                    maxArea = area
                }
            }
        }
    }
    return maxArea
}
