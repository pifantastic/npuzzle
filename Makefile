
srcfiles := main.go puzzle.go board.go priority_queue.go api.go middleware.go bindata_assetfs.go

run:
	go run $(srcfiles)

build-dev:
	go-bindata-assetfs -ignore=.DS_Store -dev static/...

build:
	go-bindata-assetfs -ignore=.DS_Store static/...
	go build $(srcfiles)
