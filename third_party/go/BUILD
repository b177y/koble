go_module(
    name = "validator.v10",
    module = "github.com/go-playground/validator/v10",
    version = "v10.11.1",
    visibility = ["PUBLIC"],
    deps = [
        ":crypto",
        ":go-urn",
        ":text",
        ":universal-translator",
    ],
)

go_module(
    name = "crypto",
    install = ["sha3"],
    module = "golang.org/x/crypto",
    version = "v0.0.0-20211215153901-e495a2d5b3d3",
    visibility = ["PUBLIC"],
)

go_module(
    name = "prettyjson",
    module = "github.com/hokaccha/go-prettyjson",
    version = "v0.0.0-20211117102719-0474bc63780f",
    deps = ["//third_party/go/github.com/fatih/color"],
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
    name = "go-urn",
    module = "github.com/leodido/go-urn",
    version = "v1.2.1",
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
    version = "v0.3.7",
    visibility = ["PUBLIC"],
)

go_module(
    name = "tablewriter",
    module = "github.com/olekukonko/tablewriter",
    version = "v0.0.5",
    visibility = ["PUBLIC"],
    deps = [":go-runewidth"],
)

go_module(
    name = "go-runewidth",
    module = "github.com/mattn/go-runewidth",
    version = "v0.0.9",
    visibility = ["PUBLIC"],
)
