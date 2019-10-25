package cli

import (
	troubleshootv1beta1 "github.com/replicatedhq/troubleshoot/pkg/apis/troubleshoot/v1beta1"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

const VersionFilename = "version.yaml"

func writeVersionFile(path string) error {
	version := troubleshootv1beta1.SupportBundleVersion{
		ApiVersion: "troubleshoot.replicated.com/v1beta1",
		Kind:       "SupportBundle",
	}
	b, err := yaml.Marshal(version)
	if err != nil {
		return err
	}

	filename := filepath.Join(path, VersionFilename)
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}
