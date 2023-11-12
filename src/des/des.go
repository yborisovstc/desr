package des

import (
    "github.com/yborisovstc/desr/logr"
    "fmt"
    "os"
    "os/exec"
    "strings"
)


const (
    SpecFileName = "spec.chs"
)


type Des struct {
    Spec string
    Outp string
}


func (des *Des) Run() error {
    // Save spec as a file first
    specFile, err := os.OpenFile(SpecFileName, os.O_CREATE|os.O_RDWR, 0666)
    if (err != nil) {
	logr.LogE("Failed to create spec file: ", err)
	return fmt.Errorf("Failed to create spec file: ", err)
    }
    specFile.Truncate(0)
    _, err = specFile.WriteString(des.Spec)
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
    des.Outp = out.String();
    fmt.Printf("out: %q\n", des.Outp)
    if err != nil {
	logr.LogE("Failed to run DES: ", err)
	return fmt.Errorf("Failed to run DES: ", err)
    }
    logr.LogE("DES run completed, outp: ", des.Outp)
    return nil

}

