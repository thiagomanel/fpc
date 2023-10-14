# 4.1 Producer-consumer problem

In multithreaded programs there is often a division of labor between 
threads. In one common pattern, some threads are producers and some 
are consumers. Producers create items of some kind and add them to a 
data structure; consumers remove the items and process them.

###  Basic producer code
Assume that producers perform the following operations over and over:
```
event = waitForEvent()
buffer.add(event)
```

### Basic consumer code
Also, assume that consumers perform the following operations:
```
event = buffer.get()
event.process()
```

Puzzle: Add synchronization statements to the producer and consumer code
to enforce the synchronization constraints.