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
