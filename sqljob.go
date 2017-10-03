package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strconv"

	_ "github.com/lib/pq"
)

type analyzer struct {
	wconn *connection
}

type connection struct {
	db *sql.DB
}

func initiateconnection(connectionstring string, database string) (*connection, error) {
	db, err := sql.Open(database, connectionstring)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	return &connection{db}, nil
}

type SQLJob struct {
	query      string
	connection *connection
}

type ResultMap struct {
	ColumnName   string
	Kind         reflect.Kind
	ExampleValue interface{}
}

// Analyze uses data from the SqlJob to call the query and
// returns an ResultMap
func (s SQLJob) Analyze() ([]ResultMap, error) {

	rows, err := s.connection.db.Query(s.query)
	if err != nil {
		return nil, err
	}

	columns, _ := rows.Columns()
	count := len(columns)
	res := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for i := range columns {

		fmt.Println(i)
		valuePtrs[i] = &res[i]
	}

	rows.Next()
	fmt.Println(len(valuePtrs))
	err = rows.Scan(valuePtrs...)
	if err != nil {
		return nil, err
	}

	result := make([]ResultMap, count)

	for i, _ := range columns {
		result[i] = ResultMap{columns[i], reflect.TypeOf(res[i]).Kind(), res[i]}
	}

	return result, nil

}

func Dump(i interface{}) {
	r := reflect.TypeOf(i)
	val := reflect.ValueOf(i).Elem()
	//v := reflect.ValueOf(r)

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		//tag := typeField.Tag

		fmt.Printf("%s  %v\n", typeField.Name, GetValueOrDefault(valueField))
	}

	fmt.Println(r.NumMethod())

	for i := 0; i < r.NumMethod(); i++ {
		valueField := r.Method(i)
		nm := valueField.Name
		fmt.Println(valueField.Type.Name())
		fmt.Println(nm)

	} // for index := 0; index < fields; index++ {
	// 	field := v.Type().Field(index)
	//alueAsString := GetValueAsString(valueField)

	// 	fmt.Printf("%v %v %v \n", field.Name, field.Type.Kind(), valueAsString)
	// }
}

func GetValueOrDefault(value reflect.Value) string {
	if !value.CanInterface() {
		return "[UNK]"
	}

	return fmt.Sprintf("%v", value.Interface())
}

func GetValueAsString(v reflect.Value) string {
	var valueInterface interface{}

	if !v.CanInterface() {
		valueInterface = v
	}

	valueInterface = v.Interface()
	val := reflect.ValueOf(valueInterface)

	switch val.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
		return strconv.FormatInt(val.Int(), 10)
	case reflect.String:
		return val.String()
	case reflect.Bool:
		return strconv.FormatBool(val.Bool())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(val.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return "Unkown right now"
	default:
		return fmt.Sprintf("<%v>", val.Kind())
	}
}

func newJob(parameters jobParameters) SQLJob {
	conn, err := initiateconnection(parameters.connectionstring, parameters.database)
	if err != nil {
		log.Panicln("Failed to establish connection", err)
	}

	log.Print("Initated connection to database")

	return SQLJob{parameters.query, conn}
}

func createJobFromFlags() SQLJob {

	parameters := createJobParametersFromFlags()

	return createJobFromParameters(parameters)
}

func createJobFromParameters(parameters jobParameters) SQLJob {
	return newJob(parameters)
}
