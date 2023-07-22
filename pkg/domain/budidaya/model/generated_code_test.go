package model_test

import (
	"testing"

	"github.com/e-fish/api/pkg/domain/budidaya/model"
	"github.com/stretchr/testify/assert"
)

func TestCodeBudidaya_GeneratedCode(t *testing.T) {
	type Args struct {
		Name      string
		PondName  string
		CodeExist string
		Setup     func(pondName, code string)
	}

	args := []Args{
		{
			Name:      "SuccesWithCodeExist",
			PondName:  "Agung",
			CodeExist: "Agung/2023/1",
			Setup: func(pondName, code string) {
				code, err := model.GeneratedCodeBudidaya(pondName, code)
				assert.NoError(t, err)
				assert.NotNil(t, code)
			},
		}, {
			Name:      "FailedCodeNotSuportedTypeCode",
			PondName:  "Agung",
			CodeExist: "Agung/1",
			Setup: func(pondName, code string) {
				code, err := model.GeneratedCodeBudidaya(pondName, code)
				assert.Error(t, err)
				assert.Equal(t, "", code)
			},
		}, {
			Name:      "FailedCodeNotSuportedTypeCode",
			PondName:  "Agung",
			CodeExist: "Agung/2023/as",
			Setup: func(pondName, code string) {
				code, err := model.GeneratedCodeBudidaya(pondName, code)
				assert.Error(t, err)
				assert.Equal(t, "", code)
			},
		}, {
			Name:      "SuccessNewCode",
			PondName:  "Agung",
			CodeExist: "",
			Setup: func(pondName, code string) {
				code, err := model.GeneratedCodeBudidaya(pondName, code)
				assert.NoError(t, err)
				assert.NotNil(t, code)
			},
		},
	}

	for _, v := range args {
		t.Run(v.Name, func(t *testing.T) {
			v.Setup(v.PondName, v.CodeExist)
		})
	}
}
