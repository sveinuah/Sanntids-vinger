from threading import Thread

i = 0

def increment():
	global i
	for i in range(1000000):
		i = i + 1

def decrement():
	global i
	for i in range(1000000):
			i = i - 1

def main():
	global i
	thread1 = Thread(target = increment, args = (),)
	thread2 = Thread(target = decrement, args = (),)

	thread1.start()
	thread2.start()

	thread1.join()
	thread2.join()

	print("i = ", i)
	
main()