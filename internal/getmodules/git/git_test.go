package git

import (
	"fmt"
	"testing"
)

func TestFromResourceURL(t *testing.T) {
	module := "github.com/Pactionly/sample-tf-module"
	_, repoSpec, err := FromResourceURL(module)
	if repoSpec.Ref == "" {
		t.Errorf("Default Ref failed to resolve")
	}
	fmt.Print(err)
}

func TestFromResourceURLWithRef(t *testing.T) {
	module := "github.com/Pactionly/sample-tf-module?ref=main"
	_, repoSpec, err := FromResourceURL(module)
	fmt.Print(repoSpec)
	fmt.Print(err)
}
