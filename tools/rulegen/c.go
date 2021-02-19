package main

var cWorkspaceTemplate = mustTemplate(`load("@rules_proto_grpc//{{ .Lang.Dir }}:repositories.bzl", rules_proto_grpc_{{ .Lang.Name }}_repos = "{{ .Lang.Name }}_repos")

rules_proto_grpc_{{ .Lang.Name }}_repos()

load("@upb//bazel:workspace_deps.bzl", "upb_deps")

upb_deps()`)

var cProtoLibraryRuleTemplate = mustTemplate(`load("//{{ .Lang.Dir }}:{{ .Lang.Name }}_{{ .Rule.Kind }}_compile.bzl", "{{ .Lang.Name }}_{{ .Rule.Kind }}_compile")
load("//internal:compile.bzl", "proto_compile_attrs")
load("//internal:filter_files.bzl", "filter_files")
load("@rules_cc//cc:defs.bzl", "cc_library")

def {{ .Rule.Name }}(**kwargs):  # buildifier: disable=function-docstring
    # Compile protos
    name_pb = kwargs.get("name") + "_pb"
    {{ .Lang.Name }}_{{ .Rule.Kind }}_compile(
        name = name_pb,
        {{ .Common.ArgsForwardingSnippet }}
    )

    # Filter files to sources and headers
    filter_files(
        name = name_pb + "_srcs",
        target = name_pb,
        extensions = ["c"],
    )

    filter_files(
        name = name_pb + "_hdrs",
        target = name_pb,
        extensions = ["h"],
    )

    # Create {{ .Lang.Name }} library
    cc_library(
        name = kwargs.get("name"),
        srcs = [name_pb + "_srcs"],
        deps = PROTO_DEPS + (kwargs.get("deps", []) if "protos" in kwargs else []),
        hdrs = [name_pb + "_hdrs"],
        includes = [name_pb],
        visibility = kwargs.get("visibility"),
        tags = kwargs.get("tags"),
    )

PROTO_DEPS = [
    "@upb//:upb",
]`)

// For C, we need to manually generate the files for any.proto
var cProtoLibraryExampleTemplate = mustTemplate(`load("@rules_proto_grpc//{{ .Lang.Dir }}:defs.bzl", "{{ .Rule.Name }}")

{{ .Rule.Name }}(
    name = "proto_{{ .Lang.Name }}_{{ .Rule.Kind }}",
    importpath = "github.com/rules-proto-grpc/rules_proto_grpc/example/proto",
    protos = [
        "@com_google_protobuf//:any_proto",
        "@rules_proto_grpc//example/proto:person_proto",
        "@rules_proto_grpc//example/proto:place_proto",
        "@rules_proto_grpc//example/proto:thing_proto",
    ],
)`)

func makeC() *Language {
	return &Language{
		Dir:   "c",
		Name:  "c",
		DisplayName: "C",
		Notes: mustTemplate("Rules for generating C protobuf `.c` & `.h` files and libraries using [upb](https://github.com/protocolbuffers/upb). Libraries are created with the Bazel native `cc_library`"),
		Flags: commonLangFlags,
		Rules: []*Rule{
			&Rule{
				Name:             "c_proto_compile",
				Kind:             "proto",
				Implementation:   aspectRuleTemplate,
				Plugins:          []string{"//c:upb_plugin"},
				WorkspaceExample: cWorkspaceTemplate,
				BuildExample:     protoCompileExampleTemplate,
				Doc:              "Generates C protobuf `.h` & `.c` artifacts",
				Attrs:            compileRuleAttrs,
				Experimental:     true,
			},
			&Rule{
				Name:             "c_proto_library",
				Kind:             "proto",
				Implementation:   cProtoLibraryRuleTemplate,
				WorkspaceExample: cWorkspaceTemplate,
				BuildExample:     cProtoLibraryExampleTemplate,
				Doc:              "Generates a C protobuf library using `cc_library`, with dependencies linked",
				Attrs:            libraryRuleAttrs,
				Experimental:     true,
			},
		},
	}
}
