# ModernC SQLite Time issue

The goal of this repo is to reproduce an issue with timezones when using the [modernC SQLite package](https://gitlab.com/cznic/sqlite).

The issue is about running the following query:

```sql
select datetime('now','utc') as UTC, datetime('now','localtime') as LOCAL;
```

The issue was raised here: <https://gitlab.com/cznic/sqlite/-/issues/141>

## Observations

The code of this repo has been run via Github actions in Linux, MacOs and inside Docker (in Linux).

The results are the following:

- Running directly in Linux: [All good](https://github.com/fr-ser/modernc-sqlite-time-bug/actions/runs/4772298815/jobs/8484723924)

    ```txt
    Run go run main.go
    UTC: 2023-04-22 02:33:39 - Local: 2023-04-22 18:33:39
    ```

- Running directly in macOs: [All good](https://github.com/fr-ser/modernc-sqlite-time-bug/actions/runs/4772298815/jobs/8484723835)

    ```txt
    Run go run main.go
    UTC: 2023-04-22 02:34:48 - Local: 2023-04-22 18:34:48
    ```

- Running in docker: [Oh no...](https://github.com/fr-ser/modernc-sqlite-time-bug/actions/runs/4772298815/jobs/8484737308)

    ```txt
    Run date && date -u
    Sat Apr 22 06:35:15 EDT 2023
    Sat Apr 22 10:35:15 UTC 2023
    ...
    Run go run main.go
    UTC: 2023-04-22 10:35:57 - Local: 2023-04-22 10:35:57
    ```

## Reproducing Locally

1. Install golang
2. Run locally: `go run main.go` -> "UTC: 2023-04-22 08:39:16 - Local: 2023-04-22 12:39:16"
3. Run in docker: `docker compose run linux-with-go sh -c "cd /app && go run main.go"` -> "UTC: 2023-04-22 10:40:53 - Local: 2023-04-22 10:40:53"
