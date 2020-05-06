package bestbuy_com

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Gssssssssy/ns-onsale/internal/site"
	"github.com/Gssssssssy/ns-onsale/internal/task"
	"github.com/avast/retry-go"
	"github.com/gocolly/colly"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

var productURL string

func init() {
	basePath := HOST + inquiryBasePath
	productURL = fmt.Sprintf(basePath, SkuIDBlack, SkuIDBlueAndRed, Token)
}

type inquiryResponse struct {
	From         int    `json:"from"`
	To           int    `json:"to"`
	CurrentPage  int    `json:"currentPage"`
	Total        int    `json:"total"`
	TotalPages   int    `json:"totalPages"`
	QueryTime    string `json:"queryTime"`
	TotalTime    string `json:"totalTime"`
	Partial      bool   `json:"partial"`
	CanonicalURL string `json:"canonicalUrl"`
	Products     []struct {
		Sku                 int     `json:"sku"`
		Name                string  `json:"name"`
		SalePrice           float64 `json:"salePrice"`
		InStoreAvailability bool    `json:"inStoreAvailability"`
	} `json:"products"`
}

func doInquiry(_ context.Context) (*task.Result, error) {
	var (
		c                  = makeColly()
		err, errOnResponse error
		decoded            = inquiryResponse{}
		result             = &task.Result{IsAlarm: false}
	)
	c.OnResponse(func(response *colly.Response) {
		errOnResponse = json.Unmarshal(response.Body, &decoded)
		if errOnResponse != nil {
			return
		}
		for _, data := range decoded.Products {
			switch data.Sku {
			case SkuIDBlack:
				result.Name1 = data.Name
				result.Price1 = strconv.FormatFloat(data.SalePrice, 'E', -1, 64)
				result.Available1 = "No"
				result.Link1 = ItemBlackDetailURL
				if data.InStoreAvailability {
					result.Available1 = "Yes"
					result.IsAlarm = true
				}
			case SkuIDBlueAndRed:
				result.Name2 = data.Name
				result.Price2 = strconv.FormatFloat(data.SalePrice, 'E', -1, 64)
				result.Available2 = "No"
				result.Link2 = ItemBlueAndRedDetailURL
				if data.InStoreAvailability {
					result.Available2 = "Yes"
					result.IsAlarm = true
				}
			default:
			}
			result.UpdatedTime = time.Now().Format("2006-01-02 15:04:05")
		}
	})
	err = retry.Do(func() error {
		return c.Visit(productURL)
	}, retry.Attempts(site.DefaultRetryTimes), retry.Delay(500*time.Millisecond))
	if err != nil {
		return result, errors.WithStack(err)
	}
	if errOnResponse != nil {
		return result, errors.WithStack(errOnResponse)
	}
	return result, nil
}
