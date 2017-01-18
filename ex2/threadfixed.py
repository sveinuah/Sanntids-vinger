from threading import Thread
from threading import Lock

var = 0
q = Queue()

def increment():
	global var
	for i in range(1000000):
		var.aquire()
		var = var + 1
		var.release()

def decrement():
	global var
	for i in range(1000000):
		var.aquire()
		var = var - 1
		var.release()

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