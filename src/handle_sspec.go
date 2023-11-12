
package main

import (
    "github.com/yborisovstc/desr/logr"
    "net/http"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func handleSspecPost(writer http.ResponseWriter, request *http.Request) {
    logr.LogI("handleSspecPost: enter")

    err := request.ParseForm()
    if err != nil {
	logr.LogE("Failed parsing sysspec form")
    }
    sspec := request.PostFormValue("sysspec")
    logr.LogI("handleSspecPost: sspec: ", sspec)

    // Run system
    model.Spec = sspec
    //err = runDes(sspec)
    err = model.Run()
    if err != nil {
	logr.LogE("Failed running DES: ", err)
    } else {
	//generateHtml(writer, sys, "layout", "navbar")
	logr.LogI("System run, outp: ", model.Outp)
	http.Redirect(writer, request, "/", 302)
    }
}

func runDes(spec string) error {
    // Save spec as a file first
    specFile, err := os.OpenFile("spec.chs", os.O_CREATE|os.O_RDWR, 0666)
    if (err != nil) {
	logr.LogE("Failed to create spec file: ", err)
	return fmt.Errorf("Failed to create spec file: ", err)
    }
    specFile.Truncate(0)
    _, err = specFile.WriteString(spec)
    if err != nil {
	logr.LogE("Failed to write to spec file: ", err)
	return fmt.Errorf("Failed to write to spec file: ", err)
    }
    err = specFile.Close()
    if err != nil {
	logr.LogE("Failed to close spec file: ", err)
	return fmt.Errorf("Failed to close spec file: ", err)
    }
    // Run DES runtime utility
    cmd := exec.Command("../../fap3/clm/fapm3", "-a", "-s", "spec.chs", "-i", "2")
    var out strings.Builder
    cmd.Stdout = &out
    err = cmd.Run()
    fmt.Printf("out: %q\n", out.String())
    if err != nil {
	logr.LogE("Failed to run DES: ", err)
	return fmt.Errorf("Failed to run DES: ", err)
    }
    logr.LogE("DES run completed")
    return nil
}
