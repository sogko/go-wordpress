package wordpress

import (
	"context"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/tomnomnom/linkheader"
)

// DiscoveredAPI is a struct containing details about a discovered WordPress REST API.
type DiscoveredAPI struct {
	BaseURL       string
	DiscoveredURL string
	ViaHeader     bool
	ViaHTML       bool
	Client        *Client
	BasicInfo     *RootInfo
}

// DiscoverAPI will discover the API root URL for the given base URL.
func DiscoverAPI(baseURL string, getRootInfo bool) (*DiscoveredAPI, error) {
	discovered := &DiscoveredAPI{
		BaseURL: baseURL,
	}
	res, httpErr := http.Get(baseURL)
	if httpErr != nil {
		return nil, httpErr
	}
	if res.Header.Get("Link") != "" {
		discoveredURL, linkErr := linkHeader(res)
		if linkErr != nil {
			return nil, linkErr
		}
		discovered.DiscoveredURL = discoveredURL
		discovered.ViaHeader = true
	} else {
		discoveredURL, linkErr := extractLinkFromHTML(res)
		if linkErr != nil {
			return nil, linkErr
		}
		discovered.DiscoveredURL = discoveredURL
		discovered.ViaHTML = true
	}
	clientOpts := &Options{
		BaseAPIURL: discovered.DiscoveredURL,
	}
	if getRootInfo {
		client := NewClient(clientOpts, nil)
		info, _, basicInfoErr := client.BasicInfo(context.Background())
		if basicInfoErr != nil {
			return nil, basicInfoErr
		}
		clientOpts.Location = info.Location
		discovered.BasicInfo = info
	}
	discovered.Client = NewClient(clientOpts, nil)
	return discovered, nil
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
