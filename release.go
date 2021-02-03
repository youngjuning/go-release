package release

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/go-version"
)

type UpdateInfo struct {
	IsUpdate bool
	Latest   string
}

// CheckUpdate 检查版本
func CheckUpdate(user string, repo string, current string) (updateInfo *UpdateInfo, err error) {
	releaseURL := fmt.Sprintf("https://github.com/%s/%s/releases/latest", user, repo)
	resp, err := http.Get(releaseURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // 为了防止内存泄漏

	current = strings.Replace(current, "v", "", 1)
	pathArr := strings.Split(resp.Request.URL.Path, "/")
	latest := strings.Replace(pathArr[len(pathArr)-1], "v", "", 1)

	currentVersion, err := version.NewVersion(current)
	if err != nil {
		return nil, err
	}
	latestVersion, err := version.NewVersion(latest)
	if err != nil {
		return nil, err
	}
	updateInfo = &UpdateInfo{
		IsUpdate: currentVersion.LessThan(latestVersion),
		Latest:   latest,
	}
	return updateInfo, nil
}
