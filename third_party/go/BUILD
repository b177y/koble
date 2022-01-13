config_setting(
    name = "__config_linux_amd64",
    values = {
        "os": "linux",
        "cpu": "amd64",
    },
)

config_setting(
    name = "__config_darwin_amd64",
    values = {
        "os": "darwin",
        "cpu": "amd64",
    },
)

go_mod_download(
    name = "github.com__BurntSushi__toml",
    _tag = "download",
    module = "github.com/BurntSushi/toml",
    version = "v0.4.1",
)

go_module(
    name = "github.com__BurntSushi__toml",
    download = ":_github.com__BurntSushi__toml#download",
    install = ["."],
    module = "github.com/BurntSushi/toml",
    visibility = ["PUBLIC"],
    deps = [":github.com__BurntSushi__toml__internal"],
)

go_module(
    name = "github.com__BurntSushi__toml__internal",
    download = ":_github.com__BurntSushi__toml#download",
    install = ["internal"],
    module = "github.com/BurntSushi/toml",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__VividCortex__ewma",
    _tag = "download",
    module = "github.com/VividCortex/ewma",
    version = "v1.2.0",
)

go_module(
    name = "github.com__VividCortex__ewma",
    download = ":_github.com__VividCortex__ewma#download",
    install = ["."],
    module = "github.com/VividCortex/ewma",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__acarl005__stripansi",
    _tag = "download",
    module = "github.com/acarl005/stripansi",
    version = "v0.0.0-20180116102854-5a71ef0e047d",
)

go_module(
    name = "github.com__acarl005__stripansi",
    download = ":_github.com__acarl005__stripansi#download",
    install = ["."],
    module = "github.com/acarl005/stripansi",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__b177y__go-uml-utilities",
    _tag = "download",
    module = "github.com/b177y/go-uml-utilities",
    version = "9882ca8ecb0d626bcdf4b02b7c48dac7de5403f6",
)

go_module(
    name = "github.com__b177y__go-uml-utilities__pkg__mconsole",
    download = ":_github.com__b177y__go-uml-utilities#download",
    install = ["pkg/mconsole"],
    module = "github.com/b177y/go-uml-utilities",
    visibility = ["PUBLIC"],
    deps = [":github.com__b177y__go-uml-utilities__pkg__mconsole__sysrq"],
)

go_module(
    name = "github.com__b177y__go-uml-utilities__pkg__mconsole__sysrq",
    download = ":_github.com__b177y__go-uml-utilities#download",
    install = ["pkg/mconsole/sysrq"],
    module = "github.com/b177y/go-uml-utilities",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__beorn7__perks",
    _tag = "download",
    module = "github.com/beorn7/perks",
    version = "v1.0.1",
)

go_module(
    name = "github.com__beorn7__perks__quantile",
    download = ":_github.com__beorn7__perks#download",
    install = ["quantile"],
    module = "github.com/beorn7/perks",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__bits-and-blooms__bitset",
    _tag = "download",
    module = "github.com/bits-and-blooms/bitset",
    version = "v1.2.0",
)

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__bits-and-blooms__bitset", download = ":_github.com__bits-and-blooms__bitset#download", install = ["."], module = "github.com/bits-and-blooms/bitset", visibility = ["PUBLIC"], deps = [])

go_mod_download(
    name = "github.com__blang__semver",
    _tag = "download",
    module = "github.com/blang/semver",
    version = "v3.5.1+incompatible",
)

go_module(
    name = "github.com__blang__semver",
    download = ":_github.com__blang__semver#download",
    install = ["."],
    module = "github.com/blang/semver",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__cespare__xxhash__v2",
    _tag = "download",
    module = "github.com/cespare/xxhash/v2",
    version = "v2.1.2",
)

go_module(
    name = "github.com__cespare__xxhash__v2",
    download = ":_github.com__cespare__xxhash__v2#download",
    install = ["."],
    module = "github.com/cespare/xxhash/v2",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__chzyer__readline",
    _tag = "download",
    module = "github.com/chzyer/readline",
    version = "v0.0.0-20180603132655-2972be24d48e",
)

go_module(
    name = "github.com__chzyer__readline",
    download = ":_github.com__chzyer__readline#download",
    install = ["."],
    module = "github.com/chzyer/readline",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__containerd__containerd",
    _tag = "download",
    module = "github.com/containerd/containerd",
    version = "v1.5.7",
)

go_module(
    name = "github.com__containerd__containerd__errdefs",
    download = ":_github.com__containerd__containerd#download",
    install = ["errdefs"],
    module = "github.com/containerd/containerd",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__pkg__errors",
        ":google.golang.org__grpc__codes",
        ":google.golang.org__grpc__status",
    ],
)

go_module(
    name = "github.com__containerd__containerd__log",
    download = ":_github.com__containerd__containerd#download",
    install = ["log"],
    module = "github.com/containerd/containerd",
    visibility = ["PUBLIC"],
    deps = [":github.com__sirupsen__logrus"],
)

go_module(
    name = "github.com__containerd__containerd__platforms",
    download = ":_github.com__containerd__containerd#download",
    install = ["platforms"],
    module = "github.com/containerd/containerd",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containerd__containerd__errdefs",
        ":github.com__containerd__containerd__log",
        ":github.com__opencontainers__image-spec__specs-go__v1",
        ":github.com__pkg__errors",
    ],
)

go_mod_download(
    name = "github.com__containerd__stargz-snapshotter__estargz",
    _tag = "download",
    module = "github.com/containerd/stargz-snapshotter/estargz",
    version = "v0.8.0",
)

go_module(
    name = "github.com__containerd__stargz-snapshotter__estargz",
    download = ":_github.com__containerd__stargz-snapshotter__estargz#download",
    install = ["."],
    module = "github.com/containerd/stargz-snapshotter/estargz",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containerd__stargz-snapshotter__estargz__errorutil",
        ":github.com__klauspost__compress__zstd",
        ":github.com__opencontainers__go-digest",
        ":github.com__pkg__errors",
        ":golang.org__x__sync__errgroup",
    ],
)

go_module(
    name = "github.com__containerd__stargz-snapshotter__estargz__errorutil",
    download = ":_github.com__containerd__stargz-snapshotter__estargz#download",
    install = ["errorutil"],
    module = "github.com/containerd/stargz-snapshotter/estargz",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__containernetworking__cni",
    _tag = "download",
    module = "github.com/containernetworking/cni",
    version = "v0.8.1",
)

go_module(
    name = "github.com__containernetworking__cni__libcni",
    download = ":_github.com__containernetworking__cni#download",
    install = ["libcni"],
    module = "github.com/containernetworking/cni",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containernetworking__cni__pkg__invoke",
        ":github.com__containernetworking__cni__pkg__types",
        ":github.com__containernetworking__cni__pkg__utils",
        ":github.com__containernetworking__cni__pkg__version",
    ],
)

go_module(
    name = "github.com__containernetworking__cni__pkg__invoke",
    download = ":_github.com__containernetworking__cni#download",
    install = ["pkg/invoke"],
    module = "github.com/containernetworking/cni",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containernetworking__cni__pkg__types",
        ":github.com__containernetworking__cni__pkg__version",
    ],
)

go_module(
    name = "github.com__containernetworking__cni__pkg__types",
    download = ":_github.com__containernetworking__cni#download",
    install = ["pkg/types"],
    module = "github.com/containernetworking/cni",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__containernetworking__cni__pkg__types__020",
    download = ":_github.com__containernetworking__cni#download",
    install = ["pkg/types/020"],
    module = "github.com/containernetworking/cni",
    visibility = ["PUBLIC"],
    deps = [":github.com__containernetworking__cni__pkg__types"],
)

go_module(
    name = "github.com__containernetworking__cni__pkg__types__current",
    download = ":_github.com__containernetworking__cni#download",
    install = ["pkg/types/current"],
    module = "github.com/containernetworking/cni",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containernetworking__cni__pkg__types",
        ":github.com__containernetworking__cni__pkg__types__020",
    ],
)

go_module(
    name = "github.com__containernetworking__cni__pkg__utils",
    download = ":_github.com__containernetworking__cni#download",
    install = ["pkg/utils"],
    module = "github.com/containernetworking/cni",
    visibility = ["PUBLIC"],
    deps = [":github.com__containernetworking__cni__pkg__types"],
)

go_module(
    name = "github.com__containernetworking__cni__pkg__version",
    download = ":_github.com__containernetworking__cni#download",
    install = ["pkg/version"],
    module = "github.com/containernetworking/cni",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containernetworking__cni__pkg__types",
        ":github.com__containernetworking__cni__pkg__types__020",
        ":github.com__containernetworking__cni__pkg__types__current",
    ],
)

go_mod_download(
    name = "github.com__containernetworking__plugins",
    _tag = "download",
    module = "github.com/containernetworking/plugins",
    version = "v0.9.1",
)

go_module(
    name = "github.com__containernetworking__plugins__pkg__ns",
    download = ":_github.com__containernetworking__plugins#download",
    install = ["pkg/ns"],
    module = "github.com/containernetworking/plugins",
    visibility = ["PUBLIC"],
    deps = select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [":golang.org__x__sys__unix"],
        "default": [],
    }),
)

go_mod_download(
    name = "github.com__containers__buildah",
    _tag = "download",
    module = "github.com/containers/buildah",
    version = "v1.23.1",
)

go_module(
    name = "github.com__containers__buildah__copier",
    download = ":_github.com__containers__buildah#download",
    install = ["copier"],
    module = "github.com/containers/buildah",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__buildah__util",
        ":github.com__containers__image__v5__pkg__compression",
        ":github.com__containers__storage__pkg__fileutils",
        ":github.com__containers__storage__pkg__idtools",
        ":github.com__containers__storage__pkg__reexec",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
        ":golang.org__x__sys__unix",
    ],
)

go_module(
    name = "github.com__containers__buildah__define",
    download = ":_github.com__containers__buildah#download",
    install = ["define"],
    module = "github.com/containers/buildah",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__manifest",
        ":github.com__containers__image__v5__types",
        ":github.com__containers__ocicrypt__config",
        ":github.com__containers__storage__pkg__archive",
        ":github.com__containers__storage__pkg__chrootarchive",
        ":github.com__containers__storage__pkg__ioutils",
        ":github.com__opencontainers__image-spec__specs-go__v1",
        ":github.com__opencontainers__runc__libcontainer__devices",
        ":github.com__opencontainers__runtime-spec__specs-go",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
        ":golang.org__x__sync__semaphore",
    ],
)

go_module(
    name = "github.com__containers__buildah__util",
    download = ":_github.com__containers__buildah#download",
    install = ["util"],
    module = "github.com/containers/buildah",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__buildah__define",
        ":github.com__containers__common__libimage",
        ":github.com__containers__common__pkg__config",
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__pkg__shortnames",
        ":github.com__containers__image__v5__pkg__sysregistriesv2",
        ":github.com__containers__image__v5__signature",
        ":github.com__containers__image__v5__transports__alltransports",
        ":github.com__containers__image__v5__types",
        ":github.com__containers__storage",
        ":github.com__docker__distribution__registry__api__errcode",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__runtime-spec__specs-go",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ] + select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [":golang.org__x__sys__unix"],
        "default": [],
    }),
)

go_mod_download(
    name = "github.com__containers__common",
    _tag = "download",
    module = "github.com/containers/common",
    version = "v0.44.2",
)

go_module(
    name = "github.com__containers__common__libimage",
    download = ":_github.com__containers__common#download",
    install = ["libimage"],
    module = "github.com/containers/common",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__common__libimage__manifests",
        ":github.com__containers__common__pkg__config",
        ":github.com__containers__common__pkg__filters",
        ":github.com__containers__common__pkg__retry",
        ":github.com__containers__common__pkg__signal",
        ":github.com__containers__common__pkg__timetype",
        ":github.com__containers__image__v5__copy",
        ":github.com__containers__image__v5__directory",
        ":github.com__containers__image__v5__docker",
        ":github.com__containers__image__v5__docker__archive",
        ":github.com__containers__image__v5__docker__daemon",
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__manifest",
        ":github.com__containers__image__v5__oci__archive",
        ":github.com__containers__image__v5__oci__layout",
        ":github.com__containers__image__v5__pkg__shortnames",
        ":github.com__containers__image__v5__pkg__sysregistriesv2",
        ":github.com__containers__image__v5__signature",
        ":github.com__containers__image__v5__storage",
        ":github.com__containers__image__v5__tarball",
        ":github.com__containers__image__v5__transports__alltransports",
        ":github.com__containers__image__v5__types",
        ":github.com__containers__ocicrypt__config",
        ":github.com__containers__storage",
        ":github.com__disiqueira__gotree__v3",
        ":github.com__docker__go-units",
        ":github.com__hashicorp__go-multierror",
        ":github.com__jinzhu__copier",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__image-spec__specs-go__v1",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
        ":golang.org__x__sync__semaphore",
    ],
)

go_module(
    name = "github.com__containers__common__libimage__manifests",
    download = ":_github.com__containers__common#download",
    install = ["libimage/manifests"],
    module = "github.com/containers/common",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__common__pkg__manifests",
        ":github.com__containers__common__pkg__supplemented",
        ":github.com__containers__image__v5__copy",
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__image",
        ":github.com__containers__image__v5__manifest",
        ":github.com__containers__image__v5__signature",
        ":github.com__containers__image__v5__storage",
        ":github.com__containers__image__v5__transports",
        ":github.com__containers__image__v5__transports__alltransports",
        ":github.com__containers__image__v5__types",
        ":github.com__containers__storage",
        ":github.com__containers__storage__pkg__lockfile",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__image-spec__specs-go__v1",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__common__pkg__apparmor",
    download = ":_github.com__containers__common#download",
    install = ["pkg/apparmor"],
    module = "github.com/containers/common",
    visibility = ["PUBLIC"],
    deps = [":github.com__containers__common__version"],
)

go_module(
    name = "github.com__containers__common__pkg__capabilities",
    download = ":_github.com__containers__common#download",
    install = ["pkg/capabilities"],
    module = "github.com/containers/common",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__pkg__errors",
        ":github.com__syndtr__gocapability__capability",
    ],
)

go_module(
    name = "github.com__containers__common__pkg__cgroupv2",
    download = ":_github.com__containers__common#download",
    install = ["pkg/cgroupv2"],
    module = "github.com/containers/common",
    visibility = ["PUBLIC"],
    deps = select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [":golang.org__x__sys__unix"],
        "default": [],
    }),
)

go_module(
    name = "github.com__containers__common__pkg__config",
    download = ":_github.com__containers__common#download",
    install = ["pkg/config"],
    module = "github.com/containers/common",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__BurntSushi__toml",
        ":github.com__containers__common__pkg__apparmor",
        ":github.com__containers__common__pkg__capabilities",
        ":github.com__containers__common__pkg__cgroupv2",
        ":github.com__containers__storage__pkg__homedir",
        ":github.com__containers__storage__pkg__unshare",
        ":github.com__containers__storage__types",
        ":github.com__docker__go-units",
        ":github.com__opencontainers__selinux__go-selinux",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ] + select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [":golang.org__x__sys__unix"],
        "default": [],
    }),
)

go_module(
    name = "github.com__containers__common__pkg__filters",
    download = ":_github.com__containers__common#download",
    install = ["pkg/filters"],
    module = "github.com/containers/common",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__common__pkg__timetype",
        ":github.com__pkg__errors",
    ],
)

go_module(
    name = "github.com__containers__common__pkg__manifests",
    download = ":_github.com__containers__common#download",
    install = ["pkg/manifests"],
    module = "github.com/containers/common",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__manifest",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__image-spec__specs-go",
        ":github.com__opencontainers__image-spec__specs-go__v1",
        ":github.com__pkg__errors",
    ],
)

go_module(
    name = "github.com__containers__common__pkg__parse",
    download = ":_github.com__containers__common#download",
    install = ["pkg/parse"],
    module = "github.com/containers/common",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__storage__pkg__unshare",
        ":github.com__opencontainers__runc__libcontainer__devices",
        ":github.com__pkg__errors",
    ],
)

go_module(
    name = "github.com__containers__common__pkg__retry",
    download = ":_github.com__containers__common#download",
    install = ["pkg/retry"],
    module = "github.com/containers/common",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__docker__distribution__registry__api__errcode",
        ":github.com__docker__distribution__registry__api__v2",
        ":github.com__hashicorp__go-multierror",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__common__pkg__signal",
    download = ":_github.com__containers__common#download",
    install = ["pkg/signal"],
    module = "github.com/containers/common",
    visibility = ["PUBLIC"],
    deps = select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [":golang.org__x__sys__unix"],
        "default": [],
    }),
)

go_module(
    name = "github.com__containers__common__pkg__supplemented",
    download = ":_github.com__containers__common#download",
    install = ["pkg/supplemented"],
    module = "github.com/containers/common",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__common__pkg__manifests",
        ":github.com__containers__image__v5__copy",
        ":github.com__containers__image__v5__image",
        ":github.com__containers__image__v5__manifest",
        ":github.com__containers__image__v5__transports",
        ":github.com__containers__image__v5__types",
        ":github.com__hashicorp__go-multierror",
        ":github.com__opencontainers__go-digest",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__common__pkg__timetype",
    download = ":_github.com__containers__common#download",
    install = ["pkg/timetype"],
    module = "github.com/containers/common",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__containers__common__version",
    download = ":_github.com__containers__common#download",
    install = ["version"],
    module = "github.com/containers/common",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__containers__image__v5",
    _tag = "download",
    module = "github.com/containers/image/v5",
    version = "v5.16.0",
)

go_module(
    name = "github.com__containers__image__v5__copy",
    download = ":_github.com__containers__image__v5#download",
    install = ["copy"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__image",
        ":github.com__containers__image__v5__internal__blobinfocache",
        ":github.com__containers__image__v5__internal__pkg__platform",
        ":github.com__containers__image__v5__internal__types",
        ":github.com__containers__image__v5__manifest",
        ":github.com__containers__image__v5__pkg__blobinfocache",
        ":github.com__containers__image__v5__pkg__compression",
        ":github.com__containers__image__v5__pkg__compression__types",
        ":github.com__containers__image__v5__signature",
        ":github.com__containers__image__v5__transports",
        ":github.com__containers__image__v5__types",
        ":github.com__containers__ocicrypt",
        ":github.com__containers__ocicrypt__config",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__image-spec__specs-go__v1",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
        ":github.com__vbauerster__mpb__v7",
        ":github.com__vbauerster__mpb__v7__decor",
        ":golang.org__x__sync__semaphore",
        ":golang.org__x__term",
    ],
)

