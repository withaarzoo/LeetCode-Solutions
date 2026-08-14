func maximumLengthSubstring(s string) int {
 // Store the number of times each lowercase letter appears
 // inside the current sliding window.
 freq := make([]int, 26)

 // left marks the start of the current window.
 left := 0

 // ans stores the maximum valid window length found so far.
 ans := 0

 // Expand the window one character at a time.
 for right := 0; right < len(s); right++ {
  // Convert the current character into an index from 0 to 25
  // and increase its frequency in the window.
  index := int(s[right] - 'a')
  freq[index]++

  // If this character appears more than two times,
  // shrink the window until it appears at most twice.
  for freq[index] > 2 {
   // Remove the character leaving the window.
   freq[int(s[left]-'a')]--

   // Move the left pointer forward.
   left++
  }

  // Calculate the current valid window length
  // and update the maximum length.
  length := right - left + 1
  if length > ans {
   ans = length
  }
 }

 // Return the longest valid substring length.
 return ans
}