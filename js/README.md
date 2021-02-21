# JavaScript rules

Rules for generating JavaScript protobuf, gRPC-node and gRPC-Web `.js` and `.d.ts` files using standard Protocol Buffers and gRPC.

| Rule | Description |
| ---: | :--- |
| [js_proto_compile](#js_proto_compile) | Generates JavaScript protobuf `.js` and `.d.ts` artifacts |
| [js_grpc_node_compile](#js_grpc_node_compile) | Generates JavaScript protobuf + gRPC-node `.js` and `.d.ts` artifacts |
| [js_grpc_web_compile](#js_grpc_web_compile) | Generates JavaScript protobuf + gRPC-Web `.js` and `.d.ts` artifacts |
| [js_proto_library](#js_proto_library) | Generates a JavaScript protobuf library using `js_library` from `rules_nodejs` |
| [js_grpc_node_library](#js_grpc_node_library) | Generates a Node.js protobuf + gRPC-node library using `js_library` from `rules_nodejs` |
| [js_grpc_web_library](#js_grpc_web_library) | Generates a JavaScript protobuf + gRPC-Web library using `js_library` from `rules_nodejs` |

---

## `js_proto_compile`

Generates JavaScript protobuf `.js` and `.d.ts` artifacts

### `WORKSPACE`

```starlark
load("@rules_proto_grpc//js:repositories.bzl", rules_proto_grpc_js_repos = "js_repos")

rules_proto_grpc_js_repos()

load("@build_bazel_rules_nodejs//:index.bzl", "yarn_install")

yarn_install(
    name = "js_modules",
    package_json = "@rules_proto_grpc//js:requirements/package.json",
    yarn_lock = "@rules_proto_grpc//js:requirements/yarn.lock",
)
```

### `BUILD.bazel`

```starlark
load("@rules_proto_grpc//js:defs.bzl", "js_proto_compile")

js_proto_compile(
    name = "person_js_proto",
    protos = ["@rules_proto_grpc//example/proto:person_proto"],
)

js_proto_compile(
    name = "place_js_proto",
    protos = ["@rules_proto_grpc//example/proto:place_proto"],
)

js_proto_compile(
    name = "thing_js_proto",
    protos = ["@rules_proto_grpc//example/proto:thing_proto"],
)
```

### Attributes

| Name | Type | Mandatory | Default | Description |
| ---: | :--- | --------- | ------- | ----------- |
| `protos` | `list<ProtoInfo>` | true | `[]`    | List of labels that provide a `ProtoInfo` (such as `rules_proto` `proto_library`)          |
| `options` | `dict<string, list(string)>` | false | `[]`    | Extra options to pass to plugins, as a dict of plugin label -> list of strings. The key * can be used exclusively to apply to all plugins          |
| `verbose` | `int` | false | `0`    | The verbosity level. Supported values and results are 1: *show command*, 2: *show command and sandbox after running protoc*, 3: *show command and sandbox before and after running protoc*, 4. *show env, command, expected outputs and sandbox before and after running protoc*          |
| `prefix_path` | `string` | false | `""`    | Path to prefix to the generated files in the output directory          |
| `extra_protoc_args` | `list<string>` | false | `[]`    | A list of extra args to pass directly to protoc, not as plugin options          |

### Plugins

- `@rules_proto_grpc//js:js_plugin`
- `@rules_proto_grpc//js:ts_plugin`

---

## `js_grpc_node_compile`

Generates JavaScript protobuf + gRPC-node `.js` and `.d.ts` artifacts

### `WORKSPACE`

```starlark
load("@rules_proto_grpc//js:repositories.bzl", rules_proto_grpc_js_repos = "js_repos")

rules_proto_grpc_js_repos()

load("@build_bazel_rules_nodejs//:index.bzl", "yarn_install")

yarn_install(
    name = "js_modules",
    package_json = "@rules_proto_grpc//js:requirements/package.json",
    yarn_lock = "@rules_proto_grpc//js:requirements/yarn.lock",
)
```

### `BUILD.bazel`

```starlark
load("@rules_proto_grpc//js:defs.bzl", "js_grpc_node_compile")

js_grpc_node_compile(
    name = "thing_js_grpc",
    protos = ["@rules_proto_grpc//example/proto:thing_proto"],
)

js_grpc_node_compile(
    name = "greeter_js_grpc",
    protos = ["@rules_proto_grpc//example/proto:greeter_grpc"],
)
```

### Attributes

| Name | Type | Mandatory | Default | Description |
| ---: | :--- | --------- | ------- | ----------- |
| `protos` | `list<ProtoInfo>` | true | `[]`    | List of labels that provide a `ProtoInfo` (such as `rules_proto` `proto_library`)          |
| `options` | `dict<string, list(string)>` | false | `[]`    | Extra options to pass to plugins, as a dict of plugin label -> list of strings. The key * can be used exclusively to apply to all plugins          |
| `verbose` | `int` | false | `0`    | The verbosity level. Supported values and results are 1: *show command*, 2: *show command and sandbox after running protoc*, 3: *show command and sandbox before and after running protoc*, 4. *show env, command, expected outputs and sandbox before and after running protoc*          |
| `prefix_path` | `string` | false | `""`    | Path to prefix to the generated files in the output directory          |
| `extra_protoc_args` | `list<string>` | false | `[]`    | A list of extra args to pass directly to protoc, not as plugin options          |

### Plugins

- `@rules_proto_grpc//js:js_plugin`
- `@rules_proto_grpc//js:grpc_node_plugin`
- `@rules_proto_grpc//js:grpc_node_ts_plugin`

---

## `js_grpc_web_compile`

Generates JavaScript protobuf + gRPC-Web `.js` and `.d.ts` artifacts

### `WORKSPACE`

```starlark
load("@rules_proto_grpc//js:repositories.bzl", rules_proto_grpc_js_repos = "js_repos")

rules_proto_grpc_js_repos()

load("@build_bazel_rules_nodejs//:index.bzl", "yarn_install")

yarn_install(
    name = "js_modules",
    package_json = "@rules_proto_grpc//js:requirements/package.json",
    yarn_lock = "@rules_proto_grpc//js:requirements/yarn.lock",
)
```

### `BUILD.bazel`

```starlark
load("@rules_proto_grpc//js:defs.bzl", "js_grpc_web_compile")

js_grpc_web_compile(
    name = "thing_js_grpc",
    protos = ["@rules_proto_grpc//example/proto:thing_proto"],
)

js_grpc_web_compile(
    name = "greeter_js_grpc",
    protos = ["@rules_proto_grpc//example/proto:greeter_grpc"],
)
```

### Attributes

| Name | Type | Mandatory | Default | Description |
| ---: | :--- | --------- | ------- | ----------- |
| `protos` | `list<ProtoInfo>` | true | `[]`    | List of labels that provide a `ProtoInfo` (such as `rules_proto` `proto_library`)          |
| `options` | `dict<string, list(string)>` | false | `[]`    | Extra options to pass to plugins, as a dict of plugin label -> list of strings. The key * can be used exclusively to apply to all plugins          |
| `verbose` | `int` | false | `0`    | The verbosity level. Supported values and results are 1: *show command*, 2: *show command and sandbox after running protoc*, 3: *show command and sandbox before and after running protoc*, 4. *show env, command, expected outputs and sandbox before and after running protoc*          |
| `prefix_path` | `string` | false | `""`    | Path to prefix to the generated files in the output directory          |
| `extra_protoc_args` | `list<string>` | false | `[]`    | A list of extra args to pass directly to protoc, not as plugin options          |

### Plugins

- `@rules_proto_grpc//js:js_plugin`
- `@rules_proto_grpc//js:grpc_web_js_plugin`

---

## `js_proto_library`

Generates a JavaScript protobuf library using `js_library` from `rules_nodejs`

### `WORKSPACE`

```starlark
load("@rules_proto_grpc//js:repositories.bzl", rules_proto_grpc_js_repos = "js_repos")

rules_proto_grpc_js_repos()

load("@build_bazel_rules_nodejs//:index.bzl", "yarn_install")

yarn_install(
    name = "js_modules",
    package_json = "@rules_proto_grpc//js:requirements/package.json",
    yarn_lock = "@rules_proto_grpc//js:requirements/yarn.lock",
)
```

### `BUILD.bazel`

```starlark
load("@rules_proto_grpc//js:defs.bzl", "js_proto_library")

js_proto_library(
    name = "person_js_proto",
    protos = ["@rules_proto_grpc//example/proto:person_proto"],
    deps = ["place_js_proto"],
)

js_proto_library(
    name = "place_js_proto",
    protos = ["@rules_proto_grpc//example/proto:place_proto"],
    deps = ["thing_js_proto"],
)

js_proto_library(
    name = "thing_js_proto",
    protos = ["@rules_proto_grpc//example/proto:thing_proto"],
)
```

### Attributes

| Name | Type | Mandatory | Default | Description |
| ---: | :--- | --------- | ------- | ----------- |
| `protos` | `list<ProtoInfo>` | true | `[]`    | List of labels that provide a `ProtoInfo` (such as `rules_proto` `proto_library`)          |
| `options` | `dict<string, list(string)>` | false | `[]`    | Extra options to pass to plugins, as a dict of plugin label -> list of strings. The key * can be used exclusively to apply to all plugins          |
| `verbose` | `int` | false | `0`    | The verbosity level. Supported values and results are 1: *show command*, 2: *show command and sandbox after running protoc*, 3: *show command and sandbox before and after running protoc*, 4. *show env, command, expected outputs and sandbox before and after running protoc*          |
| `prefix_path` | `string` | false | `""`    | Path to prefix to the generated files in the output directory          |
| `extra_protoc_args` | `list<string>` | false | `[]`    | A list of extra args to pass directly to protoc, not as plugin options          |
| `deps` | `list<Label/string>` | false | `[]`    | List of labels to pass as deps attr to underlying lang_library rule          |

---

## `js_grpc_node_library`

Generates a Node.js protobuf + gRPC-node library using `js_library` from `rules_nodejs`

### `WORKSPACE`

```starlark
load("@rules_proto_grpc//js:repositories.bzl", rules_proto_grpc_js_repos = "js_repos")

rules_proto_grpc_js_repos()

load("@build_bazel_rules_nodejs//:index.bzl", "yarn_install")

yarn_install(
    name = "js_modules",
    package_json = "@rules_proto_grpc//js:requirements/package.json",
    yarn_lock = "@rules_proto_grpc//js:requirements/yarn.lock",
)
```

### `BUILD.bazel`

```starlark
load("@rules_proto_grpc//js:defs.bzl", "js_grpc_node_library")

js_grpc_node_library(
    name = "thing_js_grpc",
    protos = ["@rules_proto_grpc//example/proto:thing_proto"],
)

js_grpc_node_library(
    name = "greeter_js_grpc",
    protos = ["@rules_proto_grpc//example/proto:greeter_grpc"],
    deps = ["thing_js_grpc"],
)
```

### Attributes

| Name | Type | Mandatory | Default | Description |
| ---: | :--- | --------- | ------- | ----------- |
| `protos` | `list<ProtoInfo>` | true | `[]`    | List of labels that provide a `ProtoInfo` (such as `rules_proto` `proto_library`)          |
| `options` | `dict<string, list(string)>` | false | `[]`    | Extra options to pass to plugins, as a dict of plugin label -> list of strings. The key * can be used exclusively to apply to all plugins          |
| `verbose` | `int` | false | `0`    | The verbosity level. Supported values and results are 1: *show command*, 2: *show command and sandbox after running protoc*, 3: *show command and sandbox before and after running protoc*, 4. *show env, command, expected outputs and sandbox before and after running protoc*          |
| `prefix_path` | `string` | false | `""`    | Path to prefix to the generated files in the output directory          |
| `extra_protoc_args` | `list<string>` | false | `[]`    | A list of extra args to pass directly to protoc, not as plugin options          |
| `deps` | `list<Label/string>` | false | `[]`    | List of labels to pass as deps attr to underlying lang_library rule          |

---

## `js_grpc_web_library`

Generates a JavaScript protobuf + gRPC-Web library using `js_library` from `rules_nodejs`

### `WORKSPACE`

```starlark
load("@rules_proto_grpc//js:repositories.bzl", rules_proto_grpc_js_repos = "js_repos")

rules_proto_grpc_js_repos()

load("@build_bazel_rules_nodejs//:index.bzl", "yarn_install")

yarn_install(
    name = "js_modules",
    package_json = "@rules_proto_grpc//js:requirements/package.json",
    yarn_lock = "@rules_proto_grpc//js:requirements/yarn.lock",
)
```

### `BUILD.bazel`

```starlark
load("@rules_proto_grpc//js:defs.bzl", "js_grpc_web_library")

js_grpc_web_library(
    name = "person_js_grpc",
    protos = ["@rules_proto_grpc//example/proto:person_proto"],
    deps = ["place_js_grpc"],
)

js_grpc_web_library(
    name = "place_js_grpc",
    protos = ["@rules_proto_grpc//example/proto:place_proto"],
    deps = ["thing_js_grpc"],
)

js_grpc_web_library(
    name = "thing_js_grpc",
    protos = ["@rules_proto_grpc//example/proto:thing_proto"],
)
```

### Attributes

| Name | Type | Mandatory | Default | Description |
| ---: | :--- | --------- | ------- | ----------- |
| `protos` | `list<ProtoInfo>` | true | `[]`    | List of labels that provide a `ProtoInfo` (such as `rules_proto` `proto_library`)          |
| `options` | `dict<string, list(string)>` | false | `[]`    | Extra options to pass to plugins, as a dict of plugin label -> list of strings. The key * can be used exclusively to apply to all plugins          |
| `verbose` | `int` | false | `0`    | The verbosity level. Supported values and results are 1: *show command*, 2: *show command and sandbox after running protoc*, 3: *show command and sandbox before and after running protoc*, 4. *show env, command, expected outputs and sandbox before and after running protoc*          |
| `prefix_path` | `string` | false | `""`    | Path to prefix to the generated files in the output directory          |
| `extra_protoc_args` | `list<string>` | false | `[]`    | A list of extra args to pass directly to protoc, not as plugin options          |
| `deps` | `list<Label/string>` | false | `[]`    | List of labels to pass as deps attr to underlying lang_library rule          |