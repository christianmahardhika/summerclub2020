**API CONTRACT**


- GET `/api/twitter/list`

REQUEST
```
NONE
```

RESPONSE
```json
{
	"errorMessage": null,
	"result": [{
		"id": 4,
		"username": "Chris",
		"post": "Coba coba coba"
	}],
	"success": true
}
```

- POST `/api/twitter/create`

REQUEST
```
{
	"username": "Chris",
	"post": "Hai semua"
}
```


RESPONSE

```json
{
	"errorMessage": null,
	"result": {
		"id": 5,
		"username": "Chris",
		"post": "Hai semua"
	},
	"success": true
}
```

- PATCH `/api/twitter/update?id=5`

REQUEST
```
{
	"id": 5,
	"username": "Chris",
	"post": "Hai semuanyanya"
}
```


RESPONSE

```json
{
	"errorMessage": null,
	"result": {
		"id": 5,
		"username": "Chris",
		"post": "Hai semuanyanya"
	},
	"success": true
}
```

- DELETE `/api/twitter/delete?id=4`

REQUEST
```
NONE
```


RESPONSE

```json
{
	"errorMessage": null,
	"result": "4",
	"success": true
}
```
