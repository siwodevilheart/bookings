package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (form *Form) Valid() bool {
	return len(form.Errors) == 0
}

func (form *Form) Required(fields ...string) {
	for _, field := range fields {
		value := form.Get(field)
		fmt.Println(value)
		if strings.TrimSpace(value) == "" {
			form.Errors.Add(field, "This field cannot be blank!")
		}
	}
}

func (form *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		form.Errors.Add(field, "This field cannot be blank!")
		return false
	}
	return true
}

func (form *Form) MinLen(field string, lenght int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < lenght {
		form.Errors.Add(field, fmt.Sprintf("This field must be atleast %d character long!", lenght))
		return false
	}
	return true
}

func (form *Form) IsEmail(field string) {
	if !govalidator.IsEmail(form.Get(field)) {
		form.Errors.Add(field, "Invalid email address!")
	}
}
