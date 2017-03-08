#include <stdio.h>
#include <pthread.h>

void* increment(void* var)
{
	int i = 0;
	for(i = 0; i < 1000000; i++)
	{
		(*((int*)var))++;
	}
	return NULL;
}

void* decrement(void* var)
{
	int i = 0;
	for(i = 0; i < 1000000; i++)
	{
		(*((int*)var))--;
	}
	return NULL;
}

int main()
{

	int var = 0;
	pthread_t thread1;
	pthread_t thread2;

	pthread_create(&thread1, NULL, increment, (void*) &var);
	pthread_create(&thread2, NULL, decrement, (void*) &var);

	pthread_join(thread1, NULL);
	pthread_join(thread2, NULL);

	printf("i = %d\n",var);


	return 0;
}