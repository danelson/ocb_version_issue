# ocb version issue

1. Open devcontainer in VSCode.
2. Build the custom collector with the [bad builder config](builder-fail-config.yaml).

    ```sh
    go clean -modcache
    ./ocb --config builder-fail-config.yaml --verbose
    ```

1. Observe that
    1. The build fails due to an API change that was made after v0.96.0.
    1. Some v0.98.0 dependencies are downloaded (should only be v0.96.0).

    ```txt
    ...
    2024-04-17T22:43:33.837Z        INFO    zapio/writer.go:146     go: finding module for package github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage
    2024-04-17T22:43:34.242Z        INFO    zapio/writer.go:146     go: downloading github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage v0.98.0
    2024-04-17T22:43:34.316Z        INFO    zapio/writer.go:146     go: downloading github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage v0.98.0
    2024-04-17T22:43:34.319Z        INFO    zapio/writer.go:146     go: downloading github.com/open-telemetry/opentelemetry-collector-contrib v0.98.0
    2024-04-17T22:43:34.413Z        INFO    zapio/writer.go:146     go: found github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage in github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage v0.98.0
    2024-04-17T22:43:34.595Z        INFO    zapio/writer.go:146     go: downloading go.opentelemetry.io/collector/component v0.98.0
    2024-04-17T22:43:34.595Z        INFO    zapio/writer.go:146     go: downloading go.opentelemetry.io/collector/extension v0.98.0
    2024-04-17T22:43:34.595Z        INFO    zapio/writer.go:146     go: downloading github.com/stretchr/testify v1.9.0
    2024-04-17T22:43:34.596Z        INFO    zapio/writer.go:146     go: downloading go.opentelemetry.io/collector/pdata v1.5.0
    2024-04-17T22:43:34.596Z        INFO    zapio/writer.go:146     go: downloading go.etcd.io/bbolt v1.3.9
    2024-04-17T22:43:34.596Z        INFO    zapio/writer.go:146     go: downloading go.opentelemetry.io/collector/config/configtelemetry v0.98.0
    2024-04-17T22:43:34.596Z        INFO    zapio/writer.go:146     go: downloading go.opentelemetry.io/collector/confmap v0.98.0
    2024-04-17T22:43:34.596Z        INFO    zapio/writer.go:146     go: downloading go.opentelemetry.io/collector/featuregate v1.5.0
    2024-04-17T22:43:34.662Z        INFO    zapio/writer.go:146     go: downloading go.opentelemetry.io/otel/metric v1.25.0
    2024-04-17T22:43:34.666Z        INFO    zapio/writer.go:146     go: downloading go.opentelemetry.io/otel/trace v1.25.0
    2024-04-17T22:43:34.667Z        INFO    zapio/writer.go:146     go: downloading github.com/prometheus/client_model v0.6.1
    2024-04-17T22:43:34.670Z        INFO    zapio/writer.go:146     go: downloading go.opentelemetry.io/otel v1.25.0
    2024-04-17T22:43:34.670Z        INFO    zapio/writer.go:146     go: downloading go.opentelemetry.io/otel/exporters/prometheus v0.47.0
    2024-04-17T22:43:34.670Z        INFO    zapio/writer.go:146     go: downloading go.opentelemetry.io/otel/sdk/metric v1.25.0
    2024-04-17T22:43:34.671Z        INFO    zapio/writer.go:146     go: downloading go.opentelemetry.io/otel/sdk v1.25.0
    2024-04-17T22:43:34.703Z        INFO    zapio/writer.go:146     go: downloading github.com/knadh/koanf/v2 v2.1.1
    2024-04-17T22:43:34.901Z        INFO    zapio/writer.go:146     go: downloading golang.org/x/sync v0.6.0
    2024-04-17T22:43:34.903Z        INFO    zapio/writer.go:146     go: downloading golang.org/x/net v0.23.0
    ...
    2024-04-17T22:50:11.310Z        INFO    zapio/writer.go:146     # github.com/open-telemetry/opentelemetry-collector-contrib/exporter/myexporter
    2024-04-17T22:50:11.310Z        INFO    zapio/writer.go:146     /workspaces/ocb_version_issue/exporter/myexporter/factory.go:17:3: cannot use typeStr (untyped string constant "myexporter") as component.Type value in argument to exporter.NewFactory
    Error: failed to compile the OpenTelemetry Collector distribution: exit status 1
    ```

1. Update the way the extension is referenced by modifying the builder config.
    
    ```diff
    - - gomod: "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage v0.96.0"
    -   import: "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage"
    + - gomod: "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage v0.96.0"
    ```

1. Build the custom collector with the [good builder config](./builder-success-config.yaml).

    ```sh
    go clean -modcache
    ./ocb --config builder-success-config.yaml --verbose
    ```
1. Observe the build succeeds.
