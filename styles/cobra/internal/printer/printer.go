package printer

import (
	"fmt"
	"os"
	"reflect"

	"github.com/olekukonko/tablewriter"
)

// PrintTable prints a table of properties for a struct using the tablewriter package.
func PrintTable(data interface{}) {
	// Get the reflect.Type and reflect.Value of the data.
	dataType := reflect.TypeOf(data)
	dataValue := reflect.ValueOf(data)

	// Ensure that the data is a struct.
	if dataType.Kind() != reflect.Struct {
		fmt.Println("Error: data is not a struct")
		return
	}

	// Get the Properties field of the struct.
	propertiesValue := dataValue.FieldByName("Properties")
	if !propertiesValue.IsValid() {
		fmt.Println("Error: struct does not have a Properties field")
		return
	}

	propertiesType := propertiesValue.Type()

	// Dereference pointers to get the actual struct type and value.
	if propertiesType.Kind() == reflect.Ptr {
		if propertiesValue.IsNil() {
			// Handle nil pointer value.
			fmt.Println("Error: Properties field is a nil pointer")
			return
		}
		propertiesType = propertiesType.Elem()
		propertiesValue = propertiesValue.Elem()
	}

	if propertiesType.Kind() != reflect.Struct {
		fmt.Println("Error: Properties field is not a struct")
		return
	}

	// Initialize a new tablewriter with tab-separated output.
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)

	// Add the headers for the table.
	headers := []string{}
	for i := 0; i < propertiesType.NumField(); i++ {
		field := propertiesType.Field(i)
		headers = append(headers, field.Name)
	}
	table.SetHeader(headers)

	// Add the row of data for the struct.
	rowValues := []string{}
	for i := 0; i < propertiesType.NumField(); i++ {
		fieldValue := propertiesValue.Field(i)
		// Dereference pointers to get the actual value.
		if fieldValue.Kind() == reflect.Ptr {
			if fieldValue.IsNil() {
				// Handle nil pointer value.
				rowValues = append(rowValues, "<nil>")
				continue
			}
			fieldValue = fieldValue.Elem()
		}
		rowValues = append(rowValues, fmt.Sprintf("%v", fieldValue.Interface()))
	}
	table.Append(rowValues)

	// Render the table.
	table.Render()
}
