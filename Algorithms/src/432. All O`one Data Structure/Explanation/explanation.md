# AllOne Data Structure: Step-by-Step Explanation

The `AllOne` class implements a data structure that can efficiently track strings and their associated counts. It supports the following operations:

- `inc(key)`: Increment the count of the key.
- `dec(key)`: Decrement the count of the key.
- `getMaxKey()`: Get the key with the maximum count.
- `getMinKey()`: Get the key with the minimum count.

Below is a step-by-step breakdown of the implementation in various programming languages: C++, Java, JavaScript, Python, and Go.

---

### C++ Implementation

1. **Data Structures**:
   - A hash map (`unordered_map`) stores the frequency of each key.
   - A set (`set<pair<int, string>>`) maintains the keys ordered by their counts.

2. **Constructor**:
   - Initializes the frequency map and clears any previous data.

3. **Increment (`inc`)**:
   - Fetches the current frequency of the key.
   - Updates the frequency and removes the old value from the set.
   - Inserts the new frequency and key into the set to keep it ordered.

4. **Decrement (`dec`)**:
   - Fetches the current frequency.
   - Decreases the frequency in the map and removes the corresponding pair from the set.
   - If the count drops to zero, the key is removed entirely.

5. **Get Maximum Key (`getMaxKey`)**:
   - Returns the key with the highest frequency from the set.

6. **Get Minimum Key (`getMinKey`)**:
   - Returns the key with the lowest frequency from the set.

---

### Java Implementation

1. **Data Structures**:
   - A `HashMap` is used to store the frequency of each key.
   - A `TreeSet` maintains the keys in sorted order based on their frequencies and lexicographical order of keys.

2. **Constructor**:
   - Initializes the frequency map and a `TreeSet` for ordering.

3. **Increment (`inc`)**:
   - Retrieves the current frequency from the map.
   - Removes the current key from the set with the old frequency and inserts it with the updated frequency.

4. **Decrement (`dec`)**:
   - Decreases the frequency or removes the key if it reaches zero.
   - Updates the `TreeSet` to reflect the change in frequency.

5. **Get Maximum Key (`getMaxKey`)**:
   - Retrieves the key with the maximum frequency from the last position in the `TreeSet`.

6. **Get Minimum Key (`getMinKey`)**:
   - Retrieves the key with the minimum frequency from the first position in the `TreeSet`.

---

### JavaScript Implementation

1. **Data Structures**:
   - A `Map` stores the current frequencies of keys.
   - A custom double-linked list structure (`dll`) organizes nodes based on frequencies, allowing efficient access.

2. **Constructor**:
   - Initializes an empty map and double-linked list.

3. **Add to Linked List (`addtodll`)**:
   - Adds a key to the double-linked list.
   - Creates a new node if needed or updates an existing node based on the key’s frequency.

4. **Increment (`inc`)**:
   - Uses `addtodll` to increment the key’s frequency and updates the node in the linked list.

5. **Remove from Linked List (`removefromdll`)**:
   - Removes the key from the linked list or decreases its frequency.
   - Deletes the node if it becomes empty.

6. **Decrement (`dec`)**:
   - Uses `removefromdll` to decrement the key’s frequency and updates the linked list.

7. **Get Maximum Key (`getMaxKey`)**:
   - Retrieves the last key in the double-linked list for the highest frequency.

8. **Get Minimum Key (`getMinKey`)**:
   - Retrieves the first key in the double-linked list for the lowest frequency.

---

### Python Implementation

1. **Data Structures**:
   - A simple dictionary stores the frequency of each key.

2. **Constructor**:
   - Initializes an empty dictionary to track frequencies.

3. **Increment (`inc`)**:
   - Increases the frequency of the key or sets it to 1 if it's new.

4. **Decrement (`dec`)**:
   - Decreases the frequency or removes the key if its frequency reaches zero.

5. **Get Maximum Key (`getMaxKey`)**:
   - Iterates through the dictionary to find the key with the highest frequency.

6. **Get Minimum Key (`getMinKey`)**:
   - Iterates through the dictionary to find the key with the lowest frequency.

---

### Go Implementation

1. **Data Structures**:
   - A `map` stores the frequency of each key.
   - A `set` (`map[int]map[string]bool`) organizes keys by their frequencies.

2. **Constructor**:
   - Initializes the frequency map and a set to track the order of keys.

3. **Increment (`Inc`)**:
   - Fetches the current frequency from the map and updates the set accordingly.
   - Removes the key from the old frequency bucket and adds it to the new bucket.

4. **Decrement (`Dec`)**:
   - Decreases the key’s frequency or removes it if its frequency reaches zero.
   - Updates the set to reflect the change.

5. **Get Maximum Key (`GetMaxKey`)**:
   - Iterates through the set in reverse to retrieve the key with the highest frequency.

6. **Get Minimum Key (`GetMinKey`)**:
   - Iterates through the set to retrieve the key with the lowest frequency.

---

By using different data structures and methodologies, each language optimizes the core operations of the `AllOne` data structure, making it efficient in terms of both time and space complexity. The use of sets and maps allows for quick updates and lookups, while linked lists or trees ensure that the elements are properly ordered for fast access to the minimum and maximum keys.
