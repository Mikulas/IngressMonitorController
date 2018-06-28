package v3

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGroups(t *testing.T) {
	uptrends := MakeUptrends("dite@mangoweb.cz", "FiDYWsF;jAg=grV779$pjKYY")

	count := func() int {
		monitors, err := uptrends.GetGroups()
		assert.Nil(t, err)
		return len(monitors)
	}

	get := func(guid string) Group {
		group, err := uptrends.GetGroup(guid)
		assert.Nil(t, err)
		return group
	}

	count1 := count()
	group := MakeGroup("example")
	err := uptrends.AddGroup(group)
	assert.Nil(t, err)
	assert.NotEmpty(t, group.Guid)

	count2 := count()
	assert.Equal(t, count1 + 1, count2, "expected exactly 1 new group")

	group.Name = "example updated"
	err = uptrends.EditGroup(group)
	assert.Nil(t, err)
	assert.Equal(t, group.Name, get(group.Guid).Name)

	err = uptrends.DeleteGroup(group.Guid)
	assert.Nil(t, err)
	count3 := count()
	assert.Equal(t, count1, count3, "expected exactly 1 deleted group")
}
