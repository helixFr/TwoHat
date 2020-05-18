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
long loadJson();
char* returnFromJson(long wPtr, char* word);
""")

if __name__ == "__main__":
    ffibuilder.compile(verbose=True)