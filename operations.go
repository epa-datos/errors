package errors

import "fmt"

type Operation string

func GetDataFromDBOperation(resourceToRetrieve string) Operation {
	return Operation(fmt.Sprintf("Retrieve %s from db", resourceToRetrieve))
}

func BindingResourceOperation(resource string) Operation {
	return Operation(fmt.Sprintf("Binding %s", resource))
}

func InsertIntoDBOperation(resource string) Operation {
	return Operation(fmt.Sprintf("Insert %s into db", resource))
}

func UpdateOnDBOperation(resource string) Operation {
	return Operation(fmt.Sprintf("Update %s on db", resource))
}

func DeleteOnDBOperation(resource string) Operation {
	return Operation(fmt.Sprintf("Delete %s on db", resource))
}
