#define Py_LIMITED_API
#include <Python.h>

static struct PyModuleDef foomodule = {
   PyModuleDef_HEAD_INIT, "foo", NULL, -1, FooMethods
};

PyMODINIT_FUNC
PyInit_foo(void)
{
    return PyModule_Create(&foomodule);
}