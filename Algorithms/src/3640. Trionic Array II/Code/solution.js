var maxSumTrionic = function(nums) {
    const n = nums.length;
    const left = Array(n).fill(0);
    const right = Array(n).fill(0);

    for(let i = 0; i < n; i++){
        left[i] = nums[i];
        if(i > 0 && nums[i-1] < nums[i] && left[i-1] > 0){
            left[i] += left[i-1];
        }
    }

    for(let i = n-1; i >= 0; i--){
        right[i] = nums[i];
        if(i+1 < n && nums[i] < nums[i+1] && right[i+1] > 0){
            right[i] += right[i+1];
        }
    }

    let parts = [];
    let l = 0, s = nums[0];
    for(let i = 1; i < n; i++){
        if(nums[i-1] <= nums[i]){
            parts.push([l, i-1, s]);
            l = i;
            s = 0;
        }
        s += nums[i];
    }
    parts.push([l, n-1, s]);

    let ans = -1e18;
    for(const [p,q,sum] of parts){
        if(p > 0 && q < n-1 && nums[p-1] < nums[p] && nums[q] < nums[q+1] && p < q){
            ans = Math.max(ans, left[p-1] + sum + right[q+1]);
        }
    }
    return ans;
};
