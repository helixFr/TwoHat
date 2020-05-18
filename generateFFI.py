#!/usr/bin/python
from cffi import FFI
ffibuilder = FFI()

ffibuilder.set_source("pyBackend",
    """ //passed to the real C compiler
        #include "backend.h"
        #include <Python.h>
    """,
    extra_objects=["backend.so"])

ffibuilder.cdef("""
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
void Call_HandleFunc(ResponseWriterPtr w, Request *r, FuncPtr *fn);
void ListenAndServe(char* p0);
void HandleFunc(char* p0, FuncPtr* p1);
int ResponseWriter_Write(unsigned int p0, char* p1, int p2);
void ResponseWriter_WriteHeader(unsigned int p0, int p1);
long loadJson();
char* returnFromJson(long wPtr, char* word);
""")

if __name__ == "__main__":
    ffibuilder.compile(verbose=True)