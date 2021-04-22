# The Little Book of Semaphores

Repository contains solutions, implemented in several programming languages, 
to the problems presented by the book Little Book of Semaphores. 

# Golang Solutions

Our Golang solutions do not use semaphores due to the simple fact that there is no native
support for it. But, you can easily find external libraries with such a feature.

We decided to only employ Golang's native [sync resources](https://golang.org/pkg/sync) in our solutions. As you will notice, these resources can replace the semaphore usage and
solve the problem as well. Here we describe some of these resources:

- ***[WaitGroup](https://golang.org/pkg/sync/#WaitGroup)***: a barrier that blocks all go routines until receive all ***N*** *Done* signals.
- ***[Channel](https://golang.org/doc/effective_go#channels)***: a low-latency and thread-safe communication channel/queue between go routines.
- ***[Mutex](https://golang.org/pkg/sync/#Mutex)***: a mutual exclusion lock.
- ***[RWLock](https://golang.org/pkg/sync/#RWMutex)***: a reader/writer mutual exclusion lock.

# Soluções em Golang [PT-BR]

Ao consultar as soluções implementadas em Golang será possível perceber que não utilizamos
semáforos. Isto porque Golang não possui suporte nativo a essa funcionalidade. Contudo, 
existem diversas bibliotecas que disponibilizam essa implementação.

Decidimos utilizar apenas [recursos de sincronização](https://golang.org/pkg/sync) nativos
da linguagem. O uso em conjunto desses recursos alternativos substituem as funcionalidades
de um semáforo. Por exemplo: 

- ***[WaitGroup](https://golang.org/pkg/sync/#WaitGroup)***: implementa uma barreira que bloqueia todas as *go routines* que esperam
por todos ***N*** sinais de *Done* necessários.
- ***[Channel](https://golang.org/doc/effective_go#channels)***: um canal/fila *thread-safe* 
de comunicação entre *go routines* de baixa latência.
- ***[Mutex](https://golang.org/pkg/sync/#Mutex)***: um *lock* de exclusão mútua.
- ***[RWLock](https://golang.org/pkg/sync/#RWMutex)***: um *lock* de exclusão mútua para
escrita e leitura.