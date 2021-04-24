package app

import "go.uber.org/zap"

type ConfigProvider interface {
	Args() string
	HelmBinary() string

	FileOrDir() string
	KubeContext() string
	Namespace() string
	Chart() string
	Selectors() []string
	StateValuesSet() map[string]interface{}
	StateValuesFiles() []string
	Env() string

	loggingConfig
}

type DeprecatedChartsConfigProvider interface {
	Values() []string

	concurrencyConfig
	loggingConfig
}

type DepsConfigProvider interface {
	Args() string
	SkipRepos() bool
}

type ReposConfigProvider interface {
	Args() string
}

type ApplyConfigProvider interface {
	Args() string

	Values() []string
	Set() []string
	SkipCRDs() bool
	SkipDeps() bool
	Wait() bool
	WaitForJobs() bool

	IncludeTests() bool

	SuppressSecrets() bool
	ShowSecrets() bool
	SuppressDiff() bool

	DetailedExitcode() bool

	NoColor() bool
	Context() int
	DiffOutput() string

	RetainValuesFiles() bool
	SkipCleanup() bool
	SkipDiffOnInstall() bool

	SkipNeeds() bool
	IncludeNeeds() bool

	concurrencyConfig
	interactive
	loggingConfig
}

type SyncConfigProvider interface {
	Args() string

	Values() []string
	Set() []string
	SkipCRDs() bool
	SkipDeps() bool
	Wait() bool
	WaitForJobs() bool

	SkipNeeds() bool
	IncludeNeeds() bool

	concurrencyConfig
	loggingConfig
}

type DiffConfigProvider interface {
	Args() string

	Values() []string
	Set() []string
	SkipCRDs() bool
	SkipDeps() bool

	IncludeTests() bool

	SuppressSecrets() bool
	ShowSecrets() bool
	SuppressDiff() bool

	SkipNeeds() bool
	IncludeNeeds() bool

	DetailedExitcode() bool
	NoColor() bool
	Context() int
	DiffOutput() string

	concurrencyConfig
}

type DeleteConfigProvider interface {
	Args() string

	Purge() bool
	SkipDeps() bool

	interactive
	loggingConfig
	concurrencyConfig
}

type DestroyConfigProvider interface {
	Args() string

	SkipDeps() bool

	interactive
	loggingConfig
	concurrencyConfig
}

type TestConfigProvider interface {
	Args() string

	SkipDeps() bool
	Timeout() int
	Cleanup() bool
	Logs() bool

	concurrencyConfig
}

type LintConfigProvider interface {
	Args() string

	Values() []string
	Set() []string
	SkipDeps() bool

	concurrencyConfig
}

type FetchConfigProvider interface {
	SkipDeps() bool
	OutputDir() string

	concurrencyConfig
}

type TemplateConfigProvider interface {
	Args() string

	Values() []string
	Set() []string
	OutputDirTemplate() string
	Validate() bool
	SkipDeps() bool
	SkipCleanup() bool
	OutputDir() string
	IncludeCRDs() bool
	IncludeNeeds() bool

	concurrencyConfig
}

type WriteValuesConfigProvider interface {
	Values() []string
	Set() []string
	OutputFileTemplate() string
	SkipDeps() bool
}

type StatusesConfigProvider interface {
	Args() string

	concurrencyConfig
}

type StateConfigProvider interface {
	EmbedValues() bool
}

type concurrencyConfig interface {
	Concurrency() int
}

type loggingConfig interface {
	Logger() *zap.SugaredLogger
}

type interactive interface {
	Interactive() bool
}

type ListConfigProvider interface {
	Output() string
}
