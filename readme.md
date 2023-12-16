## Project example for api users

view post part 1 [here](https://wiliamvj.com/posts/api-golang-parte-1)
view post part 2 [here](https://wiliamvj.com/posts/api-golang-parte-2)
view post part 3 [here](https://wiliamvj.com/posts/api-golang-parte-3)

run project:
```bash
  go run cmd/webserver/main.go
```

generate sqcl files:
```bash
  sqlc generate
```

create new migration:
```bash
  make create_migration
```

run migrations up:
```bash
  make migrate_up
```

run migrations down:
```bash
  make migrate_down
```
