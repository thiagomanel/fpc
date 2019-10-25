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
