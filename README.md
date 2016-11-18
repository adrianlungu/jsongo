
## JSONGO
A small go library to extract values from a JSON using JS-like syntax.

#### Install
`go get github.com/AdrianLungu/jsongo`

#### Example

##### JSON
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

##### GO
	var j map[string]interface{}
	err := json.Unmarshal(jsonByte, &j)
	if err != nil {
		t.Fatal(err)
	}
	
	val := jsongo.Get("users[1].work", j)
	log.Println(val)

##### Result
`Connic`

#### Additional Details
If the field is missing, the library will return an empty string by default. You can override this like this:

`jsongo.NotFound = "notFound"`
