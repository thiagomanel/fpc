# 4.4 Cigarette smokers problem

Four threads are involved: an agent and three smokers. The smokers loop
forever, first waiting for ingredients, then making and smoking cigarettes. The
ingredients are tobacco, paper, and matches.

We assume that the agent has an infinite supply of all three ingredients, and
each smoker has an infinite supply of one of the ingredients; that is, one smoker
has matches, another has paper, and the third has tobacco.

The agent repeatedly chooses two different ingredients at random and makes
them available to the smokers. Depending on which ingredients are chosen, the
smoker with the complementary ingredient should pick up both resources and
proceed.

## Agent codes

#### Agent A
```c
agentSem.wait()
tobacco.signal()
paper.signal()
```

#### Agent B
```c
agentSem.wait()
paper.signal()
match.signal()
```

#### Agent C
```c
agentSem.wait()
tobacco.signal()
match.signal()
```