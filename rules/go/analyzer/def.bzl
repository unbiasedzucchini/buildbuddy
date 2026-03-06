ANALYZERS = [
    # "any",
    # "bloop",
    # "fmtappendf",
    # "forvar",
    # "mapsloop",
    # "minmax",
    # "newexpr",
    # "omitzero",
    # "rangeint",
    # "reflecttypefor",
    # "slicescontains",
    # "slicessort",
    # "stditerators",
    # "stringscutprefix",
    # "stringsseq",
    # "stringsbuilder",
    # "testingcontext",
    # "waitgroup",
]

MODERNIZE_ANALYZERS = ["//rules/go/analyzer:" + analyzer for analyzer in ANALYZERS]
