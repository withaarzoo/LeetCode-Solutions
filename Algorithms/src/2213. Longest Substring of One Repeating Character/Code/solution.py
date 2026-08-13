from typing import List

class Solution:
    def longestRepeating(self, s: str, queryCharacters: str, queryIndices: List[int]) -> List[int]:
        # Each node is:
        # (left_char, right_char, length, prefix, suffix, best)
        tree = [None] * (4 * len(s))

        # I merge two adjacent nodes and produce their parent node.
        def merge(left, right):
            # Get the first and last characters of the merged segment.
            left_char = left[0]
            right_char = right[1]

            # The new segment length is the sum of both child lengths.
            length = left[2] + right[2]

            # Initially, the best answer is inside one of the children.
            best = max(left[5], right[5])

            # The prefix normally comes from the left child.
            prefix = left[3]

            # The left segment must be completely uniform before
            # its prefix can extend into the right segment.
            # Checking prefix == length is the important fix.
            if left[3] == left[2] and left[1] == right[0]:
                prefix = left[2] + right[3]

            # The suffix normally comes from the right child.
            suffix = right[4]

            # The right segment must be completely uniform before
            # its suffix can extend into the left segment.
            if right[4] == right[2] and left[1] == right[0]:
                suffix = left[4] + right[2]

            # If the boundary characters match, the best substring
            # may cross the boundary between the two children.
            if left[1] == right[0]:
                best = max(best, left[4] + right[3])

            # Keep all node information in a fixed order.
            return (left_char, right_char, length, prefix, suffix, best)

        # I build the segment tree from the original string.
        def build(node, l, r):
            # A leaf represents exactly one character.
            if l == r:
                tree[node] = (s[l], s[l], 1, 1, 1, 1)
                return

            mid = (l + r) // 2

            # Build the left and right children.
            build(node * 2, l, mid)
            build(node * 2 + 1, mid + 1, r)

            # Merge both children into the parent.
            tree[node] = merge(tree[node * 2], tree[node * 2 + 1])

        # I update only one root-to-leaf path.
        def update(node, l, r, idx, char):
            # Replace the target leaf with the new character.
            if l == r:
                tree[node] = (char, char, 1, 1, 1, 1)
                return

            mid = (l + r) // 2

            # Move to the child containing the updated position.
            if idx <= mid:
                update(node * 2, l, mid, idx, char)
            else:
                update(node * 2 + 1, mid + 1, r, idx, char)

            # Recalculate the parent after the update.
            tree[node] = merge(tree[node * 2], tree[node * 2 + 1])

        # Build the initial segment tree.
        build(1, 0, len(s) - 1)

        answer = []

        for i in range(len(queryCharacters)):
            # Apply the current query.
            update(
                1,
                0,
                len(s) - 1,
                queryIndices[i],
                queryCharacters[i]
            )

            # The root covers the entire string,
            # so position 5 contains the answer.
            answer.append(tree[1][5])

        return answer