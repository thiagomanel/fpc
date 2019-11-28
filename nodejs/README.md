# Programação Concorrente com NodeJS
Parte do conteúdo presente aqui são trechos retirados do material da COCIC-UFTPR. Recomendo consumir o material completo [clicando aqui](http://cocic.cm.utfpr.edu.br/progconcorrente/doku.php?id=nodejs).

Sobre NodeJS
=================================
Node.js é um interpretador de código JavaScript *open-source* e multiplataforma que executa JavaScript fora do navegador. Tipicamente a linguagem era usada apenas para adicionar dinamismo em páginas HTML, enquanto que para criar aplicações *backend* usava-se outras linguagens de programação como PHP, Java, C#, etc. No entanto, Node.js representa o paradigma *“Javascript por toda parte”*, que visa unificar o desenvolvimento de aplicações web em torno de uma única linguagem de programação, aumentando assim a produtividade dos desenvolvedores.

O Node.js foi construído utilizando utilizando dois componentes *open-source*, a *engine* JavaScript do Google chamada de V8 e a libuv. A V8 é escrita em C++ e compila o código-fonte JavaScript para o código de máquina em vez de interpretá-lo em tempo real; além disso ela é extremamente rápida com a parte de fundamentos da Internet como HTTP, DNS e TCP. A libuv, por sua vez, é utilizada para manipular eventos assíncronos e se trata de uma camada de abstração para a funcionalidade de rede e sistema de arquivos.

Modelo de Concorrência
=================================
Uma característica do JavaScript, que também está presente no Node.js, é de ter operações de I/O não bloqueantes, visando manter a concorrência, pelo fato da linguagem ser *single thread*. Além disso, Node.js utiliza a programação orientada à eventos, possibilitando o desenvolvimento de servidores web rápidos e altamente escaláveis sem a necessidade de utilizar threads explicitamente.

O **Event Loop** é o que permite que o Node.js execute operações de I/O assíncronas, enviando-as sempre que possível para o *kernel*. Quando uma operação é finalizada, o *kernel* avisa o Node.js para que o *callback* responsável por aquela operação possa ser adicionado na fila de execução. Internamente, o Node.js usa a biblioteca libuv, que tem um *pool* de *threads* (*Worker Pool* ou *Worker Threads*) fixa para manipular as operações assíncronas.

A imagem a seguir ilustra todo o fluxo de funcionamento do *event loop*.

![Event Loop Diagram](eventloop.png)


Worker Threads
=================================
A partir da versão 10.5.0 do Node foi introduzido o suporte **experimental** à *threads*, chamadas de *Worker Threads*.

Assim como os *workers* presentes no *Worker Pool*, as *worker threads* são utilizadas para tarefas de uso intenso de CPU. E devem ser utilizadas apenas para isso pois, conforme é esclarecido na [documentação oficial do Node](https://nodejs.org/docs/latest-v11.x/api/worker_threads.html#worker_threads_worker_threads), será uma perda de recursos uma vez que o mecanismo oferecido para manipular I/O é muito mais eficiente que usar uma *worker thread*.
 
O módulo `worker` fornece uma forma de criar múltiplos ambientes executando em diferentes *threads* e cria canais de mensagens entre eles. Para usar essa funcionalidade é necessário usar a *flag* `–experimental-worker` e importar o módulo no código:

```javascript
const worker = require('worker_threads');
```

Dentro desse módulo existem algumas variáveis de controle, sendo elas: `isMainThread`, `threadId` e `parentPort`.

- isMainThread: Retorna `true` se o código não estiver sendo executado em um *Worker*;
- threadId: Número inteiro que identifica a *thread* atual;
- parentPort: Canal de comunicação com a *thread* pai.

É possível compartilhar memória de forma eficiente transferindo instâncias de `ArrayBuffer` ou `SharedArrayBuffer` entre os *workers*.

Classes presentes no módulo
---------
### MessageChannel
Representa um canal de comunicação assíncrona de dois caminhos (entrada e saída). A classe não possui métodos, apenas retorna um objeto do tipo `MessagePort`, que possui os métodos para comunicação entre *Workers*.

### MessagePort
Também representa um canal de comunicação assíncrona de duas pontas, sendo uma de entrada e a outra de saída. Além disso, pode ser usado para transferir dados estruturados, regiões de memória e outras `MessagePorts` entre diferentes *Workers*.

### Worker
Representa a execução de uma *thread* JavaScript independente. A maior parte da API do Node.js está disponível nesta classe.

Exemplo
=================================
Código:
---------
```javascript	
const { Worker, isMainThread, workerData } = require('worker_threads');

console.log('Before enter the conditional check! isMainThread:' + isMainThread); // run to all threads

if (isMainThread) {
  id = 0;
  // This re-loads the current file inside a Worker instance.
  new Worker(__filename, { workerData: ++id }); // worker 1
  // the new Worker(__filename) works similar to fork, but running the code from begining
  new Worker(__filename, { workerData: ++id }); // worker 2
  new Worker(__filename, { workerData: ++id }); // worker 3

} else {
  id = workerData;
  setTimeout(function() {console.log("worker id:" + id);}, 1000 * id); // wait "id" seconds
}
```
Saída:
---------
```javascript	
Before enter the conditional check! isMainThread:true
Before enter the conditional check! isMainThread:false
Before enter the conditional check! isMainThread:false
Before enter the conditional check! isMainThread:false
worker id:1
worker id:2
worker id:3
```
Outros exemplos de concorrência com NodeJS
---------
- [problemas-de-listas](problemas-de-listas/)
- [problemas-classicos](problemas-classicos/)

Sugestões finais
---------
- [ler o material na íntegra](http://cocic.cm.utfpr.edu.br/progconcorrente/doku.php?id=nodejs)
- [ler este material alternativo também sobre threads em node](https://medium.com/@Trott/using-worker-threads-in-node-js-80494136dbb6)
- [ler este material alternativo também sobre threads em node](https://medium.com/dailyjs/threads-in-node-10-5-0-a-practical-intro-3b85a0a3c953)
- [ler este material alternativo sobre concorrencia com async/await](https://medium.com/platformer-blog/node-js-concurrency-with-async-await-and-promises-b4c4ae8f4510)
- [ler documentação oficial de Node](https://nodejs.org/docs/latest-v11.x/api/worker_threads.html#worker_threads_worker_threads)
- conferir mais exemplos em [problemas-de-listas](problemas-de-listas/) e/ou [problemas-classicos](problemas-classicos/)
