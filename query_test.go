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
	for i, v := range gotArgs {
		if v != wantArgs[i] {
			t.Errorf("Select with Where arguments[%d]: want %v, got %v", i, wantArgs[i], v)
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
	for i, v := range gotArgs {
		if v != wantArgs[i] {
			t.Errorf("Insert arguments[%d]: want %v, got %v", i, wantArgs[i], v)
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
	for i, v := range gotArgs {
		if v != wantArgs[i] {
			t.Errorf("Insert Multiple Values arguments[%d]: want %v, got %v", i, wantArgs[i], v)
		}
	}
}

func TestUpdate(t *testing.T) {
	q := NewQuery("test")
	q.Update("t1 = ?, t2 = ?, t3 = ?", "v1", 2, true).Where("id = ?", 101)

	gotStr := q.String()
	wantStr := "UPDATE test SET t1 = $1, t2 = $2, t3 = $3 WHERE id = $4"
	gotArgs := q.Args()
	wantArgs := []interface{}{"v1", 2, true, 101}

	if gotStr != wantStr {
		t.Errorf("Update string: want %q, got %q", wantStr, gotStr)
	}
	if len(gotArgs) != len(wantArgs) {
		t.Errorf("Update arguments length: want %d, got %d", len(wantArgs), len(gotArgs))
	}
	for i, v := range gotArgs {
		if v != wantArgs[i] {
			t.Errorf("Update arguments[%d]: want %v, got %v", i, wantArgs[i], v)
		}
	}

	q.SetTable("test2")
	q.Update(map[string]interface{}{
		"t1": "v1",
		"t2": 2,
	}).Where("id = ?", 101)

	gotStr = q.String()
	gotArgs = q.Args()

	wantStr = "UPDATE test2 SET t1=$1,t2=$2 WHERE id = $3"
	wantStr2 := "UPDATE test2 SET t2=$1,t1=$2 WHERE id = $3"
	wantArgs = []interface{}{"v1", 2, 101}
	wantArgs2 := []interface{}{2, "v1", 101}

	if gotStr != wantStr && gotStr != wantStr2 {
		t.Errorf("Update with map string: want %q or %q, got %q", wantStr, wantStr2, gotStr)
	}
	if len(gotArgs) != len(wantArgs) {
		t.Errorf("Update with map arguments length: want %d, got %d", len(wantArgs), len(gotArgs))
	}
	for i := range wantArgs {
		if gotArgs[i] != wantArgs[i] && gotArgs[i] != wantArgs2[i] {
			t.Errorf("Update with map arguments[%d]: want %v or %v, got %v", i, wantArgs[i], wantArgs2[i], gotArgs[i])
		}
	}
}

func TestDelete(t *testing.T) {
	q := NewQuery("test")
	q.Delete().Where("id = ?", 5)

	gotStr := q.String()
	wantStr := "DELETE FROM test WHERE id = $1"
	gotArgs := q.Args()
	wantArgs := []interface{}{5}

	if gotStr != wantStr {
		t.Errorf("Delete string: want %q, got %q", wantStr, gotStr)
	}
	if len(gotArgs) != len(wantArgs) {
		t.Errorf("Delete arguments length: want %d, got %d", len(wantArgs), len(gotArgs))
	}
	for i, v := range gotArgs {
		if v != wantArgs[i] {
			t.Errorf("Delete arguments[%d]: want %v, got %v", i, wantArgs[i], v)
		}
	}
}
