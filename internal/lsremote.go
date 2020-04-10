package internal

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type goversion string

func (g goversion) String() string {
	return string(g)
}

func (g goversion) Major() int {
	return 1
}

func (g goversion) Minor() int {
	vs := strings.SplitN(string(g), ".", 3)
	if len(vs) == 1 {
		return 0
	}

	if strings.Contains(vs[1], "rc") {
		p := strings.SplitN(vs[1], "rc", 2)
		i, err := strconv.Atoi(p[0])
		if err != nil {
			log.Panic(err)
		}

		return i
	}

	i, err := strconv.Atoi(vs[1])
	if err != nil {
		log.Panic(err)
	}

	return i
}

func (g goversion) RCVersion() int {
	vs := strings.SplitN(string(g), "rc", 2)
	if len(vs) == 1 {
		return 1000
	}

	i, err := strconv.Atoi(vs[1])
	if err != nil {
		log.Panic(err)
	}

	return i
}

func (g goversion) Patch() int {
	vs := strings.SplitN(string(g), ".", 3)
	if len(vs) != 3 {
		return 0
	}
	i, err := strconv.Atoi(vs[2])
	if err != nil {
		log.Panic(err)
	}
	return i
}

const goDLURL = "https://golang.org/dl/"

func LSRemote(c context.Context) error {
	req, err := http.NewRequestWithContext(c, "GET", goDLURL, nil)
	if err != nil {
		return fmt.Errorf("create request failed: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("http request failed: %w", err)
	}
	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		return fmt.Errorf("html parse failed: %w", err)
	}

	var f func(*html.Node)

	var versions []goversion

	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "h2" {
			maybeVersion := n.FirstChild.Data

			if strings.HasPrefix(maybeVersion, "go1") {
				vs := strings.SplitN(maybeVersion, " ", 2)
				versions = append(versions, goversion(vs[0]))
			}
		}
		for ch := n.FirstChild; ch != nil; ch = ch.NextSibling {
			f(ch)
		}
	}

	f(doc)

	versions = sortVersions(versions)

	m := make(map[goversion]struct{})

	var vs []goversion

	for _, v := range versions {
		_, ok := m[v]
		if ok {
			continue
		}
		m[v] = struct{}{}
		vs = append(vs, v)
	}

	currentMinor := vs[0].Minor()

	fmt.Println()
	fmt.Println("available versions:")
	fmt.Println()
	fmt.Println("system")

	var buf bytes.Buffer
	for i, v := range vs {
		if currentMinor != v.Minor() {
			fmt.Fprintln(&buf)
			currentMinor = v.Minor()
		} else if i != 0 {
			fmt.Fprintf(&buf, ", ")
		}

		fmt.Fprintf(&buf, "%s", v)
	}

	fmt.Println(buf.String())

	fmt.Println()

	return nil
}

func sortVersions(versions []goversion) []goversion {
	sort.Slice(versions, func(i, j int) bool {
		if versions[i].Minor() != versions[j].Minor() {
			return versions[i].Minor() < versions[j].Minor()
		}

		if versions[i].RCVersion() != versions[j].RCVersion() {
			return versions[i].RCVersion() < versions[j].RCVersion()
		}

		return versions[i].Patch() < versions[j].Patch()
	})

	return versions
}
