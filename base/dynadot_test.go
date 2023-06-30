package base_test

import (
	"testing"

	"github.com/buffreak/domain-scraper/base"
	"github.com/stretchr/testify/assert"
)

func TestGetDomainsDynadot(t *testing.T) {
	dn := &base.Dynadot{}
	domains, err := dn.GetDomains("0")
	assert.Nil(t, err)
	t.Logf("%+v", domains)
	convert, _ := base.StructToMap(dn)
	t.Logf("%+v", convert)
}
