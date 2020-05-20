import os
from pyBackend import *

_handlers = []


class ResponseWriter:
    def __init__(self, w):
        self._w = w

    def write(self, body):
        n = lib.ResponseWriter_Write(self._w, body, len(body))
        if n != len(body):
            raise IOError("Failed to write to ResponseWriter.")

    def set_status(self, code):
        lib.ResponseWriter_WriteHeader(self._w, code)


class Request:
    def __init__(self, req):
        self._req = req

    @property
    def method(self):
        return ffi.string(self._req.Method)

    @property
    def host(self):
        return ffi.string(self._req.Host)

    @property
    def url(self):
        return ffi.string(self._req.URL)

    @property
    def body(self):
        return ffi.string(self._req.Body)

    @property
    def headers(self):
        return ffi.string(self._req.Headers)

    def __repr__(self):
        return "{self.method} {self.url}".format(self=self)


def route(fn=None):
    def wrapped(fn):
        @ffi.callback("void(ResponseWriterPtr, Request*)")
        def handler(w, req):
            fn(ResponseWriter(w), Request(req))

        lib.HandleFunc(str.encode("/"), handler)
        _handlers.append(handler)

    if fn:
        return wrapped(fn)

    return wrapped