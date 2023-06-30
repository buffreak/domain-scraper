package base_test

import (
	"testing"

	"github.com/buffreak/domain-scraper/base"
	"github.com/stretchr/testify/assert"
)

func TestGetDomainsSeoDomains(t *testing.T) {
	sd := &base.SeoDomains{}
	domains, err := sd.GetDomains("1")
	assert.Nil(t, err)
	t.Logf("%+v", domains)
	convert, _ := base.StructToMap(sd)
	t.Logf("%+v", convert)
}
