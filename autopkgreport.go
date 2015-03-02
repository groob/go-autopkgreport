package main

import (
	"fmt"
	"howett.net/plist"
	"os"
)

type Failure struct {
	Message string `plist:"message"`
	Recipe  string `plist:"recipe"`
}

type NewImport struct {
	Catalog     []string `plist:"catalogs"`
	Name        string   `plist:"name"`
	PkgPath     string   `plist:"pkg_path"`
	PkgInfoPath string   `plist:"pkginfo_path"`
	Version     string   `plist:"version"`
}

type NewPackage struct {
	Id      string `plist:"id"`
	PkgPath string `plist:"pkg_path"`
	Version string `plist:"version"`
}

type AutoPkgReport struct {
	Failures     []Failure    `plist:"failures"`
	NewDownloads []string     `plist:"new_downloads"`
	NewImports   []NewImport  `plist:"new_imports"`
	NewPackages  []NewPackage `plist:"new_packages"`
}

func (r *AutoPkgReport) UnmarshalPlist(f *os.File) error {
	// Unmarshal a file into AutoPkgReport
	decoder := plist.NewDecoder(f)
	err := decoder.Decode(r)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	f, err := os.Open("report.plist")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var report AutoPkgReport
	err = report.UnmarshalPlist(f)
	if err != nil {
		panic(err)
	}
	for _, r := range report.NewImports {
		// Print munki imports.
		fmt.Printf("Imported %v, version %v into catalog %v\n", r.Name, r.Version, r.Catalog)
	}
	for _, p := range report.NewPackages {
		// Print new packages
		fmt.Printf("New Package %v, version %v\n", p.Id, p.Version)
	}
	for _, f := range report.Failures {
		// Print autopkg failures
		fmt.Printf("Recipe failed %v\n", f.Recipe)
	}
}
