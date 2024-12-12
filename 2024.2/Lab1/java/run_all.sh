cd serial
javac PasswordProcessorSerial.java  # Adicione a extensão do arquivo .java
time java PasswordProcessorSerial ../../dataset
cd ../concurrent
javac PasswordProcessorConcurrent.java  # Adicione a extensão do arquivo .java
time java PasswordProcessorConcurrent ../../dataset

