# Lista para a Dinâmica "Onde vai cada dado?"

Vou propor um modelo de **Vendas no Varejo** — é intuitivo, todo mundo já viveu essa experiência como cliente, e permite distinguir claramente atributos descritivos (dimensões) de números agregáveis (métricas/fato).

## Lista embaralhada para entregar aos participantes

*(Misture esta ordem antes de imprimir/distribuir — aqui já vai embaralhada propositalmente)*

| # | Item |
|---|------|
| 1 | Quantidade vendida |
| 2 | Nome do produto |
| 3 | Valor total da venda |
| 4 | Categoria do produto |
| 5 | Nome do vendedor |
| 6 | Dia da semana |
| 7 | Valor do desconto |
| 8 | Cidade da loja |
| 9 | Nome do cliente |
| 10 | Mês |
| 11 | Custo do frete |
| 12 | Faixa etária do cliente |
| 13 | Tamanho da loja (m²) |
| 14 | Margem de lucro (R$) |
| 15 | Marca do produto |
| 16 | Região da loja |
| 17 | Turno de trabalho do vendedor |
| 18 | Ano |
| 19 | Gênero do cliente |
| 20 | Ticket médio |
| 21 | Nome da loja |
| 22 | Trimestre |
| 23 | Tempo de empresa do vendedor |
| 24 | Quantidade de itens devolvidos |
| 25 | Preço unitário do produto |

---

# Gabarito

## Tabela Fato: **Fato_Vendas**
*(os participantes devem inferir esse nome a partir das métricas)*

| Métrica |
|---|
| Quantidade vendida |
| Valor total da venda |
| Valor do desconto |
| Custo do frete |
| Margem de lucro (R$) |
| Ticket médio |
| Quantidade de itens devolvidos |

> **Nota para o facilitador:** "Ticket médio" é um ótimo item para gerar debate — tecnicamente é uma métrica *derivada* (Valor total ÷ Quantidade de vendas), não um fato aditivo bruto. Vale usá-lo para discutir a diferença entre métricas de base e métricas calculadas.

---

## Tabela Dimensão: **Dim_Produto**

| Atributo |
|---|
| Nome do produto |
| Categoria do produto |
| Marca do produto |
| Preço unitário do produto |

> **Nota:** "Preço unitário" gera debate de propósito — pode parecer métrica, mas como atributo *descritivo* do produto (preço de tabela/catálogo) pertence à dimensão. É diferente do "valor total da venda", que é transacional.

## Tabela Dimensão: **Dim_Cliente**

| Atributo |
|---|
| Nome do cliente |
| Faixa etária do cliente |
| Gênero do cliente |

## Tabela Dimensão: **Dim_Vendedor**

| Atributo |
|---|
| Nome do vendedor |
| Turno de trabalho do vendedor |
| Tempo de empresa do vendedor |

## Tabela Dimensão: **Dim_Loja**

| Atributo |
|---|
| Nome da loja |
| Cidade da loja |
| Região da loja |
| Tamanho da loja (m²) |

## Tabela Dimensão: **Dim_Tempo**

| Atributo |
|---|
| Dia da semana |
| Mês |
| Trimestre |
| Ano |

---

## Diagrama de referência (Star Schema)Vou criar o diagrama do star schema para a dinâmica.## Sugestões de condução

- **Itens "pegadinha" propositais**: *Preço unitário do produto* (atributo, não métrica — é "preço de tabela", e fica em Dim_Produto) e *Ticket médio* (métrica derivada, não atributo) costumam gerar a melhor discussão entre os grupos.
- **Inferir nomes das tabelas**: como nenhuma tabela vem rotulada, os participantes só vão conseguir nomeá-las depois de agrupar os atributos por tema (ex.: "nome, categoria, marca, preço" → só pode ser produto).
- **Tempo de 20 minutos**: sugiro ~10 min para classificação e debate em grupo, ~5 min para correção coletiva (usando o gabarito acima) e ~5 min para desenhar as linhas ligando cada dimensão à fato e reforçar visualmente o star schema.
- Se quiser aumentar a dificuldade, pode remover as quebras de linha "óbvias" da lista (embaralhar de verdade, ex. em tiras de papel) e incluir 1–2 itens ambíguos a mais propositalmente.

Quer que eu monte isso como um documento Word ou PDF para impressão (com a lista para cortar em tirinhas e o gabarito separado)?