go_module(
    name = "github.com__containers__image__v5__directory",
    download = ":_github.com__containers__image__v5#download",
    install = ["directory"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__directory__explicitfilepath",
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__image",
        ":github.com__containers__image__v5__internal__putblobdigest",
        ":github.com__containers__image__v5__manifest",
        ":github.com__containers__image__v5__transports",
        ":github.com__containers__image__v5__types",
        ":github.com__opencontainers__go-digest",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__image__v5__directory__explicitfilepath",
    download = ":_github.com__containers__image__v5#download",
    install = ["directory/explicitfilepath"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [":github.com__pkg__errors"],
)

go_module(
    name = "github.com__containers__image__v5__docker",
    download = ":_github.com__containers__image__v5#download",
    install = ["docker"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__docker__policyconfiguration",
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__image",
        ":github.com__containers__image__v5__internal__blobinfocache",
        ":github.com__containers__image__v5__internal__iolimits",
        ":github.com__containers__image__v5__internal__putblobdigest",
        ":github.com__containers__image__v5__internal__rootless",
        ":github.com__containers__image__v5__internal__types",
        ":github.com__containers__image__v5__internal__uploadreader",
        ":github.com__containers__image__v5__manifest",
        ":github.com__containers__image__v5__pkg__blobinfocache__none",
        ":github.com__containers__image__v5__pkg__docker__config",
        ":github.com__containers__image__v5__pkg__sysregistriesv2",
        ":github.com__containers__image__v5__pkg__tlsclientconfig",
        ":github.com__containers__image__v5__transports",
        ":github.com__containers__image__v5__types",
        ":github.com__containers__image__v5__version",
        ":github.com__containers__storage__pkg__homedir",
        ":github.com__docker__distribution__registry__api__errcode",
        ":github.com__docker__distribution__registry__api__v2",
        ":github.com__docker__distribution__registry__client",
        ":github.com__docker__go-connections__tlsconfig",
        ":github.com__ghodss__yaml",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__image-spec__specs-go__v1",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__image__v5__docker__archive",
    download = ":_github.com__containers__image__v5#download",
    install = ["docker/archive"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__docker__internal__tarfile",
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__image",
        ":github.com__containers__image__v5__transports",
        ":github.com__containers__image__v5__types",
        ":github.com__pkg__errors",
    ],
)

go_module(
    name = "github.com__containers__image__v5__docker__daemon",
    download = ":_github.com__containers__image__v5#download",
    install = ["docker/daemon"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__docker__internal__tarfile",
        ":github.com__containers__image__v5__docker__policyconfiguration",
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__image",
        ":github.com__containers__image__v5__transports",
        ":github.com__containers__image__v5__types",
        ":github.com__docker__docker__client",
        ":github.com__docker__go-connections__tlsconfig",
        ":github.com__opencontainers__go-digest",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__image__v5__docker__internal__tarfile",
    download = ":_github.com__containers__image__v5#download",
    install = ["docker/internal/tarfile"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__internal__iolimits",
        ":github.com__containers__image__v5__internal__putblobdigest",
        ":github.com__containers__image__v5__internal__tmpdir",
        ":github.com__containers__image__v5__manifest",
        ":github.com__containers__image__v5__pkg__compression",
        ":github.com__containers__image__v5__types",
        ":github.com__opencontainers__go-digest",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__image__v5__docker__policyconfiguration",
    download = ":_github.com__containers__image__v5#download",
    install = ["docker/policyconfiguration"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__pkg__errors",
    ],
)

go_module(
    name = "github.com__containers__image__v5__docker__reference",
    download = ":_github.com__containers__image__v5#download",
    install = ["docker/reference"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [":github.com__opencontainers__go-digest"],
)

go_module(
    name = "github.com__containers__image__v5__image",
    download = ":_github.com__containers__image__v5#download",
    install = ["image"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__internal__iolimits",
        ":github.com__containers__image__v5__manifest",
        ":github.com__containers__image__v5__pkg__blobinfocache__none",
        ":github.com__containers__image__v5__types",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__image-spec__specs-go__v1",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__image__v5__internal__blobinfocache",
    download = ":_github.com__containers__image__v5#download",
    install = ["internal/blobinfocache"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__pkg__compression",
        ":github.com__containers__image__v5__pkg__compression__types",
        ":github.com__containers__image__v5__types",
        ":github.com__opencontainers__go-digest",
    ],
)

go_module(
    name = "github.com__containers__image__v5__internal__iolimits",
    download = ":_github.com__containers__image__v5#download",
    install = ["internal/iolimits"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [":github.com__pkg__errors"],
)

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__image__v5__internal__pkg__keyctl", download = ":_github.com__containers__image__v5#download", install = ["internal/pkg/keyctl"], module = "github.com/containers/image/v5", visibility = ["PUBLIC"], deps = select({":__config_darwin_amd64": [], ":__config_linux_amd64": [":golang.org__x__sys__unix"], "default": []}))

go_module(
    name = "github.com__containers__image__v5__internal__pkg__platform",
    download = ":_github.com__containers__image__v5#download",
    install = ["internal/pkg/platform"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__types",
        ":github.com__opencontainers__image-spec__specs-go__v1",
    ],
)

go_module(
    name = "github.com__containers__image__v5__internal__putblobdigest",
    download = ":_github.com__containers__image__v5#download",
    install = ["internal/putblobdigest"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__types",
        ":github.com__opencontainers__go-digest",
    ],
)

go_module(
    name = "github.com__containers__image__v5__internal__rootless",
    download = ":_github.com__containers__image__v5#download",
    install = ["internal/rootless"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__containers__image__v5__internal__tmpdir",
    download = ":_github.com__containers__image__v5#download",
    install = ["internal/tmpdir"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [":github.com__containers__image__v5__types"],
)

go_module(
    name = "github.com__containers__image__v5__internal__types",
    download = ":_github.com__containers__image__v5#download",
    install = ["internal/types"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__types",
    ],
)

go_module(
    name = "github.com__containers__image__v5__internal__uploadreader",
    download = ":_github.com__containers__image__v5#download",
    install = ["internal/uploadreader"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__containers__image__v5__manifest",
    download = ":_github.com__containers__image__v5#download",
    install = ["manifest"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__internal__pkg__platform",
        ":github.com__containers__image__v5__pkg__compression__types",
        ":github.com__containers__image__v5__pkg__strslice",
        ":github.com__containers__image__v5__types",
        ":github.com__containers__libtrust",
        ":github.com__containers__ocicrypt__spec",
        ":github.com__docker__docker__api__types__versions",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__image-spec__specs-go",
        ":github.com__opencontainers__image-spec__specs-go__v1",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__image__v5__oci__archive",
    download = ":_github.com__containers__image__v5#download",
    install = ["oci/archive"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__directory__explicitfilepath",
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__image",
        ":github.com__containers__image__v5__internal__tmpdir",
        ":github.com__containers__image__v5__oci__internal",
        ":github.com__containers__image__v5__oci__layout",
        ":github.com__containers__image__v5__transports",
        ":github.com__containers__image__v5__types",
        ":github.com__containers__storage__pkg__archive",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__image-spec__specs-go__v1",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__image__v5__oci__internal",
    download = ":_github.com__containers__image__v5#download",
    install = ["oci/internal"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [":github.com__pkg__errors"],
)

go_module(
    name = "github.com__containers__image__v5__oci__layout",
    download = ":_github.com__containers__image__v5#download",
    install = ["oci/layout"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__directory__explicitfilepath",
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__image",
        ":github.com__containers__image__v5__internal__putblobdigest",
        ":github.com__containers__image__v5__manifest",
        ":github.com__containers__image__v5__oci__internal",
        ":github.com__containers__image__v5__pkg__tlsclientconfig",
        ":github.com__containers__image__v5__transports",
        ":github.com__containers__image__v5__types",
        ":github.com__docker__go-connections__tlsconfig",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__image-spec__specs-go",
        ":github.com__opencontainers__image-spec__specs-go__v1",
        ":github.com__pkg__errors",
    ],
)

go_module(
    name = "github.com__containers__image__v5__openshift",
    download = ":_github.com__containers__image__v5#download",
    install = ["openshift"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__docker",
        ":github.com__containers__image__v5__docker__policyconfiguration",
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__image",
        ":github.com__containers__image__v5__internal__iolimits",
        ":github.com__containers__image__v5__manifest",
        ":github.com__containers__image__v5__transports",
        ":github.com__containers__image__v5__types",
        ":github.com__containers__image__v5__version",
        ":github.com__containers__storage__pkg__homedir",
        ":github.com__ghodss__yaml",
        ":github.com__imdario__mergo",
        ":github.com__opencontainers__go-digest",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
        ":golang.org__x__net__http2",
    ],
)

go_module(
    name = "github.com__containers__image__v5__pkg__blobinfocache",
    download = ":_github.com__containers__image__v5#download",
    install = ["pkg/blobinfocache"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__internal__rootless",
        ":github.com__containers__image__v5__pkg__blobinfocache__boltdb",
        ":github.com__containers__image__v5__pkg__blobinfocache__memory",
        ":github.com__containers__image__v5__types",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__image__v5__pkg__blobinfocache__boltdb",
    download = ":_github.com__containers__image__v5#download",
    install = ["pkg/blobinfocache/boltdb"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__internal__blobinfocache",
        ":github.com__containers__image__v5__pkg__blobinfocache__internal__prioritize",
        ":github.com__containers__image__v5__types",
        ":github.com__opencontainers__go-digest",
        ":github.com__sirupsen__logrus",
        ":go.etcd.io__bbolt",
    ],
)

go_module(
    name = "github.com__containers__image__v5__pkg__blobinfocache__internal__prioritize",
    download = ":_github.com__containers__image__v5#download",
    install = ["pkg/blobinfocache/internal/prioritize"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__internal__blobinfocache",
        ":github.com__opencontainers__go-digest",
    ],
)

go_module(
    name = "github.com__containers__image__v5__pkg__blobinfocache__memory",
    download = ":_github.com__containers__image__v5#download",
    install = ["pkg/blobinfocache/memory"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__internal__blobinfocache",
        ":github.com__containers__image__v5__pkg__blobinfocache__internal__prioritize",
        ":github.com__containers__image__v5__types",
        ":github.com__opencontainers__go-digest",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__image__v5__pkg__blobinfocache__none",
    download = ":_github.com__containers__image__v5#download",
    install = ["pkg/blobinfocache/none"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__internal__blobinfocache",
        ":github.com__containers__image__v5__types",
        ":github.com__opencontainers__go-digest",
    ],
)

go_module(
    name = "github.com__containers__image__v5__pkg__compression",
    download = ":_github.com__containers__image__v5#download",
    install = ["pkg/compression"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__pkg__compression__internal",
        ":github.com__containers__image__v5__pkg__compression__types",
        ":github.com__containers__storage__pkg__chunked__compressor",
        ":github.com__klauspost__compress__zstd",
        ":github.com__klauspost__pgzip",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
        ":github.com__ulikunitz__xz",
    ],
)

go_module(
    name = "github.com__containers__image__v5__pkg__compression__internal",
    download = ":_github.com__containers__image__v5#download",
    install = ["pkg/compression/internal"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__containers__image__v5__pkg__compression__types",
    download = ":_github.com__containers__image__v5#download",
    install = ["pkg/compression/types"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [":github.com__containers__image__v5__pkg__compression__internal"],
)

go_module(
    name = "github.com__containers__image__v5__pkg__docker__config",
    download = ":_github.com__containers__image__v5#download",
    install = ["pkg/docker/config"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__pkg__sysregistriesv2",
        ":github.com__containers__image__v5__types",
        ":github.com__containers__storage__pkg__homedir",
        ":github.com__docker__docker-credential-helpers__client",
        ":github.com__docker__docker-credential-helpers__credentials",
        ":github.com__hashicorp__go-multierror",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ] + select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [":github.com__containers__image__v5__internal__pkg__keyctl"],
        "default": [],
    }),
)

go_module(
    name = "github.com__containers__image__v5__pkg__shortnames",
    download = ":_github.com__containers__image__v5#download",
    install = ["pkg/shortnames"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__pkg__sysregistriesv2",
        ":github.com__containers__image__v5__types",
        ":github.com__manifoldco__promptui",
        ":github.com__opencontainers__go-digest",
        ":github.com__pkg__errors",
        ":golang.org__x__term",
    ],
)

go_module(
    name = "github.com__containers__image__v5__pkg__strslice",
    download = ":_github.com__containers__image__v5#download",
    install = ["pkg/strslice"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__containers__image__v5__pkg__sysregistriesv2",
    download = ":_github.com__containers__image__v5#download",
    install = ["pkg/sysregistriesv2"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__BurntSushi__toml",
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__internal__rootless",
        ":github.com__containers__image__v5__types",
        ":github.com__containers__storage__pkg__homedir",
        ":github.com__containers__storage__pkg__lockfile",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__image__v5__pkg__tlsclientconfig",
    download = ":_github.com__containers__image__v5#download",
    install = ["pkg/tlsclientconfig"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__docker__go-connections__sockets",
        ":github.com__docker__go-connections__tlsconfig",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__image__v5__signature",
    download = ":_github.com__containers__image__v5#download",
    install = ["signature"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__manifest",
        ":github.com__containers__image__v5__transports",
        ":github.com__containers__image__v5__types",
        ":github.com__containers__image__v5__version",
        ":github.com__containers__storage__pkg__homedir",
        ":github.com__mtrmac__gpgme",
        ":github.com__opencontainers__go-digest",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
        ":golang.org__x__crypto__openpgp",
    ],
)

go_module(
    name = "github.com__containers__image__v5__storage",
    download = ":_github.com__containers__image__v5#download",
    install = ["storage"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__image",
        ":github.com__containers__image__v5__internal__putblobdigest",
        ":github.com__containers__image__v5__internal__tmpdir",
        ":github.com__containers__image__v5__internal__types",
        ":github.com__containers__image__v5__manifest",
        ":github.com__containers__image__v5__pkg__blobinfocache__none",
        ":github.com__containers__image__v5__transports",
        ":github.com__containers__image__v5__types",
        ":github.com__containers__storage",
        ":github.com__containers__storage__drivers",
        ":github.com__containers__storage__pkg__archive",
        ":github.com__containers__storage__pkg__chunked",
        ":github.com__containers__storage__pkg__idtools",
        ":github.com__containers__storage__pkg__ioutils",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__image-spec__specs-go__v1",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__image__v5__tarball",
    download = ":_github.com__containers__image__v5#download",
    install = ["tarball"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__image",
        ":github.com__containers__image__v5__transports",
        ":github.com__containers__image__v5__types",
        ":github.com__klauspost__pgzip",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__image-spec__specs-go",
        ":github.com__opencontainers__image-spec__specs-go__v1",
    ],
)

go_module(
    name = "github.com__containers__image__v5__transports",
    download = ":_github.com__containers__image__v5#download",
    install = ["transports"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [":github.com__containers__image__v5__types"],
)

go_module(
    name = "github.com__containers__image__v5__transports__alltransports",
    download = ":_github.com__containers__image__v5#download",
    install = ["transports/alltransports"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__directory",
        ":github.com__containers__image__v5__docker",
        ":github.com__containers__image__v5__docker__archive",
        ":github.com__containers__image__v5__docker__daemon",
        ":github.com__containers__image__v5__oci__archive",
        ":github.com__containers__image__v5__oci__layout",
        ":github.com__containers__image__v5__openshift",
        ":github.com__containers__image__v5__storage",
        ":github.com__containers__image__v5__tarball",
        ":github.com__containers__image__v5__transports",
        ":github.com__containers__image__v5__types",
        ":github.com__pkg__errors",
    ],
)

go_module(
    name = "github.com__containers__image__v5__types",
    download = ":_github.com__containers__image__v5#download",
    install = ["types"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__docker__reference",
        ":github.com__containers__image__v5__pkg__compression__types",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__image-spec__specs-go__v1",
    ],
)

go_module(
    name = "github.com__containers__image__v5__version",
    download = ":_github.com__containers__image__v5#download",
    install = ["version"],
    module = "github.com/containers/image/v5",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__containers__libtrust",
    _tag = "download",
    module = "github.com/containers/libtrust",
    version = "v0.0.0-20190913040956-14b96171aa3b",
)

go_module(
    name = "github.com__containers__libtrust",
    download = ":_github.com__containers__libtrust#download",
    install = ["."],
    module = "github.com/containers/libtrust",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__containers__ocicrypt",
    _tag = "download",
    module = "github.com/containers/ocicrypt",
    version = "v1.1.2",
)

go_module(
    name = "github.com__containers__ocicrypt",
    download = ":_github.com__containers__ocicrypt#download",
    install = ["."],
    module = "github.com/containers/ocicrypt",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__ocicrypt__blockcipher",
        ":github.com__containers__ocicrypt__config",
        ":github.com__containers__ocicrypt__config__keyprovider-config",
        ":github.com__containers__ocicrypt__keywrap",
        ":github.com__containers__ocicrypt__keywrap__jwe",
        ":github.com__containers__ocicrypt__keywrap__keyprovider",
        ":github.com__containers__ocicrypt__keywrap__pgp",
        ":github.com__containers__ocicrypt__keywrap__pkcs11",
        ":github.com__containers__ocicrypt__keywrap__pkcs7",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__image-spec__specs-go__v1",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
        ":golang.org__x__crypto__openpgp",
        ":golang.org__x__crypto__openpgp__packet",
        ":golang.org__x__term",
    ],
)

go_module(
    name = "github.com__containers__ocicrypt__blockcipher",
    download = ":_github.com__containers__ocicrypt#download",
    install = ["blockcipher"],
    module = "github.com/containers/ocicrypt",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__ocicrypt__utils",
        ":github.com__opencontainers__go-digest",
        ":github.com__pkg__errors",
    ],
)

go_module(
    name = "github.com__containers__ocicrypt__config",
    download = ":_github.com__containers__ocicrypt#download",
    install = ["config"],
    module = "github.com/containers/ocicrypt",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__ocicrypt__crypto__pkcs11",
        ":github.com__pkg__errors",
        ":gopkg.in__yaml.v2",
    ],
)

go_module(
    name = "github.com__containers__ocicrypt__config__keyprovider-config",
    download = ":_github.com__containers__ocicrypt#download",
    install = ["config/keyprovider-config"],
    module = "github.com/containers/ocicrypt",
    visibility = ["PUBLIC"],
    deps = [":github.com__pkg__errors"],
)

go_module(
    name = "github.com__containers__ocicrypt__crypto__pkcs11",
    download = ":_github.com__containers__ocicrypt#download",
    install = ["crypto/pkcs11"],
    module = "github.com/containers/ocicrypt",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__miekg__pkcs11",
        ":github.com__pkg__errors",
        ":github.com__stefanberger__go-pkcs11uri",
        ":gopkg.in__yaml.v2",
    ],
)

go_module(
    name = "github.com__containers__ocicrypt__keywrap",
    download = ":_github.com__containers__ocicrypt#download",
    install = ["keywrap"],
    module = "github.com/containers/ocicrypt",
    visibility = ["PUBLIC"],
    deps = [":github.com__containers__ocicrypt__config"],
)

go_module(
    name = "github.com__containers__ocicrypt__keywrap__jwe",
    download = ":_github.com__containers__ocicrypt#download",
    install = ["keywrap/jwe"],
    module = "github.com/containers/ocicrypt",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__ocicrypt__config",
        ":github.com__containers__ocicrypt__keywrap",
        ":github.com__containers__ocicrypt__utils",
        ":github.com__pkg__errors",
        ":gopkg.in__square__go-jose.v2",
    ],
)

go_module(
    name = "github.com__containers__ocicrypt__keywrap__keyprovider",
    download = ":_github.com__containers__ocicrypt#download",
    install = ["keywrap/keyprovider"],
    module = "github.com/containers/ocicrypt",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__ocicrypt__config",
        ":github.com__containers__ocicrypt__config__keyprovider-config",
        ":github.com__containers__ocicrypt__keywrap",
        ":github.com__containers__ocicrypt__utils",
        ":github.com__containers__ocicrypt__utils__keyprovider",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
        ":google.golang.org__grpc",
    ],
)

go_module(
    name = "github.com__containers__ocicrypt__keywrap__pgp",
    download = ":_github.com__containers__ocicrypt#download",
    install = ["keywrap/pgp"],
    module = "github.com/containers/ocicrypt",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__ocicrypt__config",
        ":github.com__containers__ocicrypt__keywrap",
        ":github.com__pkg__errors",
        ":golang.org__x__crypto__openpgp",
        ":golang.org__x__crypto__openpgp__packet",
    ],
)

go_module(
    name = "github.com__containers__ocicrypt__keywrap__pkcs11",
    download = ":_github.com__containers__ocicrypt#download",
    install = ["keywrap/pkcs11"],
    module = "github.com/containers/ocicrypt",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__ocicrypt__config",
        ":github.com__containers__ocicrypt__crypto__pkcs11",
        ":github.com__containers__ocicrypt__keywrap",
        ":github.com__containers__ocicrypt__utils",
        ":github.com__pkg__errors",
    ],
)

go_module(
    name = "github.com__containers__ocicrypt__keywrap__pkcs7",
    download = ":_github.com__containers__ocicrypt#download",
    install = ["keywrap/pkcs7"],
    module = "github.com/containers/ocicrypt",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__ocicrypt__config",
        ":github.com__containers__ocicrypt__keywrap",
        ":github.com__containers__ocicrypt__utils",
        ":github.com__pkg__errors",
        ":go.mozilla.org__pkcs7",
    ],
)

go_module(
    name = "github.com__containers__ocicrypt__spec",
    download = ":_github.com__containers__ocicrypt#download",
    install = ["spec"],
    module = "github.com/containers/ocicrypt",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__containers__ocicrypt__utils",
    download = ":_github.com__containers__ocicrypt#download",
    install = ["utils"],
    module = "github.com/containers/ocicrypt",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__ocicrypt__crypto__pkcs11",
        ":github.com__pkg__errors",
        ":golang.org__x__crypto__openpgp",
        ":gopkg.in__square__go-jose.v2",
    ],
)

go_module(
    name = "github.com__containers__ocicrypt__utils__keyprovider",
    download = ":_github.com__containers__ocicrypt#download",
    install = ["utils/keyprovider"],
    module = "github.com/containers/ocicrypt",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__golang__protobuf__proto",
        ":google.golang.org__grpc",
        ":google.golang.org__grpc__codes",
        ":google.golang.org__grpc__status",
    ],
)

go_mod_download(
    name = "github.com__containers__podman__v3",
    _tag = "download",
    module = "github.com/containers/podman/v3",
    version = "v3.4.0",
)

go_module(
    name = "github.com__containers__podman__v3__libpod__define",
    download = ":_github.com__containers__podman__v3#download",
    install = ["libpod/define"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__manifest",
        ":github.com__containers__podman__v3__version",
        ":github.com__containers__storage__pkg__idtools",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__podman__v3__libpod__events",
    download = ":_github.com__containers__podman__v3#download",
    install = ["libpod/events"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__podman__v3__pkg__util",
        ":github.com__containers__storage__pkg__lockfile",
        ":github.com__hpcloud__tail",
        ":github.com__pkg__errors",
    ] + select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [":github.com__sirupsen__logrus"],
        "default": [],
    }),
)

go_module(
    name = "github.com__containers__podman__v3__libpod__network__types",
    download = ":_github.com__containers__podman__v3#download",
    install = ["libpod/network/types"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__api__handlers",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/api/handlers"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__common__libimage",
        ":github.com__containers__podman__v3__libpod__define",
        ":github.com__containers__podman__v3__pkg__domain__entities",
        ":github.com__containers__podman__v3__pkg__util",
        ":github.com__docker__docker__api__types",
        ":github.com__docker__docker__api__types__container",
        ":github.com__docker__docker__api__types__network",
        ":github.com__docker__go-connections__nat",
        ":github.com__gorilla__schema",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__api__handlers__types",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/api/handlers/types"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [":github.com__containers__podman__v3__pkg__domain__entities"],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__auth",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/auth"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__pkg__docker__config",
        ":github.com__containers__image__v5__types",
        ":github.com__docker__docker__api__types",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__bindings",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/bindings"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__blang__semver",
        ":github.com__containers__podman__v3__pkg__errorhandling",
        ":github.com__containers__podman__v3__pkg__terminal",
        ":github.com__containers__podman__v3__version",
        ":github.com__json-iterator__go",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
        ":golang.org__x__crypto__ssh",
        ":golang.org__x__crypto__ssh__agent",
    ],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__bindings__containers",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/bindings/containers"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__podman__v3__libpod__define",
        ":github.com__containers__podman__v3__pkg__api__handlers",
        ":github.com__containers__podman__v3__pkg__bindings",
        ":github.com__containers__podman__v3__pkg__bindings__internal__util",
        ":github.com__containers__podman__v3__pkg__copy",
        ":github.com__containers__podman__v3__pkg__domain__entities",
        ":github.com__containers__podman__v3__pkg__domain__entities__reports",
        ":github.com__containers__podman__v3__pkg__signal",
        ":github.com__containers__podman__v3__pkg__specgen",
        ":github.com__containers__podman__v3__utils",
        ":github.com__containers__storage__pkg__archive",
        ":github.com__json-iterator__go",
        ":github.com__moby__term",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
        ":golang.org__x__crypto__ssh__terminal",
    ],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__bindings__images",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/bindings/images"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__buildah__define",
        ":github.com__containers__podman__v3__pkg__api__handlers__types",
        ":github.com__containers__podman__v3__pkg__auth",
        ":github.com__containers__podman__v3__pkg__bindings",
        ":github.com__containers__podman__v3__pkg__bindings__internal__util",
        ":github.com__containers__podman__v3__pkg__domain__entities",
        ":github.com__containers__podman__v3__pkg__domain__entities__reports",
        ":github.com__containers__podman__v3__pkg__errorhandling",
        ":github.com__containers__storage__pkg__archive",
        ":github.com__containers__storage__pkg__fileutils",
        ":github.com__containers__storage__pkg__ioutils",
        ":github.com__docker__go-units",
        ":github.com__hashicorp__go-multierror",
        ":github.com__json-iterator__go",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__bindings__internal__util",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/bindings/internal/util"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [":github.com__json-iterator__go"],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__bindings__network",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/bindings/network"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__podman__v3__pkg__bindings",
        ":github.com__containers__podman__v3__pkg__bindings__internal__util",
        ":github.com__containers__podman__v3__pkg__domain__entities",
        ":github.com__json-iterator__go",
    ],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__cgroups",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/cgroups"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__podman__v3__pkg__rootless",
        ":github.com__coreos__go-systemd__v22__dbus",
        ":github.com__godbus__dbus__v5",
        ":github.com__opencontainers__runtime-spec__specs-go",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ] + select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [":golang.org__x__sys__unix"],
        "default": [],
    }),
)

go_module(
    name = "github.com__containers__podman__v3__pkg__copy",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/copy"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__podman__v3__libpod__define",
        ":github.com__pkg__errors",
    ],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__domain__entities",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/domain/entities"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containernetworking__cni__libcni",
        ":github.com__containers__buildah__define",
        ":github.com__containers__common__pkg__config",
        ":github.com__containers__image__v5__manifest",
        ":github.com__containers__image__v5__types",
        ":github.com__containers__podman__v3__libpod__define",
        ":github.com__containers__podman__v3__libpod__events",
        ":github.com__containers__podman__v3__libpod__network__types",
        ":github.com__containers__podman__v3__pkg__domain__entities__reports",
        ":github.com__containers__podman__v3__pkg__domain__entities__types",
        ":github.com__containers__podman__v3__pkg__errorhandling",
        ":github.com__containers__podman__v3__pkg__inspect",
        ":github.com__containers__podman__v3__pkg__ps__define",
        ":github.com__containers__podman__v3__pkg__specgen",
        ":github.com__containers__podman__v3__pkg__trust",
        ":github.com__containers__podman__v3__pkg__util",
        ":github.com__containers__storage__pkg__archive",
        ":github.com__cri-o__ocicni__pkg__ocicni",
        ":github.com__docker__docker__api__types",
        ":github.com__docker__docker__api__types__container",
        ":github.com__docker__docker__api__types__events",
        ":github.com__docker__docker__api__types__volume",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__image-spec__specs-go__v1",
        ":github.com__opencontainers__runtime-spec__specs-go",
        ":github.com__pkg__errors",
        ":github.com__spf13__cobra",
        ":github.com__spf13__pflag",
    ],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__domain__entities__reports",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/domain/entities/reports"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__domain__entities__types",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/domain/entities/types"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__errorhandling",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/errorhandling"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__hashicorp__go-multierror",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__inspect",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/inspect"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__manifest",
        ":github.com__containers__podman__v3__libpod__define",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__image-spec__specs-go__v1",
    ],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__namespaces",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/namespaces"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [":github.com__containers__storage__types"],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__ps__define",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/ps/define"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__rootless",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/rootless"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__storage__pkg__idtools",
        ":github.com__containers__storage__pkg__lockfile",
        ":github.com__opencontainers__runc__libcontainer__user",
        ":github.com__opencontainers__runtime-spec__specs-go",
        ":github.com__pkg__errors",
    ] + select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [
            ":github.com__containers__podman__v3__pkg__errorhandling",
            ":github.com__containers__storage__pkg__mount",
            ":github.com__containers__storage__pkg__unshare",
            ":github.com__sirupsen__logrus",
            ":golang.org__x__sys__unix",
        ],
        "default": [],
    }),
)

