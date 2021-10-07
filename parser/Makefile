.PHONY: build clean install-deps

install-deps:
	opam install . --deps-only --locked

clean:
	rm -f fintracts.exe
	dune clean

build:
	rm -f fintracts.exe
	dune build
	cp ./_build/default/fintracts.exe ./fintracts.exe
