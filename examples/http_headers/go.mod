module github.com/skyhunter555/proxy_wasm_go/examples/http_headers

go 1.16

replace github.com/tetratelabs/proxy-wasm-go-sdk => ../..

require (
	github.com/stretchr/testify v1.8.2
	github.com/tetratelabs/proxy-wasm-go-sdk v0.0.0-00010101000000-000000000000
	github.com/tidwall/gjson v1.14.3
)

require github.com/envoyproxy/go-control-plane v0.10.1
