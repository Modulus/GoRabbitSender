package sender

import (
	"testing"
)

func TestConstructorEmptyFilePathShouldSetDefault(t *testing.T) {
	sender := NewSender("")

	if sender.Config.ElasticsearchConnectionString != "" && len(sender.Config.ElasticsearchConnectionString) < 0 {
		t.Fatalf("Should have default configuration and ElasticsearchConnectionString value was %v\n", sender.Config.ElasticsearchConnectionString)
	}

	if sender.Config.RabbitmqConnectionString != "" && len(sender.Config.RabbitmqConnectionString) < 0 {
		t.Fatalf("Should have default configuration and RabbitmqConnectionString")
	}

	if sender.Config.RabbitmqExchangeName != "" && len(sender.Config.RabbitmqExchangeName) < 0 {
		t.Fatalf("Should have default configuration and RabbitmqExchangeName")
	}

	if sender.Config.RabbitmqQueueName != "" && len(sender.Config.RabbitmqQueueName) < 0 {
		t.Fatalf("Should have default configuration and RabbitmqQueueName")
	}
}
