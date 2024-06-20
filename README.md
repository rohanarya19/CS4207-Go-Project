# Go Concurrent Algorithms

This repository contains implementations of various concurrent algorithms in Go for my project of the moduele Advanced Programming Concepts. The provided Go files demonstrate different methods of sorting using concurrency.

## Algorithms

### 1. Bubble Sort
**File:** `bubble.go`

Bubble Sort is a simple comparison-based sorting algorithm. In this implementation, the algorithm is enhanced with concurrency to improve performance.

### 2. Quick Sort
**File:** `quick.go`

Quick Sort is a highly efficient sorting algorithm that uses a divide-and-conquer approach. This concurrent version splits the work across multiple goroutines.

### 3. Rohan's Bonus Algorithm
**File:** `RohanBonus.go`

This file contains a custom concurrent algorithm designed by me where I combined Radix sort with quick sort to make the algorithm faster which resulted in me winning the bonus competition and being awarded an A1 grade in the module. It showcases unique methods of leveraging Go's concurrency features.

## Getting Started

### Prerequisites
- Go (version 1.16 or higher)

### Running the Algorithms

1. **Clone the repository:**
   ```sh
   git clone https://github.com/yourusername/concurrent-algorithms.git
   cd concurrent-algorithms
   ```

2. **Run each algorithm:**

   - **Bubble Sort:**
     ```sh
     go run bubble.go
     ```

   - **Quick Sort:**
     ```sh
     go run quick.go
     ```

   - **Rohan's Bonus Algorithm:**
     ```sh
     go run RohanBonus.go
     ```

## Contributing

Contributions are welcome! If you have any improvements or additional concurrent algorithms, feel free to open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
