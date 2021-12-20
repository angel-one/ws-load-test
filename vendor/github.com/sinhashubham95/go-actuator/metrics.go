package actuator

import (
	"net/http"
	"runtime"
)

// BySizeElement reports per-size class allocation statistics.
// BySize[N] gives statistics for allocations of size S where
// BySize[N-1].Size < S ≤ BySize[N].Size.
// This does not report allocations larger than BySize[60].Size.
type BySizeElement struct {
	// Size is the maximum byte size of an object in this
	// size class.
	Size uint32

	// M-allocations is the cumulative count of heap objects
	// allocated in this size class. The cumulative bytes
	// of allocation is Size * M-allocations. The number of live
	// objects in this size class is M-allocations - Frees.
	MAllocations uint64

	// Frees is the cumulative count of heap objects freed
	// in this size class.
	Frees uint64
}

// MemStats is the memory statistics for the current running application
type MemStats struct {
	// Alloc is bytes of allocated heap objects.
	Alloc uint64 `json:"alloc"`

	// TotalAlloc is cumulative bytes allocated for heap objects.
	// This increases as heap objects are allocated, but
	// unlike Alloc, it does not decrease when
	// objects are freed.
	TotalAlloc uint64 `json:"totalAlloc"`

	// Sys is the total bytes of memory obtained from the OS.
	// It is the sum of the XSys fields below. It measures the
	// virtual address space reserved by the Go runtime for the
	// heap, stacks, and other internal data structures. It's
	// likely that not all of the virtual address space is backed
	// by physical memory at any given moment, though in general
	// it all was at some point.
	Sys uint64 `json:"sys"`

	// Lookups is the number of pointer lookups performed by the
	// runtime.
	// This is primarily useful for debugging runtime internals.
	Lookups uint64 `json:"lookups"`

	// MAllocations is the cumulative count of heap objects allocated.
	// The number of live objects is MAllocations - Frees.
	MAllocations uint64 `json:"mAllocations"`

	// Frees is the cumulative count of heap objects freed.
	Frees uint64 `json:"frees"`

	// HeapAlloc is bytes of allocated heap objects.
	// "Allocated" heap objects include all reachable objects, as
	// well as unreachable objects that the garbage collector has
	// not yet freed. Specifically, HeapAlloc increases as heap
	// objects are allocated and decreases as the heap is swept
	// and unreachable objects are freed. Sweeping occurs
	// incrementally between GC cycles, so these two processes
	// occur simultaneously, and as a result HeapAlloc tends to
	// change smoothly (in contrast with the sawtooth that is
	// typical of stop-the-world garbage collectors).
	HeapAlloc uint64 `json:"heapAlloc"`

	// HeapSys is bytes of heap memory obtained from the OS.
	// It measures the amount of virtual address space
	// reserved for the heap. This includes virtual address space
	// that has been reserved but not yet used, which consumes no
	// physical memory, but tends to be small, as well as virtual
	// address space for which the physical memory has been
	// returned to the OS after it became unused (see HeapReleased
	// for a measure of the latter).
	// It estimates the largest size the heap has had.
	HeapSys uint64 `json:"heapSys"`

	// HeapIdle is bytes in idle (unused) spans.
	// Idle spans have no objects in them. These spans could be
	// (and may already have been) returned to the OS, or they can
	// be reused for heap allocations, or they can be reused as
	// stack memory.
	// HeapIdle minus HeapReleased estimates the amount of memory
	// that could be returned to the OS, but is being retained by
	// the runtime, so it can grow the heap without requesting more
	// memory from the OS. If this difference is significantly
	// larger than the heap size, it indicates there was a recent
	// transient spike in live heap size.
	HeapIdle uint64 `json:"heapIdle"`

	// HeapInuse is bytes in in-use spans.
	// In-use spans have at least one object in them. These spans
	// can only be used for other objects of roughly the same
	// size.
	// HeapInuse minus HeapAlloc estimates the amount of memory
	// that has been dedicated to particular size classes, but is
	// not currently being used. This is an upper bound on
	// fragmentation, but in general this memory can be reused
	// efficiently.
	HeapInuse uint64 `json:"heapInUse"`

	// HeapReleased is bytes of physical memory returned to the OS.
	// This counts heap memory from idle spans that was returned
	// to the OS and has not yet been reacquired for the heap.
	HeapReleased uint64 `json:"heapReleased"`

	// HeapObjects is the number of allocated heap objects.
	// Like HeapAlloc, this increases as objects are allocated and
	// decreases as the heap is swept and unreachable objects are
	// freed.
	HeapObjects uint64 `json:"heapObjects"`

	// StackInuse is bytes in stack spans.
	// In-use stack spans have at least one stack in them. These
	// spans can only be used for other stacks of the same size.
	// There is no StackIdle because unused stack spans are
	// returned to the heap (and hence counted toward HeapIdle).
	StackInuse uint64 `json:"stackInUse"`

	// StackSys is bytes of stack memory obtained from the OS.
	// StackSys is StackInuse, plus any memory obtained directly
	// from the OS for OS thread stacks (which should be minimal).
	StackSys uint64 `json:"stackSys"`

	// MSpanInuse is bytes of allocated m-span structures.
	MSpanInuse uint64 `json:"mSpanInUse"`

	// MSpanSys is bytes of memory obtained from the OS for m-span
	// structures.
	MSpanSys uint64 `json:"mSpanSys"`

	// MCacheInuse is bytes of allocated m-cache structures.
	MCacheInuse uint64 `json:"MCacheInUse"`

	// MCacheSys is bytes of memory obtained from the OS for
	// m-cache structures.
	MCacheSys uint64 `json:"mCacheSys"`

	// BuckHashSys is bytes of memory in profiling bucket hash tables.
	BuckHashSys uint64 `json:"buckHashSys"`

	// GCSys is bytes of memory in garbage collection metadata.
	GCSys uint64 `json:"gcSys"`

	// OtherSys is bytes of memory in miscellaneous off-heap
	// runtime allocations.
	OtherSys uint64 `json:"otherSys"`

	// NextGC is the target heap size of the next GC cycle.
	// The garbage collector's goal is to keep HeapAlloc ≤ NextGC.
	// At the end of each GC cycle, the target for the next cycle
	// is computed based on the amount of reachable data and the
	// value of GO GC.
	NextGC uint64 `json:"nextGC"`

	// LastGC is the time the last garbage collection finished, as
	// nanoseconds since 1970 (the UNIX epoch).
	LastGC uint64 `json:"lastGC"`

	// PauseTotalNs is the cumulative nanoseconds in GC
	// stop-the-world pauses since the program started.
	// During a stop-the-world pause, all goroutines are paused
	// and only the garbage collector can run.
	PauseTotalNs uint64 `json:"pauseTotalNs"`

	// PauseNs is a circular buffer of recent GC stop-the-world
	// pause times in nanoseconds.
	//
	// The most recent pause is at PauseNs[(NumGC+255)%256]. In
	// general, PauseNs[N%256] records the time paused in the most
	// recent N%256th GC cycle. There may be multiple pauses per
	// GC cycle; this is the sum of all pauses during a cycle.
	PauseNs [256]uint64 `json:"pauseNs"`

	// PauseEnd is a circular buffer of recent GC pause end times,
	// as nanoseconds since 1970 (the UNIX epoch).
	// This buffer is filled the same way as PauseNs. There may be
	// multiple pauses per GC cycle; this records the end of the
	// last pause in a cycle.
	PauseEnd [256]uint64 `json:"pauseEnd"`

	// NumGC is the number of completed GC cycles.
	NumGC uint32 `json:"numGC"`

	// NumForcedGC is the number of GC cycles that were forced by
	// the application calling the GC function.
	NumForcedGC uint32 `json:"numForcedGC"`

	// GCCPUFraction is the fraction of this program's available
	// CPU time used by the GC since the program started.
	//
	// GCCPUFraction is expressed as a number between 0 and 1,
	// where 0 means GC has consumed none of this program's CPU. A
	// program's available CPU time is defined as the integral of
	// GO MAX PROCESSES since the program started. That is, if
	// GO MAX PROCESSES is 2 and a program has been running for 10
	// seconds, its "available CPU" is 20 seconds. GCCPUFraction
	// does not include CPU time used for write barrier activity.
	//
	// This is the same as the fraction of CPU reported by
	// GO DEBUG = gc trace = 1.
	GCCPUFraction float64 `json:"gcCPUFraction"`

	// EnableGC indicates that GC is enabled. It is always true,
	// even if GO GC = off.
	EnableGC bool `json:"enableGC"`

	// DebugGC is currently unused.
	DebugGC bool `json:"debugGC"`

	// BySize reports per-size class allocation statistics.
	// BySize[N] gives statistics for allocations of size S where
	// BySize[N-1].Size < S ≤ BySize[N].Size.
	// This does not report allocations larger than BySize[60].Size.
	BySize []BySizeElement
}

