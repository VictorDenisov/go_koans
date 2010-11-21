GOROOT=/home/victor/sources/go
include $(GOROOT)/src/Make.inc

TARG=mypackage
GOFILES=\
	about_asserts_test.go

include $(GOROOT)/src/Make.pkg

