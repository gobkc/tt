package main

import (
	"fmt"
	"github.com/ying32/govcl/vcl"
	"strconv"
	"time"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl/types"
)

type TMainForm struct {
	*vcl.TForm
	Button1     *vcl.TButton
	Text1       *vcl.TLabeledEdit
	ConvertText *vcl.TLabeledEdit
}

type TForm1 struct {
	*vcl.TForm
	Button1 *vcl.TButton
}

var (
	mainForm *TMainForm
	form1    *TForm1
)

func main() {
	vcl.DEBUG = true
	vcl.RunApp(&mainForm, &form1)
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.SetCaption("Timestamp Convert")
	f.EnabledMaximize(false)
	f.SetWidth(600)
	f.SetHeight(600)
	f.ScreenCenter()

	f.Text1 = vcl.NewLabeledEdit(f)
	f.Text1.SetParent(f)
	f.Text1.SetLeft(250)
	f.Text1.SetWidth(200)
	f.Text1.SetTop(100)
	f.Text1.SetLabelPosition(types.LpLeft)
	f.Text1.SetLabelSpacing(6)
	f.Text1.EditLabel().SetCaption("输入时间戳或时间:")

	f.Button1 = vcl.NewButton(f)
	f.Button1.SetParent(f)
	f.Button1.SetCaption("转换为为时间戳或时间")
	f.Button1.SetLeft(250)
	f.Button1.SetWidth(200)
	f.Button1.SetHeight(40)
	f.Button1.SetTop(150)
	f.Button1.SetOnClick(f.OnButton1Click)
	f.Button1.Font().SetStyle(types.NewSet(types.FsBold)) //f.Button1.Font().Style().Include(types.FsBold))

	f.ConvertText = vcl.NewLabeledEdit(f)
	f.ConvertText.SetParent(f)
	f.ConvertText.SetLeft(250)
	f.ConvertText.SetWidth(200)
	f.ConvertText.SetTop(200)
	f.ConvertText.SetLabelPosition(types.LpLeft)
	f.ConvertText.SetLabelSpacing(6)
	f.ConvertText.EditLabel().SetCaption("结果:")
}

func (f *TMainForm) OnFormCloseQuery(Sender vcl.IObject, CanClose *bool) {
	//*CanClose = vcl.MessageDlg("是否退出？", types.MtConfirmation, types.MbYes, types.MbNo) == types.IdYes
}

func (f *TMainForm) OnButton1Click(object vcl.IObject) {
	var buf string
	f.Text1.GetTextBuf(&buf, 25)
	if buf == `` {
		f.ConvertText.SetText(`输入结果为空`)
		return
	}
	num, err := strconv.ParseInt(buf, 10, 64)
	if err != nil {
		t, _ := time.Parse("2006-01-02 15:04:05", buf)
		f.ConvertText.SetText(fmt.Sprintf(`%v`, t.Unix()))
		return
	}
	t := time.Unix(num, 0).Format("2006-01-02 15:04:05")
	f.ConvertText.SetText(t)
}
