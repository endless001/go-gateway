// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-gateway/ent/tcprule"
	"strings"

	"github.com/facebook/ent/dialect/sql"
)

// TcpRule is the model entity for the TcpRule schema.
type TcpRule struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// ServiceID holds the value of the "service_id" field.
	ServiceID int64 `json:"service_id,omitempty"`
	// Port holds the value of the "Port" field.
	Port int `json:"Port,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TcpRule) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullInt64{}, // service_id
		&sql.NullInt64{}, // Port
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TcpRule fields.
func (tr *TcpRule) assignValues(values ...interface{}) error {
	if m, n := len(values), len(tcprule.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	tr.ID = int64(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field service_id", values[0])
	} else if value.Valid {
		tr.ServiceID = value.Int64
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field Port", values[1])
	} else if value.Valid {
		tr.Port = int(value.Int64)
	}
	return nil
}

// Update returns a builder for updating this TcpRule.
// Note that, you need to call TcpRule.Unwrap() before calling this method, if this TcpRule
// was returned from a transaction, and the transaction was committed or rolled back.
func (tr *TcpRule) Update() *TcpRuleUpdateOne {
	return (&TcpRuleClient{config: tr.config}).UpdateOne(tr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (tr *TcpRule) Unwrap() *TcpRule {
	tx, ok := tr.config.driver.(*txDriver)
	if !ok {
		panic("ent: TcpRule is not a transactional entity")
	}
	tr.config.driver = tx.drv
	return tr
}

// String implements the fmt.Stringer.
func (tr *TcpRule) String() string {
	var builder strings.Builder
	builder.WriteString("TcpRule(")
	builder.WriteString(fmt.Sprintf("id=%v", tr.ID))
	builder.WriteString(", service_id=")
	builder.WriteString(fmt.Sprintf("%v", tr.ServiceID))
	builder.WriteString(", Port=")
	builder.WriteString(fmt.Sprintf("%v", tr.Port))
	builder.WriteByte(')')
	return builder.String()
}

// TcpRules is a parsable slice of TcpRule.
type TcpRules []*TcpRule

func (tr TcpRules) config(cfg config) {
	for _i := range tr {
		tr[_i].config = cfg
	}
}
