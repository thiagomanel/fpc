#  Barbeiro Dorminhoco
Todo o conteúdo presente aqui são trechos retirados do material da COCIC-UFTPR. Recomendo consumir o material completo [clicando aqui](http://cocic.cm.utfpr.edu.br/progconcorrente/doku.php?id=nodejs).

A tecnologia NodeJS é conhecida por ser single threaded, ou seja, sua aplicação roda em cima de uma única thread. Existem algumas bibliotecas e módulos que já possibilitam o uso de várias threads, mas ainda estão em fase experimental. Portanto, nestes exemplos vamos simular o uso da programação concorrente, utilizando a função setTimeout do javaScript, que agenda uma função para ser executada em um determinado tempo que for especificado.

Código Principal
=================================
```javascript
import Barbearia from './Barbearia';
 
const barbearia = new Barbearia(4); // insere o numero maximo de cadeiras
let barbeiroEstaDormindo = null;

function criaTempoAleatorio(maximo) {
    return Math.round(Math.random() * maximo);
}
 
function atendeCliente() {
 
    const cliente = barbearia.removeCliente(); // remove da lista de espera e coloca na cadeira do barbeiro
    if (!cliente) {
        const soneca = criaTempoAleatorio(30000);
        console.log(`Não há clientes. O barbeiro vai dormir por ${soneca}.`);
        barbeiroEstaDormindo = setTimeout(atendeCliente, soneca);
    } else {
        setTimeout(atendeCliente, cliente.tempo); // Agenda o próximo atendimento ao término do cliente atual
    }
}
 
function criaCliente() {
 
    const tempo = criaTempoAleatorio(22000);
    barbearia.adicionaCliente(tempo);
 
    if (barbeiroEstaDormindo) {
        clearTimeout(barbeiroEstaDormindo);
        setImmediate(atendeCliente);
    }
 
    setTimeout(criaCliente, criaTempoAleatorio(10000));
}
 
criaCliente();
atendeCliente();
```

Buffer compartilhado
=================================
```javascript
export default class Barbearia {
 
    constructor(numeroMaximoCadeiras) {
        this.numeroTotalClientes = 1;
        this.numeroMaximoCadeiras = numeroMaximoCadeiras;
 
        this.clientes = [];
    }
 
    adicionaCliente(tempo) {
        const id = this.numeroTotalClientes++;
 
        if (this.clientes.length >= this.numeroMaximoCadeiras) {
            console.log(`Não há cadeiras disponíveis. O cliente #${id} foi embora.`);
            return;
        }
 
        console.log(`O cliente #${id} chegou com tempo de ${tempo} para atendimento.`);
        this.clientes.push({ id, tempo });
    }
 
    removeCliente() {
        if (this.clientes.length < 1) {
            return false;
        }
 
        const cliente = this.clientes.shift();
        console.log(`O cliente #${cliente.id} está sendo atendido...`);
        return cliente;
    }
}
```
