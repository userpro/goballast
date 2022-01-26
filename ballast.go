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
	runtime.SetFinalizer(f, finalizerHandler)
}

var f *finalizer

// New mem目标内存大小 单位bytes
func New(mem int) {
	f = &finalizer{
		ballast: make([]byte, mem),
	}

	f.ref = &finalizerRef{parent: f}
	runtime.SetFinalizer(f.ref, finalizerHandler)
	f.ref = nil
}

// NewWithDebug mem目标内存大小 单位bytes
func NewWithDebug(mem int) {
	f = &finalizer{
		ballast: make([]byte, mem),
	}

	f.ref = &finalizerRef{parent: f}
	runtime.SetFinalizer(f.ref, finalizerDebugHandler)
	f.ref = nil
}
