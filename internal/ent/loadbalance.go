// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-gateway/internal/ent/loadbalance"
	"strings"

	"github.com/facebook/ent/dialect/sql"
)

// LoadBalance is the model entity for the LoadBalance schema.
type LoadBalance struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// ServiceID holds the value of the "service_id" field.
	ServiceID int64 `json:"service_id,omitempty"`
	// CheckMethod holds the value of the "check_method" field.
	CheckMethod int `json:"check_method,omitempty"`
	// CheckTimeout holds the value of the "check_timeout" field.
	CheckTimeout int `json:"check_timeout,omitempty"`
	// CheckInterval holds the value of the "check_interval" field.
	CheckInterval int `json:"check_interval,omitempty"`
	// RoundType holds the value of the "round_type" field.
	RoundType int `json:"round_type,omitempty"`
	// IPList holds the value of the "ip_list" field.
	IPList string `json:"ip_list,omitempty"`
	// WeightList holds the value of the "weight_list" field.
	WeightList string `json:"weight_list,omitempty"`
	// ForbidList holds the value of the "forbid_list" field.
	ForbidList string `json:"forbid_list,omitempty"`
	// UpstreamConnectTimeout holds the value of the "upstream_connect_timeout" field.
	UpstreamConnectTimeout int `json:"upstream_connect_timeout,omitempty"`
	// UpstreamHeaderTimeout holds the value of the "upstream_header_timeout" field.
	UpstreamHeaderTimeout int `json:"upstream_header_timeout,omitempty"`
	// UpstreamIdleTimeout holds the value of the "upstream_idle_timeout" field.
	UpstreamIdleTimeout int `json:"upstream_idle_timeout,omitempty"`
	// UpstreamMaxIdle holds the value of the "upstream_max_idle" field.
	UpstreamMaxIdle int `json:"upstream_max_idle,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LoadBalance) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // service_id
		&sql.NullInt64{},  // check_method
		&sql.NullInt64{},  // check_timeout
		&sql.NullInt64{},  // check_interval
		&sql.NullInt64{},  // round_type
		&sql.NullString{}, // ip_list
		&sql.NullString{}, // weight_list
		&sql.NullString{}, // forbid_list
		&sql.NullInt64{},  // upstream_connect_timeout
		&sql.NullInt64{},  // upstream_header_timeout
		&sql.NullInt64{},  // upstream_idle_timeout
		&sql.NullInt64{},  // upstream_max_idle
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LoadBalance fields.
func (lb *LoadBalance) assignValues(values ...interface{}) error {
	if m, n := len(values), len(loadbalance.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	lb.ID = int64(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field service_id", values[0])
	} else if value.Valid {
		lb.ServiceID = value.Int64
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field check_method", values[1])
	} else if value.Valid {
		lb.CheckMethod = int(value.Int64)
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field check_timeout", values[2])
	} else if value.Valid {
		lb.CheckTimeout = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field check_interval", values[3])
	} else if value.Valid {
		lb.CheckInterval = int(value.Int64)
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field round_type", values[4])
	} else if value.Valid {
		lb.RoundType = int(value.Int64)
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field ip_list", values[5])
	} else if value.Valid {
		lb.IPList = value.String
	}
	if value, ok := values[6].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field weight_list", values[6])
	} else if value.Valid {
		lb.WeightList = value.String
	}
	if value, ok := values[7].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field forbid_list", values[7])
	} else if value.Valid {
		lb.ForbidList = value.String
	}
	if value, ok := values[8].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field upstream_connect_timeout", values[8])
	} else if value.Valid {
		lb.UpstreamConnectTimeout = int(value.Int64)
	}
	if value, ok := values[9].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field upstream_header_timeout", values[9])
	} else if value.Valid {
		lb.UpstreamHeaderTimeout = int(value.Int64)
	}
	if value, ok := values[10].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field upstream_idle_timeout", values[10])
	} else if value.Valid {
		lb.UpstreamIdleTimeout = int(value.Int64)
	}
	if value, ok := values[11].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field upstream_max_idle", values[11])
	} else if value.Valid {
		lb.UpstreamMaxIdle = int(value.Int64)
	}
	return nil
}

// Update returns a builder for updating this LoadBalance.
// Note that, you need to call LoadBalance.Unwrap() before calling this method, if this LoadBalance
// was returned from a transaction, and the transaction was committed or rolled back.
func (lb *LoadBalance) Update() *LoadBalanceUpdateOne {
	return (&LoadBalanceClient{config: lb.config}).UpdateOne(lb)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (lb *LoadBalance) Unwrap() *LoadBalance {
	tx, ok := lb.config.driver.(*txDriver)
	if !ok {
		panic("ent: LoadBalance is not a transactional entity")
	}
	lb.config.driver = tx.drv
	return lb
}

// String implements the fmt.Stringer.
func (lb *LoadBalance) String() string {
	var builder strings.Builder
	builder.WriteString("LoadBalance(")
	builder.WriteString(fmt.Sprintf("id=%v", lb.ID))
	builder.WriteString(", service_id=")
	builder.WriteString(fmt.Sprintf("%v", lb.ServiceID))
	builder.WriteString(", check_method=")
	builder.WriteString(fmt.Sprintf("%v", lb.CheckMethod))
	builder.WriteString(", check_timeout=")
	builder.WriteString(fmt.Sprintf("%v", lb.CheckTimeout))
	builder.WriteString(", check_interval=")
	builder.WriteString(fmt.Sprintf("%v", lb.CheckInterval))
	builder.WriteString(", round_type=")
	builder.WriteString(fmt.Sprintf("%v", lb.RoundType))
	builder.WriteString(", ip_list=")
	builder.WriteString(lb.IPList)
	builder.WriteString(", weight_list=")
	builder.WriteString(lb.WeightList)
	builder.WriteString(", forbid_list=")
	builder.WriteString(lb.ForbidList)
	builder.WriteString(", upstream_connect_timeout=")
	builder.WriteString(fmt.Sprintf("%v", lb.UpstreamConnectTimeout))
	builder.WriteString(", upstream_header_timeout=")
	builder.WriteString(fmt.Sprintf("%v", lb.UpstreamHeaderTimeout))
	builder.WriteString(", upstream_idle_timeout=")
	builder.WriteString(fmt.Sprintf("%v", lb.UpstreamIdleTimeout))
	builder.WriteString(", upstream_max_idle=")
	builder.WriteString(fmt.Sprintf("%v", lb.UpstreamMaxIdle))
	builder.WriteByte(')')
	return builder.String()
}

// LoadBalances is a parsable slice of LoadBalance.
type LoadBalances []*LoadBalance

func (lb LoadBalances) config(cfg config) {
	for _i := range lb {
		lb[_i].config = cfg
	}
}