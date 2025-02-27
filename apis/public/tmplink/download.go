package tmplink

import (
	"fmt"
	"regexp"
	"transfer/apis"
)

var (
	matcher = regexp.MustCompile("(https://)?tmp\\.link/f/[0-9a-z]{13}")
	regex   = regexp.MustCompile("[0-9a-z]{13}")
)

func (b tmpLink) LinkMatcher(v string) bool {
	return matcher.MatchString(v)
}

func (b tmpLink) DoDownload(link string, config apis.DownConfig) error {
	if config.Ticket == "" {
		return fmt.Errorf("%s: token is required.\n", link)
	}
	err := b.download(link, config)
	if err != nil {
		return fmt.Errorf("download failed on %s, returns %s\n", link, err)
	}
	return nil
}

func (b tmpLink) download(v string, config apis.DownConfig) error {
	fileID := regex.FindString(v)
	link := fmt.Sprintf("https://send.tmp.link/connect-%s-%s", config.Ticket, fileID)
	config.Parallel = 1 // force
	err := apis.DownloadFile(&apis.DownloaderConfig{
		Link:     link,
		Config:   config,
		Modifier: apis.AddHeaders,
	})
	if err != nil {
		return err
	}
	return nil
}
