from priority_queue import PriorityQueue

def menu():
    print("\n===== FILA DE PRIORIDADE COM HEAP =====")
    print("1. Inserir elementos do arquivo")
    print("2. Inserir elemento manualmente")
    print("3. Remover elemento de maior prioridade")
    print("4. Imprimir Heap")
    print("5. Alterar prioridade de um elemento")
    print("6. Sair")

def main():
    fila = PriorityQueue(capacity=10)

    while True:
        menu()
        opcao = input("Escolha uma opção: ")

        if opcao == "1":
            fila.load_from_file("input.txt")
            print("Elementos carregados do arquivo!")
        elif opcao == "2":
            tarefa = input("Nome da tarefa: ")
            prioridade = int(input("Prioridade: "))
            fila.insert(tarefa, prioridade)
        elif opcao == "3":
            elemento = fila.extract_max()
            if elemento:
                print(f"Elemento removido: {elemento[0]} com prioridade {elemento[1]}")
        elif opcao == "4":
            fila.print_heap()
        elif opcao == "5":
            tarefa = input("Nome da tarefa a modificar: ")
            nova_prioridade = int(input("Nova prioridade: "))
            fila.change_priority(tarefa, nova_prioridade)
        elif opcao == "6":
            print("Encerrando o programa...")
            break
        else:
            print("Opção inválida!")

if __name__ == "__main__":
    main()
