package worker

import (
	"context"
	//"io"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

type Job struct {
	Cancel context.CancelFunc
	// ETCD key, length : 160
	ID string
	JobVal
}

type JobVal struct {
	URL     string `json:"url"`
	Pattern string `json:"pattern"`
	// SM OracleWiter related, not sure yet
}

// TODO: add timeout
func (j Job) Scrap() (string, error) {
	log.Println("start scraping")
	res, err := http.Get(j.URL)
	if err != nil {
		return "", err
	}
	data, err := j.resolve(res)
	if err != nil {
		return "", err
	}
	return data, nil
}

// Not implemented
// TODO: add resolver
func (j Job) resolve(resp *http.Response) (string, error) {
	defer resp.Body.Close()
	//body, err := io.ReadAll(resp.Body)
	dom, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	var result string
	dom.Find(j.Pattern).Each(func(i int, selection *goquery.Selection) {
		result = selection.Text()
		
	})
	return result, nil
	// TODO: ADD RESOLVER
	//return string(body), nil
}
