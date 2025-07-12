# Data Structures

The Carrion Language includes a comprehensive set of built-in data structures providing efficient implementations of common algorithms and patterns. These data structures are available through the standard library and offer both fundamental building blocks and advanced functionality.

## Node Structures

### Node
Basic node structure for linked data structures.

**Properties:**
- `data` - The value stored in the node
- `next` - Reference to the next node in sequence

**Usage:**
```carrion
node = Node("value")
print(node.data)  # "value"
```

### TreeNode
Binary tree node for tree-based data structures.

**Properties:**
- `value` - The value stored in the tree node
- `left` - Reference to left child node
- `right` - Reference to right child node

**Usage:**
```carrion
tree_node = TreeNode(10)
tree_node.left = TreeNode(5)
tree_node.right = TreeNode(15)
```

## Stack

Last-In-First-Out (LIFO) data structure implemented using linked nodes.

**Constructor:**
```carrion
stack = Stack()
```

**Methods:**

### `push(value)`
Adds an element to the top of the stack.
```carrion
stack.push(10)
stack.push(20)
```

### `pop()`
Removes and returns the top element. Returns error message if stack is empty.
```carrion
value = stack.pop()  # Returns 20
```

### `peek()`
Returns the top element without removing it. Returns error message if stack is empty.
```carrion
top = stack.peek()  # Returns top element
```

### `is_empty()`
Returns `True` if stack is empty, `False` otherwise.
```carrion
if stack.is_empty():
    print("Stack is empty")
```

### `get_size()`
Returns the number of elements in the stack.
```carrion
size = stack.get_size()
```

### `print()`
Prints all elements in the stack from top to bottom.
```carrion
stack.print()
```

## Queue

First-In-First-Out (FIFO) data structure implemented using linked nodes.

**Constructor:**
```carrion
queue = Queue()
```

**Methods:**

### `enqueue(element)`
Adds an element to the rear of the queue.
```carrion
queue.enqueue(10)
queue.enqueue(20)
```

### `dequeue()`
Removes and returns the front element. Returns error message if queue is empty.
```carrion
value = queue.dequeue()  # Returns 10
```

### `peek()`
Returns the front element without removing it. Returns error message if queue is empty.
```carrion
front = queue.peek()
```

### `is_empty()`
Returns `True` if queue is empty, `False` otherwise.
```carrion
if queue.is_empty():
    print("Queue is empty")
```

### `get_size()`
Returns the number of elements in the queue.
```carrion
size = queue.get_size()
```

### `print()`
Prints all elements in the queue from front to rear.
```carrion
queue.print()
```

## Heap

Binary heap implementation supporting both min-heap and max-heap configurations.

**Constructor:**
```carrion
min_heap = Heap()              # Default: min-heap
max_heap = Heap(False)         # Max-heap
```

**Methods:**

### `insert(value)`
Inserts a new value into the heap while maintaining heap property.
```carrion
heap.insert(10)
heap.insert(5)
heap.insert(20)
```

### `extract()`
Removes and returns the root element (min/max). Returns error message if heap is empty.
```carrion
root = heap.extract()  # Returns min value for min-heap
```

### `peek()`
Returns the root element without removing it. Returns error message if heap is empty.
```carrion
root = heap.peek()
```

### `is_empty()`
Returns `True` if heap is empty, `False` otherwise.
```carrion
if heap.is_empty():
    print("Heap is empty")
```

### `get_size()`
Returns the number of elements in the heap.
```carrion
size = heap.get_size()
```

### `clear()`
Removes all elements from the heap.
```carrion
heap.clear()
```

### `to_array()`
Returns the heap elements as an array (internal representation).
```carrion
array = heap.to_array()
```

### `print()`
Prints the heap type and current elements.
```carrion
heap.print()  # Outputs: "Min Heap: [1, 3, 2, 7, 5]"
```

### `build_heap(array)`
Builds a heap from an existing array in O(n) time.
```carrion
heap.build_heap([5, 3, 8, 1, 9, 2])
```

## Binary Search Tree (BTree)

