# Programação Concorrente com NodeJS
Todo o conteúdo presente aqui são trechos retirados do material da COCIC-UFTPR. Recomendo consumir o material completo [clicando aqui](http://cocic.cm.utfpr.edu.br/progconcorrente/doku.php?id=nodejs).

Sobre NodeJS
=================================
Node.js é um interpretador de código JavaScript open-source e multiplataforma que executa JavaScript fora do navegador. Tipicamente a linguagem era usada apenas para adicionar dinamismo em páginas HTML, enquanto que para criar aplicações backend usava-se outras linguagens de programação como PHP, Java, C#, etc. No entanto, Node.js representa o paradigma “Javascript por toda parte”, que visa unificar o desenvolvimento de aplicações web em torno de uma única linguagem de programação, aumentando assim a produtividade dos desenvolvedores.

O Node.js foi construído utilizando utilizando dois componentes open-source, a engine JavaScript do Google chamada de V8 e a libuv. A V8 é escrita em C++ e compila o código-fonte JavaScript para o código de máquina em vez de interpretá-lo em tempo real; além disso ela é extremamente rápida com a parte de fundamentos da Internet como HTTP, DNS e TCP. A libuv, por sua vez, é utilizada para manipular eventos assíncronos e se trata de uma camada de abstração para a funcionalidade de rede e sistema de arquivos.

Modelo de Concorrência
=================================
Uma característica do JavaScript, que também está presente no Node.js, é de ter operações de I/O não bloqueantes, visando manter a concorrência, pelo fato da linguagem ser *single thread*. Além disso, Node.js utiliza a programação orientada à eventos, possibilitando o desenvolvimento de servidores web rápidos e altamente escaláveis sem a necessidade de utilizar threads explicitamente.

O **Event Loop** é o que permite que o Node.js execute operações de I/O assíncronas, enviando-as sempre que possível para o *kernel*. Quando uma operação é finalizada, o *kernel* avisa o Node.js para que o *callback* responsável por aquela operação possa ser adicionado na fila de execução. Internamente, o Node.js usa a biblioteca libuv, que tem um *pool* de *threads* (*Worker Pool* ou *Worker Threads*) fixa para manipular as operações assíncronas.

A imagem a seguir ilustra todo o fluxo de funcionamento do *event loop*.
![Event Loop Diagram](eventloop.png)
