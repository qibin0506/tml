package utils

import (
	"log"
	"github.com/PuerkitoBio/goquery"
)

func GetTagOrFatal(parent *goquery.Selection, name string) *goquery.Selection {
	tag, exists := GetTag(parent, name)
	if !exists {
		log.Fatalf("tag %s must be set.", name)
	}

	return tag
}

func GetTag(parent *goquery.Selection, name string) (*goquery.Selection, bool) {
	tags := parent.Find(name)
	if tags == nil || tags.Size() == 0 {
		return nil, false
	}

	tag := tags.Last()
	if tag == nil || tag.Size() == 0 {
		return nil, false
	}

	return tag, true
}

func GetTaskTag(dom *goquery.Document) *goquery.Selection {
	var tasks = dom.Find("task")
	if tasks == nil || tasks.Size() == 0 {
		log.Fatal("tag task must be set.")
	}

	task := tasks.Last()
	if task == nil || task.Size() == 0 {
		log.Fatal("tag task must be set.")
	}

	return task
}

func NextNode(node *goquery.Selection) *goquery.Selection {
	next := node.Next()
	if next == nil || next.Size() == 0 {
		return nil
	}

	return next
}

func FirstChildNode(node *goquery.Selection) *goquery.Selection {
	children := node.Children()
	if children == nil || children.Size() == 0 {
		return nil
	}

	return children.First()
}

func NodeName(node *goquery.Selection) string {
	if node == nil || node.Size() == 0 {
		return ""
	}

	return goquery.NodeName(node)
}