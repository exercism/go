package restapi

import (
	"maps"
	"slices"
	"testing"
)

func UserEqual(a, b User) bool {
	return a.Name == b.Name && a.Balance == b.Balance && maps.Equal(a.Owes, b.Owes) && maps.Equal(a.OwedBy, b.OwedBy)
}

func UsersEqual(a, b []User) bool {
	return slices.EqualFunc(a, b, UserEqual)
}

func getApi(data []User) RestApi {
	clone := make([]User, len(data))
	for i, u := range data {
		clone[i] = User{
			Name:    u.Name,
			Balance: u.Balance,
			Owes:    maps.Clone(u.Owes),
			OwedBy:  maps.Clone(u.OwedBy),
		}
	}
	return NewApi(clone)
}

func TestGetUsers(t *testing.T) {
	for _, tc := range getUsersTestCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := getApi(tc.database).GetUsers(tc.payload); !UsersEqual(actual.Users, tc.response.Users) {
				t.Fatalf("api = NewApi(%#v)\napi.GetUsers(%#v)\ngot:  %#v\nwant: %#v", tc.database, tc.payload, actual, tc.response)
			}
		})
	}
}

func TestAddUser(t *testing.T) {
	for _, tc := range addUserTestCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := getApi(tc.database).AddUser(tc.payload); !UserEqual(actual.User, tc.response.User) {
				t.Fatalf("api = NewApi(%#v)\napi.AddUser(%#v)\ngot:  %#v\nwant: %#v", tc.database, tc.payload, actual, tc.response)
			}
		})
	}
}

func TestAddIou(t *testing.T) {
	for _, tc := range addIouTestCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := getApi(tc.database).AddIou(tc.payload); !UsersEqual(actual.Users, tc.response.Users) {
				t.Fatalf("api = NewApi(%#v)\napi.AddIou(%#v)\ngot:  %#v\nwant: %#v", tc.database, tc.payload, actual, tc.response)
			}
		})
	}
}

func BenchmarkGetUsers(b *testing.B) {
	for range b.N {
		for _, tc := range getUsersTestCases {
			getApi(tc.database).GetUsers(tc.payload)
		}
	}
}

func BenchmarkAddUser(b *testing.B) {
	for range b.N {
		for _, tc := range addUserTestCases {
			getApi(tc.database).AddUser(tc.payload)
		}
	}
}

func BenchmarkAddIou(b *testing.B) {
	for range b.N {
		for _, tc := range addIouTestCases {
			getApi(tc.database).AddIou(tc.payload)
		}
	}
}
