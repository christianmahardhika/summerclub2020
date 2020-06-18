**API CONTRACT**


- GET `/api/todo/list`

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
		"title": "Beli sayur",
		"note": "beli sayur di pasar"
	}, {
		"id": 5,
		"title": "Beli Sikat Gigi",
		"note": "beli sikat gigi di pasar"
	}],
	"success": true
}
```

- POST `/api/todo/create`

REQUEST
```
{
	"title": "Ke kantor",
	"note": "- Bawa hape\n- Ambil paket"
}
```


RESPONSE

```json
{
	"errorMessage": null,
	"result": {
		"id": 6,
		"title": "Ke kantor",
		"note": "- Bawa hape\n- Ambil paket"
	},
	"success": true
}
```

- PATCH `/api/todo/update?id=5`

REQUEST
```
{
	"id": 4,
	"title": "Beli sayur",
	"note": "beli sayur di pasar mana"
}
```


RESPONSE

```json
{
	"errorMessage": null,
	"result": {
		"id": 4,
		"title": "Beli sayur",
		"note": "beli sayur di pasar mana"
	},
	"success": true
}
```

- DELETE `/api/todo/delete?id=4`

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
