package main

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// Show a navigable tree view of the current directory.
func main() {

	root := tview.NewTreeNode("root").SetColor(tcell.ColorLightBlue)
	nodeA := tview.NewTreeNode("a").
		SetReference("root/a").
		SetSelectable(true).
		SetColor(tcell.ColorLightBlue).
		SetExpanded(false)
	nodeB := tview.NewTreeNode("b").
		SetReference("root/b").
		SetSelectable(false).
		SetColor(tcell.ColorLightBlue)
	nodeC := tview.NewTreeNode("c").
		SetReference("root/a/c").
		SetSelectable(false).
		SetColor(tcell.ColorLightBlue)

	root.AddChild(nodeA)
	root.AddChild(nodeB)
	nodeA.AddChild(nodeC)

	tree := tview.NewTreeView().SetRoot(root).SetCurrentNode(root)
	// If a directory was selected, open it.
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}
		children := node.GetChildren()
		if len(children) == 0 {
			// Load and show files in this directory.
			path := reference.(string)
			fmt.Printf("Selected: %s", path)
		} else {
			// Collapse if visible, expand if collapsed.
			node.SetExpanded(!node.IsExpanded())
		}
	})

	if err := tview.NewApplication().SetRoot(tree, true).Run(); err != nil {
		panic(err)
	}
}
