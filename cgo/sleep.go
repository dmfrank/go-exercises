package main

//#include <unistd.h>
import "C"
import "time"

func CGO() {
	for i := 0; i < 1000; i++ {
		go func() {
			C.sleep(1 /* seconds */)
		}()
	}
	time.Sleep(2 * time.Second)
}

func GO() {
	for i := 0; i < 1000; i++ {
		go func() {
			time.Sleep(1 /* seconds */)
		}()
	}
	time.Sleep(2 * time.Second)
}
