package runtime

type mOS struct{}

//go:nosplit
func semacreate() (p uintptr) {
	print("runtime/os_xen.go:7 semacreate not implemented!")
	exit(2)
	return
}

func osyield() { /* nop */ }
