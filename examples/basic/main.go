package main

/*
// TODO: resource ID 转 GO代码
#include "ui/resource.h"
*/
import "C"
import (
	"github.com/whtiehack/wingui"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

var dlg *wingui.Dialog

func main() {
	log.Printf("resource %v %#[1]v  \n", C.IDD_DIALOG)
	var err error
	dlg, err = wingui.NewDialog(C.IDD_DIALOG, 0)
	if err != nil {
		log.Panic("main dialog create error", err)
	}
	log.Println("dlg create end", dlg)
	var btn *wingui.Button
	btn, _ = dlg.NewButton(C.IDB_OK)
	btn.OnClicked = modalBtnClicked
	closeBtn, _ := dlg.NewButton(C.IDB_CANCEL)
	closeBtn.OnClicked = func() {
		dlg.Close()
	}
	dlg.Show()
	wingui.MessageLoop()
	log.Println("stoped")
}

func modalBtnClicked() {
	log.Println("btn clicked")
	wingui.NewModalDialog(C.IDD_DIALOG_OK, dlg.Handle(), func(okdlg *wingui.Dialog) {
		okbtn, _ := okdlg.NewButton(C.IDB_OK)
		okbtn.OnClicked = func() {
			log.Println("modal btn click")
			okdlg.Close()
		}
	})
}