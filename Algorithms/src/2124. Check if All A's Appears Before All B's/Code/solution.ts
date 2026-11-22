function checkString(s: string): boolean {
    let seenB = false;
    for (const c of s) {
        if (c === 'b') {
            seenB = true;
        } else if (seenB) {
            return false;
        }
    }
    return true;
};