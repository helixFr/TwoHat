package main

// #cgo pkg-config: python3
// #define Py_LIMITED_API
// #include <Python.h>
import "C"
import (
	"fmt"
)