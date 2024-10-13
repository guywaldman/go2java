package main

/*
#cgo LDFLAGS: -L${SRCDIR}/../../out
#include <jni.h>
#include <stdlib.h>

// Helper function to get C string from jstring
static const char* GetStringUTFChars(JNIEnv *env, jstring str) {
    return (*env)->GetStringUTFChars(env, str, NULL);
}

// Helper function to release C string
static void ReleaseStringUTFChars(JNIEnv *env, jstring str, const char *chars) {
    (*env)->ReleaseStringUTFChars(env, str, chars);
}

// Helper function to create new jstring
static jstring NewStringUTF(JNIEnv *env, const char *chars) {
    return (*env)->NewStringUTF(env, chars);
}
*/
import (
	"C"
)
import (
	"unsafe"

	"github.com/guywaldman/go2java/internal/greeting"
)

//export Java_Main_GreetPerson
func Java_Main_GreetPerson(env *C.JNIEnv, clazz C.jclass, input C.jstring) C.jstring {
	// Get C string from jstring
	cInput := C.GetStringUTFChars(env, input)
	// Ensure the C string is released after use
	defer C.ReleaseStringUTFChars(env, input, cInput)

	// Convert C string to Go string
	goStr := C.GoString(cInput)

	// Act on the data
	output := *greeting.GreetPerson(goStr)

	// Convert Go string back to C string
	cResult := C.CString(output)
	defer C.free(unsafe.Pointer(cResult))

	// Create a new jstring to return
	return C.NewStringUTF(env, cResult)
}

// Generating a shared library requires a `main` function.
func main() {}
