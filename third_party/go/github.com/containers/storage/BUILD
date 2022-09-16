go_mod_download(
    name = "github.com__containers__storage",
    _tag = "download",
    module = "github.com/containers/storage",
    strip = ["pkg/unshare/unshare_cgo.go"],
    version = "v1.42.0",
)

go_module(
    name = "storage",
    download = ":_github.com__containers__storage#download",
    install = ["."],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = ["//third_party/go/golang.org/x/sys"],
)
