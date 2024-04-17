package myexporter

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/exporter"
)

const (
	typeStr = "myexporter"
)

// NewFactory creates New Relic logs exporter factory.
func NewFactory() exporter.Factory {
	return exporter.NewFactory(
		typeStr,
		createDefaultConfig,
		exporter.WithLogs(createLogsExporter, component.StabilityLevelStable))
}

func createDefaultConfig() component.Config {
	return &Config{}
}

func createLogsExporter(
	ctx context.Context,
	settings exporter.CreateSettings,
	cfg component.Config,
) (exporter.Logs, error) {
	return nil, nil
}
