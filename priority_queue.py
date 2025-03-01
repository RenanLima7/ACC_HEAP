import heapq

class PriorityQueue:
    def __init__(self, capacity):
        self.capacity = capacity
        self.heap = []  # Lista para armazenar o heap
        self.entry_finder = {}  # Dicionário para rastrear elementos
        self.counter = 0  # Contador para manter a ordem de inserção

    def insert(self, task, priority):
        if len(self.heap) >= self.capacity:
            print("A fila de prioridade está cheia!")
            return
        entry = (-priority, self.counter, task)  # Heap Máximo (invertendo a prioridade)
        self.entry_finder[task] = entry
        heapq.heappush(self.heap, entry)
        self.counter += 1

    def extract_max(self):
        if not self.heap:
            print("A fila de prioridade está vazia!")
            return None
        while self.heap:
            priority, _, task = heapq.heappop(self.heap)
            if task in self.entry_finder:
                del self.entry_finder[task]
                return (task, -priority)  # Retorna prioridade positiva
        return None

    def print_heap(self):
        print("Fila de prioridade (Heap Máximo):")
        for priority, _, task in self.heap:
            print(f"{task}: {-priority}")  # Prioridade positiva

    def change_priority(self, task, new_priority):
        if task not in self.entry_finder:
            print(f"Tarefa '{task}' não encontrada na fila!")
            return
        self.remove(task)
        self.insert(task, new_priority)

    def remove(self, task):
        if task in self.entry_finder:
            del self.entry_finder[task]

    def load_from_file(self, filename):
        try:
            with open(filename, "r") as file:
                lines = file.readlines()
                total_elements = int(lines[0].strip())
                self.capacity = int(lines[1].strip())

                for line in lines[2:]:
                    task, priority = line.strip().split(",")
                    self.insert(task, int(priority))
        except Exception as e:
            print("Erro ao carregar o arquivo:", e)
