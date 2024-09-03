#!/bin/bash

SERVER_ADRESS=150.165.85.31
SERVER_USER="aluno"

FILE_PATH=$1

# checking if the file exists
if [ ! -f $FILE_PATH ]; then
   echo "ERROR: File $FILE_PATH does not exist!"
   exit 1
fi

scp $FILE_PATH $SERVER_USER@$SERVER_ADRESS:/home/$SERVER_USER/answers/teste/
 	
EXIT_CODE=$?
if [ $EXIT_CODE -eq 0 ];
then 
	echo "OK! File submitted!"
else
	echo "ERROR: File was not submitted. Try again!"
fi
