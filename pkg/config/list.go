package config

// ListOptions is the options for the build command
type ListOptions struct {
	// Output is the output format
	Output string
	// KeepTempDir is the keep temp dir flag
	KeepTempDir bool
	// SkipCharts makes List skip `withPreparedCharts`
	SkipCharts bool
	// IncludeNeeds is the include needs flag
	IncludeNeeds bool
	// IncludeTransitiveNeeds is the include transitive needs flag
	IncludeTransitiveNeeds bool
}

// NewListOptions creates a new Apply
func NewListOptions() *ListOptions {
	return &ListOptions{}
}

// ListImpl is impl for applyOptions
type ListImpl struct {
	*GlobalImpl
	*ListOptions
}

// NewListImpl creates a new ListImpl
func NewListImpl(g *GlobalImpl, b *ListOptions) *ListImpl {
	return &ListImpl{
		GlobalImpl:  g,
		ListOptions: b,
	}
}

// Output returns the output
func (c *ListImpl) Output() string {
	return c.ListOptions.Output
}

// SkipCharts returns skipCharts flag
func (c *ListImpl) SkipCharts() bool {
	return c.ListOptions.SkipCharts
}

// IncludeNeeds returns includeNeeds flag
func (c *ListImpl) IncludeNeeds() bool {
	return c.ListOptions.IncludeNeeds
}

// IncludeTransitiveNeeds returns includeTransitiveNeeds flag
func (c *ListImpl) IncludeTransitiveNeeds() bool {
	return c.ListOptions.IncludeTransitiveNeeds
}