go_module(
    name = "github.com__containers__podman__v3__pkg__signal",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/signal"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [":golang.org__x__sys__unix"],
        "default": [],
    }),
)

go_module(
    name = "github.com__containers__podman__v3__pkg__specgen",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/specgen"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__common__pkg__parse",
        ":github.com__containers__image__v5__manifest",
        ":github.com__containers__podman__v3__libpod__define",
        ":github.com__containers__podman__v3__libpod__network__types",
        ":github.com__containers__podman__v3__pkg__cgroups",
        ":github.com__containers__podman__v3__pkg__rootless",
        ":github.com__containers__podman__v3__pkg__util",
        ":github.com__containers__storage",
        ":github.com__containers__storage__types",
        ":github.com__opencontainers__runtime-spec__specs-go",
        ":github.com__opencontainers__runtime-tools__generate",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ] + select({
        ":__config_darwin_amd64": [":github.com__containers__common__libimage"],
        ":__config_linux_amd64": [],
        "default": [],
    }),
)

go_module(
    name = "github.com__containers__podman__v3__pkg__terminal",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/terminal"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__storage__pkg__homedir",
        ":github.com__sirupsen__logrus",
        ":golang.org__x__crypto__ssh",
        ":golang.org__x__crypto__ssh__knownhosts",
        ":golang.org__x__crypto__ssh__terminal",
    ],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__timetype",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/timetype"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__trust",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/trust"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__image__v5__types",
        ":github.com__docker__docker__pkg__homedir",
        ":github.com__ghodss__yaml",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__podman__v3__pkg__util",
    download = ":_github.com__containers__podman__v3#download",
    install = ["pkg/util"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__BurntSushi__toml",
        ":github.com__containers__common__pkg__config",
        ":github.com__containers__image__v5__types",
        ":github.com__containers__podman__v3__pkg__errorhandling",
        ":github.com__containers__podman__v3__pkg__namespaces",
        ":github.com__containers__podman__v3__pkg__rootless",
        ":github.com__containers__podman__v3__pkg__signal",
        ":github.com__containers__podman__v3__pkg__timetype",
        ":github.com__containers__storage__pkg__idtools",
        ":github.com__containers__storage__types",
        ":github.com__opencontainers__image-spec__specs-go__v1",
        ":github.com__opencontainers__runtime-spec__specs-go",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
        ":golang.org__x__crypto__ssh__terminal",
    ] + select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [
            ":github.com__containers__psgo",
            ":golang.org__x__sys__unix",
        ],
        "default": [],
    }),
)

go_module(
    name = "github.com__containers__podman__v3__utils",
    download = ":_github.com__containers__podman__v3#download",
    install = ["utils"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__podman__v3__libpod__define",
        ":github.com__containers__podman__v3__pkg__cgroups",
        ":github.com__containers__podman__v3__pkg__rootless",
        ":github.com__containers__storage__pkg__archive",
        ":github.com__coreos__go-systemd__v22__dbus",
        ":github.com__godbus__dbus__v5",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_module(
    name = "github.com__containers__podman__v3__version",
    download = ":_github.com__containers__podman__v3#download",
    install = ["version"],
    module = "github.com/containers/podman/v3",
    visibility = ["PUBLIC"],
    deps = [":github.com__blang__semver"],
)

go_mod_download(
    name = "github.com__containers__psgo",
    _tag = "download",
    module = "github.com/containers/psgo",
    version = "v1.7.1",
)

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__psgo", download = ":_github.com__containers__psgo#download", install = ["."], module = "github.com/containers/psgo", visibility = ["PUBLIC"], deps = select({":__config_darwin_amd64": [], ":__config_linux_amd64": [":github.com__containers__psgo__internal__capabilities", ":github.com__containers__psgo__internal__dev", ":github.com__containers__psgo__internal__proc", ":github.com__containers__psgo__internal__process", ":golang.org__x__sys__unix"], "default": []}))

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__psgo__internal__capabilities", download = ":_github.com__containers__psgo#download", install = ["internal/capabilities"], module = "github.com/containers/psgo", visibility = ["PUBLIC"], deps = [])

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__psgo__internal__cgroups", download = ":_github.com__containers__psgo#download", install = ["internal/cgroups"], module = "github.com/containers/psgo", visibility = ["PUBLIC"], deps = [])

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__psgo__internal__dev", download = ":_github.com__containers__psgo#download", install = ["internal/dev"], module = "github.com/containers/psgo", visibility = ["PUBLIC"], deps = [])

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__psgo__internal__host", download = ":_github.com__containers__psgo#download", install = ["internal/host"], module = "github.com/containers/psgo", visibility = ["PUBLIC"], deps = [])

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__psgo__internal__proc", download = ":_github.com__containers__psgo#download", install = ["internal/proc"], module = "github.com/containers/psgo", visibility = ["PUBLIC"], deps = select({":__config_darwin_amd64": [], ":__config_linux_amd64": [":github.com__containers__psgo__internal__cgroups"], "default": []}))

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__psgo__internal__process", download = ":_github.com__containers__psgo#download", install = ["internal/process"], module = "github.com/containers/psgo", visibility = ["PUBLIC"], deps = select({":__config_darwin_amd64": [], ":__config_linux_amd64": [":github.com__containers__psgo__internal__host", ":github.com__containers__psgo__internal__proc", ":github.com__opencontainers__runc__libcontainer__user"], "default": []}))

go_mod_download(
    name = "github.com__containers__storage",
    _tag = "download",
    module = "github.com/containers/storage",
    strip = ["pkg/unshare/unshare_cgo.go"],
    version = "v1.36.0",
)

go_module(
    name = "github.com__containers__storage",
    download = ":_github.com__containers__storage#download",
    install = ["."],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__storage__drivers",
        ":github.com__containers__storage__drivers__register",
        ":github.com__containers__storage__pkg__archive",
        ":github.com__containers__storage__pkg__directory",
        ":github.com__containers__storage__pkg__idtools",
        ":github.com__containers__storage__pkg__ioutils",
        ":github.com__containers__storage__pkg__lockfile",
        ":github.com__containers__storage__pkg__mount",
        ":github.com__containers__storage__pkg__parsers",
        ":github.com__containers__storage__pkg__stringid",
        ":github.com__containers__storage__pkg__stringutils",
        ":github.com__containers__storage__pkg__system",
        ":github.com__containers__storage__pkg__tarlog",
        ":github.com__containers__storage__pkg__truncindex",
        ":github.com__containers__storage__pkg__unshare",
        ":github.com__containers__storage__types",
        ":github.com__google__go-intervals__intervalset",
        ":github.com__hashicorp__go-multierror",
        ":github.com__json-iterator__go",
        ":github.com__klauspost__pgzip",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__runc__libcontainer__user",
        ":github.com__opencontainers__selinux__go-selinux__label",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
        ":github.com__vbatts__tar-split__archive__tar",
        ":github.com__vbatts__tar-split__tar__asm",
        ":github.com__vbatts__tar-split__tar__storage",
    ],
)

go_module(
    name = "github.com__containers__storage__drivers",
    download = ":_github.com__containers__storage#download",
    install = ["drivers"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__storage__pkg__archive",
        ":github.com__containers__storage__pkg__chrootarchive",
        ":github.com__containers__storage__pkg__directory",
        ":github.com__containers__storage__pkg__idtools",
        ":github.com__containers__storage__pkg__ioutils",
        ":github.com__containers__storage__pkg__reexec",
        ":github.com__containers__storage__pkg__system",
        ":github.com__json-iterator__go",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__runc__libcontainer__userns",
        ":github.com__opencontainers__selinux__pkg__pwalk",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
        ":github.com__vbatts__tar-split__tar__storage",
    ] + select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [
            ":github.com__containers__storage__pkg__mount",
            ":golang.org__x__sys__unix",
        ],
        "default": [],
    }),
)

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__storage__drivers__aufs", download = ":_github.com__containers__storage#download", install = ["drivers/aufs"], module = "github.com/containers/storage", visibility = ["PUBLIC"], deps = select({":__config_darwin_amd64": [], ":__config_linux_amd64": [":github.com__containers__storage__drivers", ":github.com__containers__storage__pkg__archive", ":github.com__containers__storage__pkg__chrootarchive", ":github.com__containers__storage__pkg__directory", ":github.com__containers__storage__pkg__idtools", ":github.com__containers__storage__pkg__locker", ":github.com__containers__storage__pkg__mount", ":github.com__containers__storage__pkg__parsers", ":github.com__containers__storage__pkg__system", ":github.com__opencontainers__runc__libcontainer__userns", ":github.com__opencontainers__selinux__go-selinux__label", ":github.com__pkg__errors", ":github.com__sirupsen__logrus", ":github.com__vbatts__tar-split__tar__storage", ":golang.org__x__sys__unix"], "default": []}))

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__storage__drivers__btrfs", download = ":_github.com__containers__storage#download", install = ["drivers/btrfs"], module = "github.com/containers/storage", visibility = ["PUBLIC"], deps = select({":__config_darwin_amd64": [], ":__config_linux_amd64": [":github.com__containers__storage__drivers", ":github.com__containers__storage__pkg__directory", ":github.com__containers__storage__pkg__idtools", ":github.com__containers__storage__pkg__mount", ":github.com__containers__storage__pkg__parsers", ":github.com__containers__storage__pkg__system", ":github.com__docker__go-units", ":github.com__opencontainers__selinux__go-selinux__label", ":github.com__pkg__errors", ":github.com__sirupsen__logrus", ":golang.org__x__sys__unix"], "default": []}))

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__storage__drivers__copy", download = ":_github.com__containers__storage#download", install = ["drivers/copy"], module = "github.com/containers/storage", visibility = ["PUBLIC"], deps = select({":__config_darwin_amd64": [], ":__config_linux_amd64": [":github.com__containers__storage__pkg__idtools", ":github.com__containers__storage__pkg__pools", ":github.com__containers__storage__pkg__system", ":github.com__containers__storage__pkg__unshare", ":github.com__opencontainers__runc__libcontainer__userns", ":golang.org__x__sys__unix"], "default": []}))

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__storage__drivers__devmapper", download = ":_github.com__containers__storage#download", install = ["drivers/devmapper"], module = "github.com/containers/storage", visibility = ["PUBLIC"], deps = select({":__config_darwin_amd64": [], ":__config_linux_amd64": [":github.com__containers__storage__drivers", ":github.com__containers__storage__pkg__devicemapper", ":github.com__containers__storage__pkg__directory", ":github.com__containers__storage__pkg__dmesg", ":github.com__containers__storage__pkg__idtools", ":github.com__containers__storage__pkg__locker", ":github.com__containers__storage__pkg__loopback", ":github.com__containers__storage__pkg__mount", ":github.com__containers__storage__pkg__parsers", ":github.com__containers__storage__pkg__parsers__kernel", ":github.com__docker__go-units", ":github.com__json-iterator__go", ":github.com__opencontainers__selinux__go-selinux__label", ":github.com__pkg__errors", ":github.com__sirupsen__logrus", ":golang.org__x__sys__unix"], "default": []}))

go_module(
    name = "github.com__containers__storage__drivers__overlay",
    download = ":_github.com__containers__storage#download",
    install = ["drivers/overlay"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__storage__pkg__archive",
        ":github.com__containers__storage__pkg__system",
        ":github.com__json-iterator__go",
    ] + select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [
            ":github.com__containers__storage__drivers",
            ":github.com__containers__storage__drivers__overlayutils",
            ":github.com__containers__storage__drivers__quota",
            ":github.com__containers__storage__pkg__chrootarchive",
            ":github.com__containers__storage__pkg__directory",
            ":github.com__containers__storage__pkg__fsutils",
            ":github.com__containers__storage__pkg__idtools",
            ":github.com__containers__storage__pkg__ioutils",
            ":github.com__containers__storage__pkg__locker",
            ":github.com__containers__storage__pkg__mount",
            ":github.com__containers__storage__pkg__parsers",
            ":github.com__containers__storage__pkg__reexec",
            ":github.com__containers__storage__pkg__unshare",
            ":github.com__docker__go-units",
            ":github.com__hashicorp__go-multierror",
            ":github.com__opencontainers__go-digest",
            ":github.com__opencontainers__runc__libcontainer__userns",
            ":github.com__opencontainers__selinux__go-selinux",
            ":github.com__opencontainers__selinux__go-selinux__label",
            ":github.com__pkg__errors",
            ":github.com__sirupsen__logrus",
            ":github.com__vbatts__tar-split__tar__storage",
            ":golang.org__x__sys__unix",
        ],
        "default": [],
    }),
)

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__storage__drivers__overlayutils", download = ":_github.com__containers__storage#download", install = ["drivers/overlayutils"], module = "github.com/containers/storage", visibility = ["PUBLIC"], deps = select({":__config_darwin_amd64": [], ":__config_linux_amd64": [":github.com__containers__storage__drivers", ":github.com__pkg__errors"], "default": []}))

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__storage__drivers__quota", download = ":_github.com__containers__storage#download", install = ["drivers/quota"], module = "github.com/containers/storage", visibility = ["PUBLIC"], deps = select({":__config_darwin_amd64": [], ":__config_linux_amd64": [":github.com__containers__storage__pkg__directory", ":github.com__sirupsen__logrus", ":golang.org__x__sys__unix"], "default": []}))

go_module(
    name = "github.com__containers__storage__drivers__register",
    download = ":_github.com__containers__storage#download",
    install = ["drivers/register"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [":github.com__containers__storage__drivers__vfs"] + select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [
            ":github.com__containers__storage__drivers__aufs",
            ":github.com__containers__storage__drivers__btrfs",
            ":github.com__containers__storage__drivers__devmapper",
            ":github.com__containers__storage__drivers__overlay",
            ":github.com__containers__storage__drivers__zfs",
        ],
        "default": [],
    }),
)

go_module(
    name = "github.com__containers__storage__drivers__vfs",
    download = ":_github.com__containers__storage#download",
    install = ["drivers/vfs"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__storage__drivers",
        ":github.com__containers__storage__pkg__archive",
        ":github.com__containers__storage__pkg__directory",
        ":github.com__containers__storage__pkg__idtools",
        ":github.com__containers__storage__pkg__parsers",
        ":github.com__containers__storage__pkg__system",
        ":github.com__opencontainers__selinux__go-selinux__label",
        ":github.com__sirupsen__logrus",
        ":github.com__vbatts__tar-split__tar__storage",
    ] + select({
        ":__config_darwin_amd64": [":github.com__containers__storage__pkg__chrootarchive"],
        ":__config_linux_amd64": [":github.com__containers__storage__drivers__copy"],
        "default": [],
    }),
)

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__storage__drivers__zfs", download = ":_github.com__containers__storage#download", install = ["drivers/zfs"], module = "github.com/containers/storage", visibility = ["PUBLIC"], deps = select({":__config_darwin_amd64": [], ":__config_linux_amd64": [":github.com__containers__storage__drivers", ":github.com__containers__storage__pkg__directory", ":github.com__containers__storage__pkg__idtools", ":github.com__containers__storage__pkg__mount", ":github.com__containers__storage__pkg__parsers", ":github.com__mistifyio__go-zfs", ":github.com__opencontainers__selinux__go-selinux__label", ":github.com__pkg__errors", ":github.com__sirupsen__logrus", ":golang.org__x__sys__unix"], "default": []}))

go_module(
    name = "github.com__containers__storage__pkg__archive",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/archive"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__storage__pkg__fileutils",
        ":github.com__containers__storage__pkg__idtools",
        ":github.com__containers__storage__pkg__pools",
        ":github.com__containers__storage__pkg__promise",
        ":github.com__containers__storage__pkg__system",
        ":github.com__containers__storage__pkg__unshare",
        ":github.com__klauspost__compress__zstd",
        ":github.com__klauspost__pgzip",
        ":github.com__opencontainers__runc__libcontainer__userns",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
        ":github.com__ulikunitz__xz",
        ":golang.org__x__sys__unix",
    ],
)

go_module(
    name = "github.com__containers__storage__pkg__chrootarchive",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/chrootarchive"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__storage__pkg__archive",
        ":github.com__containers__storage__pkg__idtools",
        ":github.com__containers__storage__pkg__reexec",
        ":github.com__containers__storage__pkg__system",
        ":github.com__json-iterator__go",
        ":github.com__opencontainers__runc__libcontainer__userns",
        ":github.com__pkg__errors",
        ":golang.org__x__sys__unix",
    ] + select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [
            ":github.com__containers__storage__pkg__mount",
            ":github.com__syndtr__gocapability__capability",
        ],
        "default": [],
    }),
)

go_module(
    name = "github.com__containers__storage__pkg__chunked",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/chunked"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containerd__stargz-snapshotter__estargz",
        ":github.com__containers__storage",
        ":github.com__containers__storage__drivers",
        ":github.com__containers__storage__pkg__chunked__compressor",
        ":github.com__containers__storage__pkg__chunked__internal",
        ":github.com__klauspost__compress__zstd",
        ":github.com__klauspost__pgzip",
        ":github.com__opencontainers__go-digest",
        ":github.com__pkg__errors",
        ":github.com__vbatts__tar-split__archive__tar",
    ] + select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [
            ":github.com__containers__storage__drivers__copy",
            ":github.com__containers__storage__pkg__archive",
            ":github.com__containers__storage__pkg__idtools",
            ":github.com__containers__storage__pkg__system",
            ":github.com__containers__storage__types",
            ":github.com__sirupsen__logrus",
            ":golang.org__x__sys__unix",
        ],
        "default": [],
    }),
)

go_module(
    name = "github.com__containers__storage__pkg__chunked__compressor",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/chunked/compressor"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__storage__pkg__chunked__internal",
        ":github.com__containers__storage__pkg__ioutils",
        ":github.com__opencontainers__go-digest",
        ":github.com__vbatts__tar-split__archive__tar",
    ],
)

go_module(
    name = "github.com__containers__storage__pkg__chunked__internal",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/chunked/internal"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__klauspost__compress__zstd",
        ":github.com__opencontainers__go-digest",
    ],
)

go_module(
    name = "github.com__containers__storage__pkg__config",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/config"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [],
)

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__storage__pkg__devicemapper", download = ":_github.com__containers__storage#download", install = ["pkg/devicemapper"], module = "github.com/containers/storage", visibility = ["PUBLIC"], deps = select({":__config_darwin_amd64": [], ":__config_linux_amd64": [":github.com__sirupsen__logrus", ":golang.org__x__sys__unix"], "default": []}))

go_module(
    name = "github.com__containers__storage__pkg__directory",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/directory"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [],
)

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__storage__pkg__dmesg", download = ":_github.com__containers__storage#download", install = ["pkg/dmesg"], module = "github.com/containers/storage", visibility = ["PUBLIC"], deps = select({":__config_darwin_amd64": [], ":__config_linux_amd64": [":golang.org__x__sys__unix"], "default": []}))

go_module(
    name = "github.com__containers__storage__pkg__fileutils",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/fileutils"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__storage__pkg__fsutils", download = ":_github.com__containers__storage#download", install = ["pkg/fsutils"], module = "github.com/containers/storage", visibility = ["PUBLIC"], deps = select({":__config_darwin_amd64": [], ":__config_linux_amd64": [":golang.org__x__sys__unix"], "default": []}))

go_module(
    name = "github.com__containers__storage__pkg__homedir",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/homedir"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [":github.com__containers__storage__pkg__unshare"],
)

go_module(
    name = "github.com__containers__storage__pkg__idtools",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/idtools"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__storage__pkg__system",
        ":github.com__opencontainers__runc__libcontainer__user",
        ":github.com__pkg__errors",
    ],
)

go_module(
    name = "github.com__containers__storage__pkg__ioutils",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/ioutils"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__net__context"] + select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [":golang.org__x__sys__unix"],
        "default": [],
    }),
)

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__storage__pkg__locker", download = ":_github.com__containers__storage#download", install = ["pkg/locker"], module = "github.com/containers/storage", visibility = ["PUBLIC"], deps = [])

go_module(
    name = "github.com__containers__storage__pkg__lockfile",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/lockfile"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__storage__pkg__stringid",
        ":github.com__containers__storage__pkg__system",
        ":github.com__pkg__errors",
        ":golang.org__x__sys__unix",
    ],
)

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__storage__pkg__loopback", download = ":_github.com__containers__storage#download", install = ["pkg/loopback"], module = "github.com/containers/storage", visibility = ["PUBLIC"], deps = select({":__config_darwin_amd64": [], ":__config_linux_amd64": [":github.com__sirupsen__logrus"], "default": []}))

go_module(
    name = "github.com__containers__storage__pkg__mount",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/mount"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__moby__sys__mountinfo",
        ":golang.org__x__sys__unix",
    ],
)

go_module(
    name = "github.com__containers__storage__pkg__parsers",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/parsers"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [],
)

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__containers__storage__pkg__parsers__kernel", download = ":_github.com__containers__storage#download", install = ["pkg/parsers/kernel"], module = "github.com/containers/storage", visibility = ["PUBLIC"], deps = select({":__config_darwin_amd64": [], ":__config_linux_amd64": [":github.com__sirupsen__logrus", ":golang.org__x__sys__unix"], "default": []}))

go_module(
    name = "github.com__containers__storage__pkg__pools",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/pools"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [":github.com__containers__storage__pkg__ioutils"],
)

go_module(
    name = "github.com__containers__storage__pkg__promise",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/promise"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__containers__storage__pkg__reexec",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/reexec"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__containers__storage__pkg__stringid",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/stringid"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__containers__storage__pkg__stringutils",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/stringutils"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__containers__storage__pkg__system",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/system"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__storage__pkg__mount",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
        ":golang.org__x__sys__unix",
    ] + select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [":github.com__docker__go-units"],
        "default": [],
    }),
)

