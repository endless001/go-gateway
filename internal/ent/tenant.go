// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-gateway/internal/ent/tenant"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
)

// Tenant is the model entity for the Tenant schema.
type Tenant struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// AppID holds the value of the "app_id" field.
	AppID string `json:"app_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Secret holds the value of the "secret" field.
	Secret string `json:"secret,omitempty"`
	// WhiteIps holds the value of the "white_ips" field.
	WhiteIps string `json:"white_ips,omitempty"`
	// Qpd holds the value of the "qpd" field.
	Qpd int `json:"qpd,omitempty"`
	// QPS holds the value of the "Qps" field.
	QPS int `json:"Qps,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt time.Time `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt time.Time `json:"update_at,omitempty"`
	// IsDelete holds the value of the "is_delete" field.
	IsDelete int8 `json:"is_delete,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Tenant) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // app_id
		&sql.NullString{}, // name
		&sql.NullString{}, // secret
		&sql.NullString{}, // white_ips
		&sql.NullInt64{},  // qpd
		&sql.NullInt64{},  // Qps
		&sql.NullTime{},   // create_at
		&sql.NullTime{},   // update_at
		&sql.NullInt64{},  // is_delete
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Tenant fields.
func (t *Tenant) assignValues(values ...interface{}) error {
	if m, n := len(values), len(tenant.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	t.ID = int64(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field app_id", values[0])
	} else if value.Valid {
		t.AppID = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[1])
	} else if value.Valid {
		t.Name = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field secret", values[2])
	} else if value.Valid {
		t.Secret = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field white_ips", values[3])
	} else if value.Valid {
		t.WhiteIps = value.String
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field qpd", values[4])
	} else if value.Valid {
		t.Qpd = int(value.Int64)
	}
	if value, ok := values[5].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field Qps", values[5])
	} else if value.Valid {
		t.QPS = int(value.Int64)
	}
	if value, ok := values[6].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_at", values[6])
	} else if value.Valid {
		t.CreateAt = value.Time
	}
	if value, ok := values[7].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field update_at", values[7])
	} else if value.Valid {
		t.UpdateAt = value.Time
	}
	if value, ok := values[8].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field is_delete", values[8])
	} else if value.Valid {
		t.IsDelete = int8(value.Int64)
	}
	return nil
}

// Update returns a builder for updating this Tenant.
// Note that, you need to call Tenant.Unwrap() before calling this method, if this Tenant
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Tenant) Update() *TenantUpdateOne {
	return (&TenantClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (t *Tenant) Unwrap() *Tenant {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Tenant is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Tenant) String() string {
	var builder strings.Builder
	builder.WriteString("Tenant(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", app_id=")
	builder.WriteString(t.AppID)
	builder.WriteString(", name=")
	builder.WriteString(t.Name)
	builder.WriteString(", secret=")
	builder.WriteString(t.Secret)
	builder.WriteString(", white_ips=")
	builder.WriteString(t.WhiteIps)
	builder.WriteString(", qpd=")
	builder.WriteString(fmt.Sprintf("%v", t.Qpd))
	builder.WriteString(", Qps=")
	builder.WriteString(fmt.Sprintf("%v", t.QPS))
	builder.WriteString(", create_at=")
	builder.WriteString(t.CreateAt.Format(time.ANSIC))
	builder.WriteString(", update_at=")
	builder.WriteString(t.UpdateAt.Format(time.ANSIC))
	builder.WriteString(", is_delete=")
	builder.WriteString(fmt.Sprintf("%v", t.IsDelete))
	builder.WriteByte(')')
	return builder.String()
}

// Tenants is a parsable slice of Tenant.
type Tenants []*Tenant

func (t Tenants) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
