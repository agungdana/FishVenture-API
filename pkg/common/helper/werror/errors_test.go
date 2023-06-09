package werror_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/e-fish/api/pkg/common/helper/werror"
	"github.com/stretchr/testify/assert"
)

var exampleErr = werror.Error{
	Code:    "ErrExample",
	Message: "err example",
}

var newExampleErr = werror.Error{
	Code:    "NewErrExample",
	Message: "new err example",
}

func Test_ErrorIsSuitableWihterrors(t *testing.T) {
	assert.Error(t, exampleErr)
	fmt.Printf("exampleErr: %v\n", exampleErr)
}

func Test_AttactDetailErr(t *testing.T) {
	detail := map[string]any{"causes": "example err"}
	err := exampleErr.AttacthDetail(detail)

	assert.Equal(t, detail, err.Details)
	fmt.Printf("err: %v\n", err)
}

func Test_ErrorIs(t *testing.T) {
	ok := exampleErr.Is(DoSomething(t))
	assert.True(t, ok, "status : %v", ok)
}

func DoSomething(t *testing.T) error {
	return exampleErr
}

func Test_ErrorsIsSuitable(t *testing.T) {
	err := werror.Errors{
		Message: "example errors",
		Errors: []werror.Error{
			exampleErr, newExampleErr,
		},
	}

	assert.Error(t, err, "err: %v", err)
	fmt.Printf("err: %v\n", err)
}

func Test_AddErrors(t *testing.T) {
	errs := werror.NewError("multiple errors example")

	errs.Add(exampleErr.AttacthDetail(map[string]any{"causes": "example 1"}))
	errs.Add(exampleErr.AttacthDetail(map[string]any{"causes": "example 2", "causes2": "example 3"}))
	errs.Add(errors.New("example 3"))

	err := errs.Return()

	assert.Error(t, err, "err: %v", err)
	fmt.Printf("err: %v\n", err)

}

func Test_AddErrorsFromErrors(t *testing.T) {
	errs := werror.NewError("multiple errors example")

	errs.Add(errors.New("example 1"))
	errs.Add(DoSomethingAndReturnErrors(t))
	errs.Add(errors.New("example 1"))

	err := errs.Return()

	assert.Error(t, err, "err: %v", err)
	fmt.Println(err)

}

func DoSomethingAndReturnErrors(t *testing.T) error {
	errs := werror.NewError("multiple errors return")

	errs.Add(exampleErr.AttacthDetail(map[string]any{"causes": "example 1"}))
	errs.Add(exampleErr.AttacthDetail(map[string]any{"causes": "example 2"}))

	return errs.Return()
}