go_module(
    name = "github.com__containers__storage__pkg__tarlog",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/tarlog"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__sirupsen__logrus",
        ":github.com__vbatts__tar-split__archive__tar",
    ],
)

go_module(
    name = "github.com__containers__storage__pkg__truncindex",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/truncindex"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [":github.com__tchap__go-patricia__patricia"],
)

go_module(
    name = "github.com__containers__storage__pkg__unshare",
    download = ":_github.com__containers__storage#download",
    install = ["pkg/unshare"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containers__storage__pkg__idtools",
        ":github.com__opencontainers__runtime-spec__specs-go",
        ":github.com__pkg__errors",
        ":github.com__syndtr__gocapability__capability",
    ] + select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [
            ":github.com__containers__storage__pkg__reexec",
            ":github.com__sirupsen__logrus",
        ],
        "default": [],
    }),
)

go_module(
    name = "github.com__containers__storage__types",
    download = ":_github.com__containers__storage#download",
    install = ["types"],
    module = "github.com/containers/storage",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__BurntSushi__toml",
        ":github.com__containers__storage__drivers__overlay",
        ":github.com__containers__storage__pkg__config",
        ":github.com__containers__storage__pkg__homedir",
        ":github.com__containers__storage__pkg__idtools",
        ":github.com__containers__storage__pkg__system",
        ":github.com__pkg__errors",
        ":github.com__sirupsen__logrus",
    ],
)

go_mod_download(
    name = "github.com__coreos__go-systemd__v22",
    _tag = "download",
    module = "github.com/coreos/go-systemd/v22",
    version = "v22.3.2",
)

go_module(
    name = "github.com__coreos__go-systemd__v22__dbus",
    download = ":_github.com__coreos__go-systemd__v22#download",
    install = ["dbus"],
    module = "github.com/coreos/go-systemd/v22",
    visibility = ["PUBLIC"],
    deps = [":github.com__godbus__dbus__v5"],
)

go_mod_download(
    name = "github.com__cpuguy83__go-md2man__v2",
    _tag = "download",
    module = "github.com/cpuguy83/go-md2man/v2",
    version = "v2.0.1",
)

go_module(
    name = "github.com__cpuguy83__go-md2man__v2__md2man",
    download = ":_github.com__cpuguy83__go-md2man__v2#download",
    install = ["md2man"],
    module = "github.com/cpuguy83/go-md2man/v2",
    visibility = ["PUBLIC"],
    deps = [":github.com__russross__blackfriday__v2"],
)

go_mod_download(
    name = "github.com__creack__pty",
    _tag = "download",
    module = "github.com/creack/pty",
    version = "v1.1.17",
)

go_module(
    name = "github.com__creack__pty",
    download = ":_github.com__creack__pty#download",
    install = ["."],
    module = "github.com/creack/pty",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__creasty__defaults",
    _tag = "download",
    module = "github.com/creasty/defaults",
    version = "v1.5.2",
)

go_module(
    name = "github.com__creasty__defaults",
    download = ":_github.com__creasty__defaults#download",
    install = ["."],
    module = "github.com/creasty/defaults",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__cri-o__ocicni",
    _tag = "download",
    module = "github.com/cri-o/ocicni",
    version = "v0.2.1-0.20210621164014-d0acc7862283",
)

go_module(
    name = "github.com__cri-o__ocicni__pkg__ocicni",
    download = ":_github.com__cri-o__ocicni#download",
    install = ["pkg/ocicni"],
    module = "github.com/cri-o/ocicni",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containernetworking__cni__libcni",
        ":github.com__containernetworking__cni__pkg__invoke",
        ":github.com__containernetworking__cni__pkg__types",
        ":github.com__containernetworking__cni__pkg__types__current",
        ":github.com__containernetworking__cni__pkg__version",
        ":github.com__fsnotify__fsnotify",
        ":github.com__sirupsen__logrus",
    ] + select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [
            ":github.com__containernetworking__plugins__pkg__ns",
            ":github.com__vishvananda__netlink",
        ],
        "default": [],
    }),
)

go_mod_download(
    name = "github.com__disiqueira__gotree__v3",
    _tag = "download",
    module = "github.com/disiqueira/gotree/v3",
    version = "v3.0.2",
)

go_module(
    name = "github.com__disiqueira__gotree__v3",
    download = ":_github.com__disiqueira__gotree__v3#download",
    install = ["."],
    module = "github.com/disiqueira/gotree/v3",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__docker__distribution",
    _tag = "download",
    module = "github.com/docker/distribution",
    version = "v2.7.1+incompatible",
)

go_module(
    name = "github.com__docker__distribution",
    download = ":_github.com__docker__distribution#download",
    install = ["."],
    module = "github.com/docker/distribution",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__docker__distribution__reference",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__image-spec__specs-go__v1",
    ],
)

go_module(
    name = "github.com__docker__distribution__digestset",
    download = ":_github.com__docker__distribution#download",
    install = ["digestset"],
    module = "github.com/docker/distribution",
    visibility = ["PUBLIC"],
    deps = [":github.com__opencontainers__go-digest"],
)

go_module(
    name = "github.com__docker__distribution__metrics",
    download = ":_github.com__docker__distribution#download",
    install = ["metrics"],
    module = "github.com/docker/distribution",
    visibility = ["PUBLIC"],
    deps = [":github.com__docker__go-metrics"],
)

go_module(
    name = "github.com__docker__distribution__reference",
    download = ":_github.com__docker__distribution#download",
    install = ["reference"],
    module = "github.com/docker/distribution",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__docker__distribution__digestset",
        ":github.com__opencontainers__go-digest",
    ],
)

go_module(
    name = "github.com__docker__distribution__registry__api__errcode",
    download = ":_github.com__docker__distribution#download",
    install = ["registry/api/errcode"],
    module = "github.com/docker/distribution",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__docker__distribution__registry__api__v2",
    download = ":_github.com__docker__distribution#download",
    install = ["registry/api/v2"],
    module = "github.com/docker/distribution",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__docker__distribution__reference",
        ":github.com__docker__distribution__registry__api__errcode",
        ":github.com__gorilla__mux",
        ":github.com__opencontainers__go-digest",
    ],
)

go_module(
    name = "github.com__docker__distribution__registry__client",
    download = ":_github.com__docker__distribution#download",
    install = ["registry/client"],
    module = "github.com/docker/distribution",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__docker__distribution",
        ":github.com__docker__distribution__reference",
        ":github.com__docker__distribution__registry__api__errcode",
        ":github.com__docker__distribution__registry__api__v2",
        ":github.com__docker__distribution__registry__client__auth__challenge",
        ":github.com__docker__distribution__registry__client__transport",
        ":github.com__docker__distribution__registry__storage__cache",
        ":github.com__docker__distribution__registry__storage__cache__memory",
        ":github.com__opencontainers__go-digest",
    ],
)

go_module(
    name = "github.com__docker__distribution__registry__client__auth__challenge",
    download = ":_github.com__docker__distribution#download",
    install = ["registry/client/auth/challenge"],
    module = "github.com/docker/distribution",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__docker__distribution__registry__client__transport",
    download = ":_github.com__docker__distribution#download",
    install = ["registry/client/transport"],
    module = "github.com/docker/distribution",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__docker__distribution__registry__storage__cache",
    download = ":_github.com__docker__distribution#download",
    install = ["registry/storage/cache"],
    module = "github.com/docker/distribution",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__docker__distribution",
        ":github.com__docker__distribution__metrics",
        ":github.com__opencontainers__go-digest",
    ],
)

go_module(
    name = "github.com__docker__distribution__registry__storage__cache__memory",
    download = ":_github.com__docker__distribution#download",
    install = ["registry/storage/cache/memory"],
    module = "github.com/docker/distribution",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__docker__distribution",
        ":github.com__docker__distribution__reference",
        ":github.com__docker__distribution__registry__storage__cache",
        ":github.com__opencontainers__go-digest",
    ],
)

go_mod_download(
    name = "github.com__docker__docker",
    _tag = "download",
    module = "github.com/docker/docker",
    version = "v20.10.10+incompatible",
)

go_module(
    name = "github.com__docker__docker__api",
    download = ":_github.com__docker__docker#download",
    install = ["api"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__docker__docker__api__types",
    download = ":_github.com__docker__docker#download",
    install = ["api/types"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__docker__docker__api__types__container",
        ":github.com__docker__docker__api__types__filters",
        ":github.com__docker__docker__api__types__mount",
        ":github.com__docker__docker__api__types__network",
        ":github.com__docker__docker__api__types__registry",
        ":github.com__docker__docker__api__types__swarm",
        ":github.com__docker__go-connections__nat",
        ":github.com__docker__go-units",
        ":github.com__opencontainers__image-spec__specs-go__v1",
    ],
)

go_module(
    name = "github.com__docker__docker__api__types__blkiodev",
    download = ":_github.com__docker__docker#download",
    install = ["api/types/blkiodev"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__docker__docker__api__types__container",
    download = ":_github.com__docker__docker#download",
    install = ["api/types/container"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__docker__docker__api__types__blkiodev",
        ":github.com__docker__docker__api__types__mount",
        ":github.com__docker__docker__api__types__strslice",
        ":github.com__docker__go-connections__nat",
        ":github.com__docker__go-units",
    ],
)

go_module(
    name = "github.com__docker__docker__api__types__events",
    download = ":_github.com__docker__docker#download",
    install = ["api/types/events"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__docker__docker__api__types__filters",
    download = ":_github.com__docker__docker#download",
    install = ["api/types/filters"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [":github.com__docker__docker__api__types__versions"],
)

go_module(
    name = "github.com__docker__docker__api__types__image",
    download = ":_github.com__docker__docker#download",
    install = ["api/types/image"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__docker__docker__api__types__mount",
    download = ":_github.com__docker__docker#download",
    install = ["api/types/mount"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__docker__docker__api__types__network",
    download = ":_github.com__docker__docker#download",
    install = ["api/types/network"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [":github.com__docker__docker__api__types__filters"],
)

go_module(
    name = "github.com__docker__docker__api__types__registry",
    download = ":_github.com__docker__docker#download",
    install = ["api/types/registry"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [":github.com__opencontainers__image-spec__specs-go__v1"],
)

go_module(
    name = "github.com__docker__docker__api__types__strslice",
    download = ":_github.com__docker__docker#download",
    install = ["api/types/strslice"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__docker__docker__api__types__swarm",
    download = ":_github.com__docker__docker#download",
    install = ["api/types/swarm"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__docker__docker__api__types__container",
        ":github.com__docker__docker__api__types__mount",
        ":github.com__docker__docker__api__types__network",
        ":github.com__docker__docker__api__types__swarm__runtime",
        ":github.com__docker__go-units",
    ],
)

go_module(
    name = "github.com__docker__docker__api__types__swarm__runtime",
    download = ":_github.com__docker__docker#download",
    install = ["api/types/swarm/runtime"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [":github.com__gogo__protobuf__proto"],
)

go_module(
    name = "github.com__docker__docker__api__types__time",
    download = ":_github.com__docker__docker#download",
    install = ["api/types/time"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__docker__docker__api__types__versions",
    download = ":_github.com__docker__docker#download",
    install = ["api/types/versions"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__docker__docker__api__types__volume",
    download = ":_github.com__docker__docker#download",
    install = ["api/types/volume"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [":github.com__docker__docker__api__types"],
)

go_module(
    name = "github.com__docker__docker__client",
    download = ":_github.com__docker__docker#download",
    install = ["client"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containerd__containerd__platforms",
        ":github.com__docker__distribution__reference",
        ":github.com__docker__docker__api",
        ":github.com__docker__docker__api__types",
        ":github.com__docker__docker__api__types__container",
        ":github.com__docker__docker__api__types__events",
        ":github.com__docker__docker__api__types__filters",
        ":github.com__docker__docker__api__types__image",
        ":github.com__docker__docker__api__types__network",
        ":github.com__docker__docker__api__types__registry",
        ":github.com__docker__docker__api__types__swarm",
        ":github.com__docker__docker__api__types__time",
        ":github.com__docker__docker__api__types__versions",
        ":github.com__docker__docker__api__types__volume",
        ":github.com__docker__docker__errdefs",
        ":github.com__docker__go-connections__sockets",
        ":github.com__docker__go-connections__tlsconfig",
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__image-spec__specs-go__v1",
        ":github.com__pkg__errors",
    ],
)

go_module(
    name = "github.com__docker__docker__errdefs",
    download = ":_github.com__docker__docker#download",
    install = ["errdefs"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__containerd__containerd__errdefs",
        ":github.com__docker__distribution__registry__api__errcode",
        ":github.com__sirupsen__logrus",
        ":google.golang.org__grpc__codes",
        ":google.golang.org__grpc__status",
    ],
)

go_module(
    name = "github.com__docker__docker__pkg__broadcaster",
    download = ":_github.com__docker__docker#download",
    install = ["pkg/broadcaster"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__docker__docker__pkg__homedir",
    download = ":_github.com__docker__docker#download",
    install = ["pkg/homedir"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__docker__docker__pkg__reexec",
    download = ":_github.com__docker__docker#download",
    install = ["pkg/reexec"],
    module = "github.com/docker/docker",
    visibility = ["PUBLIC"],
    deps = select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [":golang.org__x__sys__unix"],
        "default": [],
    }),
)

go_mod_download(
    name = "github.com__docker__docker-credential-helpers",
    _tag = "download",
    module = "github.com/docker/docker-credential-helpers",
    version = "v0.6.4",
)

go_module(
    name = "github.com__docker__docker-credential-helpers__client",
    download = ":_github.com__docker__docker-credential-helpers#download",
    install = ["client"],
    module = "github.com/docker/docker-credential-helpers",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__docker__docker-credential-helpers__credentials",
        ":golang.org__x__sys__execabs",
    ],
)

go_module(
    name = "github.com__docker__docker-credential-helpers__credentials",
    download = ":_github.com__docker__docker-credential-helpers#download",
    install = ["credentials"],
    module = "github.com/docker/docker-credential-helpers",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__docker__go-connections",
    _tag = "download",
    module = "github.com/docker/go-connections",
    version = "v0.4.0",
)

go_module(
    name = "github.com__docker__go-connections__nat",
    download = ":_github.com__docker__go-connections#download",
    install = ["nat"],
    module = "github.com/docker/go-connections",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__docker__go-connections__sockets",
    download = ":_github.com__docker__go-connections#download",
    install = ["sockets"],
    module = "github.com/docker/go-connections",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__net__proxy"],
)

go_module(
    name = "github.com__docker__go-connections__tlsconfig",
    download = ":_github.com__docker__go-connections#download",
    install = ["tlsconfig"],
    module = "github.com/docker/go-connections",
    visibility = ["PUBLIC"],
    deps = [":github.com__pkg__errors"],
)

go_mod_download(
    name = "github.com__docker__go-metrics",
    _tag = "download",
    module = "github.com/docker/go-metrics",
    version = "v0.0.1",
)

go_module(
    name = "github.com__docker__go-metrics",
    download = ":_github.com__docker__go-metrics#download",
    install = ["."],
    module = "github.com/docker/go-metrics",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__prometheus__client_golang__prometheus",
        ":github.com__prometheus__client_golang__prometheus__promhttp",
    ],
)

go_mod_download(
    name = "github.com__docker__go-units",
    _tag = "download",
    module = "github.com/docker/go-units",
    version = "v0.4.0",
)

go_module(
    name = "github.com__docker__go-units",
    download = ":_github.com__docker__go-units#download",
    install = ["."],
    module = "github.com/docker/go-units",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__fatih__color",
    _tag = "download",
    module = "github.com/fatih/color",
    version = "v1.13.0",
)

go_module(
    name = "github.com__fatih__color",
    download = ":_github.com__fatih__color#download",
    install = ["."],
    module = "github.com/fatih/color",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__mattn__go-colorable",
        ":github.com__mattn__go-isatty",
    ],
)

go_mod_download(
    name = "github.com__fsnotify__fsnotify",
    _tag = "download",
    module = "github.com/fsnotify/fsnotify",
    version = "v1.5.1",
)

go_module(
    name = "github.com__fsnotify__fsnotify",
    download = ":_github.com__fsnotify__fsnotify#download",
    install = ["."],
    module = "github.com/fsnotify/fsnotify",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__sys__unix"],
)

go_mod_download(
    name = "github.com__ghodss__yaml",
    _tag = "download",
    module = "github.com/ghodss/yaml",
    version = "v1.0.0",
)

go_module(
    name = "github.com__ghodss__yaml",
    download = ":_github.com__ghodss__yaml#download",
    install = ["."],
    module = "github.com/ghodss/yaml",
    visibility = ["PUBLIC"],
    deps = [":gopkg.in__yaml.v2"],
)

go_mod_download(
    name = "github.com__go-playground__locales",
    _tag = "download",
    module = "github.com/go-playground/locales",
    version = "v0.14.0",
)

go_module(
    name = "github.com__go-playground__locales",
    download = ":_github.com__go-playground__locales#download",
    install = ["."],
    module = "github.com/go-playground/locales",
    visibility = ["PUBLIC"],
    deps = [":github.com__go-playground__locales__currency"],
)

go_module(
    name = "github.com__go-playground__locales__currency",
    download = ":_github.com__go-playground__locales#download",
    install = ["currency"],
    module = "github.com/go-playground/locales",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__go-playground__universal-translator",
    _tag = "download",
    module = "github.com/go-playground/universal-translator",
    version = "v0.18.0",
)

go_module(
    name = "github.com__go-playground__universal-translator",
    download = ":_github.com__go-playground__universal-translator#download",
    install = ["."],
    module = "github.com/go-playground/universal-translator",
    visibility = ["PUBLIC"],
    deps = [":github.com__go-playground__locales"],
)

go_mod_download(
    name = "github.com__go-playground__validator__v10",
    _tag = "download",
    module = "github.com/go-playground/validator/v10",
    version = "v10.9.0",
)

go_module(
    name = "github.com__go-playground__validator__v10",
    download = ":_github.com__go-playground__validator__v10#download",
    install = ["."],
    module = "github.com/go-playground/validator/v10",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__go-playground__universal-translator",
        ":github.com__leodido__go-urn",
        ":golang.org__x__crypto__sha3",
        ":golang.org__x__text__language",
    ],
)

go_mod_download(
    name = "github.com__godbus__dbus__v5",
    _tag = "download",
    module = "github.com/godbus/dbus/v5",
    version = "v5.0.5",
)

go_module(
    name = "github.com__godbus__dbus__v5",
    download = ":_github.com__godbus__dbus__v5#download",
    install = ["."],
    module = "github.com/godbus/dbus/v5",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__gogo__protobuf",
    _tag = "download",
    module = "github.com/gogo/protobuf",
    version = "v1.3.2",
)

go_module(
    name = "github.com__gogo__protobuf__proto",
    download = ":_github.com__gogo__protobuf#download",
    install = ["proto"],
    module = "github.com/gogo/protobuf",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__golang__protobuf",
    _tag = "download",
    module = "github.com/golang/protobuf",
    version = "v1.5.2",
)

go_module(
    name = "github.com__golang__protobuf__proto",
    download = ":_github.com__golang__protobuf#download",
    install = ["proto"],
    module = "github.com/golang/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__encoding__prototext",
        ":google.golang.org__protobuf__encoding__protowire",
        ":google.golang.org__protobuf__proto",
        ":google.golang.org__protobuf__reflect__protodesc",
        ":google.golang.org__protobuf__reflect__protoreflect",
        ":google.golang.org__protobuf__reflect__protoregistry",
        ":google.golang.org__protobuf__runtime__protoiface",
        ":google.golang.org__protobuf__runtime__protoimpl",
    ],
)

go_module(
    name = "github.com__golang__protobuf__ptypes",
    download = ":_github.com__golang__protobuf#download",
    install = ["ptypes"],
    module = "github.com/golang/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__golang__protobuf__proto",
        ":github.com__golang__protobuf__ptypes__any",
        ":github.com__golang__protobuf__ptypes__duration",
        ":github.com__golang__protobuf__ptypes__timestamp",
        ":google.golang.org__protobuf__reflect__protoreflect",
        ":google.golang.org__protobuf__reflect__protoregistry",
    ],
)

go_module(
    name = "github.com__golang__protobuf__ptypes__any",
    download = ":_github.com__golang__protobuf#download",
    install = ["ptypes/any"],
    module = "github.com/golang/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__reflect__protoreflect",
        ":google.golang.org__protobuf__runtime__protoimpl",
        ":google.golang.org__protobuf__types__known__anypb",
    ],
)

go_module(
    name = "github.com__golang__protobuf__ptypes__duration",
    download = ":_github.com__golang__protobuf#download",
    install = ["ptypes/duration"],
    module = "github.com/golang/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__reflect__protoreflect",
        ":google.golang.org__protobuf__runtime__protoimpl",
        ":google.golang.org__protobuf__types__known__durationpb",
    ],
)

go_module(
    name = "github.com__golang__protobuf__ptypes__timestamp",
    download = ":_github.com__golang__protobuf#download",
    install = ["ptypes/timestamp"],
    module = "github.com/golang/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__reflect__protoreflect",
        ":google.golang.org__protobuf__runtime__protoimpl",
        ":google.golang.org__protobuf__types__known__timestamppb",
    ],
)

go_mod_download(
    name = "github.com__google__go-intervals",
    _tag = "download",
    module = "github.com/google/go-intervals",
    version = "v0.0.2",
)

go_module(
    name = "github.com__google__go-intervals__intervalset",
    download = ":_github.com__google__go-intervals#download",
    install = ["intervalset"],
    module = "github.com/google/go-intervals",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__google__uuid",
    _tag = "download",
    module = "github.com/google/uuid",
    version = "v1.3.0",
)

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__google__uuid", download = ":_github.com__google__uuid#download", install = ["."], module = "github.com/google/uuid", visibility = ["PUBLIC"], deps = [])

go_mod_download(
    name = "github.com__gorilla__mux",
    _tag = "download",
    module = "github.com/gorilla/mux",
    version = "v1.8.0",
)

go_module(
    name = "github.com__gorilla__mux",
    download = ":_github.com__gorilla__mux#download",
    install = ["."],
    module = "github.com/gorilla/mux",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__gorilla__schema",
    _tag = "download",
    module = "github.com/gorilla/schema",
    version = "v1.2.0",
)

go_module(
    name = "github.com__gorilla__schema",
    download = ":_github.com__gorilla__schema#download",
    install = ["."],
    module = "github.com/gorilla/schema",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__gosuri__uilive",
    _tag = "download",
    module = "github.com/gosuri/uilive",
    version = "v0.0.4",
)

go_module(
    name = "github.com__gosuri__uilive",
    download = ":_github.com__gosuri__uilive#download",
    install = ["."],
    module = "github.com/gosuri/uilive",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__hashicorp__errwrap",
    _tag = "download",
    module = "github.com/hashicorp/errwrap",
    version = "v1.0.0",
)

go_module(
    name = "github.com__hashicorp__errwrap",
    download = ":_github.com__hashicorp__errwrap#download",
    install = ["."],
    module = "github.com/hashicorp/errwrap",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__hashicorp__go-multierror",
    _tag = "download",
    module = "github.com/hashicorp/go-multierror",
    version = "v1.1.1",
)

go_module(
    name = "github.com__hashicorp__go-multierror",
    download = ":_github.com__hashicorp__go-multierror#download",
    install = ["."],
    module = "github.com/hashicorp/go-multierror",
    visibility = ["PUBLIC"],
    deps = [":github.com__hashicorp__errwrap"],
)

go_mod_download(
    name = "github.com__hashicorp__hcl",
    _tag = "download",
    module = "github.com/hashicorp/hcl",
    version = "v1.0.0",
)

go_module(
    name = "github.com__hashicorp__hcl",
    download = ":_github.com__hashicorp__hcl#download",
    install = ["."],
    module = "github.com/hashicorp/hcl",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__hashicorp__hcl__hcl__ast",
        ":github.com__hashicorp__hcl__hcl__parser",
        ":github.com__hashicorp__hcl__hcl__token",
        ":github.com__hashicorp__hcl__json__parser",
    ],
)

go_module(
    name = "github.com__hashicorp__hcl__hcl__ast",
    download = ":_github.com__hashicorp__hcl#download",
    install = ["hcl/ast"],
    module = "github.com/hashicorp/hcl",
    visibility = ["PUBLIC"],
    deps = [":github.com__hashicorp__hcl__hcl__token"],
)

go_module(
    name = "github.com__hashicorp__hcl__hcl__parser",
    download = ":_github.com__hashicorp__hcl#download",
    install = ["hcl/parser"],
    module = "github.com/hashicorp/hcl",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__hashicorp__hcl__hcl__ast",
        ":github.com__hashicorp__hcl__hcl__scanner",
        ":github.com__hashicorp__hcl__hcl__token",
    ],
)

go_module(
    name = "github.com__hashicorp__hcl__hcl__printer",
    download = ":_github.com__hashicorp__hcl#download",
    install = ["hcl/printer"],
    module = "github.com/hashicorp/hcl",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__hashicorp__hcl__hcl__ast",
        ":github.com__hashicorp__hcl__hcl__parser",
        ":github.com__hashicorp__hcl__hcl__token",
    ],
)

go_module(
    name = "github.com__hashicorp__hcl__hcl__scanner",
    download = ":_github.com__hashicorp__hcl#download",
    install = ["hcl/scanner"],
    module = "github.com/hashicorp/hcl",
    visibility = ["PUBLIC"],
    deps = [":github.com__hashicorp__hcl__hcl__token"],
)

go_module(
    name = "github.com__hashicorp__hcl__hcl__strconv",
    download = ":_github.com__hashicorp__hcl#download",
    install = ["hcl/strconv"],
    module = "github.com/hashicorp/hcl",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__hashicorp__hcl__hcl__token",
    download = ":_github.com__hashicorp__hcl#download",
    install = ["hcl/token"],
    module = "github.com/hashicorp/hcl",
    visibility = ["PUBLIC"],
    deps = [":github.com__hashicorp__hcl__hcl__strconv"],
)

go_module(
    name = "github.com__hashicorp__hcl__json__parser",
    download = ":_github.com__hashicorp__hcl#download",
    install = ["json/parser"],
    module = "github.com/hashicorp/hcl",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__hashicorp__hcl__hcl__ast",
        ":github.com__hashicorp__hcl__hcl__token",
        ":github.com__hashicorp__hcl__json__scanner",
        ":github.com__hashicorp__hcl__json__token",
    ],
)

