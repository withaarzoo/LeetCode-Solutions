func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
 // I store all information required to merge two segments.
 type Node struct {
  leftChar  byte // First character of the segment.
  rightChar byte // Last character of the segment.
  len       int  // Total length of the segment.
  prefix    int  // Longest same-character prefix.
  suffix    int  // Longest same-character suffix.
  best      int  // Longest same-character substring.
 }

 n := len(s)

 // I allocate enough space for the whole segment tree.
 tree := make([]Node, 4*n)

 // I merge two neighboring segments into one parent segment.
 merge := func(left, right Node) Node {
  var res Node

  // The merged segment keeps the outer boundaries of its children.
  res.leftChar = left.leftChar
  res.rightChar = right.rightChar
  res.len = left.len + right.len

  // The best answer starts as the best answer from either child.
  res.best = left.best
  if right.best > res.best {
   res.best = right.best
  }

  // By default, the prefix is the left child's prefix.
  res.prefix = left.prefix

  // The entire left segment must be uniform before
  // its prefix can continue into the right segment.
  if left.prefix == left.len &&
   left.rightChar == right.leftChar {
   res.prefix = left.len + right.prefix
  }

  // By default, the suffix is the right child's suffix.
  res.suffix = right.suffix

  // The entire right segment must be uniform before
  // its suffix can continue into the left segment.
  if right.suffix == right.len &&
   left.rightChar == right.leftChar {
   res.suffix = left.suffix + right.len
  }

  // Check whether a repeating substring crosses the boundary.
  if left.rightChar == right.leftChar {
   cross := left.suffix + right.prefix
   if cross > res.best {
    res.best = cross
   }
  }

  return res
 }

 // I build the initial tree recursively.
 var build func(int, int, int)

 build = func(node, l, r int) {
  // A leaf contains one character, so all lengths are 1.
  if l == r {
   tree[node] = Node{
    leftChar:  s[l],
    rightChar: s[l],
    len:       1,
    prefix:    1,
    suffix:    1,
    best:      1,
   }
   return
  }

  mid := (l + r) / 2

  // Build both children.
  build(node*2, l, mid)
  build(node*2+1, mid+1, r)

  // Combine them into the current node.
  tree[node] = merge(tree[node*2], tree[node*2+1])
 }

 // I update only the path containing the modified index.
 var update func(int, int, int, int, byte)

 update = func(node, l, r, idx int, c byte) {
  // At the target leaf, replace the character.
  if l == r {
   tree[node] = Node{
    leftChar:  c,
    rightChar: c,
    len:       1,
    prefix:    1,
    suffix:    1,
    best:      1,
   }
   return
  }

  mid := (l + r) / 2

  // Move into the correct half.
  if idx <= mid {
   update(node*2, l, mid, idx, c)
  } else {
   update(node*2+1, mid+1, r, idx, c)
  }

  // Recalculate this parent after its child changes.
  tree[node] = merge(tree[node*2], tree[node*2+1])
 }

 // Build the tree once.
 build(1, 0, n-1)

 answer := make([]int, len(queryCharacters))

 for i := 0; i < len(queryCharacters); i++ {
  // Apply the current character update.
  update(
   1,
   0,
   n-1,
   queryIndices[i],
   queryCharacters[i],
  )

  // The root represents the entire string.
  answer[i] = tree[1].best
 }

 return answer
}