```sh
touch go.mod
find . -name go.mod -exec dirname {} \; | xargs -I {} go work use {}
```
