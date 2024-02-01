package crud
import (
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// Update a details based on some feilds  OR  Update All details 


func Update_Tuple(db *sql.DB,str string)(bool){

res, err := db.Exec(str)

if err != nil {
    panic(err)
}

//this query returns no. of rows afftected after query defintaly for not insertion it will be 0

isUpdated, err := res.RowsAffected()
if err != nil {
    panic(err)
}


if isUpdated ==0{
    return false
}else{
    log.Printf("Successfully UPDATED %d rows",isUpdated);
}

return true

}