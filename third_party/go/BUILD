go_module(
    name = "logrus",
    module = "github.com/sirupsen/logrus",
    version = "v1.8.1",
    visibility = ["PUBLIC"],
    deps = [":sys"],
)

go_module(
    name = "sys",
    install = [
        "internal/unsafeheader",
        "unix",
    ],
    module = "golang.org/x/sys",
    version = "v0.0.0-20210616094352-59db8d763f22",
    visibility = ["PUBLIC"],
)

go_module(
    name = "uilive",
    module = "github.com/gosuri/uilive",
    version = "v0.0.4",
    visibility = ["PUBLIC"],
)

go_module(
    name = "term",
    module = "github.com/moby/term",
    version = "v0.0.0-20210619224110-3f7ff695adc6",
    visibility = ["PUBLIC"],
    deps = [":sys"],
)

go_module(
    name = "plugins",
    install = ["pkg/ns"],
    module = "github.com/containernetworking/plugins",
    version = "v1.0.1",
    visibility = ["PUBLIC"],
    deps = [":sys"],
)

go_module(
    name = "cobra",
    module = "github.com/spf13/cobra",
    version = "v1.3.0",
    visibility = ["PUBLIC"],
    deps = [":pflag"],
)

go_module(
    name = "pflag",
    module = "github.com/spf13/pflag",
    version = "v1.0.5",
    visibility = ["PUBLIC"],
)

go_module(
    name = "crypto",
    install = [
        "sha3",
        "ssh/terminal",
    ],
    module = "golang.org/x/crypto",
    version = "v0.0.0-20220112180741-5e0467b6c7ce",
    visibility = ["PUBLIC"],
    deps = [":x.term"],
)

go_module(
    name = "x.term",
    module = "golang.org/x/term",
    version = "v0.0.0-20201126162022-7de9c90e9dd1",
    visibility = ["PUBLIC"],
    deps = [":sys"],
)

go_module(
    name = "color",
    module = "github.com/fatih/color",
    version = "v1.13.0",
    visibility = ["PUBLIC"],
    deps = [
        ":go-colorable",
        ":go-isatty",
    ],
)

go_module(
    name = "go-isatty",
    module = "github.com/mattn/go-isatty",
    version = "v0.0.14",
    visibility = ["PUBLIC"],
    deps = [":sys"],
)

go_module(
    name = "go-colorable",
    module = "github.com/mattn/go-colorable",
    version = "v0.1.9",
    visibility = ["PUBLIC"],
    deps = [":go-isatty"],
)

go_module(
    name = "ocicni",
    install = ["pkg/ocicni"],
    module = "github.com/cri-o/ocicni",
    version = "v0.2.0",
    visibility = ["PUBLIC"],
    deps = [
        ":cni",
        ":fsnotify",
        ":logrus",
    ],
)

go_module(
    name = "cni",
    install = [
        "libcni",
        "pkg/invoke",
        "pkg/types",
        "pkg/types/020",
        "pkg/types/current",
        "pkg/utils",
        "pkg/version",
    ],
    module = "github.com/containernetworking/cni",
    version = "v0.7.2-0.20190904153231-83439463f784",
    visibility = ["PUBLIC"],
)

go_module(
    name = "fsnotify",
    module = "github.com/fsnotify/fsnotify",
    version = "v1.4.7",
    visibility = ["PUBLIC"],
    deps = [":sys"],
)

go_module(
    name = "runtime-spec",
    install = ["specs-go"],
    module = "github.com/opencontainers/runtime-spec",
    version = "v1.0.2",
    visibility = ["PUBLIC"],
)

go_module(
    name = "go-urn",
    module = "github.com/leodido/go-urn",
    version = "v1.2.1",
    visibility = ["PUBLIC"],
)

go_module(
    name = "universal-translator",
    module = "github.com/go-playground/universal-translator",
    version = "v0.18.0",
    visibility = ["PUBLIC"],
    deps = [":locales"],
)

go_module(
    name = "locales",
    install = [
        ".",
        "currency",
    ],
    module = "github.com/go-playground/locales",
    version = "v0.14.0",
    visibility = ["PUBLIC"],
)

go_module(
    name = "text",
    install = [
        "internal/language",
        "internal/language/compact",
        "internal/tag",
        "language",
    ],
    module = "golang.org/x/text",
    version = "v0.3.6",
    visibility = ["PUBLIC"],
)

go_module(
    name = "validator.v10",
    module = "github.com/go-playground/validator/v10",
    version = "v10.10.0",
    visibility = ["PUBLIC"],
    deps = [
        ":crypto",
        ":go-urn",
        ":text",
        ":universal-translator",
    ],
)

go_module(
    name = "mapstructure",
    module = "github.com/mitchellh/mapstructure",
    version = "v1.4.1",
    visibility = ["PUBLIC"],
)

go_module(
    name = "copystructure",
    module = "github.com/mitchellh/copystructure",
    version = "v1.2.0",
    visibility = ["PUBLIC"],
    deps = [":reflectwalk"],
)

go_module(
    name = "koanf",
    install = [
        ".",
        "maps",
        "parsers/yaml",
        "providers/confmap",
        "providers/file",
    ],
    module = "github.com/knadh/koanf",
    version = "v1.4.0",
    visibility = ["PUBLIC"],
    deps = [
        ":copystructure",
        ":fsnotify",
        ":mapstructure",
        ":yaml.v2",
    ],
)

go_module(
    name = "reflectwalk",
    module = "github.com/mitchellh/reflectwalk",
    version = "v1.0.2",
    visibility = ["PUBLIC"],
)

go_module(
    name = "yaml.v2",
    module = "gopkg.in/yaml.v2",
    version = "v2.2.8",
    visibility = ["PUBLIC"],
)
