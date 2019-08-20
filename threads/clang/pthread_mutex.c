#include <stdio.h>
#include <assert.h>
#include <pthread.h>
#include <unistd.h>

//shared variable
long count = 0;
pthread_mutex_t mutex;

void *inc_count(void *t) {
	//here we have a critical region. it lacks proper protection
 	int i;
 	long my_id = (long)t;
	printf("I am a counter %ld\n", my_id);

    //we are protecting more than needed (critical region is line 19)
    pthread_mutex_lock(&mutex);
	for (i = 0; i < 1e7; i++) {
		count = count + 1;
	}
    pthread_mutex_unlock(&mutex);
	//we might have used a parameter in below call to be collected
	//by the pthread_join call
 	pthread_exit(&my_id);
}

int main (int argc, char *argv[]) {
 	int i;
	//args to be passed to the inc_count threaded call
 	long t1=1, t2=2, t3=3;

    //init mutex
    pthread_mutex_init(&mutex, NULL);

	//declare threads
 	pthread_t threads[3];

	//create threads. we do not change attributes, so NULL
	//last arg will be passed as parameter to the inc_count funcion
	pthread_create(&threads[0], NULL, inc_count, (void *)t1);
 	pthread_create(&threads[1], NULL, inc_count, (void *)t2);
 	pthread_create(&threads[2], NULL, inc_count, (void *)t3);

	//wait for thread termination
    long ret;
 	for (i=0; i<3; i++) {
 		pthread_join(threads[i], (void*) &ret);
        printf("foo %ld\n", ret);
 	}

	//with proper syncronization, we would have 3*1e7 printed
	printf("count %ld\n", count);
 	pthread_exit(NULL);
}
