class Solution {
    // I store all information needed to merge two adjacent segments.
    static class Node {
        char leftChar;   // First character of the segment.
        char rightChar;  // Last character of the segment.
        int len;         // Total segment length.
        int prefix;      // Longest same-character prefix.
        int suffix;      // Longest same-character suffix.
        int best;        // Longest same-character substring.

        Node() {
            leftChar = 0;
            rightChar = 0;
            len = 0;
            prefix = 0;
            suffix = 0;
            best = 0;
        }
    }

    private Node[] tree;

    // I merge two neighboring nodes into one parent node.
    private Node merge(Node left, Node right) {
        Node res = new Node();

        // The boundaries of the merged segment come from the children.
        res.leftChar = left.leftChar;
        res.rightChar = right.rightChar;
        res.len = left.len + right.len;

        // The best substring is initially inside one child.
        res.best = Math.max(left.best, right.best);

        // By default, the prefix is the left child's prefix.
        res.prefix = left.prefix;

        // The entire left segment must be uniform before
        // the prefix can continue into the right segment.
        if (left.prefix == left.len &&
            left.rightChar == right.leftChar) {
            res.prefix = left.len + right.prefix;
        }

        // By default, the suffix is the right child's suffix.
        res.suffix = right.suffix;

        // The entire right segment must be uniform before
        // the suffix can continue into the left segment.
        if (right.suffix == right.len &&
            left.rightChar == right.leftChar) {
            res.suffix = left.suffix + right.len;
        }

        // A valid repeating substring may cross the boundary.
        if (left.rightChar == right.leftChar) {
            res.best = Math.max(
                res.best,
                left.suffix + right.prefix
            );
        }

        return res;
    }

    // I build the initial segment tree.
    private void build(int node, int l, int r, char[] s) {
        // A leaf contains exactly one character.
        if (l == r) {
            tree[node].leftChar = s[l];
            tree[node].rightChar = s[l];
            tree[node].len = 1;
            tree[node].prefix = 1;
            tree[node].suffix = 1;
            tree[node].best = 1;
            return;
        }

        int mid = l + (r - l) / 2;

        // Build both children.
        build(node * 2, l, mid, s);
        build(node * 2 + 1, mid + 1, r, s);

        // Merge the children into this node.
        tree[node] = merge(tree[node * 2], tree[node * 2 + 1]);
    }

    // I update only one root-to-leaf path.
    private void update(int node, int l, int r, int idx, char c) {
        // At the target leaf, the segment contains the new character.
        if (l == r) {
            tree[node].leftChar = c;
            tree[node].rightChar = c;
            tree[node].len = 1;
            tree[node].prefix = 1;
            tree[node].suffix = 1;
            tree[node].best = 1;
            return;
        }

        int mid = l + (r - l) / 2;

        // Move into the correct half.
        if (idx <= mid) {
            update(node * 2, l, mid, idx, c);
        } else {
            update(node * 2 + 1, mid + 1, r, idx, c);
        }

        // Rebuild the current node after the child changes.
        tree[node] = merge(tree[node * 2], tree[node * 2 + 1]);
    }

    public int[] longestRepeating(String s, String queryCharacters, int[] queryIndices) {
        int n = s.length();
        int k = queryCharacters.length();

        // I allocate enough space for the segment tree.
        tree = new Node[4 * n];

        // Every tree position starts with an empty Node object.
        for (int i = 0; i < tree.length; i++) {
            tree[i] = new Node();
        }

        char[] chars = s.toCharArray();

        // Build the initial tree.
        build(1, 0, n - 1, chars);

        int[] answer = new int[k];

        for (int i = 0; i < k; i++) {
            // Apply the current point update.
            update(
                1,
                0,
                n - 1,
                queryIndices[i],
                queryCharacters.charAt(i)
            );

            // The root covers the complete string.
            answer[i] = tree[1].best;
        }

        return answer;
    }
}