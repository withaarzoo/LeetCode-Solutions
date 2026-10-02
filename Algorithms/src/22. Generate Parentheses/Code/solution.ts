function generateParenthesis(n: number): string[] { 
    const ans: string[] = [];
    const cur: string[] = [];
    function dfs(open: number, close: number): void {
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