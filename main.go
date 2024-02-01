package main

import (
	crud "github.com/shubhammishra-1/database/database"
)

func main() {

	db, err := crud.Open_DB()

	if err != nil {
		panic(err)
	}

	//                   Create Database if doesnot created & using it

	crud.Create_DB(db, "MMT")

	//as i already excuted this command  PRIMARY KEY is Roll number

	//                   Create Table

	// crud.Create_Table(db,"CREATE TABLE DETAILS (NAME VARCHAR(20),ROLL_number VARCHAR(20),University VARCHAR(20), Course VARCHAR(20), Branch VARCHAR(30), Start_Year int, End_Year int, PRIMARY KEY (ROLL_number)) ;")

	//                    Inserting some values into tables

	// str := fmt.Sprintf("INSERT INTO DETAILS VALUES('%s','%s','%s','%s','%s',%d,%d);", "Shubham Mishra", "2020UEE4597", "NSUT", "Btech.", "Electrical Engineering", 2020, 2024)

	// crud.Insert(db, str)

	//                    Reading multiple files

	// var mul []model.Details;

	// mul=crud.Read_Multiple(db,"SELECT * FROM DETAILS;");

	// for i:=0;i<len(mul);i++{
	//     fmt.Println(mul[i]);
	// }

	//                   Reading single file

	// var data model.Details;

	// data=crud.Read_Single(db,"SELECT * FROM DETAILS WHERE Roll_number='2020UEE4597' ;");

	// fmt.Println(data)

	//                 Deleting a single tuple {row}

	//    crud.Delete_Tuple(db,"DELETE FROM DETAILS WHERE Roll_number='2020UEE4599' ;")

	//                   Updating a single tuple {row}

	//   isUpdated:=crud.Update_Tuple(db,"UPDATE DETAILS SET Roll_number='2020UEE4616' where Roll_number='2020UEE4516';")

	//   if isUpdated{
	//   fmt.Println("Successfully updated the tuple")
	//   }else{
	//     fmt.Println("Can't updated this tuple")
	//   }

	defer crud.Close_DB(db)

}
