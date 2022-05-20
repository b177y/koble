# antora
docker = "podman"

# output module directory for antora
genrule(
    name = "adoc-manpages",
    srcs = ["//cmd/kob/generate:manpages"],
    outs = ["MAN-MODULE"],
    cmd = [
        # convert
        "mv cmd/kob/generate/out out",
        "mkdir -p MAN-MODULE/pages",
        "export HOME=\"/home/billy\"",
        docker + " run --rm -v ./out:/documents/ docker.io/asciidoctor/docker-asciidoctor find /documents -name '*.md' -exec sh -c 'kramdoc --heading-offset=-1 {}' {} \\;",
        # nav.adoc
        "echo .Manpages > MAN-MODULE/nav.adoc",
        "find ./out -name '*.adoc' -execdir echo \"* xref:{}[]\" >> navtmp.adoc \\;",
        "sort navtmp.adoc >> MAN-MODULE/nav.adoc",
        "cp ./out/*.adoc MAN-MODULE/pages",
    ],
)

genrule(
    name = "site",
    srcs = [
        "./docs/modules/",
        "./docs/supplemental-ui",
        "./docs/antora.yml",
        "antora-playbook.yml",
        ".git",
        ":adoc-manpages",
    ],
    outs = ["site"],
    cmd = [
        "mv MAN-MODULE docs/modules/MAN",
        "antora antora-playbook.yml --fetch",
        "mv build/site .",
        "echo koble.b177y.dev > site/CNAME",
        "touch site/.nojekyll",
    ],
)
