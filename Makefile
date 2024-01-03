clean:
	-rm -rf ./build
gobuild:
	go build -o ./build/ntt
outdir:
	mkdir ./build
build: clean outdir gobuild
