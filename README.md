# Time Clock

Na necessidade do controle dos horários que os participantes do projeto de capacitação da eruda chegam
foi pensado o desenvolvimento de um sistema em GO para apenas automatizar esse controle, a princípio o
sistema vai capturar o dia e horário que for acionado tanto quanto o nome da pessoa, após isso o sistema irá
escrever em um csv o *NOME,DATA/HORÁRIO* que o participante apareceu na empresa.


A solução foi pensada de forma simples para que possamos ter nosso controle sem ter que investir em licenças
de relógio de ponto

### Para utilizar

Primeiramente, instale GO:
https://go.dev/dl/

```
git clone https://github.com/eruda-inst/time_clock
go run main.go
```


> Necessário resolução do gitignore que não deve subir o arquivo de log
