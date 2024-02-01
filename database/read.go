package crud

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/shubhammishra-1/database/models"
)

//Fetch Single Data

func Read_Single(db *sql.DB, str string) (single model.Details) {

	results, err := db.Query(str)

	if err != nil {
		panic(err.Error())
	}

	//results stores all rows

	for results.Next() {

		err = results.Scan(&single.Name, &single.Roll_number, &single.University, &single.Course, &single.Branch, &single.Start_Year, &single.End_Year)
		if err != nil {
			panic(err.Error())
		}

	}

	return single

}

// Fetch Multiples Data

func Read_Multiple(db *sql.DB, str string) (multiples []model.Details) {

	results, err := db.Query(str)

	if err != nil {
		panic(err.Error())
	}

	//results stores all rows

	for results.Next() {
		var det model.Details // based on model fields will be mapped

		// for each row, scan the result into our tag variable [basically names will be get]

		err = results.Scan(&det.Name, &det.Roll_number, &det.University, &det.Course, &det.Branch, &det.Start_Year, &det.End_Year)
		if err != nil {
			panic(err.Error())
		}

		multiples = append(multiples, det)

	}

	return multiples

}