go_module(
    name = "github.com__hashicorp__hcl__json__scanner",
    download = ":_github.com__hashicorp__hcl#download",
    install = ["json/scanner"],
    module = "github.com/hashicorp/hcl",
    visibility = ["PUBLIC"],
    deps = [":github.com__hashicorp__hcl__json__token"],
)

go_module(
    name = "github.com__hashicorp__hcl__json__token",
    download = ":_github.com__hashicorp__hcl#download",
    install = ["json/token"],
    module = "github.com/hashicorp/hcl",
    visibility = ["PUBLIC"],
    deps = [":github.com__hashicorp__hcl__hcl__token"],
)

go_mod_download(
    name = "github.com__hokaccha__go-prettyjson",
    _tag = "download",
    module = "github.com/hokaccha/go-prettyjson",
    version = "v0.0.0-20211117102719-0474bc63780f",
)

go_module(
    name = "github.com__hokaccha__go-prettyjson",
    download = ":_github.com__hokaccha__go-prettyjson#download",
    install = ["."],
    module = "github.com/hokaccha/go-prettyjson",
    visibility = ["PUBLIC"],
    deps = [":github.com__fatih__color"],
)

go_mod_download(
    name = "github.com__hpcloud__tail",
    _tag = "download",
    module = "github.com/hpcloud/tail",
    version = "v1.0.0",
)

go_module(
    name = "github.com__hpcloud__tail",
    download = ":_github.com__hpcloud__tail#download",
    install = ["."],
    module = "github.com/hpcloud/tail",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__hpcloud__tail__ratelimiter",
        ":github.com__hpcloud__tail__util",
        ":github.com__hpcloud__tail__watch",
        ":gopkg.in__tomb.v1",
    ],
)

go_module(
    name = "github.com__hpcloud__tail__ratelimiter",
    download = ":_github.com__hpcloud__tail#download",
    install = ["ratelimiter"],
    module = "github.com/hpcloud/tail",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__hpcloud__tail__util",
    download = ":_github.com__hpcloud__tail#download",
    install = ["util"],
    module = "github.com/hpcloud/tail",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__hpcloud__tail__watch",
    download = ":_github.com__hpcloud__tail#download",
    install = ["watch"],
    module = "github.com/hpcloud/tail",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__hpcloud__tail__util",
        ":gopkg.in__fsnotify.v1",
        ":gopkg.in__tomb.v1",
    ],
)

go_mod_download(
    name = "github.com__imdario__mergo",
    _tag = "download",
    module = "github.com/imdario/mergo",
    version = "v0.3.12",
)

go_module(
    name = "github.com__imdario__mergo",
    download = ":_github.com__imdario__mergo#download",
    install = ["."],
    module = "github.com/imdario/mergo",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__jinzhu__copier",
    _tag = "download",
    module = "github.com/jinzhu/copier",
    version = "v0.3.2",
)

go_module(
    name = "github.com__jinzhu__copier",
    download = ":_github.com__jinzhu__copier#download",
    install = ["."],
    module = "github.com/jinzhu/copier",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__json-iterator__go",
    _tag = "download",
    module = "github.com/json-iterator/go",
    version = "v1.1.12",
)

go_module(
    name = "github.com__json-iterator__go",
    download = ":_github.com__json-iterator__go#download",
    install = ["."],
    module = "github.com/json-iterator/go",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__modern-go__concurrent",
        ":github.com__modern-go__reflect2",
    ],
)

go_mod_download(
    name = "github.com__juju__ansiterm",
    _tag = "download",
    module = "github.com/juju/ansiterm",
    version = "v0.0.0-20180109212912-720a0952cc2a",
)

go_module(
    name = "github.com__juju__ansiterm",
    download = ":_github.com__juju__ansiterm#download",
    install = ["."],
    module = "github.com/juju/ansiterm",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__juju__ansiterm__tabwriter",
        ":github.com__mattn__go-colorable",
        ":github.com__mattn__go-isatty",
    ],
)

go_module(
    name = "github.com__juju__ansiterm__tabwriter",
    download = ":_github.com__juju__ansiterm#download",
    install = ["tabwriter"],
    module = "github.com/juju/ansiterm",
    visibility = ["PUBLIC"],
    deps = [":github.com__lunixbochs__vtclean"],
)

go_mod_download(
    name = "github.com__klauspost__compress",
    _tag = "download",
    module = "github.com/klauspost/compress",
    version = "v1.13.5",
)

go_module(
    name = "github.com__klauspost__compress",
    download = ":_github.com__klauspost__compress#download",
    install = ["."],
    module = "github.com/klauspost/compress",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__klauspost__compress__flate",
    download = ":_github.com__klauspost__compress#download",
    install = ["flate"],
    module = "github.com/klauspost/compress",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__klauspost__compress__fse",
    download = ":_github.com__klauspost__compress#download",
    install = ["fse"],
    module = "github.com/klauspost/compress",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__klauspost__compress__huff0",
    download = ":_github.com__klauspost__compress#download",
    install = ["huff0"],
    module = "github.com/klauspost/compress",
    visibility = ["PUBLIC"],
    deps = [":github.com__klauspost__compress__fse"],
)

go_module(
    name = "github.com__klauspost__compress__internal__snapref",
    download = ":_github.com__klauspost__compress#download",
    install = ["internal/snapref"],
    module = "github.com/klauspost/compress",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__klauspost__compress__zstd",
    download = ":_github.com__klauspost__compress#download",
    install = ["zstd"],
    module = "github.com/klauspost/compress",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__klauspost__compress",
        ":github.com__klauspost__compress__huff0",
        ":github.com__klauspost__compress__internal__snapref",
        ":github.com__klauspost__compress__zstd__internal__xxhash",
    ],
)

go_module(
    name = "github.com__klauspost__compress__zstd__internal__xxhash",
    download = ":_github.com__klauspost__compress#download",
    install = ["zstd/internal/xxhash"],
    module = "github.com/klauspost/compress",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__klauspost__pgzip",
    _tag = "download",
    module = "github.com/klauspost/pgzip",
    version = "v1.2.5",
)

go_module(
    name = "github.com__klauspost__pgzip",
    download = ":_github.com__klauspost__pgzip#download",
    install = ["."],
    module = "github.com/klauspost/pgzip",
    visibility = ["PUBLIC"],
    deps = [":github.com__klauspost__compress__flate"],
)

go_mod_download(
    name = "github.com__leodido__go-urn",
    _tag = "download",
    module = "github.com/leodido/go-urn",
    version = "v1.2.1",
)

go_module(
    name = "github.com__leodido__go-urn",
    download = ":_github.com__leodido__go-urn#download",
    install = ["."],
    module = "github.com/leodido/go-urn",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__lunixbochs__vtclean",
    _tag = "download",
    module = "github.com/lunixbochs/vtclean",
    version = "v0.0.0-20180621232353-2d01aacdc34a",
)

go_module(
    name = "github.com__lunixbochs__vtclean",
    download = ":_github.com__lunixbochs__vtclean#download",
    install = ["."],
    module = "github.com/lunixbochs/vtclean",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__magiconair__properties",
    _tag = "download",
    module = "github.com/magiconair/properties",
    version = "v1.8.5",
)

go_module(
    name = "github.com__magiconair__properties",
    download = ":_github.com__magiconair__properties#download",
    install = ["."],
    module = "github.com/magiconair/properties",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__manifoldco__promptui",
    _tag = "download",
    module = "github.com/manifoldco/promptui",
    version = "v0.8.0",
)

go_module(
    name = "github.com__manifoldco__promptui",
    download = ":_github.com__manifoldco__promptui#download",
    install = ["."],
    module = "github.com/manifoldco/promptui",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__chzyer__readline",
        ":github.com__juju__ansiterm",
        ":github.com__manifoldco__promptui__list",
        ":github.com__manifoldco__promptui__screenbuf",
    ],
)

go_module(
    name = "github.com__manifoldco__promptui__list",
    download = ":_github.com__manifoldco__promptui#download",
    install = ["list"],
    module = "github.com/manifoldco/promptui",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__manifoldco__promptui__screenbuf",
    download = ":_github.com__manifoldco__promptui#download",
    install = ["screenbuf"],
    module = "github.com/manifoldco/promptui",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__mattn__go-colorable",
    _tag = "download",
    module = "github.com/mattn/go-colorable",
    version = "v0.1.12",
)

go_module(
    name = "github.com__mattn__go-colorable",
    download = ":_github.com__mattn__go-colorable#download",
    install = ["."],
    module = "github.com/mattn/go-colorable",
    visibility = ["PUBLIC"],
    deps = [":github.com__mattn__go-isatty"],
)

go_mod_download(
    name = "github.com__mattn__go-isatty",
    _tag = "download",
    module = "github.com/mattn/go-isatty",
    version = "v0.0.14",
)

go_module(
    name = "github.com__mattn__go-isatty",
    download = ":_github.com__mattn__go-isatty#download",
    install = ["."],
    module = "github.com/mattn/go-isatty",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__sys__unix"],
)

go_mod_download(
    name = "github.com__mattn__go-runewidth",
    _tag = "download",
    module = "github.com/mattn/go-runewidth",
    version = "v0.0.13",
)

go_module(
    name = "github.com__mattn__go-runewidth",
    download = ":_github.com__mattn__go-runewidth#download",
    install = ["."],
    module = "github.com/mattn/go-runewidth",
    visibility = ["PUBLIC"],
    deps = [":github.com__rivo__uniseg"],
)

go_mod_download(
    name = "github.com__matttproud__golang_protobuf_extensions",
    _tag = "download",
    module = "github.com/matttproud/golang_protobuf_extensions",
    version = "v1.0.2-0.20181231171920-c182affec369",
)

go_module(
    name = "github.com__matttproud__golang_protobuf_extensions__pbutil",
    download = ":_github.com__matttproud__golang_protobuf_extensions#download",
    install = ["pbutil"],
    module = "github.com/matttproud/golang_protobuf_extensions",
    visibility = ["PUBLIC"],
    deps = [":github.com__golang__protobuf__proto"],
)

go_mod_download(
    name = "github.com__miekg__pkcs11",
    _tag = "download",
    module = "github.com/miekg/pkcs11",
    version = "v1.0.3",
)

go_module(
    name = "github.com__miekg__pkcs11",
    download = ":_github.com__miekg__pkcs11#download",
    install = ["."],
    module = "github.com/miekg/pkcs11",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__mistifyio__go-zfs",
    _tag = "download",
    module = "github.com/mistifyio/go-zfs",
    version = "v2.1.2-0.20190413222219-f784269be439+incompatible",
)

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__mistifyio__go-zfs", download = ":_github.com__mistifyio__go-zfs#download", install = ["."], module = "github.com/mistifyio/go-zfs", visibility = ["PUBLIC"], deps = select({":__config_darwin_amd64": [], ":__config_linux_amd64": [":github.com__google__uuid"], "default": []}))

go_mod_download(
    name = "github.com__mitchellh__mapstructure",
    _tag = "download",
    module = "github.com/mitchellh/mapstructure",
    version = "v1.4.3",
)

go_module(
    name = "github.com__mitchellh__mapstructure",
    download = ":_github.com__mitchellh__mapstructure#download",
    install = ["."],
    module = "github.com/mitchellh/mapstructure",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__moby__sys__mountinfo",
    _tag = "download",
    module = "github.com/moby/sys/mountinfo",
    version = "v0.4.1",
)

go_module(
    name = "github.com__moby__sys__mountinfo",
    download = ":_github.com__moby__sys__mountinfo#download",
    install = ["."],
    module = "github.com/moby/sys/mountinfo",
    visibility = ["PUBLIC"],
    deps = select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [":golang.org__x__sys__unix"],
        "default": [],
    }),
)

go_mod_download(
    name = "github.com__moby__term",
    _tag = "download",
    module = "github.com/moby/term",
    version = "v0.0.0-20210619224110-3f7ff695adc6",
)

go_module(
    name = "github.com__moby__term",
    download = ":_github.com__moby__term#download",
    install = ["."],
    module = "github.com/moby/term",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__sys__unix"],
)

go_mod_download(
    name = "github.com__modern-go__concurrent",
    _tag = "download",
    module = "github.com/modern-go/concurrent",
    version = "v0.0.0-20180306012644-bacd9c7ef1dd",
)

go_module(
    name = "github.com__modern-go__concurrent",
    download = ":_github.com__modern-go__concurrent#download",
    install = ["."],
    module = "github.com/modern-go/concurrent",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__modern-go__reflect2",
    _tag = "download",
    module = "github.com/modern-go/reflect2",
    version = "v1.0.2",
)

go_module(
    name = "github.com__modern-go__reflect2",
    download = ":_github.com__modern-go__reflect2#download",
    install = ["."],
    module = "github.com/modern-go/reflect2",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__mtrmac__gpgme",
    _tag = "download",
    module = "github.com/mtrmac/gpgme",
    version = "v0.1.2",
)

go_module(
    name = "github.com__mtrmac__gpgme",
    download = ":_github.com__mtrmac__gpgme#download",
    install = ["."],
    module = "github.com/mtrmac/gpgme",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__nxadm__tail",
    _tag = "download",
    module = "github.com/nxadm/tail",
    version = "v1.4.8",
)

go_module(
    name = "github.com__nxadm__tail",
    download = ":_github.com__nxadm__tail#download",
    install = ["."],
    module = "github.com/nxadm/tail",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__nxadm__tail__ratelimiter",
        ":github.com__nxadm__tail__util",
        ":github.com__nxadm__tail__watch",
        ":gopkg.in__tomb.v1",
    ],
)

go_module(
    name = "github.com__nxadm__tail__ratelimiter",
    download = ":_github.com__nxadm__tail#download",
    install = ["ratelimiter"],
    module = "github.com/nxadm/tail",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__nxadm__tail__util",
    download = ":_github.com__nxadm__tail#download",
    install = ["util"],
    module = "github.com/nxadm/tail",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__nxadm__tail__watch",
    download = ":_github.com__nxadm__tail#download",
    install = ["watch"],
    module = "github.com/nxadm/tail",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__fsnotify__fsnotify",
        ":github.com__nxadm__tail__util",
        ":gopkg.in__tomb.v1",
    ],
)

go_mod_download(
    name = "github.com__olekukonko__tablewriter",
    _tag = "download",
    module = "github.com/olekukonko/tablewriter",
    version = "v0.0.5",
)

go_module(
    name = "github.com__olekukonko__tablewriter",
    download = ":_github.com__olekukonko__tablewriter#download",
    install = ["."],
    module = "github.com/olekukonko/tablewriter",
    visibility = ["PUBLIC"],
    deps = [":github.com__mattn__go-runewidth"],
)

go_mod_download(
    name = "github.com__onsi__ginkgo",
    _tag = "download",
    module = "github.com/onsi/ginkgo",
    version = "v1.16.5",
)

go_module(
    name = "github.com__onsi__ginkgo",
    download = ":_github.com__onsi__ginkgo#download",
    install = ["."],
    module = "github.com/onsi/ginkgo",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__onsi__ginkgo__config",
        ":github.com__onsi__ginkgo__internal__codelocation",
        ":github.com__onsi__ginkgo__internal__global",
        ":github.com__onsi__ginkgo__internal__remote",
        ":github.com__onsi__ginkgo__internal__testingtproxy",
        ":github.com__onsi__ginkgo__internal__writer",
        ":github.com__onsi__ginkgo__reporters",
        ":github.com__onsi__ginkgo__reporters__stenographer",
        ":github.com__onsi__ginkgo__reporters__stenographer__support__go-colorable",
        ":github.com__onsi__ginkgo__types",
    ],
)

go_module(
    name = "github.com__onsi__ginkgo__config",
    download = ":_github.com__onsi__ginkgo#download",
    install = ["config"],
    module = "github.com/onsi/ginkgo",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__onsi__ginkgo__formatter",
    download = ":_github.com__onsi__ginkgo#download",
    install = ["formatter"],
    module = "github.com/onsi/ginkgo",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__onsi__ginkgo__internal__codelocation",
    download = ":_github.com__onsi__ginkgo#download",
    install = ["internal/codelocation"],
    module = "github.com/onsi/ginkgo",
    visibility = ["PUBLIC"],
    deps = [":github.com__onsi__ginkgo__types"],
)

go_module(
    name = "github.com__onsi__ginkgo__internal__containernode",
    download = ":_github.com__onsi__ginkgo#download",
    install = ["internal/containernode"],
    module = "github.com/onsi/ginkgo",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__onsi__ginkgo__internal__leafnodes",
        ":github.com__onsi__ginkgo__types",
    ],
)

go_module(
    name = "github.com__onsi__ginkgo__internal__failer",
    download = ":_github.com__onsi__ginkgo#download",
    install = ["internal/failer"],
    module = "github.com/onsi/ginkgo",
    visibility = ["PUBLIC"],
    deps = [":github.com__onsi__ginkgo__types"],
)

go_module(
    name = "github.com__onsi__ginkgo__internal__global",
    download = ":_github.com__onsi__ginkgo#download",
    install = ["internal/global"],
    module = "github.com/onsi/ginkgo",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__onsi__ginkgo__internal__failer",
        ":github.com__onsi__ginkgo__internal__suite",
    ],
)

go_module(
    name = "github.com__onsi__ginkgo__internal__leafnodes",
    download = ":_github.com__onsi__ginkgo#download",
    install = ["internal/leafnodes"],
    module = "github.com/onsi/ginkgo",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__onsi__ginkgo__internal__codelocation",
        ":github.com__onsi__ginkgo__internal__failer",
        ":github.com__onsi__ginkgo__types",
    ],
)

go_module(
    name = "github.com__onsi__ginkgo__internal__remote",
    download = ":_github.com__onsi__ginkgo#download",
    install = ["internal/remote"],
    module = "github.com/onsi/ginkgo",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__nxadm__tail",
        ":github.com__onsi__ginkgo__config",
        ":github.com__onsi__ginkgo__internal__spec_iterator",
        ":github.com__onsi__ginkgo__internal__writer",
        ":github.com__onsi__ginkgo__reporters",
        ":github.com__onsi__ginkgo__reporters__stenographer",
        ":github.com__onsi__ginkgo__types",
        ":golang.org__x__sys__unix",
    ],
)

go_module(
    name = "github.com__onsi__ginkgo__internal__spec",
    download = ":_github.com__onsi__ginkgo#download",
    install = ["internal/spec"],
    module = "github.com/onsi/ginkgo",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__onsi__ginkgo__internal__containernode",
        ":github.com__onsi__ginkgo__internal__leafnodes",
        ":github.com__onsi__ginkgo__types",
    ],
)

go_module(
    name = "github.com__onsi__ginkgo__internal__spec_iterator",
    download = ":_github.com__onsi__ginkgo#download",
    install = ["internal/spec_iterator"],
    module = "github.com/onsi/ginkgo",
    visibility = ["PUBLIC"],
    deps = [":github.com__onsi__ginkgo__internal__spec"],
)

go_module(
    name = "github.com__onsi__ginkgo__internal__specrunner",
    download = ":_github.com__onsi__ginkgo#download",
    install = ["internal/specrunner"],
    module = "github.com/onsi/ginkgo",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__onsi__ginkgo__config",
        ":github.com__onsi__ginkgo__internal__leafnodes",
        ":github.com__onsi__ginkgo__internal__spec",
        ":github.com__onsi__ginkgo__internal__spec_iterator",
        ":github.com__onsi__ginkgo__internal__writer",
        ":github.com__onsi__ginkgo__reporters",
        ":github.com__onsi__ginkgo__types",
    ],
)

go_module(
    name = "github.com__onsi__ginkgo__internal__suite",
    download = ":_github.com__onsi__ginkgo#download",
    install = ["internal/suite"],
    module = "github.com/onsi/ginkgo",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__onsi__ginkgo__config",
        ":github.com__onsi__ginkgo__internal__containernode",
        ":github.com__onsi__ginkgo__internal__failer",
        ":github.com__onsi__ginkgo__internal__leafnodes",
        ":github.com__onsi__ginkgo__internal__spec",
        ":github.com__onsi__ginkgo__internal__spec_iterator",
        ":github.com__onsi__ginkgo__internal__specrunner",
        ":github.com__onsi__ginkgo__internal__writer",
        ":github.com__onsi__ginkgo__reporters",
        ":github.com__onsi__ginkgo__types",
    ],
)

