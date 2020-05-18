#define Py_LIMITED_API
#include <Python.h>

PyObject * loadJson(PyObject *, PyObject *);

static PyMethodDef backendMethods[] = {
    {"loadJson", loadJson, METH_VARARGS, "Add two numbers."},
    {NULL, NULL, 0, NULL}
};

static struct PyModuleDef backendmodule = {
   PyModuleDef_HEAD_INIT, "backend", NULL, -1, backendMethods
};

PyMODINIT_FUNC
PyInit_backend(void)
{
    return PyModule_Create(&backendmodule);
}