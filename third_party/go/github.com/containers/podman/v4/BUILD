go_module(
    name = "v4",
    install = [
        "pkg/bindings",
        "pkg/errorhandling",
        "pkg/terminal",
        "version",
    ],
    module = "github.com/containers/podman/v4",
    version = "v4.2.1",
    visibility = ["PUBLIC"],
    deps = [
        "//third_party/go/github.com/blang/semver/v4",
        "//third_party/go/github.com/containers/storage",
        "//third_party/go/github.com/hashicorp/go-multierror",
        "//third_party/go/github.com/sirupsen/logrus",
        "//third_party/go/golang.org/x/crypto",
        "//third_party/go/golang.org/x/term",
    ],
)
