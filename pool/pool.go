package main

import "fmt"

// limit capacity of the pool
const poolCap = 100

// limit []byte cap
const maxCap = 1024

// create pool
var chanPool = make(chan []byte, poolCap)

func getBytes() (b []byte) {
	select {
	case bt, ok := <-chanPool:
		if ok {
			return bt
		}
		// non-normal behaivor, need fix!
		panic("Who, the fuck, closed the cahnnel of the chanPool")
	default:
	}
	return // nil<[]byte>
}

func putBytes(b []byte) {
	if cap(b) > maxCap {
		return
	}
	b = b[:0]
	select {
	case chanPool <- b:
	default:
	}
	return
}

func main() {
	t := getBytes()
	t = append(t, []byte{'H', 'i', '!'}...)
	putBytes(t)
	y := getBytes()
	if cap(y) > 3 {
		fmt.Println("len:", len(y))
		fmt.Println(string(y[:3]))
	}

}
