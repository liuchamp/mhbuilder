package builder

type ScopeOuter interface {
	GetImports() (string, error)
	GetBody() (string, error)
	GetOuter() (string, error)
}
