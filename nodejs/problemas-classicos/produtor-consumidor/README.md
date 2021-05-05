# Produtor/Consumidor
Todo o conteúdo presente aqui são trechos retirados do material da COCIC-UFTPR. Recomendo consumir o material completo [clicando aqui](http://cocic.cm.utfpr.edu.br/progconcorrente/doku.php?id=nodejs).

A tecnologia NodeJS é conhecida por ser single threaded, ou seja, sua aplicação roda em cima de uma única thread. Existem algumas bibliotecas e módulos que já possibilitam o uso de várias threads, mas ainda estão em fase experimental. Portanto, nestes exemplos vamos simular o uso da programação concorrente, utilizando a função setTimeout do javaScript, que agenda uma função para ser executada em um determinado tempo que for especificado.

No código abaixo, o programa faz a primeira chamada da função “producer” que por sua vez tenta requerir a permissão para executar os próximos comandos para inserção de um item no buffer. Cada vez que função é invocada, é criada uma Promise, que o javaScript coloca em uma fila para ser executado posteriormente. Este bloco de código fica aguardando até que o callback “release” seja invocado, indicando que a thread do JavaScript pode recuperar o bloco e continuar a execução.

Após a liberação do processo, o JavaScript então executa o comando de inserção do item no buffer e vai para o próximo bloco de execução.

Cada função “producer” e “consumer” fazem o agendamento delas mesmas utilizando a função explicada acima (setTimeout) com um tempo randômico, de no máximo 3 segundos, podendo então em determinado momento serem chamadas na mesma hora. A concorrência então é controlada pelas classe “Semaphore”, uma biblioteca de terceiros que tem funcionamento semelhante a classe “Semaphore” do Java, que controla a quantidade de tarefas que podem ser executadas em paralelo. Este valor é indicado na hora da criação do objeto.

Principal
=================================
```javascript
global.Promise = require('bluebird');
import SharedBuffer from './shared-buffer';
 
const buffer = new SharedBuffer();
let count = 0;
 
function producer() {
    console.log('producer');
 
    buffer.put(++count)
        .then(_ => {
            const timeout = Math.random() * 1000;
            setTimeout(producer, timeout);
        });
}
 
function consumer() {
    console.log('consumer');
 
    buffer.get()
        .then(value => {
            console.log('read', value);
 
            const timeout = Math.random() * 1000;
            setTimeout(consumer, timeout);
        });
}
 
producer();
consumer();
```

Buffer compartilhado
=================================
```javascript
import { Semaphore, Mutex } from 'await-semaphore';
 
export default class SharedBuffer {
 
    constructor() {
        this.buffer = [];
        this.mutex = new Mutex(); // não deixa ler enquanto escreve
        this.read = new Semaphore(1); 
        this.write = new Semaphore(8);
    }
 
    put(value) {
        return new Promise((resolve, reject) => {
            this.write.acquire() // pergunta se pode escrever, pois se tiver 8 não pode escrever, se tiver cheio deixa esperando
                .then(release1 => {
 
                    this.mutex.acquire() // garante que nao tem ninguem lendo
                        .then(release2 => {
 
                            this.buffer.push(value);
 
                            release2(); // semaforo libera novamente para leitura
                            release1(); // semaforo libera novamente para escreve
 
                            resolve();
                        })
                        .catch(reject);
                })
                .catch(reject);
        });
    }
 
    get() {
        return new Promise((resolve, reject) => {
            this.read.acquire() // ve se nao tem ninguem lendo
                .then(release1 => {
 
                    this.mutex.acquire() // ve se nao tem ninguem escrevendo
                        .then(release2 => {
 
                            const value = this.buffer.pop(); // le e excluir o ultimo item do array
 
                            release2(); // libera para escrever
                            release1(); // libera para ler
 
                            resolve(value);
                        })
                        .catch(reject);
                })
                .catch(reject);
        });
    }
}
```
