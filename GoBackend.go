package main

/*
#cgo pkg-config: python3
#define Py_LIMITED_API
#include <Python.h>
#include<stdlib.h>
typedef struct Request_
{
    const char *Method;
    const char *Host;
    const char *URL;
    const char *Body;
    const char *Headers;
} Request;

typedef unsigned int ResponseWriterPtr;

typedef void FuncPtr(ResponseWriterPtr w, Request *r);

extern void Call_HandleFunc(ResponseWriterPtr w, Request *r, FuncPtr *fn);
extern void HandleFunc(char* cpattern, FuncPtr* fn);
extern void Hello();
*/
import "C"
import (
    "net/http"
    "unsafe"
    "bytes"
    "context"
    "os"
    "io/ioutil"
    "encoding/json"
    "log"
)

var cpointers = PtrProxy()
var srv http.Server = http.Server{}

//export loadJson
func loadJson() C.long {

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
    return wPtr
}

//export returnFromJson
func returnFromJson(wPtr C.long, word *C.char) *C.char {
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
    return C.CString(string(jsonString))
}

//export ListenAndServe
func ListenAndServe(caddr *C.char) {
    addr := C.GoString(caddr)
    srv.Addr = addr
    srv.ListenAndServe()
}

//export Shutdown
func Shutdown() {
    srv.Shutdown(context.Background())
}

//export HandleFunc
func HandleFunc(cpattern *C.char, cfn *C.FuncPtr) {
    // C-friendly wrapping for our http.HandleFunc call.
    pattern := C.GoString(cpattern)
    http.HandleFunc(pattern, func(w http.ResponseWriter, req *http.Request) {
        // Convert the headers to a String
        headerBuffer := new(bytes.Buffer)
        req.Header.Write(headerBuffer)
        headersString := headerBuffer.String()
        // Convert the request body to a String
        bodyBuffer := new(bytes.Buffer)
        bodyBuffer.ReadFrom(req.Body)
        bodyString := bodyBuffer.String()
        // Wrap relevant request fields in a C-friendly datastructure.
        creq := C.Request{
            Method:  C.CString(req.Method),
            Host:    C.CString(req.Host),
            URL:     C.CString(req.URL.String()),
            Body:    C.CString(bodyString),
            Headers: C.CString(headersString),
        }
        wPtr := cpointers.Ref(unsafe.Pointer(&w))
        C.Call_HandleFunc(C.ResponseWriterPtr(wPtr), &creq, cfn)
        // release the C memory
        C.free(unsafe.Pointer(creq.Method))
        C.free(unsafe.Pointer(creq.Host))
        C.free(unsafe.Pointer(creq.URL))
        C.free(unsafe.Pointer(creq.Body))
        C.free(unsafe.Pointer(creq.Headers))
        cpointers.Free(wPtr)
    })
}

//export ResponseWriter_Write
func ResponseWriter_Write(wPtr C.long, cbuf *C.char, length C.int) C.int {
    buf := C.GoBytes(unsafe.Pointer(cbuf), length)

    w, ok := cpointers.Deref(wPtr)
    if !ok {
        return 0
    }
    n, err := (*(*http.ResponseWriter)(w)).Write(buf)
    if err != nil {
        return 0
    }
    return C.int(n)
}

//export ResponseWriter_WriteHeader
func ResponseWriter_WriteHeader(wPtr C.long, header C.int) {
    w, ok := cpointers.Deref(wPtr)
    if !ok {
        return
    }
    (*(*http.ResponseWriter)(w)).WriteHeader(int(header))
}

func main() {}