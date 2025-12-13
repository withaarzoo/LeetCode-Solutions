# Coupon Code Validator (LeetCode 3606)

## Table of Contents

* Problem Summary
* Constraints
* Intuition
* Approach
* Data Structures Used
* Operations & Behavior Summary
* Complexity
* Multi-language Solutions

  * C++
  * Java
  * JavaScript
  * Python3
  * Go
* Step-by-step Detailed Explanation
* Examples
* How to use / Run locally
* Notes & Optimizations
* Author

---

## Problem Summary

I am given three arrays:

* `code[]` → coupon codes
* `businessLine[]` → category of each coupon
* `isActive[]` → whether the coupon is active

Each coupon is **valid** only if:

1. The code is **not empty**
2. The code contains **only letters, digits, or underscore**
3. The business line is one of:

   * electronics
   * grocery
   * pharmacy
   * restaurant
4. The coupon is **active**

After filtering valid coupons, I must:

* Sort them **by business line priority**
* Then sort **by code lexicographically (ASCII order)**

Finally, return **only the coupon codes**.

---

## Constraints

* `1 ≤ n ≤ 100`
* All three arrays have the same length
* Codes and business lines contain printable ASCII characters
* `isActive[i]` is either `true` or `false`

---

## Intuition

When I read the problem, I immediately realized this is a **filter + sort** problem.

My thinking was:

* First, remove all invalid coupons
* Then, define a fixed order for business categories
* Finally, sort valid coupons using that order and the coupon code

Once the rules are clear, the implementation becomes straightforward.

---

## Approach

1. I created a **priority map** for business lines

   ```
   electronics → 0
   grocery     → 1
   pharmacy    → 2
   restaurant  → 3
   ```

2. I looped through all coupons and checked validity:

   * Active status
   * Non-empty code
   * Code contains only allowed characters
   * Business line exists in the priority map

3. For every valid coupon, I stored:

   ```
   (businessPriority, couponCode)
   ```

4. I sorted the list:

   * First by business priority
   * Then by coupon code (ASCII-based)

5. I extracted only the coupon codes and returned them.

---

## Data Structures Used

* **Hash Map / Dictionary** → for business line priority
* **Array / List** → to store valid coupons
* **Tuple / Pair** → `(priority, code)` for sorting

---

## Operations & Behavior Summary

* Single pass validation → fast filtering
* Controlled sorting using fixed priority
* ASCII-based comparison to match LeetCode expectations
* Stable and predictable output

---

## Complexity

* **Time Complexity:** `O(n log n)`

  * `n` = number of coupons
  * Sorting dominates the complexity

* **Space Complexity:** `O(n)`

  * Stores valid coupons temporarily

---

## Multi-language Solutions

---

### C++

```cpp
class Solution {
public:
    vector<string> validateCoupons(vector<string>& code,
                                   vector<string>& businessLine,
                                   vector<bool>& isActive) {

        unordered_map<string, int> priority = {
            {"electronics", 0},
            {"grocery", 1},
            {"pharmacy", 2},
            {"restaurant", 3}
        };

        vector<pair<int, string>> valid;

        for (int i = 0; i < code.size(); i++) {
            if (!isActive[i]) continue;
            if (!priority.count(businessLine[i])) continue;
            if (code[i].empty()) continue;

            bool ok = true;
            for (char c : code[i]) {
                if (!isalnum(c) && c != '_') {
                    ok = false;
                    break;
                }
            }
            if (!ok) continue;

            valid.push_back({priority[businessLine[i]], code[i]});
        }

        sort(valid.begin(), valid.end());
        vector<string> res;
        for (auto &p : valid) res.push_back(p.second);
        return res;
    }
};
```

---

### Java

