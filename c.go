package tox

/*
#cgo CFLAGS: -g -O2 -std=c99 -Wall
// #cgo LDFLAGS: -ltoxcore -lvpx -lopus -lsodium -lm
// #cgo pkg-config: toxcore
// Needed because toxcore declares -lvpx in its Libs.private:
// #cgo LDFLAGS: -lvpx
// #cgo LDFLAGS: -L/home/gzleo/oss/toxcore/build/.libs/

   #cgo CFLAGS: -I${SRCDIR}/include
   #cgo LDFLAGS: -L${SRCDIR}/lib
   #cgo LDFLAGS: -ltoxcore -ltoxav -ltoxencryptsave -lvpx -lopus -lsodium -lm
*/
import "C"

// TODO what about Windows/MacOS?
