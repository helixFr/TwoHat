package main

import (
    "C"
    "sync"
    "unsafe"
)

func PtrProxy() *ptrProxy {
    return &ptrProxy{
        lookup: map[int]unsafe.Pointer{},
    }
}

type ptrProxy struct {
    sync.Mutex
    count  int
    lookup map[int]unsafe.Pointer
}

// Ref registers the given pointer and returns a corresponding id that can be
// used to retrieve it later.
func (p *ptrProxy) Ref(ptr unsafe.Pointer) C.long {
    p.Lock()
    id := p.count
    p.count++
    p.lookup[id] = ptr
    p.Unlock()
    return C.long(id)
}

// Deref takes an id and returns the corresponding pointer if it exists.
func (p *ptrProxy) Deref(id C.long) (unsafe.Pointer, bool) {
    p.Lock()
    val, ok := p.lookup[int(id)]
    p.Unlock()
    return val, ok
}

// Free releases a registered pointer by its id.
func (p *ptrProxy) Free(id C.long) {
    p.Lock()
    delete(p.lookup, int(id))
    p.Unlock()
}