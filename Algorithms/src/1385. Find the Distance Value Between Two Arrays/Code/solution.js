/**
 * @param {number[]} arr1
 * @param {number[]} arr2
 * @param {number} d
 * @return {number}
 */
var findTheDistanceValue = function(arr1, arr2, d) {
    arr2.sort((a, b) => a - b);
    let ans = 0;
    for (let x of arr1) {
        let left = 0, right = arr2.length;
        while (left < right) {
            let mid = (left + right) >> 1;
            if (arr2[mid] < x) left = mid + 1;
            else right = mid;
        }
        let ok = true;
        if (left < arr2.length && Math.abs(arr2[left] - x) <= d) ok = false;
        if (left > 0 && Math.abs(arr2[left - 1] - x) <= d) ok = false;
        if (ok) ans++;
    }
    return ans;
};