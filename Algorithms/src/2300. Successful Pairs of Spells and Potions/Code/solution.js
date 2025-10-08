/**
 * @param {number[]} spells
 * @param {number[]} potions
 * @param {number} success
 * @return {number[]}
 */
var successfulPairs = function(spells, potions, success) {
    potions.sort((a,b) => a - b);
    const m = potions.length;
    const res = new Array(spells.length);

    for (let i = 0; i < spells.length; ++i) {
        const s = spells[i];
        // compute ceil(success / s) using integer math to avoid floating rounding issues
        const need = Math.floor((success + s - 1) / s);
        // binary search for first potion >= need
        let l = 0, r = m;
        while (l < r) {
            const mid = Math.floor((l + r) / 2);
            if (potions[mid] < need) l = mid + 1;
            else r = mid;
        }
        res[i] = m - l;
    }
    return res;
};
