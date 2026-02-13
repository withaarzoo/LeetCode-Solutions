# Problem Title

1. Longest Balanced Substring II

---

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

You are given a string `s` containing only characters 'a', 'b', and 'c'.

A substring is called **balanced** if all distinct characters in that substring appear the same number of times.

Return the length of the longest balanced substring.

---

## Constraints

* 1 <= s.length <= 10^5
* s contains only 'a', 'b', and 'c'

---

## Intuition

When I first read the problem, I understood that a substring is balanced if all characters inside it appear equal number of times.

So I divided the problem into three possible cases:

1. Substring has only one distinct character.
2. Substring has exactly two distinct characters.
3. Substring has all three characters.

If I can handle all three cases efficiently, I just take the maximum length.

To compare character frequencies quickly inside substrings, I decided to use **prefix counts**.

---

## Approach

1. Maintain prefix counts of 'a', 'b', and 'c'.
2. Track the longest single-character run.
3. Use hash maps to store earliest prefix states.
4. Use differences of counts to detect balanced substrings.
5. Update maximum length whenever I find a valid balanced substring.

Key idea:
If two prefix states have the same frequency difference pattern, then the substring between them is balanced.

---

## Data Structures Used

* Integer counters for prefix counts.
* HashMap / unordered_map / Map / dictionary for storing prefix states.
* Tuple or string key to store count differences.

---

## Operations & Behavior Summary

For each index:

* Update prefix counts.
* Generate keys using count differences.
* If key seen before → update max length.
* Else store earliest occurrence.

---

## Complexity

**Time Complexity:** O(n)

Each character is processed once. Hash operations are constant time on average.

**Space Complexity:** O(n)

We store prefix states in hash maps.

---

## Multi-language Solutions

### C++

```cpp
#include <bits/stdc++.h>
using namespace std;

struct PairHash {
    size_t operator()(const pair<int,int>&p) const noexcept {
        return ((uint64_t)(p.first) << 32) ^ (uint32_t)(p.second);
    }
};

class Solution {
public:
    int longestBalanced(string s) {
        int n = s.size();
        int a=0,b=0,c=0;
        int ans=0;
        
        int run=0;
        char prev=0;
        for(int i=0;i<n;i++){
            if(i==0 || s[i]!=prev) run=1;
            else run++;
            prev=s[i];
            ans=max(ans,run);
        }

        unordered_map<pair<int,int>,int,PairHash> map3;
        unordered_map<pair<int,int>,int,PairHash> map_ab_c;
        unordered_map<pair<int,int>,int,PairHash> map_ac_b;
        unordered_map<pair<int,int>,int,PairHash> map_bc_a;

        map3[{0,0}]=0;
        map_ab_c[{0,0}]=0;
        map_ac_b[{0,0}]=0;
        map_bc_a[{0,0}]=0;

        for(int p=1;p<=n;p++){
            char ch=s[p-1];
            if(ch=='a') a++;
            else if(ch=='b') b++;
            else c++;

            pair<int,int> k3={b-a,c-a};
            if(map3.count(k3)) ans=max(ans,p-map3[k3]);
            else map3[k3]=p;

            pair<int,int> kabc={b-a,c};
            if(map_ab_c.count(kabc)) ans=max(ans,p-map_ab_c[kabc]);
            else map_ab_c[kabc]=p;

            pair<int,int> kacb={c-a,b};
            if(map_ac_b.count(kacb)) ans=max(ans,p-map_ac_b[kacb]);
            else map_ac_b[kacb]=p;

            pair<int,int> kbc={c-b,a};
            if(map_bc_a.count(kbc)) ans=max(ans,p-map_bc_a[kbc]);
            else map_bc_a[kbc]=p;
        }
        return ans;
    }
};
```

---

### Java

```java
import java.util.*;

class Solution {
    private String key(int x,int y){
        return x+"#"+y;
    }
    public int longestBalanced(String s) {
        int n=s.length();
        int a=0,b=0,c=0;
        int ans=0;

        int run=0;
        char prev=0;
        for(int i=0;i<n;i++){
            if(i==0||s.charAt(i)!=prev) run=1;
            else run++;
            prev=s.charAt(i);
            ans=Math.max(ans,run);
        }

        Map<String,Integer> map3=new HashMap<>();
        Map<String,Integer> map_ab_c=new HashMap<>();
        Map<String,Integer> map_ac_b=new HashMap<>();
        Map<String,Integer> map_bc_a=new HashMap<>();

        map3.put(key(0,0),0);
        map_ab_c.put(key(0,0),0);
        map_ac_b.put(key(0,0),0);
        map_bc_a.put(key(0,0),0);

        for(int p=1;p<=n;p++){
            char ch=s.charAt(p-1);
            if(ch=='a') a++;
            else if(ch=='b') b++;
            else c++;

            String k3=key(b-a,c-a);
            if(map3.containsKey(k3)) ans=Math.max(ans,p-map3.get(k3));
            else map3.put(k3,p);

            String kabc=key(b-a,c);
            if(map_ab_c.containsKey(kabc)) ans=Math.max(ans,p-map_ab_c.get(kabc));
            else map_ab_c.put(kabc,p);

            String kacb=key(c-a,b);
            if(map_ac_b.containsKey(kacb)) ans=Math.max(ans,p-map_ac_b.get(kacb));
            else map_ac_b.put(kacb,p);

            String kbc=key(c-b,a);
            if(map_bc_a.containsKey(kbc)) ans=Math.max(ans,p-map_bc_a.get(kbc));
            else map_bc_a.put(kbc,p);
        }
        return ans;
    }
}
```

