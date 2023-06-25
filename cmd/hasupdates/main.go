package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/fatih/color"
	"golang.org/x/mod/semver"
)

func main() {
	var verbose bool
	flag.BoolVar(&verbose, "v", false, "include modules that are current")
	flag.Parse()
	if verbose {
		println("verbose...")
	}
	println("running go list...")
	out, err := exec.Command("go", "list", "-json", "-u", "-m", "all").Output()
	if err != nil {
		log.Fatal(err)
	}

	src := bytes.NewReader(out)
	decoder := json.NewDecoder(src)
	for decoder.More() {
		var m Module

		if err := decoder.Decode(&m); err != nil {
			log.Fatal(err)
		}
		if m.Error != nil {
			color.Red(fmt.Sprintf("module: %s  error: %s\n", m.Path, m.Error.Err))
			continue
		}
		if m.Main { // don't report the main module
			continue
		}
		if !m.Indirect {
			printmodule(m, verbose)
		}
	}
}

// printmodule may print a module update information
func printmodule(m Module, verbose bool) {
	major := false
	msg := m.Path

	if verbose || m.Update != nil {
		msg += fmt.Sprintf(" => %s", m.Version)
		if m.Time != nil {
			msg += fmt.Sprintf(" (%s)", m.Time.Format("2006-01-02"))
		}
	}

	if m.Update != nil {
		msg += fmt.Sprintf(" => %s", m.Update.Version)
		if m.Update.Time != nil {
			msg += fmt.Sprintf(" (%s)", m.Update.Time.Format("2006-01-02"))
		}
		existing := semver.Major(m.Version)
		new := semver.Major(m.Update.Version)
		if existing != new {
			major = true
		}
	}
	if m.Update == nil {
		if verbose {
			color.Green(msg)
		}
	} else {
		if major {
			color.Red(msg)
		} else {
			color.Yellow(msg)
		}
	}
}

// Module is module in a repository.  This structure comes from the GO docs
// as of 25 June 2023.
type Module struct {
	Path       string       // module path
	Query      string       // version query corresponding to this version
	Version    string       // module version
	Versions   []string     // available module versions
	Replace    *Module      // replaced by this module
	Time       *time.Time   // time version was created
	Update     *Module      // available update (with -u)
	Main       bool         // is this the main module?
	Indirect   bool         // module is only indirectly needed by main module
	Dir        string       // directory holding local copy of files, if any
	GoMod      string       // path to go.mod file describing module, if any
	GoVersion  string       // go version used in module
	Retracted  []string     // retraction information, if any (with -retracted or -u)
	Deprecated string       // deprecation message, if any (with -u)
	Error      *ModuleError // error loading module
	Origin     any          // provenance of module
	Reuse      bool         // reuse of old module info is safe
}

// ModuleError holds any errors loading the module
type ModuleError struct {
	Err string // the error itself
}
