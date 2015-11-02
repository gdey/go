package runtime

import "unsafe"

// Don't split the stack.
//go:nosplit
func sysAlloc(u uintptr, sysStat *uint64) unsafe.Pointer {
	print("runtime/mem_xen.go:7 sysAlloc not implemented.\n")
	exit(2)
	return nil
}

//go:nosplit
func sysFree(v unsafe.Pointer, n uintptr, sysStat *uint64) {
	print("runtime/mem_xen.go:15 sysFree not implemented.\n")
	exit(2)
}

//go:nosplit
func sysReserve(v unsafe.Pointer, n uintptr, reserved *bool) (p unsafe.Pointer) {
	print("runtime/mem_xen.go:21 sysReserve not implemented.\n")
	exit(2)
	return p
}

func sysMap(v unsafe.Pointer, n uintptr, reserved bool, sysStat *uint64) {
	print("runtime/mem_xen.go:27 sysMap not implemented.\n")
	exit(2)
}

func sysFault(v unsafe.Pointer, n uintptr) {
	print("runtime/mem_xen.go:32 sysFault not implemented.\n")
	exit(2)
}

func sysUsed(v unsafe.Pointer, n uintptr) {
	print("runtime/mem_xen.go:37 sysUsed not implemented.\n")
	exit(2)
}

func sysUnused(v unsafe.Pointer, n uintptr) {
	print("runtime/mem_xen.go:42 sysUnused not implemented.\n")
	exit(2)
}
