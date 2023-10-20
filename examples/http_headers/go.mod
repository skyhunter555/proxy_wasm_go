module github.com/tetratelabs/proxy-wasm-go-sdk/examples/http_headers

go 1.19

replace github.com/tetratelabs/proxy-wasm-go-sdk => ../..

require (
	github.com/stretchr/testify v1.8.2
	github.com/tetratelabs/proxy-wasm-go-sdk v0.0.0-00010101000000-000000000000
	github.com/tidwall/gjson v1.14.3
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/tetratelabs/wazero v1.0.1 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	github.com/google/uuid v1.1.2
	github.com/envoyproxy/go-control-plane v0.10.1
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529
)