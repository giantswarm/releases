package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/giantswarm/apiextensions/pkg/apis/release/v1alpha1"
	"sigs.k8s.io/yaml"
)

func handleDocument(provider string, document []byte) {
	var release v1alpha1.Release
	err := yaml.Unmarshal(document, &release)
	if err != nil {
		fmt.Println(err)
		os.Exit(4)
	}

	if release.Kind == "Release" {
		release.Annotations["giantswarm.io/release-notes"] = fmt.Sprintf("https://github.com/giantswarm/releases/tree/master/%s/%s", provider, release.Name)

		r, err := yaml.Marshal(release)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		fmt.Printf("%s\n---\n", r)
	}
}

func main() {
	provider := os.Args[2]

	var buf bytes.Buffer
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				buf.WriteString(line)
				handleDocument(provider, buf.Bytes())
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		if strings.TrimSpace(line) == "---" {
			handleDocument(provider, buf.Bytes())
			buf.Reset()
		} else {
			buf.WriteString(line)
		}
	}
}
