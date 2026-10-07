function removeInvalidParentheses(s: string): string[] {
    let left = 0, right = 0;
    for (const c of s) {
        if (c === '(') left++;
        else if (c === ')') {
            if (left > 0) left--;
            else right++;
        }
    }
    const res = new Set<string>();
    const path: string[] = [];
    function dfs(i: number, leftRem: number, rightRem: number, open: number): void {
        if (i === s.length) {
            if (leftRem === 0 && rightRem === 0 && open === 0) res.add(path.join(''));
            return;
        }
        const c = s[i];
        if (c !== '(' && c !== ')') {
            path.push(c);
            dfs(i + 1, leftRem, rightRem, open);
            path.pop();
            return;
        }
        if (c === '(') {
            if (leftRem > 0) {
                dfs(i + 1, leftRem - 1, rightRem, open);
            }
            path.push(c);
            dfs(i + 1, leftRem, rightRem, open + 1);
            path.pop();
        } else {
            if (rightRem > 0) {
                dfs(i + 1, leftRem, rightRem - 1, open);
            }
            if (open > 0) {
                path.push(c);
                dfs(i + 1, leftRem, rightRem, open - 1);
                path.pop();
            }
        }
    }
    dfs(0, left, right, 0);
    return Array.from(res);
}