# 4.4 Dining philosophers

Each philosopher must alternately think and eat. However, a philosopher can only eat spaghetti when they have both left and right forks. Each fork can be held by only one philosopher and so a philosopher can use the fork only if it is not being used by another philosopher.

## Statistics for nerds

### Testing machine
Computador Intel(R) Core(TM) i7-8550U 1.80GHz de 8 núcleos,
com 16GB de memória RAM executando o Ubuntu 18.04.5 LTS com
o kernel Linux v4.15.0-142-generic-x86_64.

O experimento executou durante 15 segundos, sem logs e sem espera por operações
como *eat* e *thinking*. A versão da runtime utilizada foi golang v1.16.3.

### Results
``` shell
Philosopher (0) ate 9530710 times
Philosopher (1) ate 8549688 times
Philosopher (2) ate 9122172 times
Philosopher (3) ate 8380622 times
Philosopher (4) ate 9466254 times
```