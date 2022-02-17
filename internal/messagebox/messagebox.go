package messagebox

type IMessageBox interface {
	ShowMessage(msg string) error
}
