package main

import (
    "sort"
)

// survivedRobotsHealths calculates the remaining health of robots after potential collisions
func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
    n := len(positions) // Number of robots
    indices := make([]int, n) // Array to store original indices of robots
    stack := []int{} // Stack to store indices of robots moving to the right
    result := []int{} // Array to store the health of robots that survive

    // Initialize the indices array with values from 0 to n-1
    for i := range indices {
        indices[i] = i
    }

    // Sort the indices based on the positions of the robots
    sort.Slice(indices, func(a, b int) bool {
        return positions[a] < positions[b]
    })

    // Process each robot in the sorted order of their positions
    for _, currentIndex := range indices {
        if directions[currentIndex] == 'R' {
            // If the current robot is moving to the right, push its index onto the stack
            stack = append(stack, currentIndex)
        } else {
            // If the current robot is moving to the left
            for len(stack) > 0 && healths[currentIndex] > 0 {
                // While there are robots in the stack and the current robot's health is greater than 0
                topIndex := stack[len(stack)-1] // Get the index of the robot at the top of the stack
                stack = stack[:len(stack)-1] // Pop the top robot from the stack

                if healths[topIndex] > healths[currentIndex] {
                    // If the robot from the stack has more health than the current robot
                    healths[topIndex] -= 1 // Decrease the health of the robot from the stack by 1
                    healths[currentIndex] = 0 // The current robot's health becomes 0 (it is destroyed)
                    stack = append(stack, topIndex) // Push the robot from the stack back onto the stack
                } else if healths[topIndex] < healths[currentIndex] {
                    // If the robot from the stack has less health than the current robot
                    healths[currentIndex] -= 1 // Decrease the health of the current robot by 1
                    healths[topIndex] = 0 // The robot from the stack's health becomes 0 (it is destroyed)
                } else {
                    // If both robots have equal health
                    healths[currentIndex] = 0 // Both robots' health becomes 0 (both are destroyed)
                    healths[topIndex] = 0 // Both robots' health becomes 0 (both are destroyed)
                }
            }
        }
    }

    // Collect the healths of the robots that survived (health > 0)
    for i := 0; i < n; i++ {
        if healths[i] > 0 {
            result = append(result, healths[i])
        }
    }

    return result // Return the healths of the surviving robots
}
