package variables

const Help = `
MTSound - Um tocador de música simples em Go

Ajuda:

O MTSound é um tocador de música simples escrito em Go. Ele permite reproduzir músicas no formato MP3, MP4 e M4A usando o utilitário ffplay.

🎶 Comandos Disponíveis:

1. start: Inicia a reprodução da música atual. Se uma música estiver em reprodução, ela será interrompida. Se nenhum número de índice for fornecido, a reprodução começa a partir da primeira música da lista.
   
2. stop: Para a reprodução da música atual.

3. break: Pausa a reprodução da música atual. Para retomar a reprodução, use o comando start novamente.

4. next-: Reproduz a música anterior na lista.

5. next+: Reproduz a próxima música na lista.

📝 Uso:

Para utilizar o MTSound, siga estas instruções:

1. Compile o programa usando o Go compiler.
   
2. Execute o programa com o comando mtsound <comando>, onde <comando> é um dos comandos listados acima.

Exemplo:

mtsound start


Isso iniciará a reprodução da primeira música na lista.

Para mais informações, consulte a documentação ou o código-fonte do programa.

`
var Indicator int = 0