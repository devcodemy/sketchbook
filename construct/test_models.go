package construct

import (
	"encoding/json"
	"log"
)

func TestThemAll() {
	dt1 := map[string]string{
		"user1": "password1",
		"user2": "password2",
	}

	dt2 := map[string]int32{
		"someEntry1": 42,
	}

	dt3 := map[int]string{
		1: "some-host-1.env.dc",
		2: "some-host-2.env.dc",
		3: "some-host-3.env.dc",
		4: "some-host-4.env.dc",
	}

	// dt4 := []map[string]Record{
	dt4 := map[string]Record{
		"some-host-1.env.dc": {
			Name: "Some Specs",
			Data: map[string]int{
				"cpu":  8,
				"ram":  16,
				"sdb1": 20,
				"sdb2": 100,
			},
		},
		"some-host-2.env.dc": {
			Name: "Some Specs",
			Data: map[string]interface{}{
				"cpu":  8,
				"ram":  16,
				"sdb1": 20,
				"sdb2": 200,
				"distro": map[string]interface{}{
					"name":    "ubuntu",
					"version": "20.04",
					"lts":     true,
				},
			},
		},
	}

	rc1 := Record{
		Name: "Some Data - rc1",
		Data: dt1,
	}

	rc2 := Record{
		Name: "Some Number - rc2",
		Data: dt2,
	}

	rc3 := Record{
		Name: "List of hosts - rc3",
		Data: dt3,
	}

	rc4 := Record{
		Name: "Map Approach - rc4",
		Data: dt4,
	}

	ap1 := Application{
		AppName:     "Some Cool Application",
		Description: "Some Long Long Description",
		Records: map[string]Record{
			rc1.Name: rc1,
			rc2.Name: rc2,
			rc3.Name: rc3,
			rc4.Name: rc4,
		},
	}

	in1 := Inventory{
		Applications: []Application{ap1},
	}

	// res1, _ := json.Marshal(in1)
	res1, _ := json.MarshalIndent(in1, "", "  ")
	log.Printf("result:\n%v\n", string(res1))
}
