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
		t.Errorf("Select arguments: want <nil>, got %q", gotArgs)
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
		t.Errorf("Select with Where arguments: want %q, got %q", wantArgs, gotArgs)
	}
	for i := range wantArgs {
		if wantArgs[i] != gotArgs[i] {
			t.Errorf("Select with Where arguments[%d]: want %q, got %q", i, wantArgs[i], gotArgs[i])
		}
	}
}