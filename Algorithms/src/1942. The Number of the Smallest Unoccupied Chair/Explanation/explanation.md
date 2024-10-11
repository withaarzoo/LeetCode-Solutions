# Smallest Chair Problem - Step-by-Step Explanation

This README provides a detailed breakdown of the logic and steps used in solving the "Smallest Chair" problem across multiple programming languages: C++, Java, JavaScript, Python, and Go. Each language implements the same logic, which will be described here.

---

## Problem Overview

We are given a list of arrival and leaving times for friends attending a party. Each friend will occupy the smallest available chair upon their arrival. The goal is to determine which chair will be assigned to a specific target friend.

## General Approach

1. **Input Parsing:**
   - You are provided with a list `times`, where each entry contains two elements: arrival and leaving times of each friend.
   - You are also provided with the index of the `targetFriend`, whose assigned chair needs to be found.

2. **Sorting Arrivals:**
   - First, the friends' arrival times are sorted in ascending order to process each friend in the order they arrive.

3. **Tracking Available Chairs:**
   - Use a **Min-Heap** (priority queue) to track the smallest available chairs. All chairs start off as available.

4. **Tracking Chair Usage:**
   - Use another Min-Heap to track when chairs are vacated (i.e., when a friend leaves). This is done by storing the leaving time and chair number as pairs.

5. **Assigning Chairs:**
   - For each arriving friend, first check if any chairs are vacated (i.e., their leaving time is less than or equal to the arrival time of the current friend). If so, free the chair and return it to the available pool.
   - The smallest available chair is then assigned to the friend.

6. **Target Friend's Chair:**
   - If the current arriving friend is the target friend, return the chair number assigned to them.

7. **Edge Case:**
   - The solution guarantees that the target friend will always be seated, so no need for a failure return case.

---

## C++ Code - Step-by-Step Explanation

1. **Input Parsing:**
   - Convert the `times` array into a list of pairs, where each pair contains the arrival time and the index of the friend.

2. **Sorting:**
   - Sort the list based on arrival times.

3. **Available Chairs Tracking:**
   - Use a Min-Heap (`std::priority_queue`) to keep track of available chairs.

4. **Vacating Chairs:**
   - Use another Min-Heap to track the leaving time of friends, and release chairs when a friend leaves before or during the current friend's arrival.

5. **Assigning Chairs:**
   - Assign the smallest available chair to the current friend and, if it’s the target friend, return the chair number.

6. **Returning the Result:**
   - Return the chair number assigned to the target friend.

---

## Java Code - Step-by-Step Explanation

1. **Input Parsing:**
   - Convert the `times` array into a list of arrays where each entry contains the arrival time and the index of the friend.

2. **Sorting:**
   - Sort the list by arrival time using `Comparator`.

3. **Available Chairs Tracking:**
   - Use a Min-Heap (`PriorityQueue`) to keep track of available chairs.

4. **Vacating Chairs:**
   - Another Min-Heap tracks friends who are leaving and the chair they free up.

5. **Assigning Chairs:**
   - Assign the smallest chair from the Min-Heap to the arriving friend, and return the chair number if it’s the target friend.

6. **Returning the Result:**
   - Return the chair number when the target friend arrives.

---

## JavaScript Code - Step-by-Step Explanation

1. **Input Parsing:**
   - Convert the `times` array into a list of arrays containing the arrival time and index.

2. **Sorting:**
   - Sort the list based on the arrival time using a comparison function.

3. **Tracking Available Chairs:**
   - Use a Min-Heap (`MinPriorityQueue`) to track available chairs.

4. **Vacating Chairs:**
   - Another Min-Heap tracks when chairs become available as friends leave.

5. **Assigning Chairs:**
   - Assign the smallest available chair to the current friend and return the chair number if it's the target friend.

6. **Returning the Result:**
   - Return the assigned chair when the target friend arrives.

---

## Python Code - Step-by-Step Explanation

1. **Input Parsing:**
   - Create a list of tuples containing the arrival time and index of each friend.

2. **Sorting:**
   - Sort the list of friends by arrival time.

3. **Tracking Available Chairs:**
   - Use a Min-Heap (`heapq`) to track the available chairs.

4. **Vacating Chairs:**
   - Another Min-Heap is used to track when a friend leaves, freeing up their chair.

5. **Assigning Chairs:**
   - Assign the smallest available chair to the current friend and check if this friend is the target. If so, return the chair number.

6. **Returning the Result:**
   - Return the assigned chair for the target friend.

---

## Go Code - Step-by-Step Explanation

1. **Input Parsing:**
   - Convert the `times` array into a slice containing pairs of arrival time and index.

2. **Sorting:**
   - Sort the friends based on arrival times using the `sort.Slice` function.

3. **Tracking Available Chairs:**
   - Implement a Min-Heap (`IntHeap`) to track available chairs.

4. **Vacating Chairs:**
   - Another Min-Heap (`EventHeap`) tracks when chairs are freed up as friends leave.

5. **Assigning Chairs:**
   - Assign the smallest available chair to the current friend and return the chair number if it’s the target friend.

6. **Returning the Result:**
   - Return the chair number for the target friend.
