# N-Puzzle Solver

This is a web implementation of an n-puzzle solver. More information about
n-puzzles can be found [here](https://en.wikipedia.org/wiki/15_puzzle).

## Dependencies

The server is implemented in `Go`. More information on installing `Go` can be
found [here](https://golang.org/doc/install).

If you'd like to run the tests, you'll also need `node`. More information on
installing node can be found [here](https://nodejs.org/download/).

## Installation

If you have `go` installed and have set your `$GOPATH` to your workspace, you
you should be able to run the following:

`go get github.com/pifantastic/npuzzle`
`npuzzle`

Alternatively, you can check out the source and run `make` in the root directory.

## Tests

* `npm install`
* `make test`

## Packaging the binary

The static assets for the web client need to be bundled into the binary. This is
done using the [go-bindata-assetfs](https://github.com/elazarl/go-bindata-assetfs)
project.

* `go get github.com/elazarl/go-bindata-assetfs/...`
* `make build`

This will product an `npuzzle` binary in the root of the project.
