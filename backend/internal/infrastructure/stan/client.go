package stan

import (
	"encoding/json"
	"go.uber.org/zap"
	"social-network/internal/config"

	"github.com/nats-io/stan.go"
)

const (
	pingInterval = 60
	pingAttempts = 2 * 60
)

// Client for Nats-streaming.
type Client struct {
	conn stan.Conn
}

// NewClient gets Client instance.
func NewClient(cfg config.StanConfig, logger *zap.Logger) (*Client, error) {
	conn, err := stan.Connect(cfg.ClusterID, "backend-client",
		stan.Pings(pingInterval, pingAttempts),
		stan.NatsURL(cfg.Addr),
		stan.SetConnectionLostHandler(func(_ stan.Conn, reason error) {
			logger.Fatal("connection lost, reason", zap.Error(reason))
		}))
	if err != nil {
		return nil, err
	}

	return &Client{conn: conn}, nil
}

// Publish measurement nats-streaming.
func (c *Client) Publish(subject string, message json.Marshaler) error {
	payload, err := message.MarshalJSON()
	if err != nil {
		return err
	}

	_, err = c.conn.PublishAsync(subject, payload, func(_ string, _ error) {})

	return err
}

// Close connection.
func (c *Client) Close() error {
	return c.conn.Close()
}
