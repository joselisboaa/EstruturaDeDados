import inspect


# LIFO (Last In First Out)

def funcao1():
    print("Função 1 - Início")
    funcao2()
    print("Função 1 - Fim")

def funcao2():
    print("Função 2 - Início")
    funcao3()
    print("Função 2 - Fim")


def funcao3():
    print("Função 3 - Início")
    # Simulando alguma lógica
    print("Função 3 - Fim")


# Chamada inicial
funcao1()
