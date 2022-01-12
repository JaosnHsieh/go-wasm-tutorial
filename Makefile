run ::
	npx live-server ./site

build ::
	GOOS=js GOARCH=wasm go build -o ./site/main.wasm ./main-wasm