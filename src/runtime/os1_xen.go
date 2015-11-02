package runtime

var temp_data = []byte("getRandomData not implemented yet for xen…")

func getRandomData(r []byte) {
	n := copy(r, temp_data)
	extendRandom(r, n)
	return
}

//go:nosplit
func semasleep(ns int64) int32 {
	var r int32
	print("runtime/os1_xen.go:14 semasleep not implemented")
	exit(2)
	return r
}

//go:nosplit
func semawakeup(mp *m) {
	print("runtime/os1_xen.go:21 semawakeup not implemented")
	exit(2)
}

// We may want to handle this ourselfs. What does it mean to have environment
// variables for XEN. Maybe a location in xenstore?
func goenvs() {
	goenvs_unix()
}

func mpreinit(mp *m) {
	print("runtime/os1_xen.go:32 mpreinit not implemented")
	exit(2)
}
