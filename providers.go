package torrentlookup

// ProviderTPB -
var ProviderTPB = &Provider{
	Name:           "The pirate bay",
	SearchURL:      "https://thepiratebay.org/search/%s/0/99/0",
	RowQuery:       "#searchResult tr",
	NameSubQuery:   ".detName a",
	MagnetSubQuery: "td:nth-child(2) > a:nth-child(2)",
	SeedsSubQuery:  "td:nth-child(3)",
}