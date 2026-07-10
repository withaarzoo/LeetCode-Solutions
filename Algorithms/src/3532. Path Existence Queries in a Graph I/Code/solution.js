/**
 * @param {number} n
 * @param {number[]} nums
 * @param {number} maxDiff
 * @param {number[][]} queries
 * @return {boolean[]}
 */
var pathExistenceQueries = function(n, nums, maxDiff, queries) {
    // component[i] stores which connected component node i belongs to.
    const component = new Array(n).fill(0);

    // Start with component 0 for the first node.
    let componentId = 0;

    // Check every gap between two consecutive sorted values.
    for (let i = 1; i < n; i++) {
        // A gap larger than maxDiff separates the graph into two parts.
        if (nums[i] - nums[i - 1] > maxDiff) {
            componentId++;
        }

        // Store the component of the current node.
        component[i] = componentId;
    }

    // Create one answer for every query.
    const answer = new Array(queries.length);

    // Two nodes have a path exactly when their component IDs are equal.
    for (let i = 0; i < queries.length; i++) {
        const u = queries[i][0];
        const v = queries[i][1];

        answer[i] = component[u] === component[v];
    }

    return answer;
}; 