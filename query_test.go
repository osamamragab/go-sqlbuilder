package sqlbuilder

import "testing"

func TestSelect(t *testing.T) {
	q := NewQuery("test")
	q.Select("t1", "t2", "t3")

	gotStr := q.String()
	wantStr := "SELECT t1,t2,t3 FROM test"
	gotArgs := q.Args()

	if gotStr != wantStr {
		t.Errorf("Select string: want %q, got %q", wantStr, gotStr)
	}
	if gotArgs != nil {
		t.Errorf("Select arguments: want <nil>, got %v", gotArgs)
	}

	q.SetTable("test2")
	q.Select("t1", "t2", "t3").Where("id = ? AND name = ?", 5, "myname")

	gotStr = q.String()
	wantStr = "SELECT t1,t2,t3 FROM test2 WHERE id = $1 AND name = $2"
	gotArgs = q.Args()
	wantArgs := []interface{}{5, "myname"}

	if gotStr != wantStr {
		t.Errorf("Select with Where string: want %q, got %q", wantStr, gotStr)
	}
	if len(gotArgs) != len(wantArgs) {
		t.Errorf("Select with Where arguments length: want %d, got %d", len(wantArgs), len(gotArgs))
	}
	for i := range wantArgs {
		if wantArgs[i] != gotArgs[i] {
			t.Errorf("Select with Where arguments[%d]: want %v, got %v", i, wantArgs[i], gotArgs[i])
		}
	}
}

func TestInsert(t *testing.T) {
	q := NewQuery("test")
	q.Insert([]string{"t1", "t2", "t3"}, 50, -100, "v1")

	gotStr := q.String()
	wantStr := "INSERT INTO test (t1,t2,t3) VALUES ($1,$2,$3)"
	gotArgs := q.Args()
	wantArgs := []interface{}{50, -100, "v1"}

	if gotStr != wantStr {
		t.Errorf("Insert string: want %q, got %q", wantStr, gotStr)
	}
	if len(gotArgs) != len(wantArgs) {
		t.Errorf("Insert arguments length: want %d, got %d", len(wantArgs), len(gotArgs))
	}
	for i := range wantArgs {
		if wantArgs[i] != gotArgs[i] {
			t.Errorf("Insert arguments[%d]: want %v, got %v", i, wantArgs[i], gotArgs[i])
		}
	}

	q = NewQuery("test2")
	q.Insert(
		[]string{"t1", "t2", "t3"},
		[]interface{}{"12", 10, "abcdefg"},
		[]interface{}{"hijk", uint8(1), 21.12},
	)

	gotStr = q.String()
	wantStr = "INSERT INTO test2 (t1,t2,t3) VALUES ($1,$2,$3),($4,$5,$6)"
	gotArgs = q.Args()
	wantArgs = []interface{}{"12", 10, "abcdefg", "hijk", uint8(1), 21.12}

	if gotStr != wantStr {
		t.Errorf("Insert Multiple Values string: want %q, got %q", wantStr, gotStr)
	}
	if len(gotArgs) != len(wantArgs) {
		t.Errorf("Insert Multiple Values arguments length: want %v, got %v", len(wantArgs), len(gotArgs))
	}
	for i := range wantArgs {
		if wantArgs[i] != gotArgs[i] {
			t.Errorf("Insert Multiple Values arguments[%d]: want %v, got %v", i, wantArgs[i], gotArgs[i])
		}
	}
}
