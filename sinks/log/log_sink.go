// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logsink

import (
	"bytes"
	"fmt"

	kubecommon "github.com/AliyunContainerService/kube-eventer/common/kubernetes"
	"github.com/AliyunContainerService/kube-eventer/core"
	corev1 "k8s.io/api/core/v1"

	"k8s.io/klog"
)

// LogSink ...
type LogSink struct {
}

// Name sink name
func (sink *LogSink) Name() string {
	return "LogSink"
}

// Stop do nothing
func (sink *LogSink) Stop() {
	// Do nothing.
}

func batchToString(batch *core.EventBatch) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("EventBatch     Timestamp: %s\n", batch.Timestamp))
	for _, event := range batch.Events {
		buffer.WriteString(fmt.Sprintf("%++v   %s (cnt:%d): %s\n", event, event.LastTimestamp, event.Count, event.Message))
		if event.InvolvedObject.Kind == "Pod" {
			i, ok, _ := kubecommon.Cache().PodIndexer().GetByKey(fmt.Sprintf("%s/%s", event.Namespace, event.InvolvedObject.Name))
			if ok {
				pod := i.(*corev1.Pod)
				podIP := pod.Status.PodIP
				buffer.WriteString(fmt.Sprintf("PodName: %s, Namespace: %s, PodIP: %s\n", event.InvolvedObject.Name, event.Namespace, podIP))
			}
		}
	}
	return buffer.String()
}

// ExportEvents ...
func (sink *LogSink) ExportEvents(batch *core.EventBatch) {
	klog.Info(batchToString(batch))
}

// CreateLogSink create log sink
func CreateLogSink() (*LogSink, error) {
	return &LogSink{}, nil
}
