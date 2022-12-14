package ir

type ConnectorType string
type Lang string

const (
	GoLang Lang = "golang"

	ConnectorSource      ConnectorType = "source"
	ConnectorDestination ConnectorType = "destination"
)

type DeploymentSpec struct {
	Secrets    map[string]string `json:"secrets,omitempty"`
	Connectors []ConnectorSpec   `json:"connectors"`
	Functions  []FunctionSpec    `json:"functions,omitempty"`
	Definition DefinitionSpec    `json:"definition"`
}

type ConnectorSpec struct {
	Type       ConnectorType          `json:"type"`
	Resource   string                 `json:"resource"`
	Collection string                 `json:"collection"`
	Config     map[string]interface{} `json:"config,omitempty"`
}

type FunctionSpec struct {
	Name    string                 `json:"name"`
	Image   string                 `json:"image"`
	EnvVars map[string]interface{} `json:"env_vars,omitempty"`
}

type DefinitionSpec struct {
	GitSha   string       `json:"git_sha"`
	Metadata MetadataSpec `json:"metadata"`
}

type MetadataSpec struct {
	Turbine     TurbineSpec `json:"turbine"`
	SpecVersion string      `json:"spec_version"`
}

type TurbineSpec struct {
	Language Lang   `json:"language"`
	Version  string `json:"version"`
}
