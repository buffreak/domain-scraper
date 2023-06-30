package base_test

import (
	"testing"

	"github.com/buffreak/domain-scraper/base"
	"github.com/stretchr/testify/assert"
)

func TestGetDomainsSedo(t *testing.T) {
	sedo := &base.Sedo{}
	domains, err := sedo.GetDomains("apa", "1")
	assert.Nil(t, err)
	t.Logf("%+v", domains)
	convert, _ := sedo.StructToMap(sedo)
	t.Logf("%+v", convert)
}
