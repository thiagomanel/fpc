#include <stdio.h>
#include <assert.h>
#include <pthread.h>

long count = 0;
long COUNT_LIMIT = 57;

pthread_mutex_t count_mutex;
pthread_cond_t count_threshold_cv;

void *inc_count(void *t) {
 	int i;
 	long my_id = (long)t;
 	for (i=0; i<1e7; i++) {
 	 	pthread_mutex_lock(&count_mutex);
 	 	count++;
 	 	printf("thread_id %ld working on count %ld\n", my_id, count);
 	 	if (count >= COUNT_LIMIT) {
 	 		printf("thread_id %ld reach limit %ld\n", my_id, count);
 		 	pthread_cond_signal(&count_threshold_cv);
 	 	}
 	 	pthread_mutex_unlock(&count_mutex);
 		sleep(1);
 	}
 	pthread_exit(NULL);
}

void *watch_count(void *t) {

 	long my_id = (long)t;

 	//condition variables are always used with a mutex
 	pthread_mutex_lock(&count_mutex);
 	//protect against spurious wake-up
 	while (count<COUNT_LIMIT) {
 		pthread_cond_wait(&count_threshold_cv, &count_mutex);
 	}
 	pthread_mutex_unlock(&count_mutex);
 	pthread_exit(NULL);
}


int main (int argc, char *argv[]) {
 	int i;
 	long t1=1, t2=2, t3=3;

 	pthread_t threads[3];
 
 	//init data structures
 	pthread_mutex_init(&count_mutex, NULL);
 	pthread_cond_init (&count_threshold_cv, NULL);
	
	//create threads
	pthread_create(&threads[0], NULL, watch_count, (void *)t1);
 	pthread_create(&threads[1], NULL, inc_count, (void *)t2);
 	pthread_create(&threads[2], NULL, inc_count, (void *)t3);

 	//wait for threads termination
 	for (i=0; i<3; i++) {
 		pthread_join(threads[i], NULL);
 	}

 	//finalize data structures
	pthread_mutex_destroy(&count_mutex);
 	pthread_cond_destroy(&count_threshold_cv);

	printf("count %ld\n", count);
 	pthread_exit(NULL);
}
