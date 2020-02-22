package bun

import (
	"fmt"

	structs "../structs"
	"github.com/fogleman/gg"
)

func RenderTree(ctx *gg.Context, tree *structs.NodeDOM) {
	//tree.Children[0] is head
	body := tree.Children[1]

	body.Style.Width = float64(ctx.Width())
	body.Style.Height = float64(ctx.Height())

	layoutDOM(ctx, body, 0)
}

func getNodeContent(NodeDOM *structs.NodeDOM) string {
	return NodeDOM.Content
}

func getElementName(NodeDOM *structs.NodeDOM) string {
	return NodeDOM.Element
}

func getNodeChildren(NodeDOM *structs.NodeDOM) []*structs.NodeDOM {
	return NodeDOM.Children
}

func walkDOM(TreeDOM *structs.NodeDOM, d string) {
	fmt.Println(d, getElementName(TreeDOM))
	nodeChildren := getNodeChildren(TreeDOM)

	for i := 0; i < len(nodeChildren); i++ {
		walkDOM(nodeChildren[i], d+"-")
	}
}

func layoutDOM(ctx *gg.Context, node *structs.NodeDOM, childIdx int) {
	nodeChildren := getNodeChildren(node)

	if node.Style.Display == "block" {
		calculateBlockLayout(ctx, node, childIdx)

		for i := 0; i < len(nodeChildren); i++ {
			layoutDOM(ctx, nodeChildren[i], i)
		}

		ctx.SetRGBA(node.Style.Color.R, node.Style.Color.G, node.Style.Color.B, node.Style.Color.A)
		ctx.DrawString(node.Content, 1, node.Style.Top+node.Style.Height)
		ctx.Fill()
	}
}

func calculateBlockLayout(ctx *gg.Context, node *structs.NodeDOM, childIdx int) {
	node.Style.Width = node.Parent.Style.Width
	if len(node.Content) > 0 {
		_, height := ctx.MeasureMultilineString(node.Content, 2)

		node.Style.Height = height
	} else {
		node.Style.Height = 0
	}

	if childIdx > 0 {
		prev := node.Parent.Children[childIdx-1]

		if prev.Style.Display == "block" {
			node.Style.Top = prev.Style.Top + prev.Style.Height
		} else {
			node.Style.Top = prev.Style.Top
		}
	}
}

func GetPageTitle(TreeDOM *structs.NodeDOM) string {
	nodeChildren := getNodeChildren(TreeDOM)
	pageTitle := "Sem Titulo"

	if getElementName(TreeDOM) == "title" {
		return getNodeContent(TreeDOM)
	}

	for i := 0; i < len(nodeChildren); i++ {
		nPageTitle := GetPageTitle(nodeChildren[i])

		if nPageTitle != "Sem Titulo" {
			pageTitle = nPageTitle
		}
	}

	return pageTitle
}