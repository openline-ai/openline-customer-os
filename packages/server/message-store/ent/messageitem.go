// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store/ent/messagefeed"
	"github.com/openline-ai/openline-customer-os/packages/server/message-store/ent/messageitem"
)

// MessageItem is the model entity for the MessageItem schema.
type MessageItem struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type messageitem.Type `json:"type,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// Channel holds the value of the "channel" field.
	Channel messageitem.Channel `json:"channel,omitempty"`
	// Direction holds the value of the "direction" field.
	Direction messageitem.Direction `json:"direction,omitempty"`
	// Time holds the value of the "time" field.
	Time time.Time `json:"time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MessageItemQuery when eager-loading is set.
	Edges                     MessageItemEdges `json:"edges"`
	message_feed_message_item *int
}

// MessageItemEdges holds the relations/edges for other nodes in the graph.
type MessageItemEdges struct {
	// MessageFeed holds the value of the message_feed edge.
	MessageFeed *MessageFeed `json:"message_feed,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MessageFeedOrErr returns the MessageFeed value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MessageItemEdges) MessageFeedOrErr() (*MessageFeed, error) {
	if e.loadedTypes[0] {
		if e.MessageFeed == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: messagefeed.Label}
		}
		return e.MessageFeed, nil
	}
	return nil, &NotLoadedError{edge: "message_feed"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MessageItem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case messageitem.FieldID:
			values[i] = new(sql.NullInt64)
		case messageitem.FieldType, messageitem.FieldUsername, messageitem.FieldMessage, messageitem.FieldChannel, messageitem.FieldDirection:
			values[i] = new(sql.NullString)
		case messageitem.FieldTime:
			values[i] = new(sql.NullTime)
		case messageitem.ForeignKeys[0]: // message_feed_message_item
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MessageItem", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MessageItem fields.
func (mi *MessageItem) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case messageitem.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mi.ID = int(value.Int64)
		case messageitem.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				mi.Type = messageitem.Type(value.String)
			}
		case messageitem.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				mi.Username = value.String
			}
		case messageitem.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				mi.Message = value.String
			}
		case messageitem.FieldChannel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field channel", values[i])
			} else if value.Valid {
				mi.Channel = messageitem.Channel(value.String)
			}
		case messageitem.FieldDirection:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field direction", values[i])
			} else if value.Valid {
				mi.Direction = messageitem.Direction(value.String)
			}
		case messageitem.FieldTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field time", values[i])
			} else if value.Valid {
				mi.Time = value.Time
			}
		case messageitem.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field message_feed_message_item", value)
			} else if value.Valid {
				mi.message_feed_message_item = new(int)
				*mi.message_feed_message_item = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryMessageFeed queries the "message_feed" edge of the MessageItem entity.
func (mi *MessageItem) QueryMessageFeed() *MessageFeedQuery {
	return (&MessageItemClient{config: mi.config}).QueryMessageFeed(mi)
}

// Update returns a builder for updating this MessageItem.
// Note that you need to call MessageItem.Unwrap() before calling this method if this MessageItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (mi *MessageItem) Update() *MessageItemUpdateOne {
	return (&MessageItemClient{config: mi.config}).UpdateOne(mi)
}

// Unwrap unwraps the MessageItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mi *MessageItem) Unwrap() *MessageItem {
	_tx, ok := mi.config.driver.(*txDriver)
	if !ok {
		panic("ent: MessageItem is not a transactional entity")
	}
	mi.config.driver = _tx.drv
	return mi
}

// String implements the fmt.Stringer.
func (mi *MessageItem) String() string {
	var builder strings.Builder
	builder.WriteString("MessageItem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mi.ID))
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", mi.Type))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(mi.Username)
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(mi.Message)
	builder.WriteString(", ")
	builder.WriteString("channel=")
	builder.WriteString(fmt.Sprintf("%v", mi.Channel))
	builder.WriteString(", ")
	builder.WriteString("direction=")
	builder.WriteString(fmt.Sprintf("%v", mi.Direction))
	builder.WriteString(", ")
	builder.WriteString("time=")
	builder.WriteString(mi.Time.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// MessageItems is a parsable slice of MessageItem.
type MessageItems []*MessageItem

func (mi MessageItems) config(cfg config) {
	for _i := range mi {
		mi[_i].config = cfg
	}
}
