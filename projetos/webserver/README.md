# Servidor HTTP Concorrente

Neste projeto, você implementará um servidor HTTP concorrente. Teremos três objetivos principais:
* Aprender as primitivas de concorrência de Java;
* Saber aplicar os módulos utilitários de concorrência em Java
* Desenvolver um projeto de software concorrente.

## Visão Geral

Neste projeto, você implementará um servidor HTTP concorrente. Mais especificamente, você tornará concorrente uma implementação serial do protocolo HTTP. Concorrência será adicional ao código base com o objetivo específico aumentar a vazão do servidor HTTP. Esta implementação base considera somente uma parte do protocolo. Por esta razão, o projeto a ser desenvolvido considerará somente requisições GET.

A avaliação do projeto considerará duas partes: i) uma descrição do projeto concorrente a ser adotado (que justifique as decisões tomadas); e ii) o código fonte desenvolvido.

Para guiar as decisões de projeto, o desempenho de sua implementação será aferido (e comparado com as demais implementações de seus colegas) através de um benchmark (https://github.com/wg/wrk). Note que sua nota não será baseada no resultados do benchmark.

Mudanças na implementação do protocolo usado como base são permitidas, desde que justificadas.

## Especificação do projeto e avaliação

* iremos avaliar o desempenho de suas implementações do protocolo HTTP
* A avaliação "oficial" será feita pelo professor da disciplina; isso é necessário para termos um ambiente único de experimentação, assim, os resultados podem ser comparados mais facilmente. Para que sua implementação seja avaliada, é necessário que vocês escrevam dois scripts bash, que deve estar no diretório raiz do seu projeto:
buid.sh - esse script não deve receber nenhuma parâmetro. Ao executá-lo, deve compilar seu código fonte e gerar todos os binários necessários para a execução;
run.sh - esse script deve ser usado para iniciar e parar seu servidor web. Para iniciar, ele será usado da seguinte maneira: i) ./run.sh start port (após a execução do comando, seu servidor web deverá ter iniciado e estar escutando na porta especificada); ii) ./run stop (seu servidor deve parar de executar, após o comando ter terminado).

Suas implementações serão avaliadas em termos de latência e vazão usado a ferramenta wrk (https://github.com/wg/wrk). Seu professor divulgará os dados coletados durante a execução bem como o comando e/ou script usado com a ferramenta wrk.

Embora os resultados oficiais sejam dados pelo professor, você também pode executar a ferramenta como auxílio ao desenvolvimento (a página da ferramenta mostra um exemplo básico neste link https://github.com/wg/wrk#basic-usage). Caso não consiga compilar a ferramenta, é possível usar uma imagem docker pré-configurada para os testes (https://www.digitalocean.com/community/tutorials/how-to-benchmark-http-latency-with-wrk-on-ubuntu-14-04).
