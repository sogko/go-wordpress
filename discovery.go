package wordpress

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/tomnomnom/linkheader"
)

type DiscoveredAPI struct {
	BaseURL       string
	DiscoveredURL string
	ViaHeader     bool
	ViaHTML       bool
}

// DiscoverAPI will discover the API root URL for the given base URL.
func DiscoverAPI(baseURL string) (*DiscoveredAPI, error) {
	res, httpErr := http.Get(baseURL)
	if httpErr != nil {
		return nil, httpErr
	}
	if res.Header.Get("Link") != "" {
		discoveredURL, linkErr := linkHeader(res)
		if linkErr != nil {
			return nil, linkErr
		}
		return &DiscoveredAPI{
			BaseURL:       baseURL,
			DiscoveredURL: discoveredURL,
			ViaHeader:     true,
		}, nil
	}
	discoveredURL, linkErr := extractLinkFromHTML(res)
	if linkErr != nil {
		return nil, linkErr
	}
	return &DiscoveredAPI{
		BaseURL:       baseURL,
		DiscoveredURL: discoveredURL,
		ViaHTML:       true,
	}, nil
}

func linkHeader(resp *http.Response) (string, error) {
	for _, link := range linkheader.Parse(resp.Header.Get("Link")) {
		if link.Rel == "https://api.w.org/" {
			return link.URL, nil
		}
	}
	return "", nil
}

func extractLinkFromHTML(resp *http.Response) (string, error) {
	doc, docErr := goquery.NewDocumentFromResponse(resp)
	if docErr != nil {
		return "", docErr
	}
	href, hrefExists := doc.Find(`link[rel="https://api.w.org/"]`).Attr("href")
	if hrefExists {
		return href, nil
	}
	return "", nil
}
