package main

// #cgo pkg-config: python3
// #define Py_LIMITED_API
// #include <Python.h>
// int PyArg_ParseTuple_Ls(PyObject *, long *, char **);
// int PyArg_ParseTuple_L(PyObject *, long *);
import "C"
import (
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
		log.Fatal(err)
	    panic("Error: Json load")
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var objmap map[string]interface{}
	if err := json.Unmarshal(byteValue, &objmap); err != nil {
	    log.Fatal(err)
	    panic("Error: Json unmarshal")
	}
	var wPtr C.long
	wPtr = cpointers.Ref(unsafe.Pointer(&objmap))
	return C.PyLong_FromLong(wPtr)
}

//export returnFromJson
func returnFromJson(self, args *C.PyObject) *C.PyObject {
	var wPtr C.long
	var word *C.char
	if C.PyArg_ParseTuple_Ls(args, &wPtr, &word) == 0 {
		return nil
	}
	w, ok := cpointers.Deref(wPtr)

	if !ok {
		log.Fatal("ptrProxy: pointer not found")
		panic("Error: pointer deferencing")
	}

	objmap := (*(*map[string]interface{})(w))
	jsonString, err := json.Marshal(objmap[C.GoString(word)].(map[string]interface{}))

	if err != nil {
		log.Fatal(err)
		panic("Error: Json marshal")
	}
    return C.PyBytes_FromString(C.CString(string(jsonString)))
}

func main() {
}