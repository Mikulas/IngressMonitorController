package v3

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGroupMembers(t *testing.T) {
	uptrends := MakeUptrends("dite@mangoweb.cz", "FiDYWsF;jAg=grV779$pjKYY")

	// setup

	group := MakeGroup("membership test")
	err := uptrends.AddGroup(group)
	assert.Nil(t, err)

	monitor := MakeMonitor("example", "https://example.com/")
	err = uptrends.AddMonitor(monitor)
	assert.Nil(t, err)
	assert.NotEmpty(t, monitor.Guid)

	// memberships

	err = uptrends.AddMember(group, monitor)
	assert.Nil(t, err)

	members, err := uptrends.GetMembers(group)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(members))

	err = uptrends.RemoveMember(group, monitor)
	assert.Nil(t, err)

	members, err = uptrends.GetMembers(group)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(members))

	// cleanup

	err = uptrends.DeleteMonitor(monitor.Guid)
	assert.Nil(t, err)

	err = uptrends.DeleteGroup(group.Guid)
	assert.Nil(t, err)
}
