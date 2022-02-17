package messagebox

import "os/exec"

type MessageBox struct{}

func NewMessageBox() IMessageBox {
	return &MessageBox{}
}

func (m *MessageBox) ShowMessage(msg string) error {
	return exec.Command("notify-send", msg).Run()
}
