// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package kafka

import (
	"testing"

	"github.com/dapr/components-contrib/pubsub"
	"github.com/stretchr/testify/assert"
)

func TestParseMetadata(t *testing.T) {
	m := pubsub.Metadata{}
	m.Properties = map[string]string{"consumerGroup": "a", "brokers": "a", "topics": "a"}
	k := Kafka{}
	meta, err := k.getKafkaMetadata(m)
	assert.Nil(t, err)
	assert.Equal(t, "a", meta.Brokers[0])
	assert.Equal(t, "a", meta.ConsumerGroup)
	assert.Equal(t, "a", meta.Topics[0])
}
