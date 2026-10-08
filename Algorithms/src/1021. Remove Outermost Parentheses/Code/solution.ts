function removeOuterParentheses(s: string): string {
    let res = "";
    let bal = 0;
    for (let c of s) {
        if (c === '(') {
            if (bal > 0) res += c;
            bal++;
        } else {
            bal--;
            if (bal > 0) res += c;
        }
    }
    return res;
};