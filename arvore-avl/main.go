package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value  int
	left   *Node
	right  *Node
	height int
	parent *Node
}

type AVLTree struct {
	root *Node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (tree *AVLTree) remove(value int) {
	tree.root = removeNode(tree.root, value)
}

func removeNode(node *Node, value int) *Node {
	if node == nil {
		return node
	}

	if value < node.value {
		node.left = removeNode(node.left, value)
	} else if value > node.value {
		node.right = removeNode(node.right, value)
	} else {
		if node.left == nil || node.right == nil {
			var temp *Node
			if node.left != nil {
				temp = node.left
			} else {
				temp = node.right
			}

			if temp == nil {
				node = nil
			} else {
				*node = *temp
			}
		} else {
			temp := minValueNode(node.right)
			node.value = temp.value
			node.right = removeNode(node.right, temp.value)
		}
	}

	if node == nil {
		return node
	}

	node.height = max(height(node.left), height(node.right)) + 1
	balance := getBalance(node)

	if balance > 1 && getBalance(node.left) >= 0 {
		return rotateRight(node)
	}

	if balance > 1 && getBalance(node.left) < 0 {
		node.left = rotateLeft(node.left)
		return rotateRight(node)
	}

	if balance < -1 && getBalance(node.right) <= 0 {
		return rotateLeft(node)
	}

	if balance < -1 && getBalance(node.right) > 0 {
		node.right = rotateRight(node.right)
		return rotateLeft(node)
	}

	return node
}

func height(n *Node) int {
	if n == nil {
		return -1
	}
	return n.height
}

func getBalance(n *Node) int {
	if n == nil {
		return 0
	}
	return height(n.left) - height(n.right)
}

func rotateRight(y *Node) *Node {
	x := y.left
	T2 := x.right
	x.right = y
	y.left = T2

	if T2 != nil {
		T2.parent = y
	}
	x.parent = y.parent
	y.parent = x

y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1

	return x
}

func rotateLeft(x *Node) *Node {
	y := x.right
	T2 := y.left

y.left = x
	x.right = T2

	if T2 != nil {
		T2.parent = x
	}
	y.parent = x.parent
	x.parent = y

	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1

	return y
}

func (tree *AVLTree) insert(value int) {
	tree.root = insertNode(tree.root, value, nil)
}

func insertNode(node *Node, value int, parent *Node) *Node {
	if node == nil {
		return &Node{value: value, height: 0, parent: parent}
	}

	if value < node.value {
		node.left = insertNode(node.left, value, node)
	} else if value > node.value {
		node.right = insertNode(node.right, value, node)
	} else {
		return node
	}

	node.height = 1 + max(height(node.left), height(node.right))

	balance := getBalance(node)

	if balance > 1 && value < node.left.value {
		return rotateRight(node)
	}

	if balance < -1 && value > node.right.value {
		return rotateLeft(node)
	}

	if balance > 1 && value > node.left.value {
		node.left = rotateLeft(node.left)
		return rotateRight(node)
	}

	if balance < -1 && value < node.right.value {
		node.right = rotateRight(node.right)
		return rotateLeft(node)
	}

	return node
}

func (tree *AVLTree) buildFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		values := strings.Split(strings.Trim(scanner.Text(), "."), ";")
		for _, v := range values {
			num, err := strconv.Atoi(strings.TrimSpace(v))
			if err == nil {
				tree.insert(num)
			}
		}
	}
}

func (tree *AVLTree) printTree() {
	printTreeRecursive(tree.root, "", true)
}

func printTreeRecursive(node *Node, indent string, last bool) {
	if node != nil {
		fmt.Printf("%s%s (pai: %v – FB: %d)\n", indent, strconv.Itoa(node.value), 
			func() string {
				if node.parent != nil {
					return strconv.Itoa(node.parent.value)
				}
				return "raiz"
			}(), getBalance(node))

		newIndent := indent + "  "
		printTreeRecursive(node.right, newIndent, false)
		printTreeRecursive(node.left, newIndent, true)
	}
}

func minValueNode(node *Node) *Node {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current
}

func (tree *AVLTree) printNodeDetails(value int) {
	node := findNode(tree.root, value)
	if node == nil {
		fmt.Println("Nó não encontrado.")
		return
	}
	fmt.Printf("Nó: %d (FB: %d)\n", node.value, getBalance(node))
	if node.parent != nil {
		fmt.Printf("Pai: %d (FB: %d)\n", node.parent.value, getBalance(node.parent))
	} else {
		fmt.Println("Pai: Nenhum (raiz)")
	}
	if node.left != nil {
		fmt.Printf("Filho Esquerdo: %d (FB: %d)\n", node.left.value, getBalance(node.left))
	} else {
		fmt.Println("Filho Esquerdo: Nenhum")
	}
	if node.right != nil {
		fmt.Printf("Filho Direito: %d (FB: %d)\n", node.right.value, getBalance(node.right))
	} else {
		fmt.Println("Filho Direito: Nenhum")
	}
}

func findNode(node *Node, value int) *Node {
	if node == nil || node.value == value {
		return node
	}
	if value < node.value {
		return findNode(node.left, value)
	}
	return findNode(node.right, value)
}


func main() {
	tree := &AVLTree{}
	tree.buildFromFile("input.txt")

	var choice int
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Inserir nó")
		fmt.Println("2. Remover nó")
		fmt.Println("3. Localizar nó")
		fmt.Println("4. Imprimir Árvore")
		fmt.Println("5. Sair")
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Print("Digite o valor: ")
			var val int
			fmt.Scan(&val)
			tree.insert(val)
		case 2:
			fmt.Print("Digite o valor a remover: ")
			var val int
			fmt.Scan(&val)
			tree.remove(val)
		case 3:
			fmt.Print("Digite o valor a localizar: ")
			var val int
			fmt.Scan(&val)
			tree.printNodeDetails(val)
		case 4:
			tree.printTree()
		case 5:
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
}
