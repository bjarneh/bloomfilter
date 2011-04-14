include $(GOROOT)/src/Make.inc

TARG=github.com/bjarneh/bloomfilter
GOFILES=\
	bloomfilter.go

include $(GOROOT)/src/Make.pkg
