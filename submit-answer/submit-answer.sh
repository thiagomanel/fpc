#!/bin/bash

BASE_DIR=$(dirname -- "$( readlink -f -- "$0"; )")

SERVER_ADRESS=150.165.85.31
SERVER_USER="concurrent"
PUBLIC_KEY_PATH="$BASE_DIR/keys/public_key.pem"
SSH_KEY="$BASE_DIR/keys/concurrent_ssh_key"

chmod 600 $SSH_KEY

for FILE_PATH in $(ls $BASE_DIR/answers/);
do
	echo "Processing file $FILE_PATH"
	echo "Encrypting your the file..."

	# Encrypting your answer file using the Public Key
	FILE_NAME=$(basename $FILE_PATH .txt)
	openssl rsautl -encrypt -inkey $PUBLIC_KEY_PATH -pubin -in $BASE_DIR/answers/$FILE_PATH -out "$BASE_DIR/answers/$FILE_NAME"-encrypted.bin

	echo "Submitting the file..."
	# Sending your encrypted answer file to the server
	scp -i $SSH_KEY "$BASE_DIR/answers/$FILE_NAME"-encrypted.bin $SERVER_USER@$SERVER_ADRESS:/home/$SERVER_USER/answers/prova2/
 	
	
	EXIT_CODE=$?
	if [ $EXIT_CODE -eq 0 ];
	then 
		echo "File $FILE_PATH was submitted successfully."
	else
		echo "There was an error while submitting the file $FILE_PATH"
	fi

	rm "$BASE_DIR/answers/$FILE_NAME"-encrypted.bin
done
