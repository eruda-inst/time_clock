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
mkdir /home/[YOUR USER]/go
mkdir /home/[YOUR USER]/go/bin
mkdir /home/[YOUR USER]/go/pkg
mkdir /home/[YOUR USER]/go/src
cd /home/[YOUR USER]/go/src
git clone https://github.com/eruda-inst/time_clock
go run main.go
```


### v1.0

- Sistema funciona no CLI
- Faz registro de nome e data em um arquivo csv
