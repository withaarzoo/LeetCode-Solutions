func nextGreatestLetter(letters []byte, target byte) byte {
    left := 0
    right := len(letters) - 1
    answer := letters[0] // wrap-around default

    for left <= right {
        mid := left + (right-left)/2

        if letters[mid] > target {
            answer = letters[mid]
            right = mid - 1
        } else {
            left = mid + 1
        }
    }

    return answer
}
