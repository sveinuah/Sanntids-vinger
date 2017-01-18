#include <stdio.h>
#include <pthread.h>

static pthread_mutex_t mutex;

void* increment(void* var)
{
	int i = 0;
	for(i = 0; i < 1000000; i++)
	{
		pthread_mutex_lock(&mutex);
		(*((int*)var))++;
		pthread_mutex_unlock(&mutex);
	}
	return NULL;
}

void* decrement(void* var)
{
	int i = 0;
	for(i = 0; i < 1000000; i++)
	{
		pthread_mutex_lock(&mutex);
		(*((int*)var))--;
		pthread_mutex_unlock(&mutex);
	}
	return NULL;
}

int main()
{

	int var = 0;
	pthread_t thread1;
	pthread_t thread2;

	pthread_mutex_init(&mutex, NULL);

	pthread_create(&thread1, NULL, increment, (void*) &var);
	pthread_create(&thread2, NULL, decrement, (void*) &var);

	pthread_join(thread1, NULL);
	pthread_join(thread2, NULL);

	printf("i = %d\n",var);


	return 0;
}