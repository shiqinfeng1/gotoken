package gotoken

type TokenBuilder interface {
	Build(TokenOptions) Token
}

type Token interface {
	Generate()
	Refresh()
	Revoke()
	Prase()
	Verify()
}

func New(builder TokenBuilder, opts ...SetOptsFunc) Token {
	options := defaultTokenOptions
	for _, f := range opts {
		f(&options)
	}
	return builder.Build(options)
}
