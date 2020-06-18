**API CONTRACT**


- GET `/api/message/list`

REQUEST
```
NONE
```

RESPONSE
```json
{
	"errorMessage": null,
	"result": [{
		"id": 1,
		"from": "Salman",
		"to": "Everyone",
		"message": "Please follow my instagram: @salman.isme  ;)"
	}, {
		"id": 20,
		"from": "",
		"to": "kak Salman",
		"message": "bo ong dosa lo kakk\r\n"
	}],
	"success": true
}
```

- POST `/api/message/post`

REQUEST
```
{
	"from": "Salman",
	"to": "Bawah",
	"message": "Gak bohong kok :)"
}
```


RESPONSE

```json
{
	"errorMessage": null,
	"result": [{
		"id": 1,
		"from": "Salman",
		"to": "Everyone",
		"message": "Please follow my instagram: @salman.isme  ;)"
	}, {
		"id": 20,
		"from": "",
		"to": "kak Salman",
		"message": "bo ong dosa lo kakk\r\n"
	}],
	"success": true
}
```
