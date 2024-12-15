package gogg

type Form struct {
	Name string
}

func NewForm(name string) *Form {
	return &Form{Name: name}
}

func (f *Form) ShowModal() {

}