---

### JavaScript

```javascript
var longestBalanced = function(s) {
    let n=s.length;
    let a=0,b=0,c=0;
    let ans=0;

    let run=0,prev='';
    for(let i=0;i<n;i++){
        if(i===0||s[i]!==prev) run=1;
        else run++;
        prev=s[i];
        ans=Math.max(ans,run);
    }

    const map3=new Map();
    const map_ab_c=new Map();
    const map_ac_b=new Map();
    const map_bc_a=new Map();

    function key(x,y){return x+'#'+y;}

    map3.set(key(0,0),0);
    map_ab_c.set(key(0,0),0);
    map_ac_b.set(key(0,0),0);
    map_bc_a.set(key(0,0),0);

    for(let p=1;p<=n;p++){
        let ch=s[p-1];
        if(ch==='a') a++;
        else if(ch==='b') b++;
        else c++;

        let k3=key(b-a,c-a);
        if(map3.has(k3)) ans=Math.max(ans,p-map3.get(k3));
        else map3.set(k3,p);

        let kabc=key(b-a,c);
        if(map_ab_c.has(kabc)) ans=Math.max(ans,p-map_ab_c.get(kabc));
        else map_ab_c.set(kabc,p);

        let kacb=key(c-a,b);
        if(map_ac_b.has(kacb)) ans=Math.max(ans,p-map_ac_b.get(kacb));
        else map_ac_b.set(kacb,p);

        let kbc=key(c-b,a);
        if(map_bc_a.has(kbc)) ans=Math.max(ans,p-map_bc_a.get(kbc));
        else map_bc_a.set(kbc,p);
    }
    return ans;
};
```

---

### Python3

```python
class Solution:
    def longestBalanced(self, s: str) -> int:
        n=len(s)
        a=b=c=0
        ans=0

        run=0
        prev=''
        for i,ch in enumerate(s):
            if i==0 or ch!=prev:
                run=1
            else:
                run+=1
            prev=ch
            ans=max(ans,run)

        map3={(0,0):0}
        map_ab_c={(0,0):0}
        map_ac_b={(0,0):0}
        map_bc_a={(0,0):0}

        for p in range(1,n+1):
            ch=s[p-1]
            if ch=='a': a+=1
            elif ch=='b': b+=1
            else: c+=1

            k3=(b-a,c-a)
            if k3 in map3:
                ans=max(ans,p-map3[k3])
            else:
                map3[k3]=p

            kabc=(b-a,c)
            if kabc in map_ab_c:
                ans=max(ans,p-map_ab_c[kabc])
            else:
                map_ab_c[kabc]=p

            kacb=(c-a,b)
            if kacb in map_ac_b:
                ans=max(ans,p-map_ac_b[kacb])
            else:
                map_ac_b[kacb]=p

            kbc=(c-b,a)
            if kbc in map_bc_a:
                ans=max(ans,p-map_bc_a[kbc])
            else:
                map_bc_a[kbc]=p

        return ans
```

---

### Go

```go
func longestBalanced(s string) int {
    n:=len(s)
    a,b,c:=0,0,0
    ans:=0

    run:=0
    var prev byte
    for i:=0;i<n;i++{
        if i==0||s[i]!=prev{
            run=1
        }else{
            run++
        }
        prev=s[i]
        if run>ans{ans=run}
    }

    map3:=map[string]int{"0#0":0}
    map_ab_c:=map[string]int{"0#0":0}
    map_ac_b:=map[string]int{"0#0":0}
    map_bc_a:=map[string]int{"0#0":0}

    key:=func(x,y int) string{
        return fmt.Sprintf("%d#%d",x,y)
    }

    for p:=1;p<=n;p++{
        ch:=s[p-1]
        if ch=='a'{a++}else if ch=='b'{b++}else{c++}

        k3:=key(b-a,c-a)
        if v,ok:=map3[k3];ok{if p-v>ans{ans=p-v}}else{map3[k3]=p}

        kab:=key(b-a,c)
        if v,ok:=map_ab_c[kab];ok{if p-v>ans{ans=p-v}}else{map_ab_c[kab]=p}

        kac:=key(c-a,b)
        if v,ok:=map_ac_b[kac];ok{if p-v>ans{ans=p-v}}else{map_ac_b[kac]=p}

        kbc:=key(c-b,a)
        if v,ok:=map_bc_a[kbc];ok{if p-v>ans{ans=p-v}}else{map_bc_a[kbc]=p}
    }
    return ans
}
```

---

## Step-by-step Detailed Explanation (C++, Java, JavaScript, Python3, Go)

The logic is the same in every language.

1. Initialize counters for a, b, c.
2. Track longest single-character run.
3. Store prefix states in maps.
4. For every prefix, compute difference keys.
5. If key exists, update answer.
6. Otherwise, store earliest index.

This ensures we find longest valid substring in O(n).

---

## Examples

Input: "abbac"
Output: 4

Input: "aabcc"
Output: 3

Input: "aba"
Output: 2

---

## How to use / Run locally

1. Copy the desired language solution.
2. Paste into your IDE or LeetCode editor.
3. Compile and run with sample test cases.

---

## Notes & Optimizations

* Prefix difference hashing is the key optimization.
* Always store earliest index to maximize substring length.
* Works efficiently up to 10^5 length.

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
