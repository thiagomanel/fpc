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

if [ $TYPE = "lab2" ]
then
        echo "Submiting lab1: $FILE_PATH"
	scp $FILE_PATH $SERVER_USER@$SERVER_ADRESS:/home/$SERVER_USER/2025-1/$TYPE/

else 
	echo "ERROR: Invalid test option!"
	exit

fi

EXIT_CODE=$?
if [ $EXIT_CODE -eq 0 ];
then 
	echo "OK! File submitted!"
else
	echo "ERROR: File was not submitted. Try again!"
fi
