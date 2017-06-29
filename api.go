package main
	
import (
	datos "./db"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"strconv"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		// ----  WindSmart Get Request ---
		///////////////////////////////////////////////////////////////////////////////////////////////////////
		&rest.Route{"GET", "/employees", Employees_handler},
		&rest.Route{"GET", "/employee/:dni", Employee_Dni_handler},
		&rest.Route{"GET", "/employee/:enterprice_email/:password", Admin_handler},
		///////////////////////////////////////////////////////////////////////////////////////////////////////
		&rest.Route{"GET", "/employees/department/:iddepartment", Employees_Department_handler},
		&rest.Route{"GET", "/employees/area/:idarea", Employees_Area_handler},
		&rest.Route{"GET", "/employees/jobposition/:idjobposition", Employees_Jobposition_handler},
		&rest.Route{"GET", "/employees/schedule/:idschedule", Employees_Schedule_handler},
		///////////////////////////////////////////////////////////////////////////////////////////////////////
		&rest.Route{"GET", "/departments", Department_handler},
		&rest.Route{"GET", "/department/:iddepartment/areas", Areas_Department_handler},
		&rest.Route{"GET", "/department/area/:idarea/jobpositions", Jobpositions_Area_handler},
		&rest.Route{"GET", "/schedules", Schedule_handler},
		///////////////////////////////////////////////////////////////////////////////////////////////////////
		// ----  WindSmart Post Request ---
		///////////////////////////////////////////////////////////////////////////////////////////////////////
		&rest.Route{"POST", "/newemployee", New_Employee_handler},
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":9988", api.MakeHandler()))
}

// ----  WindSmart ---
//////////////////////////////////////////////////////////////////////////////////////////
func Employees_handler(w rest.ResponseWriter, r *rest.Request) {
	employees := datos.EmployeesRequest()
	w.WriteJson(employees)
}

func Employee_Dni_handler(w rest.ResponseWriter, r *rest.Request) {
	dni := r.PathParam("dni")
	employee := datos.EmployeeDNIRequest(dni)
	w.WriteJson(employee)
}

func Admin_handler(w rest.ResponseWriter, r *rest.Request) {
	enterprice_email := r.PathParam("enterprice_email")
	password := r.PathParam("password")
	employee := datos.EmployeeLoginRequest(enterprice_email, password)
	w.WriteJson(employee)
}

//////////////////////////////////////////////////////////////////////////////////////////
func Employees_Department_handler(w rest.ResponseWriter, r *rest.Request) {
	iddepartment := r.PathParam("iddepartment")
	id, _ := strconv.Atoi(iddepartment)
	employees := datos.EmployeesForDepartmentRequest(id)
	w.WriteJson(employees)
}

func Employees_Area_handler(w rest.ResponseWriter, r *rest.Request) {
	idarea := r.PathParam("idarea")
	id, _ := strconv.Atoi(idarea)
	employees := datos.EmployeesForAreaRequest(id)
	w.WriteJson(employees)
}

func Employees_Jobposition_handler(w rest.ResponseWriter, r *rest.Request) {
	idjobposition := r.PathParam("idjobposition")
	id, _ := strconv.Atoi(idjobposition)
	employees := datos.EmployeesForJobPositionRequest(id)
	w.WriteJson(employees)
}

func Employees_Schedule_handler(w rest.ResponseWriter, r *rest.Request) {
	idarea := r.PathParam("idarea")
	id, _ := strconv.Atoi(idarea)
	employees := datos.EmployeesForScheduleRequest(id)
	w.WriteJson(employees)
}

//////////////////////////////////////////////////////////////////////////////////////////
func Department_handler(w rest.ResponseWriter, r *rest.Request){
	departments := datos.DepartmentRequest()
	w.WriteJson(departments)	
}

func Areas_Department_handler(w rest.ResponseWriter, r *rest.Request) {
	iddepartment := r.PathParam("iddepartment")
	id, _ := strconv.Atoi(iddepartment)
	areas := datos.AreaRequest(id)
	w.WriteJson(areas)
}

func Jobpositions_Area_handler(w rest.ResponseWriter, r *rest.Request) {
	idarea := r.PathParam("idarea")
	id, _ := strconv.Atoi(idarea)
	jobpositions := datos.JobPositionRequest(id)
	w.WriteJson(jobpositions)
}

func Schedule_handler(w rest.ResponseWriter, r *rest.Request){
	schedules := datos.ScheduleRequest()
	w.WriteJson(schedules)	
}

//////////////////////////////////////////////////////////////////////////////////////////

func New_Employee_handler(w rest.ResponseWriter, r *rest.Request) {
	employee := new(datos.NewEmployee)
	err := r.DecodeJsonPayload(&employee)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	datos.InsertEmployee(*employee)
	w.WriteJson(&employee)
}
