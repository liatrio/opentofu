package git

import (
	"net/url"
	"strings"

	"github.com/hashicorp/go-getter"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type GitGetter struct {
	gogetter getter.GitGetter
}

// Get downloads the given URL into the given directory. This always
// assumes that we're updating and gets the latest version that it can.
//
// The directory may already exist (if we're updating). If it is in a
// format that isn't understood, an error should be returned. Get shouldn't
// simply nuke the directory.
func (g *GitGetter) Get(dst string, u *url.URL) error {
	// FromResourceURL doesnt assume main branch if ref is not provided, but we will and add it to the URL
	var repoSpec *RefMetadata
	if strings.Contains(u.String(), "?ref=") {
		_, repoSpec, _ = FromResourceURL(u.String())
	} else {
		_, repoSpec, _ = FromResourceURL(u.String() + "?ref=main")
	}

	span := trace.SpanFromContext(g.gogetter.Context())
	span.SetAttributes(attribute.String("module_commit", repoSpec.Sha))
	// span.SetAttributes(attribute.String("module_source", fmt.Sprintf("%s/%s", repoSpec.Namespace, repoSpec.Repo)))
	// span.SetAttributes(attribute.String("module_ref", repoSpec.Ref))
	// span.SetAttributes(attribute.Bool("module_is_tag_ref", repoSpec.Tag))
	return g.gogetter.Get(dst, u)
}

// GetFile downloads the give URL into the given path. The URL must
// reference a single file. If possible, the Getter should check if
// the remote end contains the same file and no-op this operation.
func (g *GitGetter) GetFile(dst string, u *url.URL) error {
	return g.gogetter.GetFile(dst, u)
}

// ClientMode returns the mode based on the given URL. This is used to
// allow clients to let the getters decide which mode to use.
func (g *GitGetter) ClientMode(u *url.URL) (getter.ClientMode, error) {
	return g.gogetter.ClientMode(u)
}

// SetClient allows a getter to know it's client
// in order to access client's Get functions or
// progress tracking.
func (g *GitGetter) SetClient(c *getter.Client) {
	g.gogetter.SetClient(c)
}
