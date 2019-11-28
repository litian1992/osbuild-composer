package jobqueue

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/google/uuid"
	"github.com/osbuild/osbuild-composer/internal/distro"
	"github.com/osbuild/osbuild-composer/internal/pipeline"
	"github.com/osbuild/osbuild-composer/internal/target"
	"github.com/osbuild/osbuild-composer/internal/upload/awsupload"
)

type Job struct {
	ID       uuid.UUID          `json:"id"`
	Pipeline *pipeline.Pipeline `json:"pipeline"`
	Targets  []*target.Target   `json:"targets"`
}

type JobStatus struct {
	Status string `json:"status"`
}

func (job *Job) Run(d distro.Distro) (error, []error) {
	build := pipeline.Build{
		Runner: d.Runner(),
	}

	buildFile, err := ioutil.TempFile("", "osbuild-worker-build-env-*")
	if err != nil {
		return err, nil
	}
	defer os.Remove(buildFile.Name())

	err = json.NewEncoder(buildFile).Encode(build)
	if err != nil {
		return err, nil
	}

	cmd := exec.Command(
		"osbuild",
		"--store", "/var/cache/osbuild-composer/store",
		"--build-env", buildFile.Name(),
		"--json", "-",
	)
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err, nil
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err, nil
	}

	err = cmd.Start()
	if err != nil {
		return err, nil
	}

	err = json.NewEncoder(stdin).Encode(job.Pipeline)
	if err != nil {
		return err, nil
	}
	stdin.Close()

	var result struct {
		TreeID   string `json:"tree_id"`
		OutputID string `json:"output_id"`
	}
	err = json.NewDecoder(stdout).Decode(&result)
	if err != nil {
		return err, nil
	}

	err = cmd.Wait()
	if err != nil {
		return err, nil
	}

	var r []error

	for _, t := range job.Targets {
		switch options := t.Options.(type) {
		case *target.LocalTargetOptions:
			err = os.MkdirAll(options.Location, 0755)
			if err != nil {
				r = append(r, err)
				continue
			}

			cp := exec.Command("cp", "-a", "-L", "/var/cache/osbuild-composer/store/refs/"+result.OutputID+"/.", options.Location)
			cp.Stderr = os.Stderr
			cp.Stdout = os.Stdout
			err = cp.Run()
			if err != nil {
				r = append(r, err)
				continue
			}
		case *target.AWSTargetOptions:
			a, err := awsupload.New(options.Region, options.AccessKeyID, options.SecretAccessKey)
			if err != nil {
				r = append(r, err)
				continue
			}

			_, err = a.Upload("/var/cache/osbuild-composer/store/refs/"+result.OutputID+"/image.ami", options.Bucket, options.Key)
			if err != nil {
				r = append(r, err)
				continue
			}

			/* TODO: communicate back the AMI */
			_, err = a.Register(t.ImageName, options.Bucket, options.Key)
			if err != nil {
				r = append(r, err)
				continue
			}
		case *target.AzureTargetOptions:
		default:
			r = append(r, fmt.Errorf("invalid target type"))
		}
		r = append(r, nil)
	}

	return nil, r
}
