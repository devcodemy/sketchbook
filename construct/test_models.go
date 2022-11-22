package construct

import (
	"encoding/json"
	"log"
)

func TestThemAll() {
	dt1 := []map[string]string{
		{"user1": "password1"},
		{"user2": "password2"},
	}

	dt2 := map[string]int32{
		"someEntry1": 42,
	}

	dt3 := []map[int]string{
		{1: "some-host-1.env.dc"},
		{2: "some-host-2.env.dc"},
		{3: "some-host-3.env.dc"},
		{4: "some-host-4.env.dc"},
	}

	rc1 := Record{
		Name: "Some Data",
		Data: dt1,
	}

	rc2 := Record{
		Name: "Some Number",
		Data: dt2,
	}

	rc3 := Record{
		Name: "List of hosts",
		Data: dt3,
	}

	ap1 := Application{
		AppName:     "Some Cool Application",
		Description: "Some Long Long Description",
		Records:     []Record{rc1, rc2, rc3},
	}

	in1 := Inventory{
		Applications: []Application{ap1},
	}

	// res1, _ := json.Marshal(in1)
	res1, _ := json.MarshalIndent(in1, "", "  ")
	log.Printf("result:\n%v\n", string(res1))
}
