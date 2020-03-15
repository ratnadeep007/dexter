package main

import (
	// "fmt"

	"log"
	"strings"

	"github.com/kr/pretty"
	"github.com/tborg/metascraper"
)

type Metadata struct {
	property string
	value    string
}

func main() {
	url := "https://github.com/PuerkitoBio/goquery"
	var twitter []Metadata
	var opengraph []Metadata

	p, err := metascraper.Scrape(url)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(p.Title)
	metadata := p.MetaData()
	for i := 0; i < len(metadata); i++ {
		containsTwitter, _ := checkSubstrings(metadata[i].Name, "twitter")
		containsOpenGraph, _ := checkSubstrings(metadata[i].Property, "og:")
		if containsTwitter {
			h := twitterMetadataProcessor(*metadata[i])
			twitter = append(twitter, h)
		}
		if containsOpenGraph {
			h := openGraphMetadataProcessor(*metadata[i])
			opengraph = append(opengraph, h)
		}
	}
	pretty.Println(twitter)
	pretty.Println(opengraph)
}

func openGraphMetadataProcessor(og metascraper.Meta) Metadata {
	var meta Metadata
	// Checking for `og:title`
	titleContains, _ := checkSubstrings(og.Property, "og:title")
	if titleContains {
		meta.property = og.Property
		meta.value = og.Content
	}
	// Checking for `og:url`
	urlContains, _ := checkSubstrings(og.Property, "og:url")
	if urlContains {
		meta.property = og.Property
		meta.value = og.Content
	}
	// Checking for `og:image`
	imageContains, _ := checkSubstrings(og.Property, "og:image")
	if imageContains {
		meta.property = og.Property
		meta.value = og.Content
	}
	// Checking for `og:type`
	typeContains, _ := checkSubstrings(og.Property, "og:type")
	if typeContains {
		meta.property = og.Property
		meta.value = og.Content
	}
	// Checking for `og:description`
	descriptionContains, _ := checkSubstrings(og.Property, "og:description")
	if descriptionContains {
		meta.property = og.Property
		meta.value = og.Content
	}
	// Checking for `og:locale`
	localeContains, _ := checkSubstrings(og.Property, "og:locale")
	if localeContains {
		meta.property = og.Property
		meta.value = og.Content
	}
	return meta
}

func twitterMetadataProcessor(twit metascraper.Meta) Metadata {
	var meta Metadata
	// Checking for `twitter:image:src`
	imgContains, _ := checkSubstrings(twit.Name, "twitter:image:src")
	if imgContains {
		meta.property = twit.Name
		meta.value = twit.Content
	}
	// Checking for `twitter:card`
	cardContains, _ := checkSubstrings(twit.Name, "twitter:card")
	if cardContains {
		meta.property = twit.Name
		meta.value = twit.Content
	}
	// Checking for `twitter:site`
	siteContains, _ := checkSubstrings(twit.Name, "twitter:site")
	if siteContains {
		meta.property = twit.Name
		meta.value = twit.Content
	}
	// Checking for `twitter:site:id`
	sideIdContains, _ := checkSubstrings(twit.Name, "twitter:site:id")
	if sideIdContains {
		meta.property = twit.Name
		meta.value = twit.Content
	}
	// Checking for `twitter:creator`
	creatorContains, _ := checkSubstrings(twit.Name, "twitter:creator")
	if creatorContains {
		meta.property = twit.Name
		meta.value = twit.Content
	}
	// Checking for `twitter:creator:id`
	creatorIdContains, _ := checkSubstrings(twit.Name, "twitter:creator:id")
	if creatorIdContains {
		meta.property = twit.Name
		meta.value = twit.Content
	}
	// Checking for `twitter:title`
	titleContains, _ := checkSubstrings(twit.Name, "twitter:title")
	if titleContains {
		meta.property = twit.Name
		meta.value = twit.Content
	}
	// Checking for `twitter:image`
	imageContains, _ := checkSubstrings(twit.Name, "twitter:image")
	if imageContains {
		meta.property = twit.Name
		meta.value = twit.Content
	}
	// Checking for `twitter:image:alt`
	imageAltContains, _ := checkSubstrings(twit.Name, "twitter:image:alt")
	if imageAltContains {
		meta.property = twit.Name
		meta.value = twit.Content
	}
	// Checking for `twitter:player`
	playerContains, _ := checkSubstrings(twit.Name, "twitter:player")
	if playerContains {
		meta.property = twit.Name
		meta.value = twit.Content
	}
	// Checking for `twitter:player:width`
	playerWidthContains, _ := checkSubstrings(twit.Name, "twitter:player:width")
	if playerWidthContains {
		meta.property = twit.Name
		meta.value = twit.Content
	}
	// Checking for `twitter:player:height`
	playerHeightContains, _ := checkSubstrings(twit.Name, "twitter:player:height")
	if playerHeightContains {
		meta.property = twit.Name
		meta.value = twit.Content
	}
	// Checking for `twitter:player:stream`
	playerStreamContains, _ := checkSubstrings(twit.Name, "twitter:player:stream")
	if playerStreamContains {
		meta.property = twit.Name
		meta.value = twit.Content
	}
	// Checking for `twitter:app:name:ipad`
	ipadAppNameContains, _ := checkSubstrings(twit.Name, "twitter:app:name:ipad")
	if ipadAppNameContains {
		meta.property = twit.Name
		meta.value = twit.Content
	}
	// Checking for `twitter:app:id:ipad`
	ipadAppIdContains, _ := checkSubstrings(twit.Name, "twitter:app:id:ipad")
	if ipadAppIdContains {
		meta.property = twit.Name
		meta.value = twit.Content
	}
	// Checking for `twitter:app:url:ipad`
	ipadAppUrlContains, _ := checkSubstrings(twit.Name, "twitter:app:url:ipad")
	if ipadAppUrlContains {
		meta.property = twit.Name
		meta.value = twit.Content
	}
	// Checking for `twitter:app:name:googleplay`
	googlePlayNameContains, _ := checkSubstrings(twit.Name, "twitter:app:name:googleplay")
	if googlePlayNameContains {
		meta.property = twit.Name
		meta.value = twit.Content
	}
	// Checking for `twitter:app:id:googleplay`
	googlePlayIdContains, _ := checkSubstrings(twit.Name, "twitter:app:id:googleplay")
	if googlePlayIdContains {
		meta.property = twit.Name
		meta.value = twit.Content
	}
	// Checking for `twitter:app:url:googleplay`
	googlePlayUrlContains, _ := checkSubstrings(twit.Name, "twitter:app:url:googleplay")
	if googlePlayUrlContains {
		meta.property = twit.Name
		meta.value = twit.Content
	}
	return meta
}

func checkSubstrings(str string, subs ...string) (bool, int) {

	matches := 0
	isCompleteMatch := true

	for _, sub := range subs {
		if strings.Contains(str, sub) {
			matches += 1
		} else {
			isCompleteMatch = false
		}
	}

	return isCompleteMatch, matches
}
