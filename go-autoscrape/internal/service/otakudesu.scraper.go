package service

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/anaskhan96/soup"
	"github.com/zeindevs/goautoscrape/internal"
	"github.com/zeindevs/goautoscrape/internal/model"
)

type Otakudesu struct {
	Title   string `json:"title"`
	Day     string `json:"day"`
	Episode string `json:"episode"`
	Image   string `json:"image"`
	Date    string `json:"date"`
	Url     string `json:"url"`
	Status  string `json:"status"`
}

func (o *Otakudesu) ToModel() *model.Otakudesu {
	return &model.Otakudesu{
		Title:   o.Title,
		Day:     o.Day,
		Date:    o.Date,
		Episode: o.Episode,
		Url:     o.Url,
		Image:   o.Image,
		Status:  o.Status,
	}
}

type OtakudesuScraper struct {
	httpClient *internal.HttpClient
}

func NewOtakudesuScraper() *OtakudesuScraper {
	return &OtakudesuScraper{
		httpClient: internal.NewHttpClient(),
	}
}

func (s *OtakudesuScraper) GetOngoing() ([]*Otakudesu, error) {
	s.httpClient.Header = http.Header{
		"User-Agent": []string{},
	}
	res, err := s.httpClient.Get(context.TODO(), "https://otakudesu.cloud/ongoing-anime")
	if err != nil {
		return nil, err
	}

	data := s.ParseAnime(res, "ongoing")

	return data, nil
}

func (s *OtakudesuScraper) GetOngoingAll() ([]*Otakudesu, error) {
	s.httpClient.Header = http.Header{
		"User-Agent": []string{},
	}
	res, err := s.httpClient.Get(context.TODO(), "https://otakudesu.cloud/ongoing-anime")
	if err != nil {
		return nil, err
	}

	pages := parsePagination(res)
	totalPage := pages[len(pages)-1]

	type datac struct {
		Page int
		Data []*Otakudesu
	}

	dataCh := make(chan datac)
	data := make(map[int][]*Otakudesu)

	data[1] = s.ParseAnime(res, "ongoing")

	if len(pages) >= 2 {
		for page := 2; page <= totalPage; page++ {
			go func(page int, dataCh chan datac) {
				ress, _ := s.httpClient.Get(context.TODO(), fmt.Sprintf("https://otakudesu.cloud/ongoing-anime/page/%d", page))
				dataCh <- datac{
					Page: page,
					Data: s.ParseAnime(ress, "ongoing"),
				}
			}(page, dataCh)
		}
	}

free:
	for {
		select {
		case d := <-dataCh:
			data[d.Page] = d.Data
		default:
			if len(data) == totalPage {
				break free
			}
		}
	}

	var resp []*Otakudesu

	for _, items := range data {
		for _, item := range items {
			resp = append(resp, item)
		}
	}

	return resp, nil
}

func (s *OtakudesuScraper) GetComplete() ([]*Otakudesu, error) {
	s.httpClient.Header = http.Header{
		"User-Agent": []string{},
	}
	res, err := s.httpClient.Get(context.TODO(), "https://otakudesu.cloud/complete-anime")
	if err != nil {
		return nil, err
	}

	data := s.ParseAnime(res, "complete")

	return data, nil
}

func (s *OtakudesuScraper) GetCompleteAll() ([]*Otakudesu, error) {
	s.httpClient.Header = http.Header{
		"User-Agent": []string{},
	}
	res, err := s.httpClient.Get(context.TODO(), "https://otakudesu.cloud/complete-anime")
	if err != nil {
		return nil, err
	}

	pages := parsePagination(res)
	totalPage := pages[len(pages)-1]

	type datac struct {
		Page int
		Data []*Otakudesu
	}

	dataCh := make(chan datac)
	data := make(map[int][]*Otakudesu)

	data[1] = s.ParseAnime(res, "complete")

	if len(pages) >= 2 {
		for page := 2; page <= totalPage; page++ {
			go func(page int, dataCh chan datac) {
				ress, _ := s.httpClient.Get(context.TODO(), fmt.Sprintf("https://otakudesu.cloud/complete-anime/page/%d", page))
				dataCh <- datac{
					Page: page,
					Data: s.ParseAnime(ress, "complete"),
				}
			}(page, dataCh)
		}
	}

free:
	for {
		select {
		case d := <-dataCh:
			data[d.Page] = d.Data
		default:
			if len(data) == totalPage {
				break free
			}
		}
	}

	var resp []*Otakudesu

	for _, items := range data {
		for _, item := range items {
			resp = append(resp, item)
		}
	}

	return resp, nil
}

func (o *OtakudesuScraper) ParseAnime(res []byte, status string) []*Otakudesu {
	html := soup.HTMLParse(string(res))
	content := html.Find("div", "class", "venz")
	posts := content.FindAll("li")

	var data []*Otakudesu
	for _, post := range posts {
		info := Otakudesu{
			Title:   post.Find("h2", "class", "jdlflm").Text(),
			Image:   post.Find("img", "class", "attachment-thumb").Attrs()["src"],
			Episode: post.Find("div", "class", "epz").Text(),
			Day:     post.Find("div", "class", "epztipe").Text(),
			Date:    post.Find("div", "class", "newnime").Text(),
			Url:     post.Find("a").Attrs()["href"],
			Status:  status,
		}
		data = append(data, &info)
	}
	return data
}

func parsePagination(res []byte) []int {
	html := soup.HTMLParse(string(res))
	pagination := html.Find("div", "class", "pagination")
	links := pagination.FindAll("a")
	pages := []int{}

	pattern := regexp.MustCompile(`\d+`)

	for _, link := range links {
		num := pattern.Find([]byte(link.Text()))
		if num != nil {
			n, _ := strconv.Atoi(string(num))
			pages = append(pages, n)
		}
	}

	return pages
}
