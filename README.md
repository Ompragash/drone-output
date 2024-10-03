This repository contains an example demonstrating how to export output variables from a Drone plugin. If the plugin writes output to the `DRONE_OUTPUT.env` file, you can reference those output variables in other steps and stages within the pipeline using expressions. For more details, refer to the official documentation: [Drone Plugin Step Settings - Output Variables](https://developer.harness.io/docs/continuous-integration/use-ci/use-drone-plugins/plugin-step-settings-reference/#output-variables).

In the sample plugin code [main.go](./main.go) , it exports an `FILES_INFO` output variable, which can be used in subsequent run steps. You can reuse the `WriteEnvToFile` function in your own Drone plugin to export output variables with values.

```go
func WriteEnvToFile(key, value string) error {
	outputFile, err := os.OpenFile(os.Getenv("DRONE_OUTPUT"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open output file: %w", err)
	}
	defer outputFile.Close()

	_, err = fmt.Fprintf(outputFile, "%s=%s\n", key, value)
	if err != nil {
		return fmt.Errorf("failed to write to env: %w", err)
	}

	return nil
}