```java
class Solution {
    public List<String> validateCoupons(String[] code,
                                        String[] businessLine,
                                        boolean[] isActive) {

        Map<String, Integer> priority = new HashMap<>();
        priority.put("electronics", 0);
        priority.put("grocery", 1);
        priority.put("pharmacy", 2);
        priority.put("restaurant", 3);

        List<String[]> valid = new ArrayList<>();

        for (int i = 0; i < code.length; i++) {
            if (!isActive[i]) continue;
            if (!priority.containsKey(businessLine[i])) continue;
            if (code[i].isEmpty()) continue;

            boolean ok = true;
            for (char c : code[i].toCharArray()) {
                if (!Character.isLetterOrDigit(c) && c != '_') {
                    ok = false;
                    break;
                }
            }
            if (!ok) continue;

            valid.add(new String[]{businessLine[i], code[i]});
        }

        valid.sort((a, b) -> {
            int p1 = priority.get(a[0]);
            int p2 = priority.get(b[0]);
            return p1 == p2 ? a[1].compareTo(b[1]) : p1 - p2;
        });

        List<String> res = new ArrayList<>();
        for (String[] v : valid) res.add(v[1]);
        return res;
    }
}
```

---

### JavaScript

```javascript
var validateCoupons = function(code, businessLine, isActive) {

    const priority = {
        electronics: 0,
        grocery: 1,
        pharmacy: 2,
        restaurant: 3
    };

    let valid = [];

    for (let i = 0; i < code.length; i++) {
        if (!isActive[i]) continue;
        if (!(businessLine[i] in priority)) continue;
        if (!code[i]) continue;

        let ok = true;
        for (let ch of code[i]) {
            if (!(/[a-zA-Z0-9_]/.test(ch))) {
                ok = false;
                break;
            }
        }
        if (!ok) continue;

        valid.push([priority[businessLine[i]], code[i]]);
    }

    valid.sort((a, b) => {
        if (a[0] === b[0]) return a[1] < b[1] ? -1 : 1;
        return a[0] - b[0];
    });

    return valid.map(v => v[1]);
};
```

---

### Python3

```python
class Solution:
    def validateCoupons(self, code, businessLine, isActive):

        priority = {
            "electronics": 0,
            "grocery": 1,
            "pharmacy": 2,
            "restaurant": 3
        }

        valid = []

        for i in range(len(code)):
            if not isActive[i]:
                continue
            if businessLine[i] not in priority:
                continue
            if not code[i]:
                continue
            if not all(c.isalnum() or c == '_' for c in code[i]):
                continue

            valid.append((priority[businessLine[i]], code[i]))

        valid.sort()
        return [c for _, c in valid]
```

---

### Go

```go
func validateCoupons(code []string, businessLine []string, isActive []bool) []string {

	priority := map[string]int{
		"electronics": 0,
		"grocery":     1,
		"pharmacy":    2,
		"restaurant":  3,
	}

	type Pair struct {
		p int
		c string
	}

	valid := []Pair{}

	for i := 0; i < len(code); i++ {
		if !isActive[i] {
			continue
		}
		p, ok := priority[businessLine[i]]
		if !ok || len(code[i]) == 0 {
			continue
		}

		validCode := true
		for _, ch := range code[i] {
			if !(ch >= 'a' && ch <= 'z' ||
				ch >= 'A' && ch <= 'Z' ||
				ch >= '0' && ch <= '9' ||
				ch == '_') {
				validCode = false
				break
			}
		}
		if !validCode {
			continue
		}

		valid = append(valid, Pair{p, code[i]})
	}

	sort.Slice(valid, func(i, j int) bool {
		if valid[i].p == valid[j].p {
			return valid[i].c < valid[j].c
		}
		return valid[i].p < valid[j].p
	})

	res := []string{}
	for _, v := range valid {
		res = append(res, v.c)
	}
	return res
}
```

---

## Step-by-step Detailed Explanation

* I validated each coupon one by one
* Invalid coupons were skipped early
* Valid coupons were paired with their category priority
* Sorting was done using priority + ASCII order
* Final output contains only coupon codes

---

## Examples

### Example 1

**Input**

```
code = ["SAVE20", "", "PHARMA5", "SAVE@20"]
businessLine = ["restaurant", "grocery", "pharmacy", "restaurant"]
isActive = [true, true, true, true]
```

**Output**

```
["PHARMA5", "SAVE20"]
```

---

## How to use / Run locally

* Copy any language solution
* Paste into LeetCode editor or local compiler
* Run with custom test cases
* No external libraries required

---

## Notes & Optimizations

* ASCII comparison is required (not locale-based)
* Early filtering improves performance
* Priority map makes sorting clean and readable
* Works perfectly for interview and contest scenarios

---

## Author

* [Md Aarzoo Islam](https://bento.me/withaarzoo)
