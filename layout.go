package main

import (
    "fmt"
    //"syscall/js"
    //"math/rand"
    //"strconv"

)


type Node struct {
    parent *Node
    previous *Node
    children []*Node
    obj  *tBtn
}
 


	/*self.node = node
        self.parent = parent
        self.previous = previous
        self.children = []*/
        
        
func DrawLayout(node *Node){
	var parent string = "none"
	if node.parent != nil {
		if node.parent.obj != nil {
			parent = node.parent.obj.caption
		}
	}
	if node.obj != nil && node.obj.visible  {
		fmt.Println("Draw: " + parent + node.obj.caption)
		node.obj.Draw()
	} else {
		fmt.Println("Draw: none")
	}
	if node.children != nil {
		for i := 0; i < len(node.children); i++ { 
			DrawLayout(node.children[i])
		}
	}
	return
}