go_module(
    name = "github.com__onsi__ginkgo__internal__testingtproxy",
    download = ":_github.com__onsi__ginkgo#download",
    install = ["internal/testingtproxy"],
    module = "github.com/onsi/ginkgo",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__onsi__ginkgo__internal__writer",
    download = ":_github.com__onsi__ginkgo#download",
    install = ["internal/writer"],
    module = "github.com/onsi/ginkgo",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__onsi__ginkgo__reporters",
    download = ":_github.com__onsi__ginkgo#download",
    install = ["reporters"],
    module = "github.com/onsi/ginkgo",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__onsi__ginkgo__config",
        ":github.com__onsi__ginkgo__reporters__stenographer",
        ":github.com__onsi__ginkgo__types",
    ],
)

go_module(
    name = "github.com__onsi__ginkgo__reporters__stenographer",
    download = ":_github.com__onsi__ginkgo#download",
    install = ["reporters/stenographer"],
    module = "github.com/onsi/ginkgo",
    visibility = ["PUBLIC"],
    deps = [":github.com__onsi__ginkgo__types"],
)

go_module(
    name = "github.com__onsi__ginkgo__reporters__stenographer__support__go-colorable",
    download = ":_github.com__onsi__ginkgo#download",
    install = ["reporters/stenographer/support/go-colorable"],
    module = "github.com/onsi/ginkgo",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__onsi__ginkgo__types",
    download = ":_github.com__onsi__ginkgo#download",
    install = ["types"],
    module = "github.com/onsi/ginkgo",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__onsi__ginkgo__config",
        ":github.com__onsi__ginkgo__formatter",
    ],
)

go_mod_download(
    name = "github.com__onsi__gomega",
    _tag = "download",
    module = "github.com/onsi/gomega",
    version = "v1.17.0",
)

go_module(
    name = "github.com__onsi__gomega",
    download = ":_github.com__onsi__gomega#download",
    install = ["."],
    module = "github.com/onsi/gomega",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__onsi__gomega__internal",
        ":github.com__onsi__gomega__matchers",
        ":github.com__onsi__gomega__types",
    ],
)

go_module(
    name = "github.com__onsi__gomega__format",
    download = ":_github.com__onsi__gomega#download",
    install = ["format"],
    module = "github.com/onsi/gomega",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__onsi__gomega__internal",
    download = ":_github.com__onsi__gomega#download",
    install = ["internal"],
    module = "github.com/onsi/gomega",
    visibility = ["PUBLIC"],
    deps = [":github.com__onsi__gomega__types"],
)

go_module(
    name = "github.com__onsi__gomega__matchers",
    download = ":_github.com__onsi__gomega#download",
    install = ["matchers"],
    module = "github.com/onsi/gomega",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__onsi__gomega__format",
        ":github.com__onsi__gomega__matchers__support__goraph__bipartitegraph",
        ":github.com__onsi__gomega__types",
        ":golang.org__x__net__html__charset",
        ":gopkg.in__yaml.v2",
    ],
)

go_module(
    name = "github.com__onsi__gomega__matchers__support__goraph__bipartitegraph",
    download = ":_github.com__onsi__gomega#download",
    install = ["matchers/support/goraph/bipartitegraph"],
    module = "github.com/onsi/gomega",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__onsi__gomega__matchers__support__goraph__edge",
        ":github.com__onsi__gomega__matchers__support__goraph__node",
        ":github.com__onsi__gomega__matchers__support__goraph__util",
    ],
)

go_module(
    name = "github.com__onsi__gomega__matchers__support__goraph__edge",
    download = ":_github.com__onsi__gomega#download",
    install = ["matchers/support/goraph/edge"],
    module = "github.com/onsi/gomega",
    visibility = ["PUBLIC"],
    deps = [":github.com__onsi__gomega__matchers__support__goraph__node"],
)

go_module(
    name = "github.com__onsi__gomega__matchers__support__goraph__node",
    download = ":_github.com__onsi__gomega#download",
    install = ["matchers/support/goraph/node"],
    module = "github.com/onsi/gomega",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__onsi__gomega__matchers__support__goraph__util",
    download = ":_github.com__onsi__gomega#download",
    install = ["matchers/support/goraph/util"],
    module = "github.com/onsi/gomega",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__onsi__gomega__types",
    download = ":_github.com__onsi__gomega#download",
    install = ["types"],
    module = "github.com/onsi/gomega",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__opencontainers__go-digest",
    _tag = "download",
    module = "github.com/opencontainers/go-digest",
    version = "v1.0.0",
)

go_module(
    name = "github.com__opencontainers__go-digest",
    download = ":_github.com__opencontainers__go-digest#download",
    install = ["."],
    module = "github.com/opencontainers/go-digest",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__opencontainers__image-spec",
    _tag = "download",
    module = "github.com/opencontainers/image-spec",
    version = "v1.0.2-0.20210819154149-5ad6f50d6283",
)

go_module(
    name = "github.com__opencontainers__image-spec__specs-go",
    download = ":_github.com__opencontainers__image-spec#download",
    install = ["specs-go"],
    module = "github.com/opencontainers/image-spec",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__opencontainers__image-spec__specs-go__v1",
    download = ":_github.com__opencontainers__image-spec#download",
    install = ["specs-go/v1"],
    module = "github.com/opencontainers/image-spec",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__opencontainers__go-digest",
        ":github.com__opencontainers__image-spec__specs-go",
    ],
)

go_mod_download(
    name = "github.com__opencontainers__runc",
    _tag = "download",
    module = "github.com/opencontainers/runc",
    version = "v1.0.2",
)

go_module(
    name = "github.com__opencontainers__runc__libcontainer__devices",
    download = ":_github.com__opencontainers__runc#download",
    install = ["libcontainer/devices"],
    module = "github.com/opencontainers/runc",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__sys__unix"],
)

go_module(
    name = "github.com__opencontainers__runc__libcontainer__user",
    download = ":_github.com__opencontainers__runc#download",
    install = ["libcontainer/user"],
    module = "github.com/opencontainers/runc",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__sys__unix"],
)

go_module(
    name = "github.com__opencontainers__runc__libcontainer__userns",
    download = ":_github.com__opencontainers__runc#download",
    install = ["libcontainer/userns"],
    module = "github.com/opencontainers/runc",
    visibility = ["PUBLIC"],
    deps = [":github.com__opencontainers__runc__libcontainer__user"],
)

go_mod_download(
    name = "github.com__opencontainers__runtime-spec",
    _tag = "download",
    module = "github.com/opencontainers/runtime-spec",
    version = "v1.0.3-0.20210326190908-1c3f411f0417",
)

go_module(
    name = "github.com__opencontainers__runtime-spec__specs-go",
    download = ":_github.com__opencontainers__runtime-spec#download",
    install = ["specs-go"],
    module = "github.com/opencontainers/runtime-spec",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__opencontainers__runtime-tools",
    _tag = "download",
    module = "github.com/opencontainers/runtime-tools",
    version = "v0.9.0",
)

go_module(
    name = "github.com__opencontainers__runtime-tools__error",
    download = ":_github.com__opencontainers__runtime-tools#download",
    install = ["error"],
    module = "github.com/opencontainers/runtime-tools",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__opencontainers__runtime-tools__filepath",
    download = ":_github.com__opencontainers__runtime-tools#download",
    install = ["filepath"],
    module = "github.com/opencontainers/runtime-tools",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__opencontainers__runtime-tools__generate",
    download = ":_github.com__opencontainers__runtime-tools#download",
    install = ["generate"],
    module = "github.com/opencontainers/runtime-tools",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__opencontainers__runtime-spec__specs-go",
        ":github.com__opencontainers__runtime-tools__generate__seccomp",
        ":github.com__opencontainers__runtime-tools__validate",
        ":github.com__syndtr__gocapability__capability",
    ],
)

go_module(
    name = "github.com__opencontainers__runtime-tools__generate__seccomp",
    download = ":_github.com__opencontainers__runtime-tools#download",
    install = ["generate/seccomp"],
    module = "github.com/opencontainers/runtime-tools",
    visibility = ["PUBLIC"],
    deps = [":github.com__opencontainers__runtime-spec__specs-go"],
)

go_module(
    name = "github.com__opencontainers__runtime-tools__specerror",
    download = ":_github.com__opencontainers__runtime-tools#download",
    install = ["specerror"],
    module = "github.com/opencontainers/runtime-tools",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__hashicorp__go-multierror",
        ":github.com__opencontainers__runtime-tools__error",
    ],
)

go_module(
    name = "github.com__opencontainers__runtime-tools__validate",
    download = ":_github.com__opencontainers__runtime-tools#download",
    install = ["validate"],
    module = "github.com/opencontainers/runtime-tools",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__blang__semver",
        ":github.com__hashicorp__go-multierror",
        ":github.com__opencontainers__runtime-spec__specs-go",
        ":github.com__opencontainers__runtime-tools__filepath",
        ":github.com__opencontainers__runtime-tools__specerror",
        ":github.com__sirupsen__logrus",
        ":github.com__syndtr__gocapability__capability",
        ":github.com__xeipuuv__gojsonschema",
    ] + select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [":github.com__opencontainers__selinux__go-selinux__label"],
        "default": [],
    }),
)

go_mod_download(
    name = "github.com__opencontainers__selinux",
    _tag = "download",
    module = "github.com/opencontainers/selinux",
    version = "v1.8.5",
)

go_module(
    name = "github.com__opencontainers__selinux__go-selinux",
    download = ":_github.com__opencontainers__selinux#download",
    install = ["go-selinux"],
    module = "github.com/opencontainers/selinux",
    visibility = ["PUBLIC"],
    deps = select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [
            ":github.com__bits-and-blooms__bitset",
            ":github.com__opencontainers__selinux__pkg__pwalkdir",
            ":golang.org__x__sys__unix",
        ],
        "default": [],
    }),
)

go_module(
    name = "github.com__opencontainers__selinux__go-selinux__label",
    download = ":_github.com__opencontainers__selinux#download",
    install = ["go-selinux/label"],
    module = "github.com/opencontainers/selinux",
    visibility = ["PUBLIC"],
    deps = [":github.com__opencontainers__selinux__go-selinux"],
)

go_module(
    name = "github.com__opencontainers__selinux__pkg__pwalk",
    download = ":_github.com__opencontainers__selinux#download",
    install = ["pkg/pwalk"],
    module = "github.com/opencontainers/selinux",
    visibility = ["PUBLIC"],
    deps = [],
)

if is_platform(
    arch = ["amd64"],
    os = ["linux"],
):
    go_module(name = "github.com__opencontainers__selinux__pkg__pwalkdir", download = ":_github.com__opencontainers__selinux#download", install = ["pkg/pwalkdir"], module = "github.com/opencontainers/selinux", visibility = ["PUBLIC"], deps = [])

go_mod_download(
    name = "github.com__pelletier__go-toml",
    _tag = "download",
    module = "github.com/pelletier/go-toml",
    version = "v1.9.4",
)

go_module(
    name = "github.com__pelletier__go-toml",
    download = ":_github.com__pelletier__go-toml#download",
    install = ["."],
    module = "github.com/pelletier/go-toml",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__pkg__errors",
    _tag = "download",
    module = "github.com/pkg/errors",
    version = "v0.9.1",
)

go_module(
    name = "github.com__pkg__errors",
    download = ":_github.com__pkg__errors#download",
    install = ["."],
    module = "github.com/pkg/errors",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__prometheus__client_golang",
    _tag = "download",
    module = "github.com/prometheus/client_golang",
    version = "v1.7.1",
)

go_module(
    name = "github.com__prometheus__client_golang__prometheus",
    download = ":_github.com__prometheus__client_golang#download",
    install = ["prometheus"],
    module = "github.com/prometheus/client_golang",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__beorn7__perks__quantile",
        ":github.com__cespare__xxhash__v2",
        ":github.com__golang__protobuf__proto",
        ":github.com__golang__protobuf__ptypes",
        ":github.com__prometheus__client_golang__prometheus__internal",
        ":github.com__prometheus__client_model__go",
        ":github.com__prometheus__common__expfmt",
        ":github.com__prometheus__common__model",
        ":github.com__prometheus__procfs",
    ],
)

go_module(
    name = "github.com__prometheus__client_golang__prometheus__internal",
    download = ":_github.com__prometheus__client_golang#download",
    install = ["prometheus/internal"],
    module = "github.com/prometheus/client_golang",
    visibility = ["PUBLIC"],
    deps = [":github.com__prometheus__client_model__go"],
)

go_module(
    name = "github.com__prometheus__client_golang__prometheus__promhttp",
    download = ":_github.com__prometheus__client_golang#download",
    install = ["prometheus/promhttp"],
    module = "github.com/prometheus/client_golang",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__prometheus__client_golang__prometheus",
        ":github.com__prometheus__client_model__go",
        ":github.com__prometheus__common__expfmt",
    ],
)

go_mod_download(
    name = "github.com__prometheus__client_model",
    _tag = "download",
    module = "github.com/prometheus/client_model",
    version = "v0.2.0",
)

go_module(
    name = "github.com__prometheus__client_model__go",
    download = ":_github.com__prometheus__client_model#download",
    install = ["go"],
    module = "github.com/prometheus/client_model",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__golang__protobuf__proto",
        ":github.com__golang__protobuf__ptypes__timestamp",
    ],
)

go_mod_download(
    name = "github.com__prometheus__common",
    _tag = "download",
    module = "github.com/prometheus/common",
    version = "v0.10.0",
)

go_module(
    name = "github.com__prometheus__common__expfmt",
    download = ":_github.com__prometheus__common#download",
    install = ["expfmt"],
    module = "github.com/prometheus/common",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__golang__protobuf__proto",
        ":github.com__golang__protobuf__ptypes",
        ":github.com__matttproud__golang_protobuf_extensions__pbutil",
        ":github.com__prometheus__client_model__go",
        ":github.com__prometheus__common__internal__bitbucket.org__ww__goautoneg",
        ":github.com__prometheus__common__model",
    ],
)

go_module(
    name = "github.com__prometheus__common__internal__bitbucket.org__ww__goautoneg",
    download = ":_github.com__prometheus__common#download",
    install = ["internal/bitbucket.org/ww/goautoneg"],
    module = "github.com/prometheus/common",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__prometheus__common__model",
    download = ":_github.com__prometheus__common#download",
    install = ["model"],
    module = "github.com/prometheus/common",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__prometheus__procfs",
    _tag = "download",
    module = "github.com/prometheus/procfs",
    version = "v0.6.0",
)

go_module(
    name = "github.com__prometheus__procfs",
    download = ":_github.com__prometheus__procfs#download",
    install = ["."],
    module = "github.com/prometheus/procfs",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__prometheus__procfs__internal__fs",
        ":github.com__prometheus__procfs__internal__util",
        ":golang.org__x__sys__unix",
    ],
)

go_module(
    name = "github.com__prometheus__procfs__internal__fs",
    download = ":_github.com__prometheus__procfs#download",
    install = ["internal/fs"],
    module = "github.com/prometheus/procfs",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__prometheus__procfs__internal__util",
    download = ":_github.com__prometheus__procfs#download",
    install = ["internal/util"],
    module = "github.com/prometheus/procfs",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__rivo__uniseg",
    _tag = "download",
    module = "github.com/rivo/uniseg",
    version = "v0.2.0",
)

go_module(
    name = "github.com__rivo__uniseg",
    download = ":_github.com__rivo__uniseg#download",
    install = ["."],
    module = "github.com/rivo/uniseg",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__russross__blackfriday__v2",
    _tag = "download",
    module = "github.com/russross/blackfriday/v2",
    version = "v2.1.0",
)

go_module(
    name = "github.com__russross__blackfriday__v2",
    download = ":_github.com__russross__blackfriday__v2#download",
    install = ["."],
    module = "github.com/russross/blackfriday/v2",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__sirupsen__logrus",
    _tag = "download",
    module = "github.com/sirupsen/logrus",
    version = "v1.8.1",
)

go_module(
    name = "github.com__sirupsen__logrus",
    download = ":_github.com__sirupsen__logrus#download",
    install = ["."],
    module = "github.com/sirupsen/logrus",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__sys__unix"],
)

go_mod_download(
    name = "github.com__spf13__afero",
    _tag = "download",
    module = "github.com/spf13/afero",
    version = "v1.6.0",
)

go_module(
    name = "github.com__spf13__afero",
    download = ":_github.com__spf13__afero#download",
    install = ["."],
    module = "github.com/spf13/afero",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__spf13__afero__mem",
        ":golang.org__x__text__transform",
        ":golang.org__x__text__unicode__norm",
    ],
)

go_module(
    name = "github.com__spf13__afero__mem",
    download = ":_github.com__spf13__afero#download",
    install = ["mem"],
    module = "github.com/spf13/afero",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__spf13__cast",
    _tag = "download",
    module = "github.com/spf13/cast",
    version = "v1.4.1",
)

go_module(
    name = "github.com__spf13__cast",
    download = ":_github.com__spf13__cast#download",
    install = ["."],
    module = "github.com/spf13/cast",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__spf13__cobra",
    _tag = "download",
    module = "github.com/spf13/cobra",
    version = "v1.3.0",
)

go_module(
    name = "github.com__spf13__cobra",
    download = ":_github.com__spf13__cobra#download",
    install = ["."],
    module = "github.com/spf13/cobra",
    visibility = ["PUBLIC"],
    deps = [":github.com__spf13__pflag"],
)

go_module(
    name = "github.com__spf13__cobra__doc",
    download = ":_github.com__spf13__cobra#download",
    install = ["doc"],
    module = "github.com/spf13/cobra",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__cpuguy83__go-md2man__v2__md2man",
        ":github.com__spf13__cobra",
        ":github.com__spf13__pflag",
        ":gopkg.in__yaml.v2",
    ],
)

go_mod_download(
    name = "github.com__spf13__jwalterweatherman",
    _tag = "download",
    module = "github.com/spf13/jwalterweatherman",
    version = "v1.1.0",
)

go_module(
    name = "github.com__spf13__jwalterweatherman",
    download = ":_github.com__spf13__jwalterweatherman#download",
    install = ["."],
    module = "github.com/spf13/jwalterweatherman",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__spf13__pflag",
    _tag = "download",
    module = "github.com/spf13/pflag",
    version = "v1.0.5",
)

go_module(
    name = "github.com__spf13__pflag",
    download = ":_github.com__spf13__pflag#download",
    install = ["."],
    module = "github.com/spf13/pflag",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__spf13__viper",
    _tag = "download",
    module = "github.com/spf13/viper",
    version = "v1.10.0",
)

go_module(
    name = "github.com__spf13__viper",
    download = ":_github.com__spf13__viper#download",
    install = ["."],
    module = "github.com/spf13/viper",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__fsnotify__fsnotify",
        ":github.com__magiconair__properties",
        ":github.com__mitchellh__mapstructure",
        ":github.com__spf13__afero",
        ":github.com__spf13__cast",
        ":github.com__spf13__jwalterweatherman",
        ":github.com__spf13__pflag",
        ":github.com__spf13__viper__internal__encoding",
        ":github.com__spf13__viper__internal__encoding__hcl",
        ":github.com__spf13__viper__internal__encoding__json",
        ":github.com__spf13__viper__internal__encoding__toml",
        ":github.com__spf13__viper__internal__encoding__yaml",
        ":github.com__subosito__gotenv",
        ":gopkg.in__ini.v1",
    ],
)

go_module(
    name = "github.com__spf13__viper__internal__encoding",
    download = ":_github.com__spf13__viper#download",
    install = ["internal/encoding"],
    module = "github.com/spf13/viper",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__spf13__viper__internal__encoding__hcl",
    download = ":_github.com__spf13__viper#download",
    install = ["internal/encoding/hcl"],
    module = "github.com/spf13/viper",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__hashicorp__hcl",
        ":github.com__hashicorp__hcl__hcl__printer",
    ],
)

go_module(
    name = "github.com__spf13__viper__internal__encoding__json",
    download = ":_github.com__spf13__viper#download",
    install = ["internal/encoding/json"],
    module = "github.com/spf13/viper",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__spf13__viper__internal__encoding__toml",
    download = ":_github.com__spf13__viper#download",
    install = ["internal/encoding/toml"],
    module = "github.com/spf13/viper",
    visibility = ["PUBLIC"],
    deps = [":github.com__pelletier__go-toml"],
)

go_module(
    name = "github.com__spf13__viper__internal__encoding__yaml",
    download = ":_github.com__spf13__viper#download",
    install = ["internal/encoding/yaml"],
    module = "github.com/spf13/viper",
    visibility = ["PUBLIC"],
    deps = [":gopkg.in__yaml.v2"],
)

go_mod_download(
    name = "github.com__stefanberger__go-pkcs11uri",
    _tag = "download",
    module = "github.com/stefanberger/go-pkcs11uri",
    version = "v0.0.0-20201008174630-78d3cae3a980",
)

go_module(
    name = "github.com__stefanberger__go-pkcs11uri",
    download = ":_github.com__stefanberger__go-pkcs11uri#download",
    install = ["."],
    module = "github.com/stefanberger/go-pkcs11uri",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__subosito__gotenv",
    _tag = "download",
    module = "github.com/subosito/gotenv",
    version = "v1.2.0",
)

go_module(
    name = "github.com__subosito__gotenv",
    download = ":_github.com__subosito__gotenv#download",
    install = ["."],
    module = "github.com/subosito/gotenv",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__syndtr__gocapability",
    _tag = "download",
    module = "github.com/syndtr/gocapability",
    version = "v0.0.0-20200815063812-42c35b437635",
)

go_module(
    name = "github.com__syndtr__gocapability__capability",
    download = ":_github.com__syndtr__gocapability#download",
    install = ["capability"],
    module = "github.com/syndtr/gocapability",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__tchap__go-patricia",
    _tag = "download",
    module = "github.com/tchap/go-patricia",
    version = "v2.3.0+incompatible",
)

go_module(
    name = "github.com__tchap__go-patricia__patricia",
    download = ":_github.com__tchap__go-patricia#download",
    install = ["patricia"],
    module = "github.com/tchap/go-patricia",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__ulikunitz__xz",
    _tag = "download",
    module = "github.com/ulikunitz/xz",
    version = "v0.5.10",
)

go_module(
    name = "github.com__ulikunitz__xz",
    download = ":_github.com__ulikunitz__xz#download",
    install = ["."],
    module = "github.com/ulikunitz/xz",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__ulikunitz__xz__internal__xlog",
        ":github.com__ulikunitz__xz__lzma",
    ],
)

go_module(
    name = "github.com__ulikunitz__xz__internal__hash",
    download = ":_github.com__ulikunitz__xz#download",
    install = ["internal/hash"],
    module = "github.com/ulikunitz/xz",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__ulikunitz__xz__internal__xlog",
    download = ":_github.com__ulikunitz__xz#download",
    install = ["internal/xlog"],
    module = "github.com/ulikunitz/xz",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__ulikunitz__xz__lzma",
    download = ":_github.com__ulikunitz__xz#download",
    install = ["lzma"],
    module = "github.com/ulikunitz/xz",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__ulikunitz__xz__internal__hash",
        ":github.com__ulikunitz__xz__internal__xlog",
    ],
)

