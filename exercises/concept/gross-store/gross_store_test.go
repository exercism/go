package gross

import (
	"testing"
)

type entry struct {
	name string
	unit string
	qty  int
}

func TestUnits(t *testing.T) {
	tests := []struct {
		name string
		qty  int
	}{
		{"quarter_of_a_dozen", 3},
		{"half_of_a_dozen", 6},
		{"dozen", 12},
		{"small_gross", 120},
		{"gross", 144},
		{"great_gross", 1728},
	}

	units := Units()
	for _, tt := range tests {
		qty, ok := units[tt.name]

		if !ok {
			t.Errorf(`Unit "%s" not found!`, tt.name)
			continue
		}

		if qty != tt.qty {
			t.Errorf(`Unit "%s" should have quantity %d, found %d`, tt.name, tt.qty, qty)
		}
	}

}

func TestNewBill(t *testing.T) {
	// Success, zero out the  bill
	t.Run("Should reset customerbill", func(t *testing.T) {
		bill := NewBill()

		if len(bill) != 0 {
			t.Error("Customer bill must be empty")
		}
	})
}

func TestAddItem(t *testing.T) {
	tests := []struct {
		name     string
		entry    []entry
		expected bool
	}{
		{
			"Invalid measurement unit",
			[]entry{
				{"pasta", "", 0},
				{"onion", "quarter", 0},
				{"pasta", "pound", 0},
			},
			false,
		},
		{
			"Valid measurement unit",
			[]entry{
				{"peas", "quarter_of_a_dozen", 3},
				{"tomato", "half_of_a_dozen", 6},
				{"chili", "dozen", 12},
				{"cucumber", "small_gross", 120},
				{"potato", "gross", 144},
				{"zucchini", "great_gross", 1728},
			},
			true,
		},
		{
			"check quantity of item added twice",
			[]entry{
				{"peas", "quarter_of_a_dozen", 3},
				{"peas", "quarter_of_a_dozen", 6},
				{"tomato", "half_of_a_dozen", 6},
				{"tomato", "quarter_of_a_dozen", 9},
			},
			true,
		},
	}
	units := Units()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bill := NewBill()
			for _, item := range tt.entry {
				ok := AddItem(bill, units, item.name, item.unit)
				if ok != tt.expected {
					t.Errorf("Expected %t from AddItem, found %t at %v", tt.expected, ok, item.name)
				}

				itemQty, ok := bill[item.name]
				if ok != tt.expected {
					t.Errorf("Unexpected item on bill: found %s with quantity %d", item.name, itemQty)
				}

				if itemQty != item.qty {
					t.Errorf("Expected %s to have quantity %d in customer bill, found %d", item.name, item.qty, itemQty)
				}
			}
		})
	}
}

func TestRemoveItem(t *testing.T) {
	type expectedItem struct {
		name   string
		unit   string
		qty    int
		exists bool
	}

	tests := []struct {
		name     string
		remove   []expectedItem
		expected bool
	}{
		{"Item Not found in bill",
			[]expectedItem{
				{"papaya", "gross", 0, false},
			},
			false,
		},
		{"Invalid measurement unit",
			[]expectedItem{
				{"peas", "pound", 3, true},
				{"tomato", "kilogram", 6, true},
				{"cucumber", "stone", 120, true},
			},
			false,
		},
		{"Resulted qty less than 0",
			[]expectedItem{
				{"peas", "half_of_a_dozen", 3, true},
				{"tomato", "dozen", 6, true},
				{"chili", "small_gross", 12, true},
				{"cucumber", "gross", 120, true},
				{"potato", "great_gross", 144, true},
			},
			false,
		},
		{"Should delete the item if 0",
			[]expectedItem{
				{"peas", "quarter_of_a_dozen", 0, false},
				{"tomato", "half_of_a_dozen", 0, false},
				{"chili", "dozen", 0, false},
				{"cucumber", "small_gross", 0, false},
				{"potato", "gross", 0, false},
				{"zucchini", "great_gross", 0, false},
			},
			true,
		},
		{"Should reduce the qty",
			[]expectedItem{
				{"chili", "half_of_a_dozen", 6, true},
				{"cucumber", "dozen", 108, true},
				{"zucchini", "gross", 1584, true},
			},
			true,
		},
	}

	units := Units()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bill := setupInitialBillData()
			for _, item := range tt.remove {
				ok := RemoveItem(bill, units, item.name, item.unit)
				if ok != tt.expected {
					t.Errorf("Expected %t from RemoveItem, found %t at %v", tt.expected, ok, item.name)
				}

				itemQty, ok := bill[item.name]
				if ok != item.exists {
					t.Errorf("Could not find item %s in customer bill", item.name)
				}
				if itemQty != item.qty {
					t.Errorf("Expected %s to have quantity %d in customer bill, found %d", item.name, item.qty, itemQty)
				}
			}
		})
	}
}

func TestGetItem(t *testing.T) {
	type expectedItem struct {
		name     string
		expected bool
		qty      int
	}

	test := []struct {
		name    string
		getItem []expectedItem
	}{
		{
			"Item Not found in bill",
			[]expectedItem{
				{"grape", false, 0},
			},
		},
		{
			"Success",
			[]expectedItem{
				{"peas", true, 3},
				{"tomato", true, 6},
				{"chili", true, 12},
				{"cucumber", true, 120},
				{"potato", true, 144},
				{"zucchini", true, 1728},
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			bill := setupInitialBillData()
			for _, item := range tt.getItem {
				itemQty, ok := GetItem(bill, item.name)

				if ok != item.expected {
					msg := "Could not find item %s in customer bill, expected %t"
					if item.expected == false {
						msg = "Found item %s in customer bill, expected %t"
					}

					t.Errorf(msg, item.name, item.expected)
				}

				if itemQty != item.qty {
					t.Errorf("Expected %s to have quantity %d in customer bill, found %d", item.name, item.qty, itemQty)
				}
			}
		})
	}
}

func setupInitialBillData() map[string]int {
	bill := NewBill()
	bill["peas"] = 3
	bill["tomato"] = 6
	bill["chili"] = 12
	bill["cucumber"] = 120
	bill["potato"] = 144
	bill["zucchini"] = 1728
	return bill
}
