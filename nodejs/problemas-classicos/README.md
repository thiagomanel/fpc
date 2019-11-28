# Problemas Clássicos de Programação Concorrente

Produtor/Consumidor
=================================
No problema do Produtor/Consumidor, há um buffer que é compartilhado entre dois processos chamados de produtor e consumidor. O produtor irá produzir e armazenar itens no buffer, enquanto o consumidor irá retirar e consumir os dados do buffer. Os dois processos devem acessar o buffer de maneira exclusiva.

Confiram [aqui](produtor-consumidor/README.md) o código solução do Produtor/Consumidor com NodeJS.

Barbeiro Dorminhoco
=================================
O Barbeiro Dorminhoco consiste em um problema clássico de programação concorrente. A descrição do problema é a seguinte: Na barbearia há um barbeiro e a cadeira do barbeiro, além de n cadeiras para os clientes esperarem atendimento. Quando não há clientes o barbeiro dorme, mas assim que chega um novo cliente o barbeiro acorda. Os novos clientes se sentaram nas cadeiras de espera caso ele esteja atendendo ou iram embora caso as cadeiras estejam ocupadas.

Confiram [aqui](barbeiro-dorminhoco/README.md) o código solução do Barbeiro Dorminhoco com NodeJS.
