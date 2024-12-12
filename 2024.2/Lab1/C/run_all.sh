cd serial
gcc cript.c -o cript
time ./cript ../../dataset
cd ../concurrent
gcc cript.c -o cript ../../dataset
time ./cript
