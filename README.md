# uuid-enrich

Assumes a stream of JSON formatted log entries, and enriches them with an eventId as a UUID.

```shell
go build -v -o uuid-enrich .
./uuid-enrich -in input.log -out output.log
```
