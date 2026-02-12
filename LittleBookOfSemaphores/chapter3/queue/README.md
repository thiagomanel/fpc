# 3.8 Queue

For example, imagine that threads represent ballroom dancers and that two
kinds of dancers, leaders and followers, wait in two queues before entering the
dance floor. When a leader arrives, it checks to see if there is a follower waiting.
If so, they can both proceed. Otherwise it waits.

Similarly, when a follower arrives, it checks for a leader and either proceeds
or waits, accordingly.

Puzzle: write code for leaders and followers that enforces these constraints.