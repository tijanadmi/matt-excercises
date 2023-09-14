#Fan-Out, Fan-In Pattern

Sometimes, stages in your pipeline can be particularly computationally expensive. Imagine you have many attributes to validate and Stage 1 (Validation) is taking longer time then expected. In fact, it turns out it can, and you can solve it by another pattern has name: fan-out, fan-in.
Here's how the fan-out, fan-in pattern works:
Fan-Out: The input data is divided into smaller chunks, and each chunk is processed concurrently by a separate goroutine or stage. This allows for parallel processing of the workload.
Fan-In: The results produced by the concurrent goroutines or stages are collected and merged into a single output channel or data structure. This consolidation combines the individual results into a final result.
The fan-out, fan-in pattern is particularly useful when you have a large amount of data or a computationally intensive task that can be divided into independent subtasks. By distributing the work across multiple goroutines or stages like validateOrders goroutine, you can take advantage of parallel processing capabilities and improve overall performance.