// MetricsResponse is the response for the metrics endpoint
type MetricsResponse struct {
	MemStats MemStats `json:"memory"`
}

func getRuntimeMetrics() MetricsResponse {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	bySize := make([]BySizeElement, 0, len(memStats.BySize))
	for _, size := range memStats.BySize {
		bySize = append(bySize, BySizeElement{
			Size:         size.Size,
			MAllocations: size.Mallocs,
			Frees:        size.Frees,
		})
	}

	return MetricsResponse{
		MemStats: MemStats{
			Alloc:         memStats.Alloc,
			TotalAlloc:    memStats.TotalAlloc,
			Sys:           memStats.Sys,
			Lookups:       memStats.Lookups,
			MAllocations:  memStats.Mallocs,
			Frees:         memStats.Frees,
			HeapAlloc:     memStats.HeapAlloc,
			HeapSys:       memStats.HeapSys,
			HeapIdle:      memStats.HeapIdle,
			HeapInuse:     memStats.HeapInuse,
			HeapReleased:  memStats.HeapReleased,
			HeapObjects:   memStats.HeapObjects,
			StackInuse:    memStats.StackInuse,
			StackSys:      memStats.StackSys,
			MSpanInuse:    memStats.MSpanInuse,
			MSpanSys:      memStats.MSpanSys,
			MCacheInuse:   memStats.MCacheInuse,
			MCacheSys:     memStats.MCacheSys,
			BuckHashSys:   memStats.BuckHashSys,
			GCSys:         memStats.GCSys,
			OtherSys:      memStats.OtherSys,
			NextGC:        memStats.NextGC,
			LastGC:        memStats.LastGC,
			PauseTotalNs:  memStats.PauseTotalNs,
			PauseNs:       memStats.PauseNs,
			PauseEnd:      memStats.PauseEnd,
			NumGC:         memStats.NumGC,
			NumForcedGC:   memStats.NumForcedGC,
			GCCPUFraction: memStats.GCCPUFraction,
			EnableGC:      memStats.EnableGC,
			DebugGC:       memStats.DebugGC,
			BySize:        bySize,
		},
	}
}

// handleMetrics is the handler function for the metrics endpoint
func handleMetrics(writer http.ResponseWriter, _ *http.Request) {
	body, err := encodeJSONFunction(getRuntimeMetrics())
	if err != nil {
		// some error occurred
		// send the error in the response
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Header().Add(contentTypeHeader, textStringContentType)
		_, _ = writer.Write([]byte(err.Error()))
		return
	}
	// now once we have the correct response
	writer.Header().Add(contentTypeHeader, applicationJSONContentType)
	_, _ = writer.Write(body)
}
