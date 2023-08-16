package pulp

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/osbuild/pulp-client/pulpclient"
)

type Client struct {
	client *pulpclient.APIClient
	ctx    context.Context
}

type Credentials struct {
	Username string
	Password string
}

func NewClient(url string, creds *Credentials) *Client {
	ctx := context.WithValue(context.Background(), pulpclient.ContextServerIndex, 0)
	transport := &http.Transport{}
	httpClient := http.Client{Transport: transport}

	pulpConfig := pulpclient.NewConfiguration()
	pulpConfig.HTTPClient = &httpClient
	pulpConfig.Servers = pulpclient.ServerConfigurations{pulpclient.ServerConfiguration{
		URL: url,
	}}
	client := pulpclient.NewAPIClient(pulpConfig)

	if creds != nil {
		ctx = context.WithValue(ctx, pulpclient.ContextBasicAuth, pulpclient.BasicAuth{
			UserName: creds.Username,
			Password: creds.Password,
		})
	}

	return &Client{
		client: client,
		ctx:    ctx,
	}
}

// readBody returns the body of a response as a string and ignores
// errors. Useful for returning details from failed requests.
func readBody(r *http.Response) string {
	if r == nil {
		return ""
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return ""
	}
	return string(b)
}

// UploadFile uploads the file at the given path and returns the href of the
// new artifact.
func (cl *Client) UploadFile(path string) (string, error) {
	fp, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fp.Close()
	create := cl.client.ArtifactsAPI.ArtifactsCreate(cl.ctx).File(fp)
	res, resp, err := create.Execute()
	if err != nil {
		return "", fmt.Errorf("failed to upload file %q: %s (%s)", path, err.Error(), readBody(resp))
	}

	return res.GetPulpHref(), nil
}

// ListOSTreeRepositories returns a map (repository name -> pulp href) of
// existing ostree repositories.
func (cl *Client) ListOSTreeRepositories() (map[string]string, error) {
	list, resp, err := cl.client.RepositoriesOstreeAPI.RepositoriesOstreeOstreeList(cl.ctx).Execute()
	if err != nil {
		return nil, fmt.Errorf("repository list request returned an error: %s (%s)", err.Error(), readBody(resp))
	}

	repos := make(map[string]string, list.GetCount())
	for _, repo := range list.GetResults() {
		name := repo.Name
		href := repo.GetPulpHref()
		repos[name] = href
	}

	return repos, nil
}

// CreateOSTreeRepository creates a new ostree repository with a name and description
// and returns the pulp href.
func (cl *Client) CreateOSTreeRepository(name, description string) (string, error) {
	req := cl.client.RepositoriesOstreeAPI.RepositoriesOstreeOstreeCreate(cl.ctx)
	repo := pulpclient.OstreeOstreeRepository{
		Name: name,
	}
	if description != "" {
		repo.Description = *pulpclient.NewNullableString(&description)
	}
	req = req.OstreeOstreeRepository(repo)
	result, resp, err := req.Execute()
	if err != nil {
		return "", fmt.Errorf("repository creation failed: %s (%s)", err.Error(), readBody(resp))
	}

	return result.GetPulpHref(), nil
}

// ImportCommit imports a commit that has already been uploaded to a given
// repository. The commitHref must reference a commit tarball artifact. This
// task is asynchronous. The returned value is the href for the import task.
func (cl *Client) ImportCommit(commitHref, repoHref string) (string, error) {
	req := cl.client.RepositoriesOstreeAPI.RepositoriesOstreeOstreeImportAll(cl.ctx, repoHref)
	importOptions := *pulpclient.NewOstreeImportAll(commitHref, "repo") // our commit archives always use the repo name "repo"

	result, resp, err := req.OstreeImportAll(importOptions).Execute()
	if err != nil {
		return "", fmt.Errorf("ostree commit import failed: %s (%s)", err.Error(), readBody(resp))
	}

	return result.Task, nil
}
