package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type HashTable struct {
	size       int
	table      [][]int
	openTable  []int  
	occupied   []bool 
	method     string 
	hashMethod string 
	count      int    
}

func NewHashTable(size int, method string, hashMethod string) *HashTable {
	ht := &HashTable{
		size:       size,
		method:     method,
		hashMethod: hashMethod,
		count:      0,
	}

	if method == "encadeamento" {
		ht.table = make([][]int, size)
	} else {
		ht.openTable = make([]int, size)
		ht.occupied = make([]bool, size)
	}
	return ht
}

func (ht *HashTable) hash(value int) int {
	if ht.hashMethod == "divisao" {
		return value % ht.size
	} else {
		A := (math.Sqrt(5) - 1) / 2
		return int(math.Floor(float64(ht.size) * (float64(value)*A - math.Floor(float64(value)*A))))
	}
}

func (ht *HashTable) Insert(value int) {
	index := ht.hash(value)

	if ht.method == "encadeamento" {
		ht.table[index] = append(ht.table[index], value)
	} else { // Endereçamento aberto (sondagem linear)
		for i := 0; i < ht.size; i++ {
			newIndex := (index + i) % ht.size
			if !ht.occupied[newIndex] {
				ht.openTable[newIndex] = value
				ht.occupied[newIndex] = true
				break
			}
		}
	}
	ht.count++
}

func (ht *HashTable) Remove(value int) {
	index := ht.hash(value)

	if ht.method == "encadeamento" {
		for i, v := range ht.table[index] {
			if v == value {
				ht.table[index] = append(ht.table[index][:i], ht.table[index][i+1:]...)
				ht.count--
				return
			}
		}
	} else {
		for i := 0; i < ht.size; i++ {
			newIndex := (index + i) % ht.size
			if ht.occupied[newIndex] && ht.openTable[newIndex] == value {
				ht.occupied[newIndex] = false
				ht.count--
				return
			}
		}
	}
}

func (ht *HashTable) Search(value int) (bool, int) {
	index := ht.hash(value)
	colisions := 0

	if ht.method == "encadeamento" {
		for _, v := range ht.table[index] {
			if v == value {
				return true, colisions
			}
			colisions++
		}
	} else {
		for i := 0; i < ht.size; i++ {
			newIndex := (index + i) % ht.size
			if ht.occupied[newIndex] {
				if ht.openTable[newIndex] == value {
					return true, colisions
				}
				colisions++
			}
		}
	}
	return false, colisions
}

func (ht *HashTable) Print() {
	fmt.Println("Estado atual da Tabela Hash:")
	if ht.method == "encadeamento" {
		for i, bucket := range ht.table {
			fmt.Printf("%d: %v\n", i, bucket)
		}
	} else {
		for i := 0; i < ht.size; i++ {
			if ht.occupied[i] {
				fmt.Printf("%d: %d\n", i, ht.openTable[i])
			} else {
				fmt.Printf("%d: []\n", i)
			}
		}
	}
	fmt.Printf("Fator de carga: %.2f\n", ht.LoadFactor())
}

// Fator de Carga
func (ht *HashTable) LoadFactor() float64 {
	return float64(ht.count) / float64(ht.size)
}

func LoadFromFile(filename string, method, hashMethod string) *HashTable {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erro ao abrir arquivo:", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		fmt.Println("Arquivo vazio ou com erro.")
		return nil
	}

	numbers := strings.Split(scanner.Text(), ";")
	size := len(numbers)

	ht := NewHashTable(size, method, hashMethod)

	for _, num := range numbers {
		num = strings.TrimSpace(num)
		if num == "" {
			continue
		}
		value, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Erro ao converter valor:", num)
			continue
		}
		ht.Insert(value)
	}

	return ht
}

func main() {
	var method, hashMethod string

	fmt.Print("Escolha o método de colisão (encadeamento/aberto): ")
	fmt.Scan(&method)

	fmt.Print("Escolha a função de dispersão (divisao/multiplicacao): ")
	fmt.Scan(&hashMethod)

	ht := LoadFromFile("input.txt", method, hashMethod)
	if ht == nil {
		fmt.Println("Erro ao carregar a tabela hash do arquivo. Terminando o programa.")
		return
	}

	var option int
	for {
		fmt.Println("\n1. Inserir Manualmente\n2. Remover\n3. Buscar\n4. Imprimir\n5. Sair")
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			var value int
			fmt.Print("Digite um valor: ")
			fmt.Scan(&value)
			ht.Insert(value)
		case 2:
			var value int
			fmt.Print("Digite um valor para remover: ")
			fmt.Scan(&value)
			ht.Remove(value)
		case 3:
			var value int
			fmt.Print("Digite um valor para buscar: ")
			fmt.Scan(&value)
			found, collisions := ht.Search(value)
			fmt.Printf("Encontrado: %v, Colisões: %d\n", found, collisions)
		case 4:
			ht.Print()
		case 5:
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida!")
		}
	}
}
