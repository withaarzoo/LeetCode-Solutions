# Number to Words Conversion Solutions

This README provides step-by-step explanations for converting numbers to their English words representation in multiple programming languages: C++, Java, JavaScript, Python, and Go.

## C++ Solution

### Explanation

1. **Initialization**:
    - Create arrays for words representing numbers below 20, tens, and thousands.
2. **Main Function (`numberToWords`)**:
    - Check if the number is zero, return "Zero".
    - Initialize an empty string for the result.
    - Initialize an index for tracking thousands, millions, etc.
    - Process each group of three digits from the number:
        - Convert the group to words using a helper function.
        - Append the appropriate thousand's place (if applicable).
        - Divide the number by 1000 to move to the next group.
3. **Helper Function (`helper`)**:
    - Convert numbers below 20 using the `below_20` array.
    - Convert tens using the `tens` array and recurse for the ones place.
    - Convert hundreds and recurse for the remaining digits.

## Java Solution

### Explanation

1. **Initialization**:
    - Create arrays for words representing numbers below 20, tens, and thousands.
2. **Main Function (`numberToWords`)**:
    - Check if the number is zero, return "Zero".
    - Initialize an empty string for the result.
    - Initialize an index for tracking thousands, millions, etc.
    - Process each group of three digits from the number:
        - Convert the group to words using a helper function.
        - Append the appropriate thousand's place (if applicable).
        - Divide the number by 1000 to move to the next group.
3. **Helper Function (`helper`)**:
    - Convert numbers below 20 using the `below_20` array.
    - Convert tens using the `tens` array and recurse for the ones place.
    - Convert hundreds and recurse for the remaining digits.

## JavaScript Solution

### Explanation

1. **Initialization**:
    - Create arrays for words representing numbers below 20, tens, and thousands.
2. **Main Function (`numberToWords`)**:
    - Check if the number is zero, return "Zero".
    - Initialize an empty string for the result.
    - Initialize an index for tracking thousands, millions, etc.
    - Process each group of three digits from the number:
        - Convert the group to words using a helper function.
        - Append the appropriate thousand's place (if applicable).
        - Divide the number by 1000 to move to the next group.
3. **Helper Function (`helper`)**:
    - Convert numbers below 20 using the `below_20` array.
    - Convert tens using the `tens` array and recurse for the ones place.
    - Convert hundreds and recurse for the remaining digits.

## Python Solution

### Explanation

1. **Initialization**:
    - Create arrays for words representing numbers below 20, tens, and thousands.
2. **Main Function (`numberToWords`)**:
    - Check if the number is zero, return "Zero".
    - Initialize an empty string for the result.
    - Initialize an index for tracking thousands, millions, etc.
    - Process each group of three digits from the number:
        - Convert the group to words using a helper function.
        - Append the appropriate thousand's place (if applicable).
        - Divide the number by 1000 to move to the next group.
3. **Helper Function (`helper`)**:
    - Convert numbers below 20 using the `below_20` array.
    - Convert tens using the `tens` array and recurse for the ones place.
    - Convert hundreds and recurse for the remaining digits.

## Go Solution

### Explanation

1. **Initialization**:
    - Create arrays for words representing numbers below 20, tens, and thousands.
2. **Main Function (`numberToWords`)**:
    - Check if the number is zero, return "Zero".
    - Initialize an empty string for the result.
    - Initialize an index for tracking thousands, millions, etc.
    - Process each group of three digits from the number:
        - Convert the group to words using a helper function.
        - Append the appropriate thousand's place (if applicable).
        - Divide the number by 1000 to move to the next group.
3. **Helper Function (`helper`)**:
    - Convert numbers below 20 using the `below_20` array.
    - Convert tens using the `tens` array and recurse for the ones place.
    - Convert hundreds and recurse for the remaining digits.
