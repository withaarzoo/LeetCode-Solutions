/** 
 * @param {number} n 
 * @return {string[]} 
 */ 
var generateParenthesis = function(n) { 
    const ans = [];
    const cur = [];
    function dfs(open, close) {
        if (open === 0 && close === 0) {
            ans.push(cur.join(''));
            return;
        }
        if (open > 0) {
            cur.push('(');
            dfs(open - 1, close);
            cur.pop();
        }
        if (close > open) {
            cur.push(')');
            dfs(open, close - 1);
            cur.pop();
        }
    }
    dfs(n, n);
    return ans;
};