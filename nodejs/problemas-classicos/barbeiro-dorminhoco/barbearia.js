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
