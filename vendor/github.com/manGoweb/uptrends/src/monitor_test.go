package v3

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMonitors(t *testing.T) {
	uptrends := MakeUptrends("dite@mangoweb.cz", "FiDYWsF;jAg=grV779$pjKYY")

	count := func() int {
		monitors, err := uptrends.GetMonitors()
		assert.Nil(t, err)
		return len(monitors)
	}

	get := func(guid string) Monitor {
		monitor, err := uptrends.GetMonitor(guid)
		assert.Nil(t, err)
		return monitor
	}

	count1 := count()
	m := MakeMonitor("example", "https://example.com/")
	err := uptrends.AddMonitor(m)
	assert.Nil(t, err)
	assert.NotEmpty(t, m.Guid)

	count2 := count()
	assert.Equal(t, count1 + 1, count2, "expected exactly 1 new monitor")

	m.URL = "https://example.com/update"
	err = uptrends.EditMonitor(m)
	assert.Nil(t, err)
	assert.Equal(t, m.URL, get(m.Guid).URL)

	err = uptrends.DeleteMonitor(m.Guid)
	assert.Nil(t, err)
	count3 := count()
	assert.Equal(t, count1, count3, "expected exactly 1 deleted monitor")
}
