package torrentlookup

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var trackers []string = []string{
	"udp://open.demonii.com:1337/announce",
	"udp://tracker.publicbt.com:80/announce",
	"udp://tracker.openbittorrent.com:80/announce",
	"udp://tracker.istole.it:80",
	// "http://www.eddie4.nl:6969/announce",
	// "http://tracker.nwps.ws:6969/announce",
	// "http://bigfoot1942.sektori.org:6969/announce",
	// "http://9.rarbg.com:2710/announce",
	// "http://torrent-tracker.ru:80/announce.php",
	// "http://bttracker.crunchbanglinux.org:6969/announce",
	// "http://explodie.org:6969/announce",
	// "http://tracker.tfile.me/announce",
	// "http://tracker.best-torrents.net:6969/announce",
	// "http://tracker1.wasabii.com.tw:6969/announce",
	// "http://bt.careland.com.cn:6969/announce",
}

func Search(term string, deepCrawl bool) (string, string) {
	infohash := ""
	name := ""
	searchUrl := fmt.Sprintf("https://torrentz.eu/verified?f=%s", url.QueryEscape(term))
	doc, err := goquery.NewDocument(searchUrl)
	if err == nil {
		doc.Find(".results dl dt").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Find("a").Attr("href")
			name = s.Find("a").Text()
			if deepCrawl == true {
				results := listResultPages("https://torrentz.eu" + link)
				for _, link := range results {
					magnets := findMagnets(link)
					if len(magnets) > 0 {
						infohash = magnets[0] // TODO Return infohash, not magnet
					}
				}
			} else {
				infohash = strings.Trim(link, "/")
			}
		})
	}
	return name, infohash
}

func listResultPages(url string) map[string]string {
	results := make(map[string]string)
	doc, err := goquery.NewDocument(url)
	if err == nil {
		doc.Find(".download dl").Each(func(i int, s *goquery.Selection) {
			dd := s.Find("dd").Text()
			if dd != "Sponsored Link" {
				link, _ := s.Find("dt a").Attr("href")
				name := s.Find("dt a span.u").Text()
				results[name] = link
			}
		})
	}
	return results
}

func findMagnets(url string) []string {
	magnets := make([]string, 0)
	doc, err := goquery.NewDocument(url)
	if err == nil {
		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			link, _ := s.Attr("href")
			if strings.Contains(string(link), "magnet:") {
				magnets = append(magnets, link)
			}
		})
	}
	return nil
}

func FakeMagnet(infohash string) string {
	var magnetUrl string = fmt.Sprintf("magnet:?xt=urn:btih:%s", infohash)
	for _, tracker := range trackers {
		magnetUrl += fmt.Sprintf("&tr=%s", url.QueryEscape(tracker))
	}
	return magnetUrl
}
