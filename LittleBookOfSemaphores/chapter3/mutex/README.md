# 3.4 Mutex

Puzzle: Add semaphores to the following example to enforce mutual 
exclusion to the shared variable count.

| Thread A | Thread B | 
| --- | --- |
| count = count + 1 | count = count + 1 |