package main

/*
#cgo windows CFLAGS: -DCGO_OS_WINDOWS=1
#cgo darwin CFLAGS: -DCGO_OS_DARWIN=1
#cgo linux CFLAGS: -DCGO_OS_LINUX=1

#include <stdio.h>
#if defined(CGO_OS_WINDOWS)
	const char* os = "windows";
#elif defined(CGO_OS_DARWIN)
	char* os = "darwin";
#elif defined(CGO_OS_LINUX)
	const char* os = "linux";
#else
#	error(unknown os);
#endif

*/
import "C"

func main() {
	println(C.GoString(C.os))
}
