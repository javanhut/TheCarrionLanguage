package update

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	repoOwner = "javanhut"
	repoName  = "TheCarrionLanguage"
	apiBase   = "https://api.github.com"
)

type releaseAsset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
	Size               int64  `json:"size"`
}

type releaseInfo struct {
	TagName     string         `json:"tag_name"`
	Name        string         `json:"name"`
	PublishedAt time.Time      `json:"published_at"`
	HTMLURL     string         `json:"html_url"`
	Prerelease  bool           `json:"prerelease"`
	Draft       bool           `json:"draft"`
	Assets      []releaseAsset `json:"assets"`
}

type commitInfo struct {
	SHA    string `json:"sha"`
	Commit struct {
		Message   string `json:"message"`
		Committer struct {
			Date time.Time `json:"date"`
			Name string    `json:"name"`
		} `json:"committer"`
	} `json:"commit"`
	HTMLURL string `json:"html_url"`
}

// httpClient is the shared client used for GitHub API and asset downloads.
// A generous timeout covers slow CI links but avoids hanging forever.
var httpClient = &http.Client{Timeout: 5 * time.Minute}

func fetchLatestRelease() (*releaseInfo, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/releases/latest", apiBase, repoOwner, repoName)
	body, err := apiGet(url)
	if err != nil {
		return nil, err
	}
	var r releaseInfo
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, fmt.Errorf("decode release: %w", err)
	}
	return &r, nil
}

func fetchLatestCommit(branch string) (*commitInfo, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/commits/%s", apiBase, repoOwner, repoName, branch)
	body, err := apiGet(url)
	if err != nil {
		return nil, err
	}
	var c commitInfo
	if err := json.Unmarshal(body, &c); err != nil {
		return nil, fmt.Errorf("decode commit: %w", err)
	}
	return &c, nil
}

func apiGet(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "carrion-update")
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("github request failed: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read github response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("github API returned %s: %s", resp.Status, truncate(string(body), 200))
	}
	return body, nil
}

// downloadTo streams a URL to dst, creating any parent directories.
func downloadTo(url, dst string) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "carrion-update")
	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("download %s: %w", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download %s: %s", url, resp.Status)
	}
	f, err := createFile(dst)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := io.Copy(f, resp.Body); err != nil {
		return fmt.Errorf("write %s: %w", dst, err)
	}
	return nil
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
