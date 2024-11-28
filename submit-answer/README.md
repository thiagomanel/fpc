# Ferramenta para submissão de respostas de atividades/provas

Instruções para usar o script _submit-answer.sh_:
 1) Faça o clone deste repositório
 2) No momento da submissão da resposta, você deve informar um tipo do teste a ser submetido dentre os possíveis (**lab1**).
 3) Para submeter sua resposta, execute o script como segue. Supondo que você está submetendo resposta para o lab1 e quer submeter o arquivo `lab1_123456_654321.txt` , execute o comando a seguir:

`bash submit-answer.sh lab1 path/to/lab1_123456_654321.txt`

O script vai submeter o arquivo passado como parâmetro para o servidor da disciplina. Para cada arquivo espera-se uma dentre as mensagens abaixo:
  * **ERROR: File <filepath> does not exist!**: verifique o _path_ passado como parâmetro para o script;
  * **ERROR: Invalid test option!**: verifique o tipo da prova de reposição informado, é permitido apenas **prova1**, **prova2** ou **prova3**;
  * **ERROR: File was not submitted. Try again!**: o arquivo não foi enviado possivelmente devido a algum problema de rede, tente novamente. Se o problema persistir, entre em contato com o professor;
  * **OK! File submitted!**: o arquivo foi submetido com sucesso.
