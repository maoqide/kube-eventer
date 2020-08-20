package wechatgroup

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGetLevel(t *testing.T) {
	warning := getLevel(v1.EventTypeWarning)
	normal := getLevel(v1.EventTypeNormal)
	none := getLevel("")
	assert.True(t, warning > normal)
	assert.True(t, warning == WARNING)
	assert.True(t, normal == NORMAL)
	assert.True(t, 0 == none)
}

func TestCreateMsgFromEvent(t *testing.T) {
	labels := make([]string, 2)
	labels[0] = "abcd"
	labels[1] = "defg"
	event := createTestEvent()
	u, _ := url.Parse("wechatgroup:?corp_id=xxxxxxxxxxxxxx&corp_secret=xxxxxxxxxx&agent_id=1111111&chat_id=1212121212121212&label=k8s&level=Warning")
	d, _ := NewWechatGroupSink(u)
	msg := createMsgFromEvent(d, event)
	text, _ := json.Marshal(msg)
	t.Log(string(text))
	// d.Send(event)
	t.Log(msg.Text)
	assert.True(t, msg != nil)
}

func createTestEvent() *v1.Event {
	now := time.Now()
	event := &v1.Event{
		Reason:         "test",
		Message:        "just for just",
		Count:          99,
		LastTimestamp:  metav1.NewTime(now),
		FirstTimestamp: metav1.NewTime(now),
		Type:           "Warning",
	}
	return event
}
