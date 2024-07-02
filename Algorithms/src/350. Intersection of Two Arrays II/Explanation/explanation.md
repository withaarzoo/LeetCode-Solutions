# Intersection of Two Arrays II

This code implements a solution to the "Intersection of Two Arrays II" problem, which is to find the intersection of two arrays and return it as an array. The intersection is defined as the elements that appear in both arrays, with their frequency taken into account.

Here's a step-by-step explanation of the code:

## C++

1. We include the necessary headers: `vector` for dynamic arrays and `unordered_map` for hash maps.
2. We use `using namespace std;` to avoid having to write `std::` before standard library entities.
3. The `intersect` function takes two `vector<int>` parameters: `nums1` and `nums2`.
4. We create an `unordered_map<int, int>` called `countMap` to store the frequency of each element in `nums1`.
5. We create a `vector<int>` called `result` to store the intersection.
6. We iterate through `nums1` and increment the count of each element in `countMap`.
7. We iterate through `nums2`:
   - For each element `num` in `nums2`, we check if it exists in `countMap` and its count is greater than 0.
   - If the condition is true, we append `num` to `result` and decrement its count in `countMap`.
8. Finally, we return the `result` vector containing the intersection.

## Java

1. We import the necessary classes: `ArrayList` for dynamic arrays, `HashMap` for hash maps, and `List` for lists.
2. The `intersect` function takes two `int[]` parameters: `nums1` and `nums2`.
3. We create a `HashMap<Integer, Integer>` called `countMap` to store the frequency of each element in `nums1`.
4. We create an `ArrayList` called `result` to store the intersection.
5. We iterate through `nums1` and put each element `num` in `countMap` with its count.
6. We iterate through `nums2`:
   - For each element `num` in `nums2`, we check if it exists in `countMap` and its count is greater than 0.
   - If the condition is true, we append `num` to `result` and decrement its count in `countMap`.
7. Finally, we convert `result` to an `int[]` array and return it.

## JavaScript

1. We define a function `intersect` that takes two parameters: `nums1` and `nums2`.
2. We create a `Map` called `countMap` to store the frequency of each element in `nums1`.
3. We create an array called `result` to store the intersection.
4. We iterate through `nums1` and put each element `num` in `countMap` with its count.
5. We iterate through `nums2`:
   - For each element `num` in `nums2`, we check if it exists in `countMap` and its count is greater than 0.
   - If the condition is true, we append `num` to `result` and decrement its count in `countMap`.
6. Finally, we return the `result` array containing the intersection.

## Python

1. We import `Counter` from the `collections` module to use a frequency map.
2. We define a class `Solution` with a method `intersect` that takes two parameters: `nums1` and `nums2`.
3. We create a `Counter` called `countMap` to store the frequency of each element in `nums1`.
4. We create an array called `result` to store the intersection.
5. We iterate through `nums2`:
   - For each element `num` in `nums2`, we check if its count in `countMap` is greater than 0.
   - If the condition is true, we append `num` to `result` and decrement its count in `countMap`.
6. Finally, we return the `result` array containing the intersection.

## Go

1. We define a function `intersect` that takes two parameters: `nums1` and `nums2`.
2. We create a `map[int]int` called `countMap` to store the frequency of each element in `nums1`.
3. We create a slice called `result` to store the intersection.
4. We iterate through `nums1` and increment the count of each element in `countMap`.
5. We iterate through `nums2`:
   - For each element `num` in `nums2`, we check if its count in `countMap` is greater than 0.
   - If the condition is true, we append `num` to `result` and decrement its count in `countMap`.
6. Finally, we return the `result` slice containing the intersection.

The time complexity of this solution is O(m+n), where m and n are the lengths of `nums1` and `nums2`, respectively. The space complexity is O(min(m,n)) since we use a hash map to store the frequency of elements from the smaller array.
