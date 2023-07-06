package jwt

type builder struct {
	opts JwtOptions
}

func (b *builder) Build(opts JwtOptions) Token {
	return &JwtToken{}
}
func NewJwtBuilder(o TokenOptions) TokenBuilder {
	return &builder{
		opts: o,
	}
}
