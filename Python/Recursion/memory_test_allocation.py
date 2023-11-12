import tracemalloc
from fatorial import fatorial


def start():
    tracemalloc.start()

    result = fatorial(20)

    tracemalloc.stop()

    malloc_stats = tracemalloc.get_traced_memory()

    print(result)

    print(f"Memória alocada durante a execução: {malloc_stats[0] / (1024 ** 2):.2f} MB")


start()
