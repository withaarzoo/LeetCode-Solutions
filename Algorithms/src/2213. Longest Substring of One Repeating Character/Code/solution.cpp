class Solution
{
public:
    // I store every piece of information needed to merge two segments.
    struct Node
    {
        char leftChar = 0;  // First character of this segment.
        char rightChar = 0; // Last character of this segment.
        int len = 0;        // Total length of this segment.
        int prefix = 0;     // Longest same-character prefix.
        int suffix = 0;     // Longest same-character suffix.
        int best = 0;       // Longest same-character substring anywhere.
    };

    vector<Node> tree;

    // I merge two adjacent segments into one larger segment.
    Node merge(Node left, Node right)
    {
        Node res;

        // The new segment starts where the left segment starts
        // and ends where the right segment ends.
        res.leftChar = left.leftChar;
        res.rightChar = right.rightChar;
        res.len = left.len + right.len;

        // The best answer is initially inside one of the two children.
        res.best = max(left.best, right.best);

        // By default, the prefix comes from the left segment.
        res.prefix = left.prefix;

        // A prefix can extend into the right segment only when
        // the entire left segment consists of one repeated character.
        // Checking prefix == len is important because
        // prefix == suffix == best alone does NOT prove uniformity.
        if (left.prefix == left.len &&
            left.rightChar == right.leftChar)
        {
            res.prefix = left.len + right.prefix;
        }

        // By default, the suffix comes from the right segment.
        res.suffix = right.suffix;

        // A suffix can extend into the left segment only when
        // the entire right segment consists of one repeated character.
        if (right.suffix == right.len &&
            left.rightChar == right.leftChar)
        {
            res.suffix = left.suffix + right.len;
        }

        // If the boundary characters match, a repeating substring
        // can cross the boundary between the two segments.
        if (left.rightChar == right.leftChar)
        {
            res.best = max(res.best, left.suffix + right.prefix);
        }

        return res;
    }

    // I build the segment tree from the original string.
    void build(int node, int l, int r, const string &s)
    {
        // A leaf represents exactly one character.
        if (l == r)
        {
            tree[node].leftChar = s[l];
            tree[node].rightChar = s[l];
            tree[node].len = 1;
            tree[node].prefix = 1;
            tree[node].suffix = 1;
            tree[node].best = 1;
            return;
        }

        int mid = l + (r - l) / 2;

        // Build the left half.
        build(node * 2, l, mid, s);

        // Build the right half.
        build(node * 2 + 1, mid + 1, r, s);

        // Merge both children to construct this node.
        tree[node] = merge(tree[node * 2], tree[node * 2 + 1]);
    }

    // I update only the path containing the changed index.
    void update(int node, int l, int r, int idx, char c)
    {
        // Once I reach the target position, update that leaf.
        if (l == r)
        {
            tree[node].leftChar = c;
            tree[node].rightChar = c;
            tree[node].len = 1;
            tree[node].prefix = 1;
            tree[node].suffix = 1;
            tree[node].best = 1;
            return;
        }

        int mid = l + (r - l) / 2;

        // Move to the half containing the target index.
        if (idx <= mid)
        {
            update(node * 2, l, mid, idx, c);
        }
        else
        {
            update(node * 2 + 1, mid + 1, r, idx, c);
        }

        // Recalculate this parent after its child changes.
        tree[node] = merge(tree[node * 2], tree[node * 2 + 1]);
    }

    vector<int> longestRepeating(string s, string queryCharacters, vector<int> &queryIndices)
    {
        int n = s.size();
        int k = queryCharacters.size();

        // I allocate enough nodes for the segment tree.
        tree.resize(4 * n);

        // Build the tree once using the original string.
        build(1, 0, n - 1, s);

        vector<int> answer;
        answer.reserve(k);

        for (int i = 0; i < k; ++i)
        {
            // Apply the current character change at the given index.
            update(1, 0, n - 1, queryIndices[i], queryCharacters[i]);

            // The root represents the whole string,
            // so root.best is the answer after this query.
            answer.push_back(tree[1].best);
        }

        return answer;
    }
};