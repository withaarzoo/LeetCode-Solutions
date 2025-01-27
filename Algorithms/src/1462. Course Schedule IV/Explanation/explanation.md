# Problem: Check If a Prerequisite Exists

This repository provides solutions to the problem "Check If a Prerequisite Exists" implemented in **C++**, **Java**, **JavaScript**, **Python**, and **Go**. Below, you'll find step-by-step explanations for each language.

---

## C++ Code Implementation: Step-by-Step

1. **Initialize the Adjacency Matrix**  
   Start by initializing a 2D matrix to store whether a course is a prerequisite of another course.

2. **Set up the Prerequisite Relationships**  
   Using the prerequisites list, update the adjacency matrix with direct relationships between courses.

3. **Floyd-Warshall Algorithm**  
   Apply the Floyd-Warshall algorithm to compute all pairs' reachability. This updates the matrix to reflect indirect prerequisites.

4. **Process the Queries**  
   For each query, check the adjacency matrix to see if one course is a prerequisite of another.

---

## Java Code Implementation: Step-by-Step

1. **Prepare the Adjacency Matrix**  
   Create a 2D array to store whether a course is a prerequisite for another.

2. **Update Relationships from Prerequisites**  
   Populate the matrix with the prerequisites directly given in the input.

3. **Use Floyd-Warshall Algorithm**  
   Implement Floyd-Warshall to ensure all transitive prerequisites are captured in the matrix.

4. **Answer the Queries**  
   For each query, use the matrix to determine if the prerequisite relationship exists.

---

## JavaScript Code Implementation: Step-by-Step

1. **Build the Graph**  
   Create a graph representation using a 2D array to represent the prerequisite relationships.

2. **Populate Direct Relationships**  
   Populate the graph with direct prerequisites using the input data.

3. **Compute Transitive Closure**  
   Utilize the Floyd-Warshall algorithm to compute transitive relationships between courses.

4. **Evaluate Queries**  
   Loop through each query and return whether the course relationship exists.

---

## Python Code Implementation: Step-by-Step

1. **Initialize a 2D Matrix**  
   Create a matrix where `matrix[i][j]` indicates whether course `i` is a prerequisite for course `j`.

2. **Update the Matrix for Direct Prerequisites**  
   Populate the matrix with the relationships provided in the prerequisites list.

3. **Apply Floyd-Warshall Algorithm**  
   Implement the Floyd-Warshall algorithm to find all indirect relationships between courses.

4. **Answer the Queries**  
   For each query, check the value in the matrix to determine if the prerequisite exists.

---

## Go Code Implementation: Step-by-Step

1. **Set up the Graph as a 2D Slice**  
   Initialize a 2D slice to represent the adjacency matrix of course relationships.

2. **Add Direct Relationships**  
   Update the graph with the direct prerequisites based on the input data.

3. **Floyd-Warshall Algorithm for Transitive Closure**  
   Implement Floyd-Warshall to propagate indirect prerequisites across the matrix.

4. **Handle Queries**  
   For each query, determine whether a course is a prerequisite by checking the matrix.

---

Each implementation shares a common approach:

1. Build a representation of prerequisites.  
2. Compute all transitive relationships using the Floyd-Warshall algorithm.  
3. Check the matrix to answer queries.

For the complete implementation, refer to the respective files:  

- `solution.cpp`  
- `solution.java`  
- `solution.js`  
- `solution.py`  
- `solution.go`

---

### Complexity Analysis

- **Time Complexity**:  
  The Floyd-Warshall algorithm takes $$O(n^3)$$ time, where \(n\) is the number of courses. Processing the queries takes $$O(q)$$, where \(q\) is the number of queries.  

- **Space Complexity**:  
  The space complexity is $$O(n^2)$$ for the adjacency matrix.  
