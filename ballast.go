package goballast

import (
	"log"
	"runtime"
	"time"
)

type finalizer struct {
	ballast []byte
	ref     *finalizerRef
}

type finalizerRef struct {
	parent *finalizer
}

func finalizerHandler(f *finalizerRef) {
	runtime.KeepAlive(f.parent.ballast)
	runtime.SetFinalizer(f, finalizerHandler)
}

func finalizerDebugHandler(f *finalizerRef) {
	m := &runtime.MemStats{}
	runtime.ReadMemStats(m)
	log.Printf("gc: %s, alloc: %d KB, nextGC: %d KB, num_gc: %d",
		time.Now().String(), m.Alloc/1024, m.NextGC/1024, m.NumGC)

	runtime.KeepAlive(f.parent.ballast)
	runtime.SetFinalizer(f, finalizerDebugHandler)
}

var f *finalizer

// New mem*2 = target memory usage for GC trigger, bytes
func New(mem int) {
	f = &finalizer{
		ballast: make([]byte, mem),
	}

	f.ref = &finalizerRef{parent: f}
	runtime.SetFinalizer(f.ref, finalizerHandler)
	f.ref = nil
}

// NewWithDebug mem*2 = target memory usage for GC trigger, bytes
func NewWithDebug(mem int) {
	f = &finalizer{
		ballast: make([]byte, mem),
	}

	f.ref = &finalizerRef{parent: f}
	runtime.SetFinalizer(f.ref, finalizerDebugHandler)
	f.ref = nil
}
