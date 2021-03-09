package release

import (
	"testing"
)

func TestRelease(t *testing.T) {
	lowVersion := "0.0.1"
	update, err := CheckUpdate("denoland", "deno", lowVersion)
	if err != nil {
		t.Errorf("check update ouccerred error, error msg is %v.\n", err.Error())
		return
	}
	if !update.IsUpdate {
		t.Errorf("Latest version is %v.\n", update.LatestVersion) // out: Latest version is 1.8.0.
		// Run upgrade command
	} else {
		t.Logf("Latest version is %v.\n", update.LatestVersion)
	}

	if update.LatestReleaseURL == "" {
		t.Errorf("Latest Release URL is empty.\n")
		// Run upgrade command
	} else {
		t.Logf("Latest Release URL is %v.\n", update.LatestReleaseURL)
	}
}
