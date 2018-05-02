# Servidor HTTP Concorrente

Neste projeto, você implementará um servidor HTTP concorrente. Teremos três objetivos principais:
* Aprender as primitivas de concorrência de Java;
* Saber aplicar os módulos utilitários de concorrência em Java
* Desenvolver um projeto de software concorrente.

## Visão Geral

Neste projeto, você implementará um servidor HTTP concorrente. Mais especificamente, você tornará concorrente uma implementação serial do protocolo HTTP. Concorrência será adicionada ao código base (https://github.com/gseguin/NanoHTTPd) com o objetivo específico aumentar a vazão do servidor HTTP. Esta implementação base considera somente uma parte do protocolo. Por esta razão, o projeto a ser desenvolvido considerará somente requisições GET.

A avaliação do projeto considerará duas partes: i) uma descrição do projeto concorrente a ser adotado (que justifique as decisões tomadas); e ii) o código fonte desenvolvido.

Para guiar as decisões de projeto, o desempenho de sua implementação será aferido (e comparado com as demais implementações de seus colegas) através de um benchmark (https://github.com/wg/wrk). Note que sua nota não será baseada no resultados do benchmark.

Mudanças na implementação do protocolo usado como base são permitidas, desde que justificadas.

## Especificação do projeto e avaliação

* Seu código fonte deve ser mantido em um repositório privado no Github. Adicione o professor da disciplina como colaborador do projeto. Você pode tanto fazer um fork do código base ou copiar o código fonte deste para um repositório independente.
* A avaliação de desempenho será feita de maneira automática e periódica. Os resultados serão divulgados publicamente. Para isso, adotaremos algumas convenções:
	* Será considerado somente o código do branch *master*
	* Deverá existir um script `build.sh` na raiz de seu repositório. Esse script não deve receber nenhuma parâmetro. Ao executá-lo, deve compilar seu código fonte e gerar todos os binários necessários para a execução;
	* Deverá existir um script `run.sh` na raiz de seu repositório. Esse script deve ser usado para iniciar e parar seu servidor web. Para iniciar, ele será usado da seguinte maneira: i) `./run.sh start port` (após a execução do comando, seu servidor web deverá ter iniciado e estar escutando na porta especificada); ii) `./run stop` (seu servidor deve parar de executar, após o comando ter terminado).
* Embora a medicação de desempenho "oficial" seja feita pelo professor da disciplina (isso é necessário para termos um ambiente único de experimentação, assim, os resultados podem ser comparados mais facilmente) nada impede que você também a execute. *TODO* indicar o procedimento de medição.
* A nota do projeto será baseada na adequação do projeto propost bem como na corretude e qualidade de sua implementação.
