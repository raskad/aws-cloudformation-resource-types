package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type awsCloudformationDocJSON struct {
	Contents []awsCloudformationDocJSONMap
}

type awsCloudformationDocJSONMap struct {
	Title    string
	Href     string
	Contents []awsCloudformationDocJSONMapInner
}

type awsCloudformationDocJSONMapInner struct {
	Title    string
	Href     string
	Contents []awsCloudformationDocJSONMapEnd
}

type awsCloudformationDocJSONMapEnd struct {
	Title string
	Href  string
}

func main() {
	cloudformationServices, err := getCloudformationServices()
	if err != nil {
		fmt.Println("Error fetching cloudformation services: ", err)
		os.Exit(1)
	}
	var resources = []string{}
	for _, service := range cloudformationServices {
		url := fmt.Sprintf("https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/toc-%v.json", service)
		r, err := getResourceTypes(url)
		if err != nil {
			fmt.Println("Error getting resource tpyes for url: ", url)
			fmt.Println("Error: ", err)
			os.Exit(1)
		}
		resources = append(resources, r...)
	}
	for _, resource := range resources {
		fmt.Println(resource)
	}
}

func getCloudformationServices() (cloudformationServices []string, err error) {
	url := "https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.partial.html"
	resp, err := http.Get(url)
	if err != nil {
		return cloudformationServices, err
	}
	defer resp.Body.Close()
	if err != nil {
		return cloudformationServices, err
	}
	z := html.NewTokenizer(resp.Body)
	for {
		token := z.Next()
		switch {
		case token == html.ErrorToken:
			return cloudformationServices, nil
		case token == html.StartTagToken:
			token := z.Token()
			if token.Data == "li" {
				_ = z.Next()
				token := z.Token()
				if token.Data == "a" {
					service := token.Attr[0].Val
					service = strings.Replace(service, "./", "", -1)
					service = strings.Replace(service, ".html", "", -1)
					if service != "cfn-reference-shared" {
						cloudformationServices = append(cloudformationServices, service)
					}
				}
			}
		}
	}
}

func getResourceTypes(url string) (resources []string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return resources, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resources, err
	}
	var doc awsCloudformationDocJSON
	err = json.Unmarshal(body, &doc)
	if err != nil {
		return resources, err
	}
	for _, resource := range doc.Contents[0].Contents {
		resources = append(resources, resource.Title)
	}
	return
}
