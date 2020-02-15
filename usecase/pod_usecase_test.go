package usecase_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test指定したネームスペースのPodの一覧を取得することができる(t *testing.T) {

	expected := []string{"sample-deployment-7bb5fc6bc6-9btmd","sample-deployment-7bb5fc6bc6-hhd2l"}
	assert.Equal(t, expected, actual)
}
