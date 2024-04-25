# db with golang

## database GUI

- 推薦1 [pgAdmin](https://hub.docker.com/r/dpage/pgadmin4/)
- 推薦2 [vscode](https://marketplace.visualstudio.com/items?itemName=cweijan.dbclient-jdbc)
- 推薦3 [jetbrains](https://www.jetbrains.com)

## Install

postgresql

```bash
docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 0.0.0.0:5432:5432 -d postgres
```