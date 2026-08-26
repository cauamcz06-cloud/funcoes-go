package moduloaulas

import "fmt"

func main()
//1-arrays homogemeos,mesmo tipo
//[1, 2, 3, 4, 5]-[]int
//[1.23, 4.55 ]-[]float
//["septh","rodrigo","paulo"]-[]string 

//2-maps heterogemeos, misturar tipos 
//estrutura chave-valor (chave tem um tipo e o valor outro)
//[key]= value  
//map[string] int
//{"steph":28, "rogerio":46}
//map [string] string
//{"steph":cardoso, "bento":cardoso} 

//array
//tamanho fixo de 0 ou mais elementos do mesmo tipo
//acessamos valores com indice a[0], a[1], ...
//função len() retorna o tamanho
//pelo tamanho fixo nao e muito usado

//slice (tipo de array sem tamanho fixo)
//acessamos valores com indice a[0], a[1], ...
//função len() retorna o tamanho
//função append() para adicionar valores

var array [2]string
array[0] = "hello"
array[1] = "world"
fmt.println(array[0],array[1])
fmt.println(array)

numPrimos :=[6]int{2,3,5,7,11,13}
fmt.println(numPrimos)
fmt.println(numprimos[0:3])
fmt.println(numprimos[1])

