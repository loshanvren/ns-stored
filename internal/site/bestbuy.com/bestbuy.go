package bestbuy_com

import (
	"context"
	"github.com/Gssssssssy/ns-stored/internal/site"
	"github.com/Gssssssssy/ns-stored/internal/task"
	"github.com/gocolly/colly"
)

// 链接
const (
	ItemBlackDetailURL      = "https://www.bestbuy.com/site/nintendo-switch-32gb-console-gray-joy-con/6364253.p?skuId=6364253"
	ItemBlueAndRedDetailURL = "https://www.bestbuy.com/site/nintendo-switch-32gb-console-neon-red-neon-blue-joy-con/6364255.p?skuId=6364255"
	inquiryBasePath         = `/v1/products(sku="%d"|sku="%d")?format=json&show=sku,name,salePrice,inStoreAvailability&apiKey=%s`
)

// 货号
const (
	SkuIDBlack      = 6364253
	SkuIDBlueAndRed = 6364255
)

func makeColly() *colly.Collector {
	c := colly.NewCollector(colly.AllowURLRevisit())
	return c
}

type collector struct{}

func (c *collector) Inquiry(ctx context.Context) (*task.Result, error) {
	return doInquiry(ctx)
}

func NewCollector() site.Collector {
	return new(collector)
}

var _ site.Collector = (*collector)(nil)
