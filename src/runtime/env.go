//go:build linux || darwin || windows || wasip1

package runtime

// Update the C environment if cgo is loaded.
// Called from Go 1.20 and above.
//
//go:linkname syscallSetenv syscall.runtimeSetenv
func syscallSetenv(key, value string) {
	keydata := cstring(key)
	valdata := cstring(value)
	setenv(&keydata[0], &valdata[0])
	if key == "GODEBUG" && godebugUpdate != nil {
		// Starting with Go 1.20, we need to call a callback (set by
		// internal/godebug) to notify the GODEBUG environment variable has
		// changed. This is necessary to get archive/zip to pass tests.
		godebugUpdate(key, value)
	}
}

// Update the C environment if cgo is loaded.
// Called from Go 1.20 and above.
//
//go:linkname syscallUnsetenv syscall.runtimeUnsetenv
func syscallUnsetenv(key string) {
	keydata := cstring(key)
	unsetenv(&keydata[0])
}

// Compatibility with Go 1.19 and below.
//
//go:linkname syscall_setenv_c syscall.setenv_c
func syscall_setenv_c(key string, val string) {
	syscallSetenv(key, val)
}

// Compatibility with Go 1.19 and below.
//
//go:linkname syscall_unsetenv_c syscall.unsetenv_c
func syscall_unsetenv_c(key string) {
	syscallUnsetenv(key)
}

// cstring converts a Go string to a C string.
// borrowed from syscall
func cstring(s string) []byte {
	data := make([]byte, len(s)+1)
	copy(data, s)
	// final byte should be zero from the initial allocation
	return data
}
