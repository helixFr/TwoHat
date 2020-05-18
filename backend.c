#define Py_LIMITED_API
#include <Python.h>

PyObject * loadJson(PyObject *, PyObject *);
PyObject * returnFromJson(PyObject *, PyObject *);

int PyArg_ParseTuple_Ls(PyObject * args, long * a, char ** word) {
    return PyArg_ParseTuple(args, "Ls", a, word);
}

int PyArg_ParseTuple_L(PyObject * args, long * a) {
    return PyArg_ParseTuple(args, "L", a);
}

static PyMethodDef backendMethods[] = {
    {"loadJson", loadJson, METH_VARARGS, "Loads json file."},
    {"returnFromJson", returnFromJson, METH_VARARGS, "Returns topics map from loaded json."},
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