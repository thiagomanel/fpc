# Ferramenta para submissão de respostas de atividades/provas

Instruções para usar o script _submit-answer.sh_:
 1) Faça o clone deste repositório
 2) No diretório **answers**, crie um arquivo **.txt** para cada questão da atividade/prova de acordo com o template já disponível em **answers**. O nome do arquivo deve ter o seguinte formato `<seu nome>_<sua matriculaz>_q<número da questão>.txt`. Além disso, no cabeçalho de cada arquivo de resposta você deve inserir seu nome e sua matrícula.
 3) No momento da submissão da resposta, você deve informar o nome da atividade/teste dentre os possíveis (**prova1** ou **prova2**).
 4) Para cada arquivo de resposta, execute o script como segue. Supondo que você está realizando a prova1 e quer submeter o arquivo `maria_123456_q1.txt` , execute o comando a seguir:
 
`bash submit-answer.sh prova1 path/to/maria_123456_q1.txt`

O script vai submeter o arquivo passado como parâmetro para o servidor da disciplina. Para cada arquivo espera-se uma dentre as mensagens abaixo:
  * **ERROR: File <filepath> does not exist!**: verifique o _path_ passado como parâmetro para o script;
  * **ERROR: Invalid test option!**: verifique o tipo da prova informado, é permitido apenas **prova2**;
  * **ERROR: File was not submitted. Try again!**: o arquivo não foi enviado possivelmente devido a algum problema de rede, tente novamente. Se o problema persistir, entre em contato com o professor;
  * **OK! File submitted!**: o arquivo foi submetido com sucesso.
