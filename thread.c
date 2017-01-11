#include <stdio.h>
#include <pthread.h>

void increment(int* var)
{
	for(int i = 0; i < 1000000; i++)
	{
		(*var)++;
	}
}

void decrement(int* var)
{
	for(int i = 0; i < 1000000; i++)
	{
		(*var)--;
	}
}

int main()
{
	int var = 0;
	pthread_t thread1;
	pthread_t thread2;

	pthread_create(&thread1, NULL, increment, NULL);
	pthread_create(&thread2, NULL, decrement, NULL);

	pthread_join(thread1, NULL);
	pthread_join(thread2, NULL);

	printf("i = %d",i);


	return 0;
}