go_mod_download(
    name = "github.com__vbatts__tar-split",
    _tag = "download",
    module = "github.com/vbatts/tar-split",
    version = "v0.11.2",
)

go_module(
    name = "github.com__vbatts__tar-split__archive__tar",
    download = ":_github.com__vbatts__tar-split#download",
    install = ["archive/tar"],
    module = "github.com/vbatts/tar-split",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "github.com__vbatts__tar-split__tar__asm",
    download = ":_github.com__vbatts__tar-split#download",
    install = ["tar/asm"],
    module = "github.com/vbatts/tar-split",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__vbatts__tar-split__archive__tar",
        ":github.com__vbatts__tar-split__tar__storage",
    ],
)

go_module(
    name = "github.com__vbatts__tar-split__tar__storage",
    download = ":_github.com__vbatts__tar-split#download",
    install = ["tar/storage"],
    module = "github.com/vbatts/tar-split",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__vbauerster__mpb__v7",
    _tag = "download",
    module = "github.com/vbauerster/mpb/v7",
    version = "v7.1.4",
)

go_module(
    name = "github.com__vbauerster__mpb__v7",
    download = ":_github.com__vbauerster__mpb__v7#download",
    install = ["."],
    module = "github.com/vbauerster/mpb/v7",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__acarl005__stripansi",
        ":github.com__mattn__go-runewidth",
        ":github.com__vbauerster__mpb__v7__cwriter",
        ":github.com__vbauerster__mpb__v7__decor",
        ":github.com__vbauerster__mpb__v7__internal",
    ],
)

go_module(
    name = "github.com__vbauerster__mpb__v7__cwriter",
    download = ":_github.com__vbauerster__mpb__v7#download",
    install = ["cwriter"],
    module = "github.com/vbauerster/mpb/v7",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__sys__unix"],
)

go_module(
    name = "github.com__vbauerster__mpb__v7__decor",
    download = ":_github.com__vbauerster__mpb__v7#download",
    install = ["decor"],
    module = "github.com/vbauerster/mpb/v7",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__VividCortex__ewma",
        ":github.com__acarl005__stripansi",
        ":github.com__mattn__go-runewidth",
        ":github.com__vbauerster__mpb__v7__internal",
    ],
)

go_module(
    name = "github.com__vbauerster__mpb__v7__internal",
    download = ":_github.com__vbauerster__mpb__v7#download",
    install = ["internal"],
    module = "github.com/vbauerster/mpb/v7",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__vishvananda__netlink",
    _tag = "download",
    module = "github.com/vishvananda/netlink",
    version = "v1.1.1-0.20210330154013-f5de75959ad5",
)

go_module(
    name = "github.com__vishvananda__netlink",
    download = ":_github.com__vishvananda__netlink#download",
    install = ["."],
    module = "github.com/vishvananda/netlink",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__vishvananda__netlink__nl",
        ":github.com__vishvananda__netns",
        ":golang.org__x__sys__unix",
    ],
)

go_module(
    name = "github.com__vishvananda__netlink__nl",
    download = ":_github.com__vishvananda__netlink#download",
    install = ["nl"],
    module = "github.com/vishvananda/netlink",
    visibility = ["PUBLIC"],
    deps = select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [
            ":github.com__vishvananda__netns",
            ":golang.org__x__sys__unix",
        ],
        "default": [],
    }),
)

go_mod_download(
    name = "github.com__vishvananda__netns",
    _tag = "download",
    module = "github.com/vishvananda/netns",
    version = "v0.0.0-20210104183010-2eb08e3e575f",
)

go_module(
    name = "github.com__vishvananda__netns",
    download = ":_github.com__vishvananda__netns#download",
    install = ["."],
    module = "github.com/vishvananda/netns",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__sys__unix"],
)

go_mod_download(
    name = "github.com__xeipuuv__gojsonpointer",
    _tag = "download",
    module = "github.com/xeipuuv/gojsonpointer",
    version = "v0.0.0-20190809123943-df4f5c81cb3b",
)

go_module(
    name = "github.com__xeipuuv__gojsonpointer",
    download = ":_github.com__xeipuuv__gojsonpointer#download",
    install = ["."],
    module = "github.com/xeipuuv/gojsonpointer",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "github.com__xeipuuv__gojsonreference",
    _tag = "download",
    module = "github.com/xeipuuv/gojsonreference",
    version = "v0.0.0-20180127040603-bd5ef7bd5415",
)

go_module(
    name = "github.com__xeipuuv__gojsonreference",
    download = ":_github.com__xeipuuv__gojsonreference#download",
    install = ["."],
    module = "github.com/xeipuuv/gojsonreference",
    visibility = ["PUBLIC"],
    deps = [":github.com__xeipuuv__gojsonpointer"],
)

go_mod_download(
    name = "github.com__xeipuuv__gojsonschema",
    _tag = "download",
    module = "github.com/xeipuuv/gojsonschema",
    version = "v1.2.0",
)

go_module(
    name = "github.com__xeipuuv__gojsonschema",
    download = ":_github.com__xeipuuv__gojsonschema#download",
    install = ["."],
    module = "github.com/xeipuuv/gojsonschema",
    visibility = ["PUBLIC"],
    deps = [":github.com__xeipuuv__gojsonreference"],
)

go_mod_download(
    name = "go.etcd.io__bbolt",
    _tag = "download",
    module = "go.etcd.io/bbolt",
    version = "v1.3.6",
)

go_module(
    name = "go.etcd.io__bbolt",
    download = ":_go.etcd.io__bbolt#download",
    install = ["."],
    module = "go.etcd.io/bbolt",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__sys__unix"],
)

go_mod_download(
    name = "go.mozilla.org__pkcs7",
    _tag = "download",
    module = "go.mozilla.org/pkcs7",
    version = "v0.0.0-20200128120323-432b2356ecb1",
)

go_module(
    name = "go.mozilla.org__pkcs7",
    download = ":_go.mozilla.org__pkcs7#download",
    install = ["."],
    module = "go.mozilla.org/pkcs7",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "golang.org__x__crypto",
    _tag = "download",
    module = "golang.org/x/crypto",
    version = "v0.0.0-20210817164053-32db794688a5",
)

