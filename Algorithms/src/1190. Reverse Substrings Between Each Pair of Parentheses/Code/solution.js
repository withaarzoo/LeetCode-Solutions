/** 
 * @param {string} s 
 * @return {string} 
 */ 
var reverseParentheses = function(s) { 
    const n = s.length;
    const pair = new Array(n);
    const st = [];
    for (let i = 0; i < n; ++i) {
        if (s[i] === '(') st.push(i);
        else if (s[i] === ')') {
            const j = st.pop();
            pair[i] = j;
            pair[j] = i;
        }
    }
    let res = '';
    let i = 0, dir = 1;
    while (i >= 0 && i < n) {
        if (s[i] === '(' || s[i] === ')') {
            i = pair[i];
            dir = -dir;
        } else {
            res += s[i];
        }
        i += dir;
    }
    return res;
};