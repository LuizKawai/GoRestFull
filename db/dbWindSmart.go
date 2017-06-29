package mysqldb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	//"strconv"
)

var CxStr string = "gcm:gcm@tcp(ec2-13-59-199-234.us-east-2.compute.amazonaws.com:3306)/windsmart"

func main() {

}
//	Consulta de todos los empleados de la empresa
func EmployeesRequest() []Employee {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select u.dni, u.firstname, u.lastname, u.age, u.address, u.personal_email, u.phone_number, u.enterprice_email,  u.is_active, s.turn, j.name, a.name, d.name from user u join jobposition j on j.idjobposition = u.idjobposition join schedule s on s.idschedule = u.idschedule join area a on j.idarea=a.idarea join department d on d.iddepartment = a.iddepartment;"
	fmt.Println(query)
	rows, err := db.Query(query) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	employee := new(Employee)
	employees := []Employee{}
	for rows.Next() {
		err1 := rows.Scan(&employee.dni, &employee.firstname, &employee.lastname, &employee.age, &employee.address, &employee.personal_email, &employee.phone_number, &employee.enterprice_email, &employee.is_active, &employee.turn, &employee.jobPositionName, &employee.areaName, &employee.departmentName)
		if err1 != nil {
			panic(err1.Error())
		} else {
			employees = append(employees, *employee)
		}
	}
	return employees
}

// Consulta del administrador segun enterpice_email y password
func EmployeeLoginRequest(enterprice_email string, password string) []Employee {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select u.dni, u.firstname, u.lastname, u.age, u.address, u.personal_email, u.phone_number, u.enterprice_email, u.is_active, j.name, s.turn, a.name, d.name from user u join jobposition j on j.idjobposition = u.idjobposition join schedule s on s.idschedule = u.idschedule join area a on j.idarea=a.idarea join department d on d.iddepartment = a.iddepartment where u.enterprice_email = ? and password = ?"
	fmt.Println(query)
	rows, err := db.Query(query, enterprice_email, password) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	employee := new(Employee)
	employees := []Employee{}
	for rows.Next() {
		err1 := rows.Scan(&employee.dni, &employee.firstname, &employee.lastname, &employee.age, &employee.address, &employee.personal_email, &employee.phone_number, &employee.enterprice_email, &employee.is_active, &employee.turn, &employee.jobPositionName, &employee.areaName, &employee.departmentName)
		if err1 != nil {
			panic(err1.Error())
		} else {
			employees = append(employees, *employee)
		}
	}
	return employees
}
// Consulta de un empleado identificado por su dni
func EmployeeDNIRequest(dni string) []Employee {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select u.dni, u.firstname, u.lastname, u.age, u.address, u.personal_email, u.phone_number, u.enterprice_email, u.is_active, j.name, s.turn, a.name, d.name from user u join jobposition j on j.idjobposition = u.idjobposition join schedule s on s.idschedule = u.idschedule join area a on j.idarea=a.idarea join department d on d.iddepartment = a.iddepartment where u.dni = ?"
	fmt.Println(query)
	rows, err := db.Query(query, dni) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	employee := new(Employee)
	employees := []Employee{}
	for rows.Next() {
		err1 := rows.Scan(&employee.dni, &employee.firstname, &employee.lastname, &employee.age, &employee.address, &employee.personal_email, &employee.phone_number, &employee.enterprice_email, &employee.is_active, &employee.turn, &employee.jobPositionName, &employee.areaName, &employee.departmentName)
		if err1 != nil {
			panic(err1.Error())
		} else {
			employees = append(employees, *employee)
		}
	}
	return employees
}

// Consulta de empleados por area de trabajo
func EmployeesForAreaRequest(idarea int) []Employee {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select u.dni, u.firstname, u.lastname, u.age, u.address, u.personal_email, u.phone_number, u.enterprice_email, u.is_active, j.name, s.turn, a.name, d.name from user u join jobposition j on j.idjobposition = u.idjobposition join schedule s on s.idschedule = u.idschedule join area a on j.idarea=a.idarea join department d on d.iddepartment = a.iddepartment where a.idarea = ?"
	fmt.Println(query)
	rows, err := db.Query(query, idarea) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	employee := new(Employee)
	employees := []Employee{}
	for rows.Next() {
		err1 := rows.Scan(&employee.dni, &employee.firstname, &employee.lastname, &employee.age, &employee.address, &employee.personal_email, &employee.phone_number, &employee.enterprice_email, &employee.is_active, &employee.turn, &employee.jobPositionName, &employee.areaName, &employee.departmentName)
		if err1 != nil {
			panic(err1.Error())
		} else {
			employees = append(employees, *employee)
		}
	}
	return employees
}

// Consulta de empleados por departamento donde trabaja
func EmployeesForDepartmentRequest(iddepartment int) []Employee {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select u.dni, u.firstname, u.lastname, u.age, u.address, u.personal_email, u.phone_number, u.enterprice_email, u.is_active, j.name, s.turn, a.name, d.name from user u join jobposition j on j.idjobposition = u.idjobposition join schedule s on s.idschedule = u.idschedule join area a on j.idarea=a.idarea join department d on d.iddepartment = a.iddepartment where d.iddepartment = ?"
	fmt.Println(query)
	rows, err := db.Query(query, iddepartment) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	employee := new(Employee)
	employees := []Employee{}
	for rows.Next() {
		err1 := rows.Scan(&employee.dni, &employee.firstname, &employee.lastname, &employee.age, &employee.address, &employee.personal_email, &employee.phone_number, &employee.enterprice_email, &employee.is_active, &employee.turn, &employee.jobPositionName, &employee.areaName, &employee.departmentName)
		if err1 != nil {
			panic(err1.Error())
		} else {
			employees = append(employees, *employee)
		}
	}
	return employees
}

// Consulta de empleados por puesto laboral
func EmployeesForJobPositionRequest(idjobposition int) []Employee {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select u.dni, u.firstname, u.lastname, u.age, u.address, u.personal_email, u.phone_number, u.enterprice_email, u.is_active, j.name, s.turn, a.name, d.name from user u join jobposition j on j.idjobposition = u.idjobposition join schedule s on s.idschedule = u.idschedule join area a on j.idarea=a.idarea join department d on d.iddepartment = a.iddepartment where j.idjobposition = ?"
	fmt.Println(query)
	rows, err := db.Query(query, idjobposition) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	employee := new(Employee)
	employees := []Employee{}
	for rows.Next() {
		err1 := rows.Scan(&employee.dni, &employee.firstname, &employee.lastname, &employee.age, &employee.address, &employee.personal_email, &employee.phone_number, &employee.enterprice_email, &employee.is_active, &employee.turn, &employee.jobPositionName, &employee.areaName, &employee.departmentName)
		if err1 != nil {
			panic(err1.Error())
		} else {
			employees = append(employees, *employee)
		}
	}
	return employees
}