Binary search tree implementation with standard BST operations and traversals.

**Constructor:**
```carrion
bst = BTree()
```

**Methods:**

### `insert(value)`
Inserts a value into the BST while maintaining BST property.
```carrion
bst.insert(10)
bst.insert(5)
bst.insert(15)
bst.insert(3)
```

### `size()`
Returns the total number of nodes in the tree.
```carrion
node_count = bst.size()
```

### `max_depth()`
Returns the maximum depth (height) of the tree.
```carrion
depth = bst.max_depth()
```

### `find(value)`
Searches for a value in the tree. Returns `True` if found, `False` otherwise.
```carrion
exists = bst.find(10)  # Returns True if 10 is in the tree
```

### Tree Traversals

#### `inorder()`
Returns array of values in inorder traversal (left, root, right).
```carrion
values = bst.inorder()  # Returns sorted values
```

#### `preorder()`
Returns array of values in preorder traversal (root, left, right).
```carrion
values = bst.preorder()
```

#### `postorder()`
Returns array of values in postorder traversal (left, right, root).
```carrion
values = bst.postorder()
```

#### `print_tree(traversal="inorder")`
Prints the tree using specified traversal method.
```carrion
bst.print_tree("inorder")    # Default
bst.print_tree("preorder")
bst.print_tree("postorder")
```

## Usage Examples

### Stack Example
```carrion
# Create and use a stack
stack = Stack()
stack.push(1)
stack.push(2)
stack.push(3)

print(stack.peek())      # 3
print(stack.pop())       # 3
print(stack.get_size())  # 2
stack.print()            # Prints: 2, 1
```

### Queue Example
```carrion
# Create and use a queue
queue = Queue()
queue.enqueue("first")
queue.enqueue("second")
queue.enqueue("third")

print(queue.peek())       # "first"
print(queue.dequeue())    # "first"
print(queue.get_size())   # 2
queue.print()             # Prints: second, third
```

### Heap Example
```carrion
# Min-heap example
min_heap = Heap()
min_heap.insert(20)
min_heap.insert(5)
min_heap.insert(15)
min_heap.insert(1)

print(min_heap.peek())    # 1 (minimum value)
print(min_heap.extract()) # 1
min_heap.print()          # Min Heap: [5, 20, 15]

# Max-heap example
max_heap = Heap(False)
max_heap.build_heap([3, 1, 4, 1, 5, 9, 2, 6])
print(max_heap.extract()) # 9 (maximum value)
```

### Binary Search Tree Example
```carrion
# Create and populate BST
bst = BTree()
values = [10, 5, 15, 3, 7, 12, 18]
for value in values:
    bst.insert(value)

print(bst.find(7))        # True
print(bst.size())         # 7
print(bst.max_depth())    # 3

# Print different traversals
bst.print_tree("inorder")    # Sorted: [3, 5, 7, 10, 12, 15, 18]
bst.print_tree("preorder")   # Root first: [10, 5, 3, 7, 15, 12, 18]
bst.print_tree("postorder")  # Root last: [3, 7, 5, 12, 18, 15, 10]
```

## Performance Characteristics

| Data Structure | Insert | Delete | Search | Access |
|---------------|--------|--------|--------|--------|
| Stack         | O(1)   | O(1)   | O(n)   | O(n)   |
| Queue         | O(1)   | O(1)   | O(n)   | O(n)   |
| Heap          | O(log n) | O(log n) | O(n) | O(1) peek |
| Binary Search Tree | O(log n) avg, O(n) worst | O(log n) avg, O(n) worst | O(log n) avg, O(n) worst | N/A |

## Error Handling

All data structures provide meaningful error messages for invalid operations:
- Attempting to pop/dequeue from empty structures
- Peeking at empty structures
- Invalid operations return descriptive error strings

## Integration with Carrion Language

These data structures integrate seamlessly with Carrion's:
- **Array system** - Many structures return or work with Carrion arrays
- **Error handling** - Consistent error reporting across all structures
- **Object system** - All structures are grimoires with proper spell methods
- **Type system** - Work with any Carrion data types as values