package optional

type String struct {
	string
	present bool
}

func OfString(s string) String {
	return String{string: s, present: true}
}

func (o *String) Set(s string) {
	o.string = s
	o.present = true
}

func (o *String) String() string {
	return o.string
}

func (o *String) Present() bool {
	return o.present
}

func (o *String) OrElse(s string) string {
	if o.present {
		return o.string
	}
	return s
}
