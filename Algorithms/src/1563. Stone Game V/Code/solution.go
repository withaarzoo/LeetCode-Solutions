func stoneGameV(stoneValue []int) int {
 n := len(stoneValue)

 // prefix[i] stores the sum of the first i stones.
 // It allows every interval sum to be calculated in O(1).
 prefix := make([]int64, n+1)

 for i := 0; i < n; i++ {
  // Build the prefix sum array.
  prefix[i+1] = prefix[i] + int64(stoneValue[i])
 }

 // dp[l][r] stores Alice's maximum score for interval [l, r].
 dp := make([][]int, n)

 // leftBest[l][r] stores the best value when the left side is kept.
 leftBest := make([][]int, n)

 // rightBest[l][r] stores the best value when the right side is kept.
 rightBest := make([][]int, n)

 for i := 0; i < n; i++ {
  // Allocate all DP tables for this starting index.
  dp[i] = make([]int, n)
  leftBest[i] = make([]int, n)
  rightBest[i] = make([]int, n)

  // A single stone cannot be split, but it can be a valid side.
  leftBest[i][i] = stoneValue[i]
  rightBest[i][i] = stoneValue[i]
 }

 // leftPtr[l] is the last split with leftSum <= rightSum.
 leftPtr := make([]int, n)

 // rightPtr[l] is the first split with leftSum >= rightSum.
 rightPtr := make([]int, n)

 for i := 0; i < n; i++ {
  // Initially there is no valid split before i.
  leftPtr[i] = i - 1

  // The first possible split starts at i.
  rightPtr[i] = i
 }

 // Process intervals from shorter to longer.
 // Every smaller interval is already solved when needed.
 for length := 2; length <= n; length++ {
  for l := 0; l+length <= n; l++ {
   r := l + length - 1

   // Calculate the total sum of [l, r].
   total := prefix[r+1] - prefix[l]

   // Move leftPtr while the left side is not larger.
   for leftPtr[l]+1 <= r-1 {
    k := leftPtr[l] + 1
    leftSum := prefix[k+1] - prefix[l]

    // 2 * leftSum <= total means leftSum <= rightSum.
    if 2*leftSum > total {
     break
    }

    leftPtr[l]++
   }

   // Move rightPtr until the left side is at least as large.
   for rightPtr[l] <= r-1 {
    k := rightPtr[l]
    leftSum := prefix[k+1] - prefix[l]

    // Stop at the first split where leftSum >= rightSum.
    if 2*leftSum >= total {
     break
    }

    rightPtr[l]++
   }

   best := 0

   // If the left side is smaller or equal, Alice keeps the left side.
   if leftPtr[l] >= l {
    best = leftBest[l][leftPtr[l]]
   }

   // If the right side is smaller or equal, Alice keeps the right side.
   if rightPtr[l] <= r-1 {
    candidate := rightBest[rightPtr[l]+1][r]

    // Keep whichever choice gives Alice the higher score.
    if candidate > best {
     best = candidate
    }
   }

   // Store the best score for the current interval.
   dp[l][r] = best

   // Use [l, r] as a possible left part of a future interval.
   currentSum := int(total)
   leftCandidate := dp[l][r] + currentSum

   if leftBest[l][r-1] > leftCandidate {
    leftBest[l][r] = leftBest[l][r-1]
   } else {
    leftBest[l][r] = leftCandidate
   }

   // Use [l, r] as a possible right part of a future interval.
   rightCandidate := dp[l][r] + currentSum

   if rightBest[l+1][r] > rightCandidate {
    rightBest[l][r] = rightBest[l+1][r]
   } else {
    rightBest[l][r] = rightCandidate
   }
  }
 }

 // The full array is the interval [0, n-1].
 return dp[0][n-1]
}