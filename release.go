package release

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/codeskyblue/go-sh"
	"github.com/hashicorp/go-version"
)

type UpdateInfo struct {
	IsUpdate         bool
	LatestVersion    string
	LatestReleaseURL string
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
		IsUpdate:         currentVersion.LessThan(latestVersion),
		LatestVersion:    latest,
		LatestReleaseURL: releaseURL,
	}
	return updateInfo, nil
}

// RunInstaller 执行安装程序安装最新版
func RunInstaller(latestReleaseURL string, shellName string, homeDirName string) {
	args := fmt.Sprintf("curl -fsSL https://raw.githubusercontent.com/youngjuning/go-release/main/install.sh | sh -s %s %s %s", latestReleaseURL, shellName, homeDirName)
	sh.Command("bash", "-c", args).Run()
}
