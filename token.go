package gotoken

type TokenBuilder interface {
	Build(TokenOptions) Token
}

type SignedToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type Token interface {
	Generate() (*SignedToken, error)
	Refresh()
	UpdateSecret()
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
