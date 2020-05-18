package main

// #cgo pkg-config: python3
// #define Py_LIMITED_API
// #include <Python.h>
import "C"
import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
	"unsafe"
	"log"
)

var cpointers = PtrProxy()

//export loadJson
func loadJson(self, args *C.PyObject) *C.PyObject {

	jsonFile, err := os.Open("data.json")
	defer jsonFile.Close()
	if err != nil {
	    fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var objmap map[string]interface{}
	if err := json.Unmarshal(byteValue, &objmap); err != nil {
	    log.Fatal(err)
	}
	var wPtr C.long
	wPtr = cpointers.Ref(unsafe.Pointer(&objmap))
	return C.PyLong_FromLong(wPtr)
}

func main() {
}