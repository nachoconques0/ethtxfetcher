package app

// Option defines the function used for
// setup an application option config
type Option func(a *Application)

// WithHTTPPort function adds the given http
// port into the application base config
func WithHTTPPort(port string) Option {
	return func(a *Application) {
		a.httpPort = port
	}
}

// WithNodeURL function adds the given ETH Node URL
// into the application base config
func WithNodeURL(url string) Option {
	return func(a *Application) {
		a.nodeURL = url
	}
}
