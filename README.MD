
# Go Struct Validation Benchmarking

## Overview

This branch contains benchmarking code for comparing the performance of reflective and non-reflective struct validation methods in Go. The benchmarks are designed to measure and compare the execution speed and efficiency of two different approaches to validating data structures:

1. **Reflective Validation**: Uses the Go `reflect` package to dynamically inspect and validate struct fields based on struct tags.
2. **Non-Reflective Validation**: Directly accesses struct fields without reflection, using predefined rules.

The purpose of these benchmarks is to provide insight into the trade-offs between flexibility and performance when using reflection in Go.

## Structs and Validation Rules

The benchmarks operate on a `UserDto` struct, which includes a single string field `Name` with validation rules defined in its struct tag. The validation rules include:
- `max-len`: The maximum allowed length of the string.
- `allow-symbol`: Whether or not symbols are permitted in the string.

## How to Run Benchmarks

To run the benchmarks included in this branch, follow these steps:

### Prerequisites

Ensure you have Go installed on your machine. The benchmarks are tested with Go version 1.15 and above. You can check your Go version using:

\`\`\`bash
go version
\`\`\`

### Running Benchmarks

1. Clone the repository and switch to this branch.
2. Execute the following command in your terminal:

\`\`\`bash
make reflect-test 
\`\`\`

This command will run all benchmarks defined in the test files and output the results, showing how many operations each method can perform per second and the average time per operation.

## Interpretation of Benchmark Results

The benchmark output will provide two main metrics:
- **ns/op**: Nanoseconds per operation, indicating the average duration of one validation operation.
- **b/op**: Bytes per operation, which shows the amount of memory allocated per operation.

Lower values in both metrics are indicative of better performance. The results will help you understand the impact of using reflection for struct validation in terms of speed and resource usage.

## Contributing

Contributions to this benchmarking suite are welcome. If you have suggestions for improving the benchmarks or adding new cases, please create a pull request or raise an issue.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
