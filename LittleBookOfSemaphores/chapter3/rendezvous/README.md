# Rendezvous Problem

Puzzle: Generalize the signal pattern so that it works both ways. Thread A has
to wait for Thread B and vice versa. In other words, given this code

| Thread A | Thread B |  
| --- | --- |
| statement a1 | statement b1 |
| statement a2 | statement b2 |

we want to guarantee that a1 happens before b2 and b1 happens before a2. In 
writing your solution, be sure to specify the names and initial values of 
your semaphores (little hint there).