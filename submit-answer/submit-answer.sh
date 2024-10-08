#!/bin/bash

SERVER_ADRESS=150.165.85.31
SERVER_USER="aluno"

TYPE=$1
FILE_PATH=$2

# checking if the file exists
if [ ! -f $FILE_PATH ]; then
   echo "ERROR: File $FILE_PATH does not exist!"
   exit 1
fi

if [ $TYPE = "prova1" ]
then
        echo "Repondo prova 1"
	scp $FILE_PATH $SERVER_USER@$SERVER_ADRESS:/home/$SERVER_USER/answers/reposicao/prova1/

else 
	if [ $TYPE = "prova2" ]
	then
        	echo "Repondo prova2"
		 scp $FILE_PATH $SERVER_USER@$SERVER_ADRESS:/home/$SERVER_USER/answers/reposicao/prova2/
	else 
		if [ $TYPE = "prova3" ]
		then 
			echo "Repondo prova3"
			 scp $FILE_PATH $SERVER_USER@$SERVER_ADRESS:/home/$SERVER_USER/answers/reposicao/prova3/
		else
			echo "ERROR: Invalid test option!"
		fi
	fi
fi

#scp $FILE_PATH $SERVER_USER@$SERVER_ADRESS:/home/$SERVERi_USER/answers/prova2/
 	
EXIT_CODE=$?
if [ $EXIT_CODE -eq 0 ];
then 
	echo "OK! File submitted!"
else
	echo "ERROR: File was not submitted. Try again!"
fi
