# pc

Instruções para usar o script _submit-answer.sh_:
 1) Faça o clone deste repositório
 2) Crie um arquivo **.txt** no diretório **answers** para cada questão da prova seguindo o template já disponível em **answers**. No cabeçalho de cada arquivo de resposta você deve inserir seu nome e matrícula.
 3) Após criar uma arquivo de resposta para cada questão, execute o comando abaixo:

`bash submit-answer.sh`

O script vai criptografar cada um de seus arquivos de respostas (armazenados no diretório **answers**) e submetê-los para o servidor da disciplina. Para cada arquivo espera-se uma mensagem **"File <filename> was submitted successfully."** Caso contrário, entre em contato com o professor.
