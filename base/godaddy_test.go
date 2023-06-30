package base_test

import (
	"testing"

	"github.com/buffreak/domain-scraper/base"
	"github.com/stretchr/testify/assert"
)

func TestGetDomainsGoDaddy(t *testing.T) {
	gd := &base.GoDaddy{}
	domains, err := gd.GetDomains("2", "ddlAdvKeyword:3|txtKeyword:|ddlCharacters:0|txtCharacters:|txtMinTraffic:|txtMaxTraffic:|txtMinDomainAge:|txtMaxDomainAge:|txtMinPrice:|txtMaxPrice:|ddlCategories:0|chkAddBuyNow:false|chkAddFeatured:false|chkAddDash:true|chkAddDigit:true|chkAddWeb:false|chkAddAppr:false|chkAddInv:false|chkAddReseller:false|ddlPattern1:|ddlPattern2:|ddlPattern3:|ddlPattern4:|chkSaleOffer:false|chkSalePublic:false|chkSaleExpired:false|chkSaleCloseouts:false|chkSaleUsed:false|chkSaleBuyNow:false|chkSaleDC:false|chkAddOnSale:false|ddlAdvBids:0|txtBids:|txtAuctionID:|ddlDateOffset:|ddlSort:currentpriceA", "-1")
	assert.Nil(t, err)
	t.Logf("%+v", domains)
	convert, _ := base.StructToMap(gd)
	t.Logf("%+v", convert)
}
