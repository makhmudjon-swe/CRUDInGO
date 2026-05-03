package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
    host = "localhost"
    port = 5432
    user = "postgres"
    password = "1916"
    dbname = "crud"
)

func main(){
    psqlInfo := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)

    if err != nil{
        fmt.Println("Not connected")
    }

    defer db.Close()

    err = db.Ping()

    if err != nil{
        panic(err)
    }

    // err = createUser(db, "Alisher", "alisher@gmail.com", 20)
    // if err != nil{
    //     fmt.Println(err.Error())
    // }else{
    //     fmt.Println("User successfuly created!")
    // }

    // err = getUsers(db)
    // if err != nil {
    //     fmt.Println(err.Error())
    // }

    // err = updateUser(db, "Jahongir", 23, 1)

    // if err != nil {
    //     fmt.Println(err.Error())
    // }else{
    //     fmt.Println("User successfuly updated")
    // }

    err = deleteUser(db,1)

    if err != nil {
        fmt.Println(err.Error())
    }else{
        fmt.Println("User successfuly deleted!")
    }


}

func createUser(db *sql.DB, name string, email string, age int) error{
    query := `
    insert into users(
    name,
    email,
    age
    ) values(
    $1, $2, $3 
    )
    `
    _, err := db.Exec(query, name, email, age)
    return  err
}

type User struct{
    Id int
    Name string
    Email string
    Age int
}

func getUsers(db *sql.DB) error{
    query := `
    Select 
        id,
        name,
        email,
        age
    From users
    `
    rows , err := db.Query(query)

    if err != nil{
        return  err
    }

    defer rows.Close()

    var users []User

    for rows.Next(){
        var user User
        err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Age)

        if err != nil{
            return err
        }

        users = append(users, user)
    }

    fmt.Println(users)

    return nil
}

func updateUser(db *sql.DB, name string, age int, id int) error{
    query := `
    UPDATE users set 
        name = $1,
        age = $2 
    where id = $3 
    `
    
    _, err := db.Exec(query, name, age, id)

    if err != nil {
        return  err
    }

    return  nil
}


func deleteUser(db *sql.DB, id int) error{
    query := `
    Delete from users where id = $1;
    `
    _, err := db.Exec(query, id)

    if err != nil {
        return err
    }

    return nil
}