# Concurrency in Carrion Language

The Carrion language provides built-in concurrency support through two primary keywords: `diverge` and `converge`. These primitives allow you to create and manage concurrent goroutines with a simple, intuitive syntax.

## Table of Contents

1. [Overview](#overview)
2. [The `diverge` Keyword](#the-diverge-keyword)
3. [The `converge` Keyword](#the-converge-keyword)
4. [Goroutine Management](#goroutine-management)
5. [Examples](#examples)
6. [Best Practices](#best-practices)
7. [Technical Details](#technical-details)
8. [Error Handling](#error-handling)

## Overview

Carrion's concurrency model is built on top of Go's goroutines, providing a higher-level abstraction that integrates seamlessly with the language's syntax and object system. The concurrency primitives are designed to be:

- **Simple**: Easy to understand and use
- **Safe**: Proper resource management and cleanup
- **Flexible**: Support for both named and anonymous goroutines
- **Efficient**: Built on proven Go concurrency patterns

## The `diverge` Keyword

The `diverge` keyword creates a new goroutine that executes code concurrently with the main program flow.

### Syntax

```carrion
# Anonymous goroutine
diverge:
    # code block

# Named goroutine
diverge name:
    # code block
```

### How it Works

1. **Creates a goroutine object** with a synchronization channel
2. **Launches a Go goroutine** with proper error handling and cleanup
3. **Returns immediately** while the goroutine runs in the background
4. **Isolates variables** using an enclosed environment

### Anonymous Goroutines

```carrion
diverge:
    print("Hello from anonymous goroutine")
    sleep(1000)  # Sleep for 1 second
    print("Anonymous goroutine finished")
```

### Named Goroutines

```carrion
diverge worker:
    print("Starting worker goroutine")
    for i in range(5):
        print("Worker step:", i)
        sleep(500)
    print("Worker completed")
```

## The `converge` Keyword

The `converge` keyword waits for goroutines to complete their execution.

### Syntax

```carrion
# Wait for all goroutines
converge

# Wait for specific named goroutines
converge name1, name2, name3
```

### How it Works

1. **Waits for completion** by reading from goroutine Done channels
2. **Blocks execution** until specified goroutines finish
3. **Cleans up resources** automatically after completion
4. **Validates goroutine names** and reports errors if not found

### Wait for All Goroutines

```carrion
diverge:
    print("First goroutine")
    
diverge:
    print("Second goroutine")
    
converge  # Waits for both goroutines to complete
print("All goroutines finished")
```

### Wait for Specific Goroutines

```carrion
diverge worker1:
    print("Worker 1 starting")
    sleep(1000)
    print("Worker 1 done")
    
diverge worker2:
    print("Worker 2 starting")
    sleep(2000)
    print("Worker 2 done")
    
converge worker1  # Wait only for worker1
print("Worker 1 completed, worker2 may still be running")

converge worker2  # Wait for worker2
print("All workers completed")
```

## Goroutine Management

Carrion uses a global `GoroutineManager` that provides:

- **Thread-safe operations** with mutex protection
- **Automatic cleanup** of completed goroutines
- **Resource limits** to prevent memory exhaustion
- **Named and anonymous** goroutine storage

### Resource Management

The goroutine manager includes configurable limits:

```carrion
# The manager automatically cleans up completed goroutines
# Maximum limits can be configured to prevent memory issues
# Default: unlimited with automatic cleanup enabled
```

## Examples

### Example 1: Basic Concurrent Processing

```carrion
main:
    print("Starting concurrent processing")
    
    diverge task1:
        print("Task 1: Processing data...")
        sleep(2000)
        print("Task 1: Complete")
    
    diverge task2:
        print("Task 2: Calculating results...")
        sleep(1500)
        print("Task 2: Complete")
    
    print("Both tasks started, doing other work...")
    sleep(500)
    print("Main thread work done, waiting for tasks...")
    
    converge task1, task2
    print("All tasks completed!")
```

### Example 2: Producer-Consumer Pattern

```carrion
main:
    print("Starting producer-consumer example")
    
    # Producer goroutine
    diverge producer:
        for i in range(10):
            print("Producing item", i)
            sleep(200)
        print("Producer finished")
    
    # Consumer goroutine
    diverge consumer:
        for i in range(10):
            print("Consuming item", i)
            sleep(300)
        print("Consumer finished")
    
    # Let them run
    converge producer, consumer
    print("Production and consumption complete")
```

### Example 3: Parallel Computation

```carrion
main:
    print("Parallel computation example")
    
    # Launch multiple workers
    for i in range(4):
        diverge:
            worker_id = i
            print("Worker", worker_id, "starting")
            
            # Simulate work
            total = 0
            for j in range(1000000):
                total += j
                
            print("Worker", worker_id, "result:", total)
    
    # Wait for all workers to complete
    converge
    print("All workers completed")
```

### Example 4: Sequential Goroutine Processing

```carrion
main:
    print("Sequential processing with goroutines")
    
    # Stage 1
    diverge stage1:
        print("Stage 1: Data collection")
        sleep(1000)
        print("Stage 1: Complete")
    
    converge stage1
    print("Stage 1 finished, starting stage 2")
    
    # Evaluation after converge works correctly
    stage1_result = "data_collected"
    
    # Stage 2
    diverge stage2:
        print("Stage 2: Data processing")
        sleep(1500)
        print("Stage 2: Complete")
    
    converge stage2
    print("Stage 2 finished, starting stage 3")
    
    # Multiple sequential converge operations work reliably
    stage2_result = stage1_result + "_processed"
    
    # Stage 3
    diverge stage3:
        print("Stage 3: Data analysis")
        sleep(800)
        print("Stage 3: Complete")
    
    converge stage3
    print("All stages completed!")
    
    # Final evaluation after all stages
    final_result = stage2_result + "_analyzed"
    print("Pipeline result:", final_result)
```

### Example 5: Multiple Workers with Individual Convergence

```carrion
main:
    print("Testing multiple workers with individual convergence")
    
    # Start multiple workers
    diverge worker_a:
        print("Worker A: Processing...")
        sleep(500)
        print("Worker A: Done")
    
    diverge worker_b:
        print("Worker B: Processing...")
        sleep(300)
        print("Worker B: Done")
    
    diverge worker_c:
        print("Worker C: Processing...")
        sleep(200)
        print("Worker C: Done")
    
    # Converge workers individually as they complete
    converge worker_c  # Fastest worker
    print("Worker C completed, evaluation works")
    result_c = 100
    
    converge worker_b  # Medium speed worker
    print("Worker B completed, evaluation works")
    result_b = result_c + 50
    
    converge worker_a  # Slowest worker
    print("Worker A completed, evaluation works")
    result_a = result_b + result_c
    
    print("Final results:", result_a, result_b, result_c)
```

## Best Practices

### 1. Use Named Goroutines for Complex Logic

```carrion
# Good: Named goroutines for clarity
diverge database_worker:
    # Database operations
    
diverge file_processor:
    # File processing
    
converge database_worker, file_processor
```

### 2. Avoid Long-Running Anonymous Goroutines

```carrion
# Avoid: Hard to track and manage
diverge:
    while True:
        # Long-running loop
        
# Better: Use named goroutines for long-running tasks
diverge background_service:
    while True:
        # Long-running loop
```

### 3. Handle Errors Appropriately

```carrion
# Goroutines should handle their own errors
diverge worker:
    attempt:
        # Risky operation
        process_data()
    ensnare error:
        print("Worker error:", error)
        # Handle error appropriately
```

### 4. Use Converge Strategically

```carrion
# Wait for critical goroutines before proceeding
diverge critical_task:
    # Important work
    
diverge optional_task:
    # Optional work
    
converge critical_task  # Wait for critical work
print("Critical work done, continuing...")

# Code evaluation works correctly after converge
important_result = process_critical_data()

# Optional: wait for remaining work
converge optional_task
final_cleanup()
```

### 5. Sequential Operations Work Reliably

```carrion
# Multiple sequential diverge/converge operations are fully supported
for i in range(5):
    worker_name = "worker" + str(i)
    
    diverge worker_name:
        print("Worker", i, "processing...")
        # Do work
        
    converge worker_name
    print("Worker", i, "completed")
    
    # Evaluation after each converge works correctly
    result = i * 10
    process_result(result)
```

## Technical Details

### Goroutine Object Structure

Each goroutine is represented by a `Goroutine` object containing:

- **Name**: Optional identifier for the goroutine
- **Done**: Buffered channel for synchronization (capacity 1)
- **Result**: Execution result (if any)
- **Error**: Error object if execution failed
- **IsRunning**: Current execution state
- **cleaned**: Cleanup status flag to prevent double cleanup

### Resource Management Improvements

As of the latest version, the concurrency system includes several important improvements:

- **Proper cleanup**: Named goroutines are cleaned up using `RemoveAndCleanupNamed()` which ensures channels are properly closed and resources released
- **Race condition protection**: Converge operations use `select` statements to handle timing issues where goroutines complete between checking `IsRunning` and reading from the `Done` channel
- **Automatic cleanup**: The goroutine manager automatically cleans up completed goroutines when adding new ones (if `AutoCleanup` is enabled)
- **Thread-safe operations**: All goroutine manager operations are protected by mutexes

### Environment Isolation

Each goroutine gets its own enclosed environment:

```carrion
x = 10

diverge:
    x = 20  # This doesn't affect the main thread's x
    print("Goroutine x:", x)  # Prints: 20

print("Main x:", x)  # Prints: 10
```

### Resource Limits

The goroutine manager supports configurable limits:

- **MaxNamedSize**: Maximum number of named goroutines
- **MaxAnonymousSize**: Maximum number of anonymous goroutines
- **AutoCleanup**: Automatic cleanup of completed goroutines

### Synchronization

- Each goroutine has a buffered `Done` channel
- `converge` blocks on channel reads
- Automatic signaling on completion
- Thread-safe manager operations

## Error Handling

### Goroutine Errors

```carrion
diverge error_prone:
    # This will cause an error
    result = 10 / 0  # Division by zero
    
converge error_prone
# The error is contained within the goroutine
print("Main thread continues normally")
```

### Validation Errors

```carrion
diverge worker:
    print("Worker running")
    
# This will cause an error - goroutine not found
converge nonexistent_worker  # Error: goroutine 'nonexistent_worker' not found
```

### Cleanup and Recovery

The system automatically:

- Recovers from panics within goroutines
- Cleans up resources when goroutines complete
- Manages memory to prevent leaks
- Provides error tracing and reporting

## Troubleshooting

### Common Issues and Solutions

#### Sequential Converge Operations

**Issue**: Evaluation not working after multiple sequential converge operations.

**Solution**: This has been fixed in the latest version. The system now properly cleans up goroutine resources and handles race conditions that could affect subsequent operations.

#### Resource Leaks

**Issue**: Long-running programs with many goroutines consuming increasing memory.

**Solution**: The goroutine manager now includes automatic cleanup (enabled by default) and proper resource management. Named goroutines are cleaned up when converged.

#### Race Conditions

**Issue**: Timing issues where converge operations occasionally fail.

**Solution**: The converge logic now uses `select` statements to handle cases where goroutines complete between checking `IsRunning` and reading from the `Done` channel.

## Recent Improvements

The concurrency system has been enhanced with the following improvements:

- **Fixed resource cleanup**: Named goroutines now use `RemoveAndCleanupNamed()` for proper resource management
- **Race condition protection**: Converge operations handle timing edge cases more robustly  
- **Better error handling**: Improved panic recovery and error propagation within goroutines
- **Sequential operation support**: Multiple sequential diverge/converge operations work reliably
- **Automatic cleanup**: Completed goroutines are automatically cleaned up to prevent memory leaks

## Conclusion

Carrion's concurrency system provides a powerful yet simple way to write concurrent programs. The `diverge` and `converge` keywords offer intuitive control over goroutine creation and synchronization, while the underlying system handles the complex details of resource management, error handling, and cleanup.

With the recent improvements, you can now reliably use sequential diverge/converge operations, multiple named workers, and complex evaluation patterns after convergence. The system is designed to be both easy to use and robust for production applications.

By following the patterns and best practices outlined in this document, you can effectively use Carrion's concurrency features to build robust, efficient concurrent applications.