package silurlshortener

// ShortenURLPayload is the input used to shorten any given URL
type ShortenURLPayload struct {
	LongURL         string   `json:"longUrl,omitempty"`
	ValidSince      string   `json:"validSince,omitempty"`
	ValidUntil      string   `json:"validUntil,omitempty"`
	MaxVisits       int      `json:"maxVisits,omitempty"`
	Tags            []string `json:"tags,omitempty"`
	Title           string   `json:"title,omitempty"`
	Crawlable       bool     `json:"crawlable,omitempty"`
	ForwardQuery    bool     `json:"forwardQuery,omitempty"`
	CustomSlug      string   `json:"customSlug,omitempty"`
	PathPrefix      string   `json:"pathPrefix,omitempty"`
	FindIfExists    bool     `json:"findIfExists,omitempty"`
	Domain          string   `json:"domain,omitempty"`
	ShortCodeLength int      `json:"shortCodeLength,omitempty"`
}
