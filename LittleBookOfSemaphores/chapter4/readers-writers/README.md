# 4.2 Readers-writers problem

The next classical problem, called the Reader-Writer Problem, pertains to any
situation where a data structure, database, or file system is read and modified
by concurrent threads. While the data structure is being written or modified
it is often necessary to bar other threads from reading, in order to prevent a
reader from interrupting a modification in progress and reading inconsistent or
invalid data.

The synchronization constraints are:
1. Any number of readers can be in the critical section simultaneously.
2. Writers must have exclusive access to the critical section.

**Puzzle**: Use semaphores to enforce these constraints, while allowing readers
and writers to access the data structure, and avoiding the possibility of deadlock.