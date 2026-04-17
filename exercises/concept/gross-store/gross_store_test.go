package gross

import (
	"testing"
)

type entry struct {
	name string
	unit string
	qty  int
}

type unitsTestCase struct {
	description string
	expected    int
}

func TestUnits(t *testing.T) {
	tests := []unitsTestCase{
		{description: "quarter_of_a_dozen", expected: 3},
		{description: "half_of_a_dozen", expected: 6},
		{description: "dozen", expected: 12},
		{description: "small_gross", expected: 120},
		{description: "gross", expected: 144},
		{description: "great_gross", expected: 1728},
	}
	for _, tc := range tests {
		units := Units()
		qty, ok := units[tc.description]

		if !ok {
			t.Errorf("Unit %q not found!", tc.description)
			continue
		}

		if qty != tc.expected {
			t.Errorf("Unit %q should have quantity %d, found %d", tc.description, tc.expected, qty)
		}
	}
}

func TestNewBill(t *testing.T) {
	// Success, zero out the  bill
	bill := NewBill()

	if len(bill) != 0 {
		t.Error("Customer bill must be empty")
	}
}

type addItemTestCase struct {
	description string
	entry       []entry
	expected    bool
}

func TestAddItem(t *testing.T) {
	tests := []addItemTestCase{
		{
			description: "Invalid measurement unit",
			entry: []entry{
				{"pasta", "", 0},
				{"onion", "quarter", 0},
				{"pasta", "pound", 0},
			},
			expected: false,
		},
		{
			description: "Valid measurement unit",
			entry: []entry{
				{"peas", "quarter_of_a_dozen", 3},
				{"tomato", "half_of_a_dozen", 6},
				{"chili", "dozen", 12},
				{"cucumber", "small_gross", 120},
				{"potato", "gross", 144},
				{"zucchini", "great_gross", 1728},
			},
			expected: true,
		},
		{
			description: "check quantity of item added twice",
			entry: []entry{
				{"peas", "quarter_of_a_dozen", 3},
				{"peas", "quarter_of_a_dozen", 6},
				{"tomato", "half_of_a_dozen", 6},
				{"tomato", "quarter_of_a_dozen", 9},
			},
			expected: true,
		},
	}
	units := Units()
	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			bill := NewBill()
			for _, item := range tc.entry {
				ok := AddItem(bill, units, item.name, item.unit)
				if ok != tc.expected {
					t.Errorf("Expected %t from AddItem, found %t at %v", tc.expected, ok, item.name)
				}

				itemQty, ok := bill[item.name]
				if ok != tc.expected {
					t.Errorf("Unexpected item on bill: found %s with quantity %d", item.name, itemQty)
				}

				if itemQty != item.qty {
					t.Errorf("Expected %s to have quantity %d in customer bill, found %d", item.name, item.qty, itemQty)
				}
			}
		})
	}
}

type expectedRemoveItem struct {
	name   string
	unit   string
	qty    int
	exists bool
}

type removeItemTestCase struct {
	description string
	remove      []expectedRemoveItem
	expected    bool
}

func TestRemoveItem(t *testing.T) {
	tests := []removeItemTestCase{
		{
			description: "Item Not found in bill",
			remove: []expectedRemoveItem{
				{"papaya", "gross", 0, false},
			},
			expected: false,
		},
		{
			description: "Invalid measurement unit",
			remove: []expectedRemoveItem{
				{"peas", "pound", 3, true},
				{"tomato", "kilogram", 6, true},
				{"cucumber", "stone", 120, true},
			},
			expected: false,
		},
		{
			description: "Resulted qty less than 0",
			remove: []expectedRemoveItem{
				{"peas", "half_of_a_dozen", 3, true},
				{"tomato", "dozen", 6, true},
				{"chili", "small_gross", 12, true},
				{"cucumber", "gross", 120, true},
				{"potato", "great_gross", 144, true},
			},
			expected: false,
		},
		{
			description: "Should delete the item if 0",
			remove: []expectedRemoveItem{
				{"peas", "quarter_of_a_dozen", 0, false},
				{"tomato", "half_of_a_dozen", 0, false},
				{"chili", "dozen", 0, false},
				{"cucumber", "small_gross", 0, false},
				{"potato", "gross", 0, false},
				{"zucchini", "great_gross", 0, false},
			},
			expected: true,
		},
		{
			description: "Should reduce the qty",
			remove: []expectedRemoveItem{
				{"chili", "half_of_a_dozen", 6, true},
				{"cucumber", "dozen", 108, true},
				{"zucchini", "gross", 1584, true},
			},
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.description, func(t *testing.T) {
			units := Units()
			bill := setupInitialBillData()
			for _, item := range tc.remove {
				ok := RemoveItem(bill, units, item.name, item.unit)
				if ok != tc.expected {
					t.Errorf("Expected %t from RemoveItem, found %t at %v", tc.expected, ok, item.name)
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

type expectedGetItem struct {
	name     string
	expected bool
	qty      int
}

type getItemTestCase struct {
	description string
	getItem     []expectedGetItem
}

func TestGetItem(t *testing.T) {
	test := []getItemTestCase{
		{
			description: "Item Not found in bill",
			getItem: []expectedGetItem{
				{"grape", false, 0},
			},
		},
		{
			description: "Success",
			getItem: []expectedGetItem{
				{"peas", true, 3},
				{"tomato", true, 6},
				{"chili", true, 12},
				{"cucumber", true, 120},
				{"potato", true, 144},
				{"zucchini", true, 1728},
			},
		},
	}

	for _, tc := range test {
		t.Run(tc.description, func(t *testing.T) {
			bill := setupInitialBillData()
			for _, item := range tc.getItem {
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
