package main

import "unsafe"

func main() {
	l := []int{9, 45, 23, 67, 78}
	t := 0
	i := 0

	var tmp int
	p := uintptr(unsafe.Pointer(&l[0]))

	if i >= 5 {
		goto end
	}
body:
	tmp = *(*int)(unsafe.Pointer(p))
	p += unsafe.Sizeof(l[0])
	i++
	t += tmp
	if i < 5 {
		goto body
	}
end:
	println(t)
}

// https://medium.com/a-journey-with-go/go-how-are-loops-translated-to-assembly-835b985309b3
