package jsongo_test

import (
	"testing"
	"adrianlungu.com/jsongo"
	"encoding/json"
)

var jsonByte []byte = []byte(`
{
  "users": [
    {
      "id": 0,
      "name": "Adam Carter",
      "work": "Unilogic",
      "email": "adam.carter@unilogic.com",
      "dob": "24/11/1978",
      "address": "83 Warner Street",
      "city": "Boston",
      "optedin": true
    },
    {
      "id": 1,
      "name": "Leanne Brier",
      "work": "Connic",
      "email": "leanne.brier@connic.org",
      "dob": "13/05/1987",
      "address": "9 Coleman Avenue",
      "city": "Toronto",
      "optedin": false
    }
  ],
  "images": [
    "img0.png",
    "img1.png",
    "img2.png"
  ],
  "coordinates": {
    "x": 35.12,
    "y": -21.49
  },
  "price": "$59,395"
}
`)

var tests = []struct {
	input string
	output string
} {
	{"users[0].id", "0"},
	{"users[1].work", "Connic"},
	{"coordinates.x", "35.12"},
	{"images[0]", "img0.png"},
	{"price", "$59,395"},
}

func TestJsonGo(t *testing.T) {

	var j map[string]interface{}
	err := json.Unmarshal(jsonByte, &j)
	if err != nil {
		t.Fatal(err)
	}

	for _, test := range tests {
		output := jsongo.Get(test.input, j)
		if output != test.output {
			t.Errorf("Expected %s, got %s", test.output, output)
		}
	}

}


func BenchmarkJsonGo(b *testing.B) {

	var j map[string]interface{}
	err := json.Unmarshal(jsonByte, &j)
	if err != nil {
		b.Fatal(err)
	}

	for i:=0; i < b.N; i++ {
		jsongo.Get("users[0].id", j)
	}

}