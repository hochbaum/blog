GO = go

.DEFAULT_GOAL = blog

clean:
	@rm -rf build/

setup: clean
	@mkdir -p build
	@cp -r web/* build/

blog: setup
	@$(GO) build -o build/$@ ./cmd/blog/