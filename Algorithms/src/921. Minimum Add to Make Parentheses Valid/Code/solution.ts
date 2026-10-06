function minAddToMakeValid(s: string): number {
    let open = 0, add = 0;
    for (const c of s) {
        if (c === '(') {
            open++;
        } else {
            if (open > 0) {
                open--;
            } else {
                add++;
            }
        }
    }
    return add + open;
};