package mysqldb


// ---- WINDSMART ---


type Department struct{
	iddepartment int
	name int
	description string
}

type Area struct{
	idarea int
	name int
	description string
	iddepartment int
}

type JobPosition struct{
	idjobposition int
	name string
	description string	
	idarea int
}

type Schedule struct{
	idschedule int
	starttime string
	finaltime string
	turn string
}

type Employee struct{
	dni string
	firstname string
	lastname string
	age int
	address string
	personal_email string
	phone_number string
	enterprice_email string
	is_active bool
	turn string
	jobPositionName string
	areaName string
	departmentName string
}

type NewEmployee struct{
	dni string 
	firstname string 
	lastname string 
	address string 
	birthdate string 
	age int 
	gender string 
	personal_email string 
	phone_number string 
	hire_date string 
	enterprice_email string 
	password string  
	is_active bool 
	is_admin bool 
	idjobposition int 
	idschedule int
}


