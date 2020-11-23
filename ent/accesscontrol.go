// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-gateway/ent/accesscontrol"
	"strings"

	"github.com/facebook/ent/dialect/sql"
)

// AccessControl is the model entity for the AccessControl schema.
type AccessControl struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// ServiceID holds the value of the "service_id" field.
	ServiceID int64 `json:"service_id,omitempty"`
	// OpenAuth holds the value of the "open_auth" field.
	OpenAuth int `json:"open_auth,omitempty"`
	// BlackList holds the value of the "black_list" field.
	BlackList string `json:"black_list,omitempty"`
	// WhiteList holds the value of the "white_list" field.
	WhiteList string `json:"white_list,omitempty"`
	// WhiteHostName holds the value of the "white_host_name" field.
	WhiteHostName string `json:"white_host_name,omitempty"`
	// ClientipFlowLimit holds the value of the "clientip_flow_limit" field.
	ClientipFlowLimit int `json:"clientip_flow_limit,omitempty"`
	// ServiceFlowLimit holds the value of the "service_flow_limit" field.
	ServiceFlowLimit int `json:"service_flow_limit,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AccessControl) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // service_id
		&sql.NullInt64{},  // open_auth
		&sql.NullString{}, // black_list
		&sql.NullString{}, // white_list
		&sql.NullString{}, // white_host_name
		&sql.NullInt64{},  // clientip_flow_limit
		&sql.NullInt64{},  // service_flow_limit
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AccessControl fields.
func (ac *AccessControl) assignValues(values ...interface{}) error {
	if m, n := len(values), len(accesscontrol.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	ac.ID = int64(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field service_id", values[0])
	} else if value.Valid {
		ac.ServiceID = value.Int64
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field open_auth", values[1])
	} else if value.Valid {
		ac.OpenAuth = int(value.Int64)
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field black_list", values[2])
	} else if value.Valid {
		ac.BlackList = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field white_list", values[3])
	} else if value.Valid {
		ac.WhiteList = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field white_host_name", values[4])
	} else if value.Valid {
		ac.WhiteHostName = value.String
	}
	if value, ok := values[5].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field clientip_flow_limit", values[5])
	} else if value.Valid {
		ac.ClientipFlowLimit = int(value.Int64)
	}
	if value, ok := values[6].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field service_flow_limit", values[6])
	} else if value.Valid {
		ac.ServiceFlowLimit = int(value.Int64)
	}
	return nil
}

// Update returns a builder for updating this AccessControl.
// Note that, you need to call AccessControl.Unwrap() before calling this method, if this AccessControl
// was returned from a transaction, and the transaction was committed or rolled back.
func (ac *AccessControl) Update() *AccessControlUpdateOne {
	return (&AccessControlClient{config: ac.config}).UpdateOne(ac)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ac *AccessControl) Unwrap() *AccessControl {
	tx, ok := ac.config.driver.(*txDriver)
	if !ok {
		panic("ent: AccessControl is not a transactional entity")
	}
	ac.config.driver = tx.drv
	return ac
}

// String implements the fmt.Stringer.
func (ac *AccessControl) String() string {
	var builder strings.Builder
	builder.WriteString("AccessControl(")
	builder.WriteString(fmt.Sprintf("id=%v", ac.ID))
	builder.WriteString(", service_id=")
	builder.WriteString(fmt.Sprintf("%v", ac.ServiceID))
	builder.WriteString(", open_auth=")
	builder.WriteString(fmt.Sprintf("%v", ac.OpenAuth))
	builder.WriteString(", black_list=")
	builder.WriteString(ac.BlackList)
	builder.WriteString(", white_list=")
	builder.WriteString(ac.WhiteList)
	builder.WriteString(", white_host_name=")
	builder.WriteString(ac.WhiteHostName)
	builder.WriteString(", clientip_flow_limit=")
	builder.WriteString(fmt.Sprintf("%v", ac.ClientipFlowLimit))
	builder.WriteString(", service_flow_limit=")
	builder.WriteString(fmt.Sprintf("%v", ac.ServiceFlowLimit))
	builder.WriteByte(')')
	return builder.String()
}

// AccessControls is a parsable slice of AccessControl.
type AccessControls []*AccessControl

func (ac AccessControls) config(cfg config) {
	for _i := range ac {
		ac[_i].config = cfg
	}
}
