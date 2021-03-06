package openshift

import (
	"testing"
	"github.com/magiconair/properties/assert"
)

const (
	name = "hello-world"
	namespace = "demo-test"
	fromNamespace = "demo-dev"
	version = "v1"
	fullName = name + ":" + version
)


func TestCreateTags(t *testing.T) {
	ist, err := NewImageStreamTags(name, version, namespace)
	assert.Equal(t, nil, err)
	is, err := ist.Create(fromNamespace)
	assert.Equal(t, nil, err)
	assert.Equal(t, fullName, is.Name)
}

func TestGetTags(t *testing.T) {
	ist, err := NewImageStreamTags(name, version, namespace)
	assert.Equal(t, nil, err)
	is, err := ist.Get()
	assert.Equal(t, nil, err)
	assert.Equal(t, fullName, is.Name)
}

func TestDeleteTag(t *testing.T)  {
	ist, err := NewImageStreamTags(name, version, namespace)
	assert.Equal(t, nil, err)
	err = ist.Delete()
	assert.Equal(t, nil, err)
}