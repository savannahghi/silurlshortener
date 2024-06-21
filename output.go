package silurlshortener

import "time"

// ShortenURLResponse models the for a shortened URL
type ShortenURLResponse struct {
	ShortURL      string        `json:"shortUrl,omitempty" mapstructure:"shortURL"`
	ShortCode     string        `json:"shortCode,omitempty" mapstructure:"shortCode"`
	LongURL       string        `json:"longUrl,omitempty" mapstructure:"longURL"`
	DateCreated   time.Time     `json:"dateCreated,omitempty" mapstructure:"dateCreated"`
	Tags          []string      `json:"tags,omitempty" mapstructure:"tags"`
	Meta          Meta          `json:"meta,omitempty" mapstructure:"meta"`
	Domain        string        `json:"domain,omitempty" mapstructure:"domain"`
	Title         string        `json:"title,omitempty" mapstructure:"title"`
	Crawlable     bool          `json:"crawlable,omitempty" mapstructure:"crawlable"`
	ForwardQuery  bool          `json:"forwardQuery,omitempty" mapstructure:"forwardQuery"`
	VisitsSummary VisitsSummary `json:"visitsSummary,omitempty" mapstructure:"visitsSummary"`
}

// Meta models the metadata associated with the shortened URL
type Meta struct {
	ValidSince string `json:"validSince,omitempty" mapstructure:"validSince"`
	ValidUntil string `json:"validUntil,omitempty" mapstructure:"validUntil"`
	MaxVisits  string `json:"maxVisits,omitempty" mapstructure:"maxVisits"`
}

// VisitsSummary is used to describe the number of visits in a given URL.
type VisitsSummary struct {
	Total   int `json:"total,omitempty" mapstructure:"total"`
	NonBots int `json:"nonBots,omitempty" mapstructure:"nonBots"`
	Bots    int `json:"bots,omitempty" mapstructure:"bots"`
}
