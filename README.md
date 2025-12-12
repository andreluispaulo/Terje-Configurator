# Terje Mod Configurator

Este é um editor de configuração avançado e fácil de usar, projetado especificamente para gerenciar os arquivos `.cfg` e `.xml` do pacote de mods **Terje** para DayZ.

Desenvolvido para oferecer uma experiência visual moderna e segura, substituindo a edição manual de arquivos de texto propensa a erros por uma interface gráfica intuitiva, dinâmica e robusta.

## 🔗 Links Oficiais

- **Desenvolvedor do Mod (Steam Workshop):** [Terje Bruøygard - Workshop Items](https://steamcommunity.com/id/terjebruoygard/myworkshopfiles/?appid=221100)

## ✨ Funcionalidades Principais

### 🖥️ Interface Moderna e Dinâmica
- **UI Responsiva:** Construída com **Vue.js** e **Tailwind CSS** (Tema Slate Dark) para uma experiência visual agradável e profissional.
- **Geração Dinâmica de Formulários:** O sistema lê os arquivos de configuração e cria os campos de edição automaticamente. Se o mod for atualizado com novos parâmetros, o configurador os exibirá sem necessidade de atualização do software.
- **Detecção Inteligente de Tipos:** Identifica automaticamente campos booleanos (switches), numéricos e de texto.
- **Suporte a Ícones Personalizados:** Utiliza ícones Lucide para uma interface limpa.

### 🛠️ Edição de Arquivos (Core)
- **Parser Cirúrgico:** Mantém a estrutura original dos arquivos `.cfg` e `.xml` intacta (comentários, espaçamento, formatação). Apenas os valores dos parâmetros são alterados.
- **Suporte a Múltiplos Formatos:**
  - **.CFG:** Edição completa de parâmetros chave=valor.
  - **.XML:** Editor hierárquico com suporte a atributos e estruturas aninhadas (Aviso: Em desenvolvimento ativo).
- **Navegação em Árvore:** Menu lateral dinâmico que reflete a estrutura de pastas e arquivos da pasta `TerjeSettings`.

### 💾 Segurança e Histórico
- **Snapshot History:** Cada salvamento gera automaticamente um ponto de restauração no banco de dados SQLite embutido.
- **Restauração Instantânea:** Permite reverter qualquer arquivo para versões anteriores com um clique, garantindo segurança contra configurações erradas.
- **Validação de Tipos:** Previne erros comuns, como inserir texto em campos numéricos ou booleanos.

### 🌍 Internacionalização (i18n)
- Suporte nativo a 3 idiomas, alteráveis em tempo real:
  - 🇧🇷 Português (Brasil)
  - 🇺🇸 Inglês (US)
  - 🇪🇸 Espanhol (ES)
- **Tradução de Descrições:** Sistema inteligente que traduz as descrições dos parâmetros (tooltips) quando disponível, com fallback para o inglês original.

### 📦 Distribuição e Instalação
- **Single EXE:** O aplicativo é distribuído como um único executável portátil (`.exe`).
- **Zero Dependências:** Não requer instalação de Java, .NET ou bibliotecas externas.
- **Frontend Embutido:** A interface web é compilada e embutida dentro do binário Go.
- **Banco de Dados Automático:** O arquivo `history.db` é criado automaticamente na primeira execução.

## 🚀 Como Usar

1. Baixe o arquivo `TerjeConfigurator_Final_v1.exe`.
2. Coloque-o na raiz do seu servidor ou pasta de projeto, **ao lado** da pasta `TerjeSettings`.
   ```
   /SeuServidor
   ├── TerjeConfigurator_Final_v1.exe
   └── TerjeSettings/
       ├── Core.cfg
       ├── Skills.cfg
       └── ...
   ```
3. Execute o aplicativo.
4. O navegador abrirá automaticamente com a interface de configuração.
5. Selecione um arquivo no menu lateral e comece a editar!

## 🔧 Tecnologias Utilizadas

- **Backend:** Go (Golang) + SQLite + GORM
- **Frontend:** Vue 3 + TypeScript + Vite + Tailwind CSS + Pinia
- **Ícones:** Lucide Vue Next
- **Compilação:** Go Embed + RSRC (para ícones do Windows)

---
*Desenvolvido para a comunidade DayZ.*
