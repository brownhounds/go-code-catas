# Go Code Katas

Small Go code katas organized as stand-alone modules. This repo is a personal
practice space for algorithm drills, data-structure patterns, and core Go
syntax fluency.

## Why code katas

- Build speed and accuracy on common patterns.
- Isolate a single concept to practice it deeply with tests.
- Keep muscle memory fresh for interviews or real-world problem solving.
- Create a growing library of reference solutions you can revisit later.

### Slices / Arrays

- `two_sum`: Find two indices whose values add to a target using a hash map
(complement lookup).
- `rotate_slice_right`: Rotate a slice to the right by `k` steps in-place using
the three-reversal trick.
- `remove_duplicates`: Remove duplicates from a sorted slice in-place using a 
two-pointer pass; returns the new logical length.
- `move_zeros`: Move all zeros to the end of a slice in-place while preserving 
the order of non-zero elements (stable compaction).
- `maximum_subarray_sum`: Compute the max sum of any contiguous subarray using 
Kadane's algorithm.
- `product_of_array_except_self`: Build a result slice where each element is 
the product of all other elements using prefix/suffix products.
- `slice_functions`: JS-style helpers for slices using Go generics. Includes:
    - `Filter` - returns a new slice with elements that pass a test.
    - `Find` - returns the first element that passes a test.
    - `Map` - returns a new slice by transforming each element.
    - `Some` - returns true if any element passes the test.
    - `Every` - returns true if all elements pass the test.
    - `Includes` - returns true if a value exists in the slice.

### Strings

- `valid_anagram`: Determine whether two strings are anagrams using character frequency counting.
- `longest_unique_substring`: Compute the length of the longest substring without repeating characters 
using a sliding window with last-seen index tracking.
- `minimum_window_substring`: Find the smallest contiguous substring containing all characters of a target
  string (including duplicates) using a sliding window and frequency map; Unicode-safe via rune processing.
- `group_anagrams`: Group strings into collections where each collection contains
  words that are anagrams of each other; Unicode-safe with character multiplicity respected.

### Simple Programs

- `parse_csv`: Read a CSV file into a struct, manipulate it in memory, and write it back to disk. Demonstrates
file I/O, string manipulation, and memory-safe data structures in Go.

### Go Generic Concepts

- `pointer_chasing`: Compare pointer-based traversal versus contiguous slice traversal to demonstrate the
pointer chasing effect, CPU cache friendliness, and hardware prefetch behaviour.
- `struct_layout`: Show how Go arranges struct fields in memory and how field alignment requirements introduce
padding between fields or at the end of the struct. This kata demonstrates how field order affects memory layout,
why pointers and slice headers increase alignment and size, and why it is generally preferable for any unavoidable
padding to appear as tail padding rather than between fields, keeping frequently accessed data tightly packed
and cache-friendly.
- `heap_escapes`: Demonstrate how Go’s escape analysis moves values from the stack to the heap when their lifetime
cannot be proven to be local, and how this directly translates into heap allocations and garbage-collector pressure.
The benchmarks contrast value returns versus pointer returns, global pointer storage, closure capture, interface
boxing, and address-taking inside loops, showing how small structural changes can turn allocation-free code into
allocation-heavy hot paths despite identical algorithmic complexity.
- `http_server`: Simple HTTP server demonstrating:
    - Handlers with middleware
    - Path parameter extraction (manual, pure net/http)
    - JSON responses
    - Safe, idiomatic Go web handling without external frameworks
- `http_client`: HTTP client demonstrating:
    - Dependency injection of *http.Client for testability
    - Sending JSON requests and decoding JSON responses
    - Mockable responses derived from the request body for unit testing
    - Safe, idiomatic Go handling of HTTP requests without external frameworks
- `goroutines`: Computes squares of numbers concurrently using multiple goroutines.
    - Uses `sync.Mutex` to safely write to a shared slice.
    - Uses `sync.WaitGroup` to wait until all goroutines finish.    
- `channels`: Computes squares of numbers using channels to pass data between goroutines. Implements a producer-consumer pattern:
    - Producer goroutine sends numbers to a channel.
    - Consumer goroutine squares numbers and sends results to a results channel.
    - Main goroutine collects results from the results channel.
