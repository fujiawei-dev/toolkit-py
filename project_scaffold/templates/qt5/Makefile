{{HASHTAG_COMMENTS}}

.PHONY: ;
#.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell
.IGNORE: clean;            # ignore all errors, keep going

all: clean

info:
	@echo $(SHELL)

init:
	git init
	git add .
	git commit -m "first commit"

clean:
	-(mkdir $(CMAKE_BUILD_PREFIX)-debug)
	-(cd $(CMAKE_BUILD_PREFIX)-debug; rm -r *)

	-(mkdir $(CMAKE_BUILD_PREFIX)-release)
	-(cd $(CMAKE_BUILD_PREFIX)-release; rm -r *)
