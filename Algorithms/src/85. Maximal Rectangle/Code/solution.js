var maximalRectangle = function(matrix) {
    if (matrix.length === 0) return 0;

    const cols = matrix[0].length;
    let heights = new Array(cols).fill(0);
    let maxArea = 0;

    const largestRectangleArea = (heights) => {
        let stack = [];
        let max = 0;
        heights.push(0);

        for (let i = 0; i < heights.length; i++) {
            while (stack.length && heights[stack[stack.length - 1]] > heights[i]) {
                let h = heights[stack.pop()];
                let w = stack.length === 0 ? i : i - stack[stack.length - 1] - 1;
                max = Math.max(max, h * w);
            }
            stack.push(i);
        }
        heights.pop();
        return max;
    };

    for (let row of matrix) {
        for (let j = 0; j < cols; j++) {
            heights[j] = row[j] === '1' ? heights[j] + 1 : 0;
        }
        maxArea = Math.max(maxArea, largestRectangleArea(heights));
    }

    return maxArea;
};
