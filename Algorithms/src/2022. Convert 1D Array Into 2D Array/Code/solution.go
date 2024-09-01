func construct2DArray(original []int, m int, n int) [][]int {
    // Step 1: Check if the total number of elements in the original array
    // matches the required size for the 2D array (m * n).
    // If not, return an empty 2D array because it's not possible to form
    // a valid 2D array with the given dimensions.
    if len(original) != m * n {
        return [][]int{} // Return an empty array if the sizes don't match
    }
    
    // Step 2: Initialize the 2D array 'result' with 'm' rows.
    // Each row will be a slice of integers (initially nil).
    result := make([][]int, m)
    
    // Step 3: Fill each row of the 2D array by slicing the original array.
    // Loop through each row index from 0 to m-1.
    for i := 0; i < m; i++ {
        // For the i-th row, take a slice from the original array starting at
        // index 'i*n' to index '(i+1)*n'. This slice forms the i-th row of the 2D array.
        result[i] = original[i*n : (i+1)*n]
    }
    
    // Step 4: Return the fully constructed 2D array.
    return result
}
