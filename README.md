# GO REPORTER
A pointless mini server which takes reports in and allows you to view them again

## Usage
Usuage is pretty limited, you can send a report which must have an **`id`**, a **`type`** which must be either, `error`, `warn` or `info` and a **`message`**. Once you have posted some reports, you can access them by [viewing reports](#viewing-reports)

### Posting new reports

```
curl --request PUT \
  --url http://localhost:1111/report \
  --header 'content-type: application/json' \
  --data '{
	"id": "abc123",
	"type": "error",
	"message": "My really important message!"
}'
```

### Viewing reports

```
curl --request GET \
  --url http://localhost:1111/reports/abc123
```

### Viewing stats
This is just to give a bit of visibility as to what is going on in the Consumer

```
curl --request GET \
  --url http://localhost:1111/stats/abc123
```
