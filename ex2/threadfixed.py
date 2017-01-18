from threading import Thread
from threading import RLock

var = 0
lock = RLock()

def increment():
	global var
	for i in range(1000000):
		with lock:
			var = var + 1

def decrement():
	global var
	for i in range(1000000):
		with lock:
			var = var - 1

def main():
	global var
	thread1 = Thread(target = increment, args = (),)
	thread2 = Thread(target = decrement, args = (),)

	thread1.start()
	thread2.start()

	thread1.join()
	thread2.join()

	print('var = ', var)
	
main()