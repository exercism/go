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
		rows  [][]ColumnValue
		want  []string
	}{
		{
			name:  "Empty rows slice",
			table: "users",
			rows:  [][]ColumnValue{},
			want:  []string{},
		},
		{
			name:  "Single row single column",
			table: "tags",
			rows: [][]ColumnValue{
				{{Column: "name", Value: "golang"}},
			},
			want: []string{
				"INSERT INTO tags (name) VALUES ('golang');",
			},
		},
		{
			name:  "Multiple rows with ordered execution",
			table: "users",
			rows: [][]ColumnValue{
				{{Column: "name", Value: "Alice"}},
				{{Column: "name", Value: "Bob"}},
			},
			want: []string{
				"INSERT INTO users (name) VALUES ('Alice');",
				"INSERT INTO users (name) VALUES ('Bob');",
			},
		},
		{
			name:  "Multiple columns preserve precise slice order",
			table: "products",
			rows: [][]ColumnValue{
				// Порядок гарантирован самим слайсом, сортировка не нужна!
				{
					{Column: "title", Value: "Book"},
					{Column: "price", Value: "100"},
				},
			},
			want: []string{
				"INSERT INTO products (title, price) VALUES ('Book', '100');",
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

// --- SELECT QUERY BENCHMARKS ---

// BenchmarkBuildSelectQuery_Efficient benchmarks the optimized solution
// that pre-allocates memory using Grow() to minimize heap allocations.
func BenchmarkBuildSelectQuery_Efficient(b *testing.B) {
	columns := []string{"id", "name", "email", "created_at", "updated_at", "status", "role"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = BuildSelectQuery("users", columns)
	}
}

// --- BATCH INSERT BENCHMARKS ---

// BenchmarkBuildBatchInserts_Efficient benchmarks the optimized solution
// that declares a single strings.Builder and reuses it across loop
// iterations via Reset() to preserve allocated capacity.
func BenchmarkBuildBatchInserts_Efficient(b *testing.B) {
	rows := [][]ColumnValue{
		{{Column: "name", Value: "Alice"}, {Column: "email", Value: "alice@example.com"}},
		{{Column: "name", Value: "Bob"}, {Column: "email", Value: "bob@example.com"}},
		{{Column: "name", Value: "Charlie"}, {Column: "email", Value: "charlie@example.com"}},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = BuildBatchInserts("users", rows)
	}
}
