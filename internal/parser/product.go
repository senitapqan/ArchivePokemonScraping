package parser

import (
	"errors"
	"fmt"
	"goPockemonParsingService/models"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"
)

func (p *parser) GetPageCount() (int, error) {
	url := "https://scrapeme.live/shop/"

	resp, err := http.Get(url)
	if err != nil {
		log.Err(err).Msg("error while getting response from http request with url: " + url)
		return -1, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Err(err).Msg("error while getting document from http response")
		return -1, err
	}

	nav := doc.Find("ul.page-numbers")
	if nav.Length() == 0 {
		return -1, errors.New("there is no navigation throw pages")
	}

	pageCount := 0
	nav.Find("a").Each(func(index int, item *goquery.Selection) {
		text := item.Text()
		if num, err := strconv.Atoi(text); err == nil {
			if num > pageCount {
				pageCount = num
			}
		}
	})

	return pageCount, nil
}

func (p *parser) GetPockemonsOnPage(page int) ([]models.Pockemon, error) {
	var pockemons []models.Pockemon

	url := fmt.Sprintf("https://scrapeme.live/shop/page/%d", page)

	resp, err := http.Get(url)
	if err != nil {
		log.Err(err).Msg("error while getting response from http request with url: " + url)
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Err(err).Msg("error while getting document from http response")
		return nil, err
	}

	ul := doc.Find("ul.products.columns-4")
	if ul.Length() == 0 {
		return nil, errors.New("there is no list of pockemons")
	}

	ul.Find("li").Each(func(index int, item *goquery.Selection) {
		current := item.Find("a.woocommerce-LoopProduct-link.woocommerce-loop-product__link")

		imageUrl, _ := current.Find("img").Attr("src")
		name := current.Find("h2.woocommerce-loop-product__title").Text()
		priceAmount := current.Find("span.woocommerce-Price-amount.amount").Text()
		priceSymbol := current.Find("span.woocommerce-Price-currencySymbol").Text()

		pockemon := models.NewPockemonBuilder().SetName(name).SetPrice(priceAmount, priceSymbol).SetImageUrl(imageUrl).Build()

		pockemons = append(pockemons, *pockemon)
	})

	return pockemons, nil
}
