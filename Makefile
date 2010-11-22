GOROOT=/home/victor/sources/go
include $(GOROOT)/src/Make.inc

TARG=mypackage
GOFILES=\
	about_testing_test.go

include $(GOROOT)/src/Make.pkg

