package base_test

import (
	"testing"

	"github.com/buffreak/domain-scraper/base"
	"github.com/stretchr/testify/assert"
)

func TestGetDomainsDomainMarket(t *testing.T) {
	dm := &base.DomainMarket{}
	domains, err := dm.GetDomains("ao", "1")
	assert.Nil(t, err)
	t.Logf("%+v", domains)
	convert, _ := dm.StructToMap(dm)
	t.Logf("%+v", convert)
}
