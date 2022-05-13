package main

import (
	"fmt"
	"reflect"
)

// ----------- Version

type Version struct {
	ID    int
	stage string
	owner string
}

func (v Version) createDraft(createdBy string) {
	v.ID = 0
	v.stage = "DRAFT"
	v.owner = createdBy
}

// createDraft()
// createShadow()
// updateShadows()
// updateColaborators()
// createFinal()

// ----------- Non-Versionables

type DocNonVersioned struct {
	ID          string
	description string
}

// ----------- Versionables

type Versionable interface {
	GetVersionData() *Version
}

type DocModel1Versioned struct {
	ID          string
	version     Version
	description string
}

func (d DocModel1Versioned) GetVersionData() *Version {
	return &d.version
}

type DocModel2Versioned struct {
	ID          string
	version     Version
	description string
}

func (d DocModel2Versioned) GetVersionData() *Version {
	return &d.version
}

// ----------- ShowTime

func main() {

	var doc1 = DocModel1Versioned{"doc1", Version{}, "doc versionado 1"}
	var doc2 = DocModel2Versioned{"doc2", Version{}, "doc versionado 2"}
	//var doc3 = DocNonVersioned{"doc3", "doc não versionado 3"}

	genericOperation(doc1)
	genericOperation(doc2)
	//genericOperation(doc3)

	fmt.Println(doc1.version)
	fmt.Println(doc2.version)

}

func genericOperation(doc Versionable) {
	fmt.Println(reflect.ValueOf(doc).Kind())
	fmt.Println(doc)
	if doc != nil {
		version := *doc.GetVersionData()
		version.stage = "DRAFT"
		version.ID = 1
		fmt.Println(version)
	}
	fmt.Println(doc)
}
