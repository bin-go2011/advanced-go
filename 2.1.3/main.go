package main

/*
#include <stdio.h>

static void SayHello(const char* s) {
	puts(s);
}
*/
//void SayHello2(const char* s);
import "C"

func main() {
	C.SayHello2(C.CString("Hello World"))
}
