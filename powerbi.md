Sua estrutura inicial está muito boa. O que eu ajustaria é pensar menos em “funcionalidades do Power BI” e mais em uma jornada de construção de um projeto real.

Analistas costumam aprender melhor quando conseguem responder à pergunta:

“Por que estou fazendo isso?”

Se você organizar os três dias como uma evolução de um único projeto, a retenção será muito maior.

⸻

Proposta de Estrutura

Dia 1 — Os Dados: Construindo a Fundação

Objetivo

Entender como os dados entram no Power BI e como organizá-los corretamente.

Tópicos

1. Introdução rápida (15 min)

* O que é BI
* Fluxo dos dados
* Conceito de ETL
* O que o Power BI faz

2. Importação de dados (20 min)

* Excel
* CSV
* Banco de dados (conceitual)

3. Power Query (50 min)

* Renomear colunas
* Alterar tipos
* Filtrar linhas
* Mesclar consultas
* Acrescentar consultas
* Colunas calculadas simples

4. Modelagem de dados (60 min)

* Relacionamentos
* Cardinalidade
* Direção de filtro

5. Star Schema (40 min)

* Tabela fato
* Tabela dimensão
* Boas práticas

6. Exercício prático (45 min)

Transformar um conjunto de planilhas “bagunçadas” em um modelo funcional.

Encerramento (10 min)

Criar:

* Tabela
* Matriz
* Gráfico de barras

Apenas para validar que o modelo funciona.

⸻

Dia 2 — Comunicação Visual

Objetivo

Transformar dados em informação.

⸻

1. Antes de abrir o Power BI (20 min)

Discussão:

Mostrar 2 dashboards:

* Um excelente
* Um horrível

Perguntar:

* O que chamou atenção?
* O que está confuso?
* Onde seus olhos foram primeiro?

Isso gera muita participação.

⸻

2. Escolha de visuais (45 min)

Quando usar:

* Barras
* Linhas
* Pizza
* Cartões
* KPI
* Tabelas
* Matrizes

Erros comuns.

⸻

3. Construção do relatório (45 min)

* Hierarquia visual
* Espaçamento
* Alinhamento
* Consistência

⸻

4. Design e identidade visual (30 min)

* Cores
* Contraste
* Tipografia
* Temas

⸻

5. Interações (30 min)

* Segmentações
* Cross-filter
* Drill-down

⸻

6. Publicação (20 min)

* Publicar
* Atualizar
* Compartilhar

⸻

Exercício final (50 min)

Cada grupo recebe os mesmos dados.

Missão:

Criar um dashboard executivo para um gerente.

Depois todos apresentam em 2 minutos.

⸻

Dia 3 — Inteligência e Refinamento

Objetivo

Adicionar lógica e tornar o relatório mais profissional.

⸻

1. Introdução ao DAX (60 min)

Diferença entre:

* Coluna calculada
* Medida

Medidas básicas:

* SUM
* COUNT
* DISTINCTCOUNT
* DIVIDE

⸻

2. Filtros e contexto (30 min)

Conceito fundamental.

Sem aprofundar demais.

⸻

3. Recursos avançados (45 min)

* Bookmarks
* Botões
* Navegação
* Tooltips

⸻

4. Formatação condicional (30 min)

* Cores
* Ícones
* Barras de dados

⸻

5. Dicas práticas (30 min)

Exemplos reais:

* Coluna texto para busca
* Como lidar com nulos
* Como identificar visuais vazios
* Como evitar erros de relacionamento
* Performance básica

⸻

6. Desafio Final (45 min)

Recebem um modelo incompleto.

Precisam:

* Corrigir modelagem
* Criar medidas
* Ajustar visual
* Entregar dashboard final

⸻

Dinâmicas Presenciais

Aqui está onde você pode ganhar a atenção da turma.

⸻

Dinâmica 1 — “Vocês são um Banco de Dados”

(Excelente para Star Schema)

Material

Papéis A4.

⸻

Distribuição

Aluno 1:

CLIENTE

Aluno 2:

PRODUTO

Aluno 3:

TEMPO

Aluno 4:

AGÊNCIA

Aluno 5:

VENDA

⸻

Os alunos-dimensão recebem informações.

Exemplo:

CLIENTE

ID Cliente: 100
Nome: João

⸻

PRODUTO

ID Produto: 200
Notebook

⸻

VENDA

ID Cliente: 100
ID Produto: 200
Valor: 5000

⸻

Fase 1

Monte o relacionamento fisicamente.

Use barbante.

Pergunte:

Como a venda sabe quem é o cliente?

Eles enxergam a chave estrangeira.

⸻

Fase 2

Crie um erro.

Dê dois clientes com mesmo ID.

Pergunte:

O que acontece agora?

Introduz cardinalidade.

⸻

Fase 3

Crie relacionamento muitos para muitos.

Eles próprios percebem a confusão.

⸻

Resultado

Em 20 minutos eles entendem mais modelagem do que em 1 hora de slides.

⸻

Dinâmica 2 — Dashboard de Papel

Antes de abrir o Power BI.

Entregue folhas A3.

Pergunta:

O diretor quer saber como está a empresa. O que deve aparecer?

Eles desenham.

Sem computador.

Depois compare os resultados.

Ensina:

* Hierarquia visual
* Storytelling
* Priorização

⸻

Dinâmica 3 — Detective de Dados

Entregue um conjunto de dados propositalmente ruim.

Problemas:

* Datas em texto
* Valores nulos
* Duplicidades
* Colunas inúteis

Missão:

Encontrar todos os problemas.

Vira uma competição.

⸻

Dinâmica 4 — Batalha dos Visuais

Mostre um indicador.

Exemplo:

Vendas por mês.

Pergunte:

Qual visual usar?

Cada grupo defende sua escolha.

Depois discuta vantagens e desvantagens.

⸻

Uma sugestão adicional

Reserve os últimos 15 minutos de cada aula para um quiz relâmpago em equipes.

Perguntas rápidas:

* Fato ou dimensão?
* Medida ou coluna?
* Qual visual usar?
* Qual relacionamento correto?

Você pode dar pontos para cada mesa/grupo e manter um ranking pelos 3 dias.

Isso ajuda bastante a reduzir a distração no Teams porque cria um elemento de participação contínua e competição saudável. O curso deixa de ser “assistir PowerPoint” e passa a ser “resolver problemas de BI”.