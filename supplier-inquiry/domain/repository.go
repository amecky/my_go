package domain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Repository interface {
	FindByMemberId(memberId int) SupplierInquiryList
	FindByMemberIdAndId(memberId int, id string) (SupplierInquiry, bool)
	SetStatus(memberId int, id string, status string)
}

type MockRepository struct {
	supplierInquiries []SupplierInquiry
}

//
// Liest eine Datei und wandelt diese in ein JSON Objekt
//
func parse_file(repo *MockRepository, filename string) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var si SupplierInquiry
	json.Unmarshal(byteValue, &si)
	//fmt.Printf("SI: %s => %s %d - %d\n",filename,si.Id,si.MemberId,si.CtoKey)
	repo.supplierInquiries = append(repo.supplierInquiries, si)
}

//
// Listet zuerst alle Dateien in einem Verzeichnis auf und liest dann jede einzelne
// Datei und wandelt den Inhalt in JSON Objekt
//
func read_json_files(repo *MockRepository) {
	var files []string

	root := "lierue"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		parse_file(repo, file)
	}
}

func NewMockRepository() *MockRepository {
	repo := MockRepository{}
	read_json_files(&repo)
	fmt.Printf("Lierues: %d\n", len(repo.supplierInquiries))
	return &repo
}

func (r *MockRepository) FindByMemberId(memberId int) SupplierInquiryList {
	var ret SupplierInquiryList
	for _, item := range r.supplierInquiries {
		if item.MemberId == memberId {
			ret.SupplierInquiries = append(ret.SupplierInquiries, item)
		}
	}
	ret.TotalCount = len(ret.SupplierInquiries)
	return ret
}

func (r *MockRepository) FindByMemberIdAndId(memberId int, id string) (SupplierInquiry, bool) {
	for _, item := range r.supplierInquiries {
		if item.MemberId == memberId && item.Id == id {
			return item, true
		}
	}
	return SupplierInquiry{}, false
}

//
// Methode um den Status einer Lierue zu setzen, aber nur falls der Status noch offen ist
//
func (r *MockRepository) SetStatus(memberId int, id string, status string) {
	for i, item := range r.supplierInquiries {
		if item.MemberId == memberId && item.Id == id && item.Status == "OPEN" {
			r.supplierInquiries[i].Status = status
		}
	}
}
