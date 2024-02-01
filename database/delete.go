package crud

import (
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// Delete Database

func Delete_Database(db *sql.DB,str string){
     
    _,err := db.Exec(str)

   if err != nil {
       panic(err)
   }


}

// Delete Table

func Delete_Table(db *sql.DB,str string){
     
    _,err := db.Exec(str)

   if err != nil {
       panic(err)
   }


}


// Delete a specific Tuple {Row}

func Delete_Tuple(db *sql.DB,str string){
     
    _,err := db.Exec(str)

   if err != nil {
       panic(err)
   }else{
  log.Println("Deleted this row ::: ", str);
   }


}


