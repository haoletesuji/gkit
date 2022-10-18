package common

import "testing"

func TestGetServerAddrs(t *testing.T) {
	tests := []struct {
		addrs string
		n     int
	}{
		{"", 0},
		{",", 0},
		{" , ", 0},
		{"cassandra-host-0", 1},
		{"cassandra-host-0,cassandra-host-1, ", 2},
		{"cassandra-host-0,cassandra-host-1,cassandra-host-2", 3},
	}

	for _, test := range tests {
		addrs := GetServerAddrs(test.addrs)
		if len(addrs) != test.n {
			t.Errorf("GetServerAddrs was incorrect, got: %d, want: %d.", len(addrs), test.n)
		}
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		inputs []string
		result string
	}{
		{[]string{""}, ""},
		{[]string{"", "a"}, "a"},
		{[]string{"a", "b"}, "ab"},
	}

	for _, test := range tests {
		s := Join(test.inputs...)
		if s != test.result {
			t.Errorf("Join was incorrect, got: %s, want: %s.", s, test.result)
		}
	}
}

func TestGetPageNumber(t *testing.T) {
	tests := []struct {
		input  string
		result int
	}{
		{"", 0},
		{"a", 0},
		{"1", 0},
		{"10", 9},
	}

	for _, test := range tests {
		s := GetPageNumber(test.input)
		if s != test.result {
			t.Errorf("GetPageNumber was incorrect, got: %d, want: %d.", s, test.result)
		}
	}
}
