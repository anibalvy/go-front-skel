package config

import (
    "fmt"
    // "os"
    "context"
    "github.com/jackc/pgx/v5"
)

// func PGconn(conf map[string]string) (string, error) {
func PGconn() ( *pgx.Conn, error) {

    // db_url := "postgres://username@server_address:5432/db_name"
    fmt.Println("pg conn config: " + Conf["db_url"].(string))
    // db_url := "postgres://kanibal:paralelepipedo@127.0.0.1:5532/kanibaldb"
    db_url := Conf["db_url"].(string)
    conn, err := pgx.Connect(context.Background(), db_url)
    if err != nil {
        // fmt.Fprint(os.Stderr, "PGconn - Unable to connect to db: %v\n", err)
        fmt.Printf("PGconn - Unable to connect to db: %v\n", err)
        // os.Exit()
        // return "NOK", err
    }
    return conn, err
    // defer conn.Close(context.Background())

    // var greeting string
    // err = conn.QueryRow(context.Background(), "select to_jsonb(row(username, email)) from tb_users").Scan(&greeting)
    // // err = conn.QueryRow(context.Background(), "select to_jsonb(row(username, email)) from tb_users limit 1 offset 1;").Scan(&greeting)
    // // err = conn.QueryRow(context.Background(), "select json_agg(to_json(row(username, email))) from tb_users limit 1 offset 1;").Scan(&greeting)
    // // err = conn.QueryRow(context.Background(), "select jsonb_agg(to_jsonb(row(username, email))) from tb_users limit 1 offset 1;").Scan(&greeting)
    // // err = conn.QueryRow(context.Background(), "select 'xxxxx' from tb_users limit 1;").Scan(&greeting)
    // if err != nil {
    //     fmt.Fprint(os.Stderr, "Query failed: %v\n", err)
    //     // os.Exit()
    //     return "NOK", err

    // }

    // fmt.Println("result:  " + greeting)
    // return greeting, err


}
