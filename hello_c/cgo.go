package main

/*
#include <stdio.h>
#include <stdlib.h>

void myprintf(char *s){
	printf("%s", s);
}
*/
import "C"

import "unsafe"

func main(){
	cs := C.CString("Hello from cgo\n");
	C.myprintf(cs);
	C.free(unsafe.Pointer(cs))
}
