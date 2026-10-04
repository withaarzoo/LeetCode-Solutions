function checkValidString(s: string): boolean {
    let low = 0, high = 0;
    for (let c of s) {
        if (c === '(') {
            low++;
            high++;
        } else if (c === ')') {
            low--;
            high--;
        } else {
            low--;
            high++;
        }
        if (high < 0) return false;
        if (low < 0) low = 0;
    }
    return low === 0;
};