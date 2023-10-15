## Lista encadeada x Array
### Array 

Desvantagens:
- É sequencial e necessita de uma sequência de memória
- Precisa ser realocado caso ele cresça e não vá caber mais na sequência de memória (crescimento lento)
Vantagens:
- Um elemento podem ser acessado sem precisar tomar o conhecimento dos outros elementos

### Lista encadeada

Caracteristicas:
- São boas se você precisar ler todos os itens (sem pular itens)

Desvantagens:
- Para ser acessado o elemento n de uma lista encadeada teremos que passar pelos n-1 elementos anteriores pois cada elemento guardará o endereço de memória do posterior
Vantagens:
- Por não ser sequencial não depende de uma sequência de endereços de memória estejam livre para ser alocado e nem precisa ser realocado pois pode ficar em qualquer endereço (no entanto que esteja livre)