// Consulta de empleados por turno laboral
func EmployeesForScheduleRequest(idschedule int) []Employee {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select u.dni, u.firstname, u.lastname, u.age, u.address, u.personal_email, u.phone_number, u.enterprice_email, u.is_active, j.name, s.turn, a.name, d.name from user u join jobposition j on j.idjobposition = u.idjobposition join schedule s on s.idschedule = u.idschedule join area a on j.idarea=a.idarea join department d on d.iddepartment = a.iddepartment where s.idschedule = ?"
	fmt.Println(query)
	rows, err := db.Query(query, idschedule) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	employee := new(Employee)
	employees := []Employee{}
	for rows.Next() {
		err1 := rows.Scan(&employee.dni, &employee.firstname, &employee.lastname, &employee.age, &employee.address, &employee.personal_email, &employee.phone_number, &employee.enterprice_email, &employee.is_active, &employee.turn, &employee.jobPositionName, &employee.areaName, &employee.departmentName)
		if err1 != nil {
			panic(err1.Error())
		} else {
			employees = append(employees, *employee)
		}
	}
	return employees
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// Consultar Departamentos
func DepartmentRequest() []Department {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select iddepartment, name, description from department"
	fmt.Println(query)
	rows, err := db.Query(query) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	department := new(Department)
	departments := []Department{}
	for rows.Next() {
		err1 := rows.Scan(&department.iddepartment, &department.name, &department.description)
		if err1 != nil {
			panic(err1.Error())
			} else {
			departments = append(departments, *department)
		}
	}
	return departments
}

// Consultar Areas
func AreaRequest(id int) []Area {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select idarea, name, description, iddepartment from area where iddepartment = ?"
	fmt.Println(query)
	rows, err := db.Query(query, id) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	area := new(Area)
	areas := []Area{}
	for rows.Next() {
		err1 := rows.Scan(&area.idarea, &area.name, &area.description, &area.iddepartment)
		if err1 != nil {
			panic(err1.Error())
		} else {
			areas = append(areas, *area)
		}
	}
	return areas
}

// Consultar Puestos Laborales
func JobPositionRequest(id int) []JobPosition {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select idjobposition, name, description, idarea from Recomendacion where idarea = ?"
	fmt.Println(query)
	rows, err := db.Query(query, id) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	jobPosition := new(JobPosition)
	jobPositions := []JobPosition{}
	for rows.Next() {
		err1 := rows.Scan(&jobPosition.idjobposition, &jobPosition.name, &jobPosition.description, &jobPosition.idarea)
		if err1 != nil {
			panic(err1.Error())
		} else {
			jobPositions = append(jobPositions, *jobPosition)
		}
	}
	return jobPositions
}

// Consultar Horarios
func ScheduleRequest() []Schedule {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	// Execute the query
	query := "select idschedule, starttime, finaltime, turn from schedule"
	fmt.Println(query)
	rows, err := db.Query(query) //SELECT * FROM table
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	schedule := new(Schedule)
	schedules := []Schedule{}
	for rows.Next() {
		err1 := rows.Scan(&schedule.idschedule, &schedule.starttime, &schedule.finaltime, &schedule.turn)
		if err1 != nil {
			panic(err1.Error())
			} else {
			schedules = append(schedules, *schedule)
		}
	}
	return schedules
}

///////////////////////////////////////////////////////////////////////////////////////////////////////

func InsertEmployee(usr NewEmployee) {
	db, err := sql.Open("mysql", CxStr)
	if err != nil {
		log.Printf(err.Error())
	}

	defer db.Close()

	res := mustExec(db, "INSERT INTO user (dni, firstname, lastname, address, birthdate, age, gender, personal_email, phone_number, hire_date, enterprice_email, password,  is_active, is_admin, idjobposition, idschedule)  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", usr.dni, usr.firstname, usr.lastname, usr.address, usr.birthdate, usr.age, usr.gender, usr.personal_email, usr.phone_number, usr.hire_date, usr.enterprice_email, usr.password, usr.is_active, usr.is_admin, usr.idjobposition, usr.idschedule)
	count, err := res.RowsAffected()
	if err != nil {
		log.Printf("res.RowsAffected() returned error: %s", err.Error())
	}
	if count != 1 {
		log.Printf("expected 1 affected row, got %d", count)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Printf("res.LastInsertId() returned error: %s", err.Error())
	}
	if id != 0 {
		log.Printf("expected InsertId 0, got %d", id)
	}
}


func mustExec(db *sql.DB, query string, args ...interface{}) (res sql.Result) {
	res, err := db.Exec(query, args...)
	if err != nil {
		fail("exec", query, err)
	}
	return res
}

func fail(method, query string, err error) {
	if len(query) > 300 {
		query = "[query too large to print]"
	}
	// dbt.Fatalf("error on %s %s: %s", method, query, err.Error())
	log.Printf("error on %s %s: %s", method, query, err.Error())
}