go_module(
    name = "golang.org__x__crypto__blowfish",
    download = ":_golang.org__x__crypto#download",
    install = ["blowfish"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__crypto__cast5",
    download = ":_golang.org__x__crypto#download",
    install = ["cast5"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__crypto__chacha20",
    download = ":_golang.org__x__crypto#download",
    install = ["chacha20"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__crypto__internal__subtle"],
)

go_module(
    name = "golang.org__x__crypto__curve25519",
    download = ":_golang.org__x__crypto#download",
    install = ["curve25519"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__crypto__curve25519__internal__field"],
)

go_module(
    name = "golang.org__x__crypto__curve25519__internal__field",
    download = ":_golang.org__x__crypto#download",
    install = ["curve25519/internal/field"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__crypto__ed25519",
    download = ":_golang.org__x__crypto#download",
    install = ["ed25519"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__crypto__internal__subtle",
    download = ":_golang.org__x__crypto#download",
    install = ["internal/subtle"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__crypto__openpgp",
    download = ":_golang.org__x__crypto#download",
    install = ["openpgp"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__crypto__openpgp__armor",
        ":golang.org__x__crypto__openpgp__errors",
        ":golang.org__x__crypto__openpgp__packet",
        ":golang.org__x__crypto__openpgp__s2k",
    ],
)

go_module(
    name = "golang.org__x__crypto__openpgp__armor",
    download = ":_golang.org__x__crypto#download",
    install = ["openpgp/armor"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__crypto__openpgp__errors"],
)

go_module(
    name = "golang.org__x__crypto__openpgp__elgamal",
    download = ":_golang.org__x__crypto#download",
    install = ["openpgp/elgamal"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__crypto__openpgp__errors",
    download = ":_golang.org__x__crypto#download",
    install = ["openpgp/errors"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__crypto__openpgp__packet",
    download = ":_golang.org__x__crypto#download",
    install = ["openpgp/packet"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__crypto__cast5",
        ":golang.org__x__crypto__openpgp__elgamal",
        ":golang.org__x__crypto__openpgp__errors",
        ":golang.org__x__crypto__openpgp__s2k",
    ],
)

go_module(
    name = "golang.org__x__crypto__openpgp__s2k",
    download = ":_golang.org__x__crypto#download",
    install = ["openpgp/s2k"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__crypto__openpgp__errors"],
)

go_module(
    name = "golang.org__x__crypto__pbkdf2",
    download = ":_golang.org__x__crypto#download",
    install = ["pbkdf2"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__crypto__poly1305",
    download = ":_golang.org__x__crypto#download",
    install = ["poly1305"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__crypto__sha3",
    download = ":_golang.org__x__crypto#download",
    install = ["sha3"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__crypto__ssh",
    download = ":_golang.org__x__crypto#download",
    install = ["ssh"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__crypto__chacha20",
        ":golang.org__x__crypto__curve25519",
        ":golang.org__x__crypto__ed25519",
        ":golang.org__x__crypto__poly1305",
        ":golang.org__x__crypto__ssh__internal__bcrypt_pbkdf",
    ],
)

go_module(
    name = "golang.org__x__crypto__ssh__agent",
    download = ":_golang.org__x__crypto#download",
    install = ["ssh/agent"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__crypto__ed25519",
        ":golang.org__x__crypto__ssh",
    ],
)

go_module(
    name = "golang.org__x__crypto__ssh__internal__bcrypt_pbkdf",
    download = ":_golang.org__x__crypto#download",
    install = ["ssh/internal/bcrypt_pbkdf"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__crypto__blowfish"],
)

go_module(
    name = "golang.org__x__crypto__ssh__knownhosts",
    download = ":_golang.org__x__crypto#download",
    install = ["ssh/knownhosts"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__crypto__ssh"],
)

go_module(
    name = "golang.org__x__crypto__ssh__terminal",
    download = ":_golang.org__x__crypto#download",
    install = ["ssh/terminal"],
    module = "golang.org/x/crypto",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__term"],
)

go_mod_download(
    name = "golang.org__x__net",
    _tag = "download",
    module = "golang.org/x/net",
    version = "v0.0.0-20210813160813-60bc85c4be6d",
)

go_module(
    name = "golang.org__x__net__context",
    download = ":_golang.org__x__net#download",
    install = ["context"],
    module = "golang.org/x/net",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__net__html",
    download = ":_golang.org__x__net#download",
    install = ["html"],
    module = "golang.org/x/net",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__net__html__atom"],
)

go_module(
    name = "golang.org__x__net__html__atom",
    download = ":_golang.org__x__net#download",
    install = ["html/atom"],
    module = "golang.org/x/net",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__net__html__charset",
    download = ":_golang.org__x__net#download",
    install = ["html/charset"],
    module = "golang.org/x/net",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__net__html",
        ":golang.org__x__text__encoding",
        ":golang.org__x__text__encoding__charmap",
        ":golang.org__x__text__encoding__htmlindex",
        ":golang.org__x__text__transform",
    ],
)

go_module(
    name = "golang.org__x__net__http__httpguts",
    download = ":_golang.org__x__net#download",
    install = ["http/httpguts"],
    module = "golang.org/x/net",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__net__idna"],
)

go_module(
    name = "golang.org__x__net__http2",
    download = ":_golang.org__x__net#download",
    install = ["http2"],
    module = "golang.org/x/net",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__net__http2__hpack",
        ":golang.org__x__net__http__httpguts",
        ":golang.org__x__net__idna",
    ],
)

go_module(
    name = "golang.org__x__net__http2__hpack",
    download = ":_golang.org__x__net#download",
    install = ["http2/hpack"],
    module = "golang.org/x/net",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__net__idna",
    download = ":_golang.org__x__net#download",
    install = ["idna"],
    module = "golang.org/x/net",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__text__secure__bidirule",
        ":golang.org__x__text__unicode__bidi",
        ":golang.org__x__text__unicode__norm",
    ],
)

go_module(
    name = "golang.org__x__net__internal__socks",
    download = ":_golang.org__x__net#download",
    install = ["internal/socks"],
    module = "golang.org/x/net",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__net__internal__timeseries",
    download = ":_golang.org__x__net#download",
    install = ["internal/timeseries"],
    module = "golang.org/x/net",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__net__proxy",
    download = ":_golang.org__x__net#download",
    install = ["proxy"],
    module = "golang.org/x/net",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__net__internal__socks"],
)

go_module(
    name = "golang.org__x__net__trace",
    download = ":_golang.org__x__net#download",
    install = ["trace"],
    module = "golang.org/x/net",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__net__internal__timeseries"],
)

go_mod_download(
    name = "golang.org__x__sync",
    _tag = "download",
    module = "golang.org/x/sync",
    version = "v0.0.0-20210220032951-036812b2e83c",
)

go_module(
    name = "golang.org__x__sync__errgroup",
    download = ":_golang.org__x__sync#download",
    install = ["errgroup"],
    module = "golang.org/x/sync",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__sync__semaphore",
    download = ":_golang.org__x__sync#download",
    install = ["semaphore"],
    module = "golang.org/x/sync",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "golang.org__x__sys",
    _tag = "download",
    module = "golang.org/x/sys",
    version = "v0.0.0-20211216021012-1d35b9e2eb4e",
)

go_module(
    name = "golang.org__x__sys__execabs",
    download = ":_golang.org__x__sys#download",
    install = ["execabs"],
    module = "golang.org/x/sys",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__sys__internal__unsafeheader",
    download = ":_golang.org__x__sys#download",
    install = ["internal/unsafeheader"],
    module = "golang.org/x/sys",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__sys__unix",
    download = ":_golang.org__x__sys#download",
    install = ["unix"],
    module = "golang.org/x/sys",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__sys__internal__unsafeheader"],
)

go_mod_download(
    name = "golang.org__x__term",
    _tag = "download",
    module = "golang.org/x/term",
    version = "v0.0.0-20210927222741-03fcf44c2211",
)

go_module(
    name = "golang.org__x__term",
    download = ":_golang.org__x__term#download",
    install = ["."],
    module = "golang.org/x/term",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__sys__unix"],
)

go_mod_download(
    name = "golang.org__x__text",
    _tag = "download",
    module = "golang.org/x/text",
    version = "v0.3.7",
)

go_module(
    name = "golang.org__x__text__encoding",
    download = ":_golang.org__x__text#download",
    install = ["encoding"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__text__encoding__internal__identifier",
        ":golang.org__x__text__transform",
    ],
)

go_module(
    name = "golang.org__x__text__encoding__charmap",
    download = ":_golang.org__x__text#download",
    install = ["encoding/charmap"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__text__encoding",
        ":golang.org__x__text__encoding__internal",
        ":golang.org__x__text__encoding__internal__identifier",
        ":golang.org__x__text__transform",
    ],
)

go_module(
    name = "golang.org__x__text__encoding__htmlindex",
    download = ":_golang.org__x__text#download",
    install = ["encoding/htmlindex"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__text__encoding",
        ":golang.org__x__text__encoding__charmap",
        ":golang.org__x__text__encoding__internal__identifier",
        ":golang.org__x__text__encoding__japanese",
        ":golang.org__x__text__encoding__korean",
        ":golang.org__x__text__encoding__simplifiedchinese",
        ":golang.org__x__text__encoding__traditionalchinese",
        ":golang.org__x__text__encoding__unicode",
        ":golang.org__x__text__language",
    ],
)

go_module(
    name = "golang.org__x__text__encoding__internal",
    download = ":_golang.org__x__text#download",
    install = ["encoding/internal"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__text__encoding",
        ":golang.org__x__text__encoding__internal__identifier",
        ":golang.org__x__text__transform",
    ],
)

go_module(
    name = "golang.org__x__text__encoding__internal__identifier",
    download = ":_golang.org__x__text#download",
    install = ["encoding/internal/identifier"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__text__encoding__japanese",
    download = ":_golang.org__x__text#download",
    install = ["encoding/japanese"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__text__encoding",
        ":golang.org__x__text__encoding__internal",
        ":golang.org__x__text__encoding__internal__identifier",
        ":golang.org__x__text__transform",
    ],
)

go_module(
    name = "golang.org__x__text__encoding__korean",
    download = ":_golang.org__x__text#download",
    install = ["encoding/korean"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__text__encoding",
        ":golang.org__x__text__encoding__internal",
        ":golang.org__x__text__encoding__internal__identifier",
        ":golang.org__x__text__transform",
    ],
)

go_module(
    name = "golang.org__x__text__encoding__simplifiedchinese",
    download = ":_golang.org__x__text#download",
    install = ["encoding/simplifiedchinese"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__text__encoding",
        ":golang.org__x__text__encoding__internal",
        ":golang.org__x__text__encoding__internal__identifier",
        ":golang.org__x__text__transform",
    ],
)

go_module(
    name = "golang.org__x__text__encoding__traditionalchinese",
    download = ":_golang.org__x__text#download",
    install = ["encoding/traditionalchinese"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__text__encoding",
        ":golang.org__x__text__encoding__internal",
        ":golang.org__x__text__encoding__internal__identifier",
        ":golang.org__x__text__transform",
    ],
)

go_module(
    name = "golang.org__x__text__encoding__unicode",
    download = ":_golang.org__x__text#download",
    install = ["encoding/unicode"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__text__encoding",
        ":golang.org__x__text__encoding__internal",
        ":golang.org__x__text__encoding__internal__identifier",
        ":golang.org__x__text__internal__utf8internal",
        ":golang.org__x__text__runes",
        ":golang.org__x__text__transform",
    ],
)

go_module(
    name = "golang.org__x__text__internal__language",
    download = ":_golang.org__x__text#download",
    install = ["internal/language"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__text__internal__tag"],
)

go_module(
    name = "golang.org__x__text__internal__language__compact",
    download = ":_golang.org__x__text#download",
    install = ["internal/language/compact"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__text__internal__language"],
)

go_module(
    name = "golang.org__x__text__internal__tag",
    download = ":_golang.org__x__text#download",
    install = ["internal/tag"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__text__internal__utf8internal",
    download = ":_golang.org__x__text#download",
    install = ["internal/utf8internal"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__text__language",
    download = ":_golang.org__x__text#download",
    install = ["language"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__text__internal__language",
        ":golang.org__x__text__internal__language__compact",
    ],
)

go_module(
    name = "golang.org__x__text__runes",
    download = ":_golang.org__x__text#download",
    install = ["runes"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__text__transform"],
)

go_module(
    name = "golang.org__x__text__secure__bidirule",
    download = ":_golang.org__x__text#download",
    install = ["secure/bidirule"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__text__transform",
        ":golang.org__x__text__unicode__bidi",
    ],
)

go_module(
    name = "golang.org__x__text__transform",
    download = ":_golang.org__x__text#download",
    install = ["transform"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__text__unicode__bidi",
    download = ":_golang.org__x__text#download",
    install = ["unicode/bidi"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "golang.org__x__text__unicode__norm",
    download = ":_golang.org__x__text#download",
    install = ["unicode/norm"],
    module = "golang.org/x/text",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__text__transform"],
)

go_mod_download(
    name = "google.golang.org__genproto",
    _tag = "download",
    module = "google.golang.org/genproto",
    version = "v0.0.0-20211208223120-3a66f561d7aa",
)

go_module(
    name = "google.golang.org__genproto__googleapis__rpc__status",
    download = ":_google.golang.org__genproto#download",
    install = ["googleapis/rpc/status"],
    module = "google.golang.org/genproto",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__reflect__protoreflect",
        ":google.golang.org__protobuf__runtime__protoimpl",
        ":google.golang.org__protobuf__types__known__anypb",
    ],
)

go_mod_download(
    name = "google.golang.org__grpc",
    _tag = "download",
    module = "google.golang.org/grpc",
    version = "v1.42.0",
)

go_module(
    name = "google.golang.org__grpc",
    download = ":_google.golang.org__grpc#download",
    install = ["."],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__net__trace",
        ":google.golang.org__grpc__backoff",
        ":google.golang.org__grpc__balancer",
        ":google.golang.org__grpc__balancer__base",
        ":google.golang.org__grpc__balancer__roundrobin",
        ":google.golang.org__grpc__codes",
        ":google.golang.org__grpc__connectivity",
        ":google.golang.org__grpc__credentials",
        ":google.golang.org__grpc__encoding",
        ":google.golang.org__grpc__encoding__proto",
        ":google.golang.org__grpc__grpclog",
        ":google.golang.org__grpc__internal",
        ":google.golang.org__grpc__internal__backoff",
        ":google.golang.org__grpc__internal__balancerload",
        ":google.golang.org__grpc__internal__binarylog",
        ":google.golang.org__grpc__internal__buffer",
        ":google.golang.org__grpc__internal__channelz",
        ":google.golang.org__grpc__internal__envconfig",
        ":google.golang.org__grpc__internal__grpcrand",
        ":google.golang.org__grpc__internal__grpcsync",
        ":google.golang.org__grpc__internal__grpcutil",
        ":google.golang.org__grpc__internal__resolver",
        ":google.golang.org__grpc__internal__resolver__dns",
        ":google.golang.org__grpc__internal__resolver__passthrough",
        ":google.golang.org__grpc__internal__resolver__unix",
        ":google.golang.org__grpc__internal__serviceconfig",
        ":google.golang.org__grpc__internal__transport",
        ":google.golang.org__grpc__keepalive",
        ":google.golang.org__grpc__metadata",
        ":google.golang.org__grpc__peer",
        ":google.golang.org__grpc__resolver",
        ":google.golang.org__grpc__serviceconfig",
        ":google.golang.org__grpc__stats",
        ":google.golang.org__grpc__status",
        ":google.golang.org__grpc__tap",
    ],
)

go_module(
    name = "google.golang.org__grpc__attributes",
    download = ":_google.golang.org__grpc#download",
    install = ["attributes"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "google.golang.org__grpc__backoff",
    download = ":_google.golang.org__grpc#download",
    install = ["backoff"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "google.golang.org__grpc__balancer",
    download = ":_google.golang.org__grpc#download",
    install = ["balancer"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__grpc__connectivity",
        ":google.golang.org__grpc__credentials",
        ":google.golang.org__grpc__internal",
        ":google.golang.org__grpc__metadata",
        ":google.golang.org__grpc__resolver",
        ":google.golang.org__grpc__serviceconfig",
    ],
)

go_module(
    name = "google.golang.org__grpc__balancer__base",
    download = ":_google.golang.org__grpc#download",
    install = ["balancer/base"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__grpc__balancer",
        ":google.golang.org__grpc__connectivity",
        ":google.golang.org__grpc__grpclog",
        ":google.golang.org__grpc__resolver",
    ],
)

go_module(
    name = "google.golang.org__grpc__balancer__grpclb__state",
    download = ":_google.golang.org__grpc#download",
    install = ["balancer/grpclb/state"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [":google.golang.org__grpc__resolver"],
)

go_module(
    name = "google.golang.org__grpc__balancer__roundrobin",
    download = ":_google.golang.org__grpc#download",
    install = ["balancer/roundrobin"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__grpc__balancer",
        ":google.golang.org__grpc__balancer__base",
        ":google.golang.org__grpc__grpclog",
        ":google.golang.org__grpc__internal__grpcrand",
    ],
)

go_module(
    name = "google.golang.org__grpc__binarylog__grpc_binarylog_v1",
    download = ":_google.golang.org__grpc#download",
    install = ["binarylog/grpc_binarylog_v1"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__golang__protobuf__proto",
        ":google.golang.org__protobuf__reflect__protoreflect",
        ":google.golang.org__protobuf__runtime__protoimpl",
        ":google.golang.org__protobuf__types__known__durationpb",
        ":google.golang.org__protobuf__types__known__timestamppb",
    ],
)

go_module(
    name = "google.golang.org__grpc__codes",
    download = ":_google.golang.org__grpc#download",
    install = ["codes"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "google.golang.org__grpc__connectivity",
    download = ":_google.golang.org__grpc#download",
    install = ["connectivity"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [":google.golang.org__grpc__grpclog"],
)

go_module(
    name = "google.golang.org__grpc__credentials",
    download = ":_google.golang.org__grpc#download",
    install = ["credentials"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__golang__protobuf__proto",
        ":google.golang.org__grpc__attributes",
        ":google.golang.org__grpc__internal__credentials",
    ],
)

go_module(
    name = "google.golang.org__grpc__encoding",
    download = ":_google.golang.org__grpc#download",
    install = ["encoding"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "google.golang.org__grpc__encoding__proto",
    download = ":_google.golang.org__grpc#download",
    install = ["encoding/proto"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__golang__protobuf__proto",
        ":google.golang.org__grpc__encoding",
    ],
)

go_module(
    name = "google.golang.org__grpc__grpclog",
    download = ":_google.golang.org__grpc#download",
    install = ["grpclog"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [":google.golang.org__grpc__internal__grpclog"],
)

go_module(
    name = "google.golang.org__grpc__internal",
    download = ":_google.golang.org__grpc#download",
    install = ["internal"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__grpc__attributes",
        ":google.golang.org__grpc__connectivity",
        ":google.golang.org__grpc__resolver",
        ":google.golang.org__grpc__serviceconfig",
    ],
)

go_module(
    name = "google.golang.org__grpc__internal__backoff",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/backoff"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__grpc__backoff",
        ":google.golang.org__grpc__internal__grpcrand",
    ],
)

go_module(
    name = "google.golang.org__grpc__internal__balancerload",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/balancerload"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [":google.golang.org__grpc__metadata"],
)

go_module(
    name = "google.golang.org__grpc__internal__binarylog",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/binarylog"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__golang__protobuf__proto",
        ":github.com__golang__protobuf__ptypes",
        ":google.golang.org__grpc__binarylog__grpc_binarylog_v1",
        ":google.golang.org__grpc__grpclog",
        ":google.golang.org__grpc__internal__grpcutil",
        ":google.golang.org__grpc__metadata",
        ":google.golang.org__grpc__status",
    ],
)

go_module(
    name = "google.golang.org__grpc__internal__buffer",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/buffer"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "google.golang.org__grpc__internal__channelz",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/channelz"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__grpc__connectivity",
        ":google.golang.org__grpc__credentials",
        ":google.golang.org__grpc__grpclog",
    ] + select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [":golang.org__x__sys__unix"],
        "default": [],
    }),
)

go_module(
    name = "google.golang.org__grpc__internal__credentials",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/credentials"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [":google.golang.org__grpc__grpclog"],
)

go_module(
    name = "google.golang.org__grpc__internal__envconfig",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/envconfig"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [":google.golang.org__grpc__internal__xds__env"],
)

go_module(
    name = "google.golang.org__grpc__internal__grpclog",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/grpclog"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "google.golang.org__grpc__internal__grpcrand",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/grpcrand"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "google.golang.org__grpc__internal__grpcsync",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/grpcsync"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "google.golang.org__grpc__internal__grpcutil",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/grpcutil"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [":google.golang.org__grpc__metadata"],
)

go_module(
    name = "google.golang.org__grpc__internal__metadata",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/metadata"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__grpc__metadata",
        ":google.golang.org__grpc__resolver",
    ],
)

go_module(
    name = "google.golang.org__grpc__internal__resolver",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/resolver"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__grpc__internal__serviceconfig",
        ":google.golang.org__grpc__metadata",
        ":google.golang.org__grpc__resolver",
    ],
)

go_module(
    name = "google.golang.org__grpc__internal__resolver__dns",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/resolver/dns"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__grpc__balancer__grpclb__state",
        ":google.golang.org__grpc__grpclog",
        ":google.golang.org__grpc__internal__backoff",
        ":google.golang.org__grpc__internal__envconfig",
        ":google.golang.org__grpc__internal__grpcrand",
        ":google.golang.org__grpc__resolver",
        ":google.golang.org__grpc__serviceconfig",
    ],
)

go_module(
    name = "google.golang.org__grpc__internal__resolver__passthrough",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/resolver/passthrough"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [":google.golang.org__grpc__resolver"],
)

go_module(
    name = "google.golang.org__grpc__internal__resolver__unix",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/resolver/unix"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__grpc__internal__transport__networktype",
        ":google.golang.org__grpc__resolver",
    ],
)

go_module(
    name = "google.golang.org__grpc__internal__serviceconfig",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/serviceconfig"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__grpc__balancer",
        ":google.golang.org__grpc__codes",
        ":google.golang.org__grpc__grpclog",
        ":google.golang.org__grpc__serviceconfig",
    ],
)

go_module(
    name = "google.golang.org__grpc__internal__status",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/status"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__golang__protobuf__proto",
        ":github.com__golang__protobuf__ptypes",
        ":google.golang.org__genproto__googleapis__rpc__status",
        ":google.golang.org__grpc__codes",
    ],
)

go_module(
    name = "google.golang.org__grpc__internal__syscall",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/syscall"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [":google.golang.org__grpc__grpclog"] + select({
        ":__config_darwin_amd64": [],
        ":__config_linux_amd64": [":golang.org__x__sys__unix"],
        "default": [],
    }),
)

go_module(
    name = "google.golang.org__grpc__internal__transport",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/transport"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":github.com__golang__protobuf__proto",
        ":golang.org__x__net__http2",
        ":golang.org__x__net__http2__hpack",
        ":google.golang.org__genproto__googleapis__rpc__status",
        ":google.golang.org__grpc__codes",
        ":google.golang.org__grpc__credentials",
        ":google.golang.org__grpc__grpclog",
        ":google.golang.org__grpc__internal__channelz",
        ":google.golang.org__grpc__internal__credentials",
        ":google.golang.org__grpc__internal__grpcrand",
        ":google.golang.org__grpc__internal__grpcutil",
        ":google.golang.org__grpc__internal__metadata",
        ":google.golang.org__grpc__internal__syscall",
        ":google.golang.org__grpc__internal__transport__networktype",
        ":google.golang.org__grpc__keepalive",
        ":google.golang.org__grpc__metadata",
        ":google.golang.org__grpc__peer",
        ":google.golang.org__grpc__resolver",
        ":google.golang.org__grpc__stats",
        ":google.golang.org__grpc__status",
        ":google.golang.org__grpc__tap",
    ],
)

go_module(
    name = "google.golang.org__grpc__internal__transport__networktype",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/transport/networktype"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [":google.golang.org__grpc__resolver"],
)

go_module(
    name = "google.golang.org__grpc__internal__xds__env",
    download = ":_google.golang.org__grpc#download",
    install = ["internal/xds/env"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "google.golang.org__grpc__keepalive",
    download = ":_google.golang.org__grpc#download",
    install = ["keepalive"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "google.golang.org__grpc__metadata",
    download = ":_google.golang.org__grpc#download",
    install = ["metadata"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "google.golang.org__grpc__peer",
    download = ":_google.golang.org__grpc#download",
    install = ["peer"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [":google.golang.org__grpc__credentials"],
)

go_module(
    name = "google.golang.org__grpc__resolver",
    download = ":_google.golang.org__grpc#download",
    install = ["resolver"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__grpc__attributes",
        ":google.golang.org__grpc__credentials",
        ":google.golang.org__grpc__serviceconfig",
    ],
)

go_module(
    name = "google.golang.org__grpc__serviceconfig",
    download = ":_google.golang.org__grpc#download",
    install = ["serviceconfig"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "google.golang.org__grpc__stats",
    download = ":_google.golang.org__grpc#download",
    install = ["stats"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [":google.golang.org__grpc__metadata"],
)

go_module(
    name = "google.golang.org__grpc__status",
    download = ":_google.golang.org__grpc#download",
    install = ["status"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__genproto__googleapis__rpc__status",
        ":google.golang.org__grpc__codes",
        ":google.golang.org__grpc__internal__status",
    ],
)

go_module(
    name = "google.golang.org__grpc__tap",
    download = ":_google.golang.org__grpc#download",
    install = ["tap"],
    module = "google.golang.org/grpc",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "google.golang.org__protobuf",
    _tag = "download",
    module = "google.golang.org/protobuf",
    version = "v1.27.1",
)

go_module(
    name = "google.golang.org__protobuf__encoding__prototext",
    download = ":_google.golang.org__protobuf#download",
    install = ["encoding/prototext"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__encoding__protowire",
        ":google.golang.org__protobuf__internal__encoding__messageset",
        ":google.golang.org__protobuf__internal__encoding__text",
        ":google.golang.org__protobuf__internal__errors",
        ":google.golang.org__protobuf__internal__flags",
        ":google.golang.org__protobuf__internal__genid",
        ":google.golang.org__protobuf__internal__order",
        ":google.golang.org__protobuf__internal__pragma",
        ":google.golang.org__protobuf__internal__set",
        ":google.golang.org__protobuf__internal__strs",
        ":google.golang.org__protobuf__proto",
        ":google.golang.org__protobuf__reflect__protoreflect",
        ":google.golang.org__protobuf__reflect__protoregistry",
    ],
)

go_module(
    name = "google.golang.org__protobuf__encoding__protowire",
    download = ":_google.golang.org__protobuf#download",
    install = ["encoding/protowire"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [":google.golang.org__protobuf__internal__errors"],
)

go_module(
    name = "google.golang.org__protobuf__internal__descfmt",
    download = ":_google.golang.org__protobuf#download",
    install = ["internal/descfmt"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__internal__detrand",
        ":google.golang.org__protobuf__internal__pragma",
        ":google.golang.org__protobuf__reflect__protoreflect",
    ],
)

go_module(
    name = "google.golang.org__protobuf__internal__descopts",
    download = ":_google.golang.org__protobuf#download",
    install = ["internal/descopts"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [":google.golang.org__protobuf__reflect__protoreflect"],
)

go_module(
    name = "google.golang.org__protobuf__internal__detrand",
    download = ":_google.golang.org__protobuf#download",
    install = ["internal/detrand"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "google.golang.org__protobuf__internal__encoding__defval",
    download = ":_google.golang.org__protobuf#download",
    install = ["internal/encoding/defval"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__internal__encoding__text",
        ":google.golang.org__protobuf__internal__errors",
        ":google.golang.org__protobuf__reflect__protoreflect",
    ],
)

go_module(
    name = "google.golang.org__protobuf__internal__encoding__messageset",
    download = ":_google.golang.org__protobuf#download",
    install = ["internal/encoding/messageset"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__encoding__protowire",
        ":google.golang.org__protobuf__internal__errors",
        ":google.golang.org__protobuf__reflect__protoreflect",
    ],
)

go_module(
    name = "google.golang.org__protobuf__internal__encoding__tag",
    download = ":_google.golang.org__protobuf#download",
    install = ["internal/encoding/tag"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__internal__encoding__defval",
        ":google.golang.org__protobuf__internal__filedesc",
        ":google.golang.org__protobuf__internal__strs",
        ":google.golang.org__protobuf__reflect__protoreflect",
    ],
)

go_module(
    name = "google.golang.org__protobuf__internal__encoding__text",
    download = ":_google.golang.org__protobuf#download",
    install = ["internal/encoding/text"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__internal__detrand",
        ":google.golang.org__protobuf__internal__errors",
        ":google.golang.org__protobuf__internal__flags",
        ":google.golang.org__protobuf__internal__strs",
    ],
)

go_module(
    name = "google.golang.org__protobuf__internal__errors",
    download = ":_google.golang.org__protobuf#download",
    install = ["internal/errors"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [":google.golang.org__protobuf__internal__detrand"],
)

go_module(
    name = "google.golang.org__protobuf__internal__filedesc",
    download = ":_google.golang.org__protobuf#download",
    install = ["internal/filedesc"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__encoding__protowire",
        ":google.golang.org__protobuf__internal__descfmt",
        ":google.golang.org__protobuf__internal__descopts",
        ":google.golang.org__protobuf__internal__encoding__defval",
        ":google.golang.org__protobuf__internal__encoding__messageset",
        ":google.golang.org__protobuf__internal__errors",
        ":google.golang.org__protobuf__internal__genid",
        ":google.golang.org__protobuf__internal__pragma",
        ":google.golang.org__protobuf__internal__strs",
        ":google.golang.org__protobuf__proto",
        ":google.golang.org__protobuf__reflect__protoreflect",
        ":google.golang.org__protobuf__reflect__protoregistry",
    ],
)

go_module(
    name = "google.golang.org__protobuf__internal__filetype",
    download = ":_google.golang.org__protobuf#download",
    install = ["internal/filetype"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__internal__descopts",
        ":google.golang.org__protobuf__internal__filedesc",
        ":google.golang.org__protobuf__internal__impl",
        ":google.golang.org__protobuf__reflect__protoreflect",
        ":google.golang.org__protobuf__reflect__protoregistry",
    ],
)

go_module(
    name = "google.golang.org__protobuf__internal__flags",
    download = ":_google.golang.org__protobuf#download",
    install = ["internal/flags"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "google.golang.org__protobuf__internal__genid",
    download = ":_google.golang.org__protobuf#download",
    install = ["internal/genid"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [":google.golang.org__protobuf__reflect__protoreflect"],
)

go_module(
    name = "google.golang.org__protobuf__internal__impl",
    download = ":_google.golang.org__protobuf#download",
    install = ["internal/impl"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__encoding__prototext",
        ":google.golang.org__protobuf__encoding__protowire",
        ":google.golang.org__protobuf__internal__descopts",
        ":google.golang.org__protobuf__internal__detrand",
        ":google.golang.org__protobuf__internal__encoding__messageset",
        ":google.golang.org__protobuf__internal__encoding__tag",
        ":google.golang.org__protobuf__internal__errors",
        ":google.golang.org__protobuf__internal__filedesc",
        ":google.golang.org__protobuf__internal__flags",
        ":google.golang.org__protobuf__internal__genid",
        ":google.golang.org__protobuf__internal__order",
        ":google.golang.org__protobuf__internal__pragma",
        ":google.golang.org__protobuf__internal__strs",
        ":google.golang.org__protobuf__proto",
        ":google.golang.org__protobuf__reflect__protoreflect",
        ":google.golang.org__protobuf__reflect__protoregistry",
        ":google.golang.org__protobuf__runtime__protoiface",
    ],
)

go_module(
    name = "google.golang.org__protobuf__internal__order",
    download = ":_google.golang.org__protobuf#download",
    install = ["internal/order"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [":google.golang.org__protobuf__reflect__protoreflect"],
)

go_module(
    name = "google.golang.org__protobuf__internal__pragma",
    download = ":_google.golang.org__protobuf#download",
    install = ["internal/pragma"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "google.golang.org__protobuf__internal__set",
    download = ":_google.golang.org__protobuf#download",
    install = ["internal/set"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "google.golang.org__protobuf__internal__strs",
    download = ":_google.golang.org__protobuf#download",
    install = ["internal/strs"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__internal__flags",
        ":google.golang.org__protobuf__reflect__protoreflect",
    ],
)

go_module(
    name = "google.golang.org__protobuf__internal__version",
    download = ":_google.golang.org__protobuf#download",
    install = ["internal/version"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "google.golang.org__protobuf__proto",
    download = ":_google.golang.org__protobuf#download",
    install = ["proto"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__encoding__protowire",
        ":google.golang.org__protobuf__internal__encoding__messageset",
        ":google.golang.org__protobuf__internal__errors",
        ":google.golang.org__protobuf__internal__flags",
        ":google.golang.org__protobuf__internal__genid",
        ":google.golang.org__protobuf__internal__order",
        ":google.golang.org__protobuf__internal__pragma",
        ":google.golang.org__protobuf__internal__strs",
        ":google.golang.org__protobuf__reflect__protoreflect",
        ":google.golang.org__protobuf__reflect__protoregistry",
        ":google.golang.org__protobuf__runtime__protoiface",
    ],
)

go_module(
    name = "google.golang.org__protobuf__reflect__protodesc",
    download = ":_google.golang.org__protobuf#download",
    install = ["reflect/protodesc"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__encoding__protowire",
        ":google.golang.org__protobuf__internal__encoding__defval",
        ":google.golang.org__protobuf__internal__errors",
        ":google.golang.org__protobuf__internal__filedesc",
        ":google.golang.org__protobuf__internal__flags",
        ":google.golang.org__protobuf__internal__genid",
        ":google.golang.org__protobuf__internal__pragma",
        ":google.golang.org__protobuf__internal__strs",
        ":google.golang.org__protobuf__proto",
        ":google.golang.org__protobuf__reflect__protoreflect",
        ":google.golang.org__protobuf__reflect__protoregistry",
        ":google.golang.org__protobuf__types__descriptorpb",
    ],
)

go_module(
    name = "google.golang.org__protobuf__reflect__protoreflect",
    download = ":_google.golang.org__protobuf#download",
    install = ["reflect/protoreflect"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__encoding__protowire",
        ":google.golang.org__protobuf__internal__pragma",
    ],
)

go_module(
    name = "google.golang.org__protobuf__reflect__protoregistry",
    download = ":_google.golang.org__protobuf#download",
    install = ["reflect/protoregistry"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__internal__encoding__messageset",
        ":google.golang.org__protobuf__internal__errors",
        ":google.golang.org__protobuf__internal__flags",
        ":google.golang.org__protobuf__reflect__protoreflect",
    ],
)

go_module(
    name = "google.golang.org__protobuf__runtime__protoiface",
    download = ":_google.golang.org__protobuf#download",
    install = ["runtime/protoiface"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__internal__pragma",
        ":google.golang.org__protobuf__reflect__protoreflect",
    ],
)

go_module(
    name = "google.golang.org__protobuf__runtime__protoimpl",
    download = ":_google.golang.org__protobuf#download",
    install = ["runtime/protoimpl"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__internal__filedesc",
        ":google.golang.org__protobuf__internal__filetype",
        ":google.golang.org__protobuf__internal__impl",
        ":google.golang.org__protobuf__internal__version",
    ],
)

go_module(
    name = "google.golang.org__protobuf__types__descriptorpb",
    download = ":_google.golang.org__protobuf#download",
    install = ["types/descriptorpb"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__reflect__protoreflect",
        ":google.golang.org__protobuf__runtime__protoimpl",
    ],
)

go_module(
    name = "google.golang.org__protobuf__types__known__anypb",
    download = ":_google.golang.org__protobuf#download",
    install = ["types/known/anypb"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__proto",
        ":google.golang.org__protobuf__reflect__protoreflect",
        ":google.golang.org__protobuf__reflect__protoregistry",
        ":google.golang.org__protobuf__runtime__protoimpl",
    ],
)

go_module(
    name = "google.golang.org__protobuf__types__known__durationpb",
    download = ":_google.golang.org__protobuf#download",
    install = ["types/known/durationpb"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__reflect__protoreflect",
        ":google.golang.org__protobuf__runtime__protoimpl",
    ],
)

go_module(
    name = "google.golang.org__protobuf__types__known__timestamppb",
    download = ":_google.golang.org__protobuf#download",
    install = ["types/known/timestamppb"],
    module = "google.golang.org/protobuf",
    visibility = ["PUBLIC"],
    deps = [
        ":google.golang.org__protobuf__reflect__protoreflect",
        ":google.golang.org__protobuf__runtime__protoimpl",
    ],
)

go_mod_download(
    name = "gopkg.in__fsnotify.v1",
    _tag = "download",
    module = "gopkg.in/fsnotify.v1",
    version = "v1.4.7",
)

go_module(
    name = "gopkg.in__fsnotify.v1",
    download = ":_gopkg.in__fsnotify.v1#download",
    install = ["."],
    module = "gopkg.in/fsnotify.v1",
    visibility = ["PUBLIC"],
    deps = [":golang.org__x__sys__unix"],
)

go_mod_download(
    name = "gopkg.in__ini.v1",
    _tag = "download",
    module = "gopkg.in/ini.v1",
    version = "v1.66.2",
)

go_module(
    name = "gopkg.in__ini.v1",
    download = ":_gopkg.in__ini.v1#download",
    install = ["."],
    module = "gopkg.in/ini.v1",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "gopkg.in__square__go-jose.v2",
    _tag = "download",
    module = "gopkg.in/square/go-jose.v2",
    version = "v2.5.1",
)

go_module(
    name = "gopkg.in__square__go-jose.v2",
    download = ":_gopkg.in__square__go-jose.v2#download",
    install = ["."],
    module = "gopkg.in/square/go-jose.v2",
    visibility = ["PUBLIC"],
    deps = [
        ":golang.org__x__crypto__ed25519",
        ":golang.org__x__crypto__pbkdf2",
        ":gopkg.in__square__go-jose.v2__cipher",
        ":gopkg.in__square__go-jose.v2__json",
    ],
)

go_module(
    name = "gopkg.in__square__go-jose.v2__cipher",
    download = ":_gopkg.in__square__go-jose.v2#download",
    install = ["cipher"],
    module = "gopkg.in/square/go-jose.v2",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "gopkg.in__square__go-jose.v2__json",
    download = ":_gopkg.in__square__go-jose.v2#download",
    install = ["json"],
    module = "gopkg.in/square/go-jose.v2",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "gopkg.in__tomb.v1",
    _tag = "download",
    module = "gopkg.in/tomb.v1",
    version = "v1.0.0-20141024135613-dd632973f1e7",
)

go_module(
    name = "gopkg.in__tomb.v1",
    download = ":_gopkg.in__tomb.v1#download",
    install = ["."],
    module = "gopkg.in/tomb.v1",
    visibility = ["PUBLIC"],
    deps = [],
)

go_mod_download(
    name = "gopkg.in__yaml.v2",
    _tag = "download",
    module = "gopkg.in/yaml.v2",
    version = "v2.4.0",
)

go_module(
    name = "gopkg.in__yaml.v2",
    download = ":_gopkg.in__yaml.v2#download",
    install = ["."],
    module = "gopkg.in/yaml.v2",
    visibility = ["PUBLIC"],
    deps = [],
)

go_module(
    name = "shellescape",
    module = "github.com/alessio/shellescape",
    version = "v1.4.1",
    visibility = ["PUBLIC"],
)

go_module(
    name = "reflectwalk",
    module = "github.com/mitchellh/reflectwalk",
    version = "v1.0.2",
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
        "providers/file",
        "providers/confmap",
        "parsers/yaml",
    ],
    module = "github.com/knadh/koanf",
    version = "v1.4.0",
    visibility = ["PUBLIC"],
    deps = [
        ":copystructure",
        ":github.com__mitchellh__mapstructure",
        ":github.com__fsnotify__fsnotify",
        ":gopkg.in__yaml.v2",
    ],
)
