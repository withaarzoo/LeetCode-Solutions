/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {
	// Store the position of the first critical point.
	first := -1

	// Store the position of the most recent critical point.
	last := -1

	// Use the largest integer value so the first valid distance can replace it.
	minDistance := int(^uint(0) >> 1)

	// The current node starts at position 1 because head is position 0.
	position := 1

	// prev starts at head because the current node needs a previous node.
	prev := head

	// curr starts from the second node.
	curr := head.Next

	// The last node cannot be critical because it has no next node.
	for curr != nil && curr.Next != nil {
		// Check whether curr is a local maximum or local minimum.
		isCritical := (curr.Val > prev.Val && curr.Val > curr.Next.Val) ||
			(curr.Val < prev.Val && curr.Val < curr.Next.Val)

		// Process the node only if it is a critical point.
		if isCritical {
			// Save this position if it is the first critical point.
			if first == -1 {
				first = position
			} else {
				// Calculate the distance from the previous critical point.
				distance := position - last

				// Keep the smallest distance found so far.
				if distance < minDistance {
					minDistance = distance
				}
			}

			// Store this critical point as the latest one.
			last = position
		}

		// Move prev forward for the next comparison.
		prev = curr

		// Move curr forward to the next node.
		curr = curr.Next

		// Move to the next position.
		position++
	}

	// Fewer than two critical points means no valid distance exists.
	if first == -1 || first == last {
		return []int{-1, -1}
	}

	// The distance between the first and last critical points is the maximum.
	maxDistance := last - first

	// Return the minimum and maximum distances.
	return []int{minDistance, maxDistance}
}