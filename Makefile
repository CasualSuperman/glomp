include $(GOROOT)/src/Make.inc

TARG=glomp
GOFILES=\
	glomp.go\
	defaults.go\
	conn.go\
	gui.go\
	actions.go

include $(GOROOT)/src/Make.cmd
