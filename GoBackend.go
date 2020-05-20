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
*/
import "C"
import (
    "github.com/valyala/fasthttp"
    "unsafe"
    "os"
    "io/ioutil"
    "encoding/json"
    "log"
)

var cpointers = PtrProxy()

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

//export HandleFunc
func HandleFunc(cpattern *C.char, cfn *C.FuncPtr) {
    fasthttp.ListenAndServe(":5050", func(ctx *fasthttp.RequestCtx) {
        creq := C.Request{
            Method:  C.CString(string(ctx.Method())),
            Host:    C.CString(string(ctx.Host())),
            URL:     C.CString(string(ctx.RequestURI())),
            Body:    C.CString(string(ctx.PostBody())),
            Headers: C.CString("Header"), //figure this out
        }
        wPtr := cpointers.Ref(unsafe.Pointer(&ctx))
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
    ctxU, ok := cpointers.Deref(wPtr)
    if !ok {
        return 0
    }
    ctx := *((**fasthttp.RequestCtx)(ctxU))
    n, err := ctx.Write(buf)
    if err != nil {
        return 0
    }
    return C.int(n)
}

//export ResponseWriter_WriteHeader
func ResponseWriter_WriteHeader(wPtr C.long, header C.int) {
    ctx, ok := cpointers.Deref(wPtr)
    if !ok {
        return
    }
    (*(**fasthttp.RequestCtx)(ctx)).Response.Header.Set("Content-type", "application/json")
    (*(**fasthttp.RequestCtx)(ctx)).Response.Header.Set("status", "1")
}

func main() {}