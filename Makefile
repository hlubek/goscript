include $(GOROOT)/src/Make.inc

TARG=goarch
GOFILES=goarch.go

install:
	cp goarch goscript $(GOROOT)/bin

include $(GOROOT)/src/Make.cmd

