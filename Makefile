include $(GOROOT)/src/Make.inc

TARG=glomp
GOFILES=\
	glomp.go\
	defaults.go

include $(GOROOT)/src/Make.cmd
