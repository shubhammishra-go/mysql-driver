package crud

import (
    "fmt"
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)


// Create Database

func Create_DB(db *sql.DB,name string) {
    
    str:=fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s ;",name);

   _,err := db.Exec(str)
   if err != nil {
       panic(err)
   }else{
  log.Println("Successfully created database..")
    }


   //to use created database

   _,err = db.Exec("USE "+name+";")
   if err != nil {
       panic(err)
   }else{
  log.Println("Currently using "+name +" database..")
   }

}


// Create Table

func Create_Table(db *sql.DB,str string){

    _,err := db.Exec(str)

   if err != nil {
       panic(err)
   }else{
  log.Println("Created a new table ",str);
   }

}



// Insert Details on Table

func Insert(db *sql.DB, str string){
   
   _,err := db.Exec(str)

   if err != nil {
       panic(err)
   }else{
  log.Println("Inserted values into table...");
   }
   

}

