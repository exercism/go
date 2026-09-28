package sqlbuilder

import (
	"reflect"
	"testing"
)

func TestBuildSelectQuery(t *testing.T) {
	tests := []struct {
		name    string
		table   string
		columns []string
		want    string
	}{
		{
			name:    "Basic table with multiple columns",
			table:   "users",
			columns: []string{"id", "name", "email"},
			want:    "SELECT id, name, email FROM users;",
		},
		{
			name:    "Single column",
			table:   "products",
			columns: []string{"price"},
			want:    "SELECT price FROM products;",
		},
		{
			name:    "Empty columns slice should fallback to wildcard",
			table:   "orders",
			columns: []string{},
			want:    "SELECT * FROM orders;",
		},
		{
			name:    "Nil columns slice should fallback to wildcard",
			table:   "logs",
			columns: nil,
			want:    "SELECT * FROM logs;",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildSelectQuery(tt.table, tt.columns)
			if got != tt.want {
				t.Errorf("BuildSelectQuery(%q, %v) = %q; want %q", tt.table, tt.columns, got, tt.want)
			}
		})
	}
}

func TestBuildBatchInserts(t *testing.T) {
	tests := []struct {
		name  string
		table string
		rows  []map[string]string
		want  []string
	}{
		{
			name:  "Empty rows slice",
			table: "users",
			rows:  []map[string]string{},
			want:  []string{},
		},
		{
			name:  "Single row single column",
			table: "tags",
			rows: []map[string]string{
				{"name": "golang"},
			},
			want: []string{
				"INSERT INTO tags (name) VALUES ('golang');",
			},
		},
		{
			name:  "Multiple rows with ordered execution",
			table: "users",
			rows: []map[string]string{
				{"name": "Alice"},
				{"name": "Bob"},
			},
			want: []string{
				"INSERT INTO users (name) VALUES ('Alice');",
				"INSERT INTO users (name) VALUES ('Bob');",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildBatchInserts(tt.table, tt.rows)
			if len(tt.want) == 0 && len(got) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildBatchInserts(%q, %v)\ngot:  %v\nwant: %v", tt.table, tt.rows, got, tt.want)
			}
		})
	}
}
