package main

// #cgo CFLAGS: -I${SRCDIR}/c_code
// #cgo LDFLAGS: ${SRCDIR}/callC.a
// #include <stdlib.h>
// #include <callC.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Going to call a C function!")
	C.cHello()

	fmt.Println("Going to call another C function!")
	myMessage := C.CString("This is Zhenya!")
	defer C.free(unsafe.Pointer(myMessage))
	C.printMessage(myMessage)

	fmt.Println("All perfectly done!")
}
