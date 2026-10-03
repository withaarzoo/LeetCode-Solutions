function longestValidParentheses(s: string): number {
    let left = 0, right = 0, maxLen = 0;
    for (let c of s) {
        if (c === '(') left++;
        else right++;
        if (left === right) maxLen = Math.max(maxLen, 2 * right);
        else if (right > left) left = right = 0;
    }
    left = right = 0;
    for (let i = s.length - 1; i >= 0; i--) {
        if (s[i] === '(') left++;
        else right++;
        if (left === right) maxLen = Math.max(maxLen, 2 * left);
        else if (left > right) left = right = 0;
    }
    return maxLen;
};