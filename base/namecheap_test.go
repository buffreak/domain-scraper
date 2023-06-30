package base_test

import (
	"testing"

	"github.com/buffreak/domain-scraper/base"
	"github.com/stretchr/testify/assert"
)

func TestGetDomainsNameCheap(t *testing.T) {
	nc := &base.NameCheap{}
	domains, err := nc.GetDomains("1")
	assert.Nil(t, err)
	t.Logf("%+v", domains)
	convert, _ := nc.StructToMap(nc)
	t.Logf("%+v", convert)
}
