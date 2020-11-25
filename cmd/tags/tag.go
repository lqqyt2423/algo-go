package tags

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"

	"github.com/lqqyt2423/algo-go/pkg/leetcode"
)

var tagTpl = template.Must(template.New("tag").Parse(tagStr))

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func Run() {
	tags, err := leetcode.GetTags()
	if err != nil {
		log.Fatal(err)
	}
	wg := sync.WaitGroup{}
	sb := strings.Builder{}
	for _, tag := range tags {
		if tag.TranslatedName == "" {
			tag.TranslatedName = tag.Name
		}
		fp := filepath.Join("./toc", tag.Slug+".md")

		sb.WriteString(fmt.Sprintf("- [%s](%s)\n", tag.TranslatedName, fp))

		tag := tag
		wg.Add(1)
		go func() {
			if fileExists(fp) {
				wg.Done()
				return
			}
			var content bytes.Buffer
			err := tagTpl.Execute(&content, tag)
			if err != nil {
				log.Fatal(err)
			}
			ioutil.WriteFile(fp, content.Bytes(), 0644)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(sb.String())
}

var tagStr = `# {{ .TranslatedName }}

<!--- table -->

`
