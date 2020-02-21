<h1>Función TLV</h1>
Es una funcion implementada con el lenguaje de programación GO: recibe una cadena de caracteres la cual contiene multiples campos en el formato TLV y genera un map[string]string con los campos TLV encontrados en la cadena. El código fuente está en el archivo TLV.go y en el archivo TLV_test.go está el Test de esta función. 

<br>
<h2>Formato de los campos TLV</h2>
Cada campo TLV esta compuesto por 3 partes:
<br>

<b>Tipo:</b> El tipo tiene un largo de 3 caracteres donde el primer caracter indica el tipo de dato (A: Alfanumérico y N: Numérico) y dos caracteres para indicar el numero de campo Ejemplo: "01" o "15".
<br>

<b>Largo:</b> 2 caracteres que indican el largo del valor, este campo es importante puesto que indica cuantos caracteres leer a continuación.
<br>

<b>Valor:</b> Este es el valor del campo, el cual corresponde al valor del campo, su largo esta determinado por la porción Largo.
<br>

Ejemplo:

Para "A0511AB398765UJ1N230200" Los campos son:

05 de tipo Alfanumérico de largo 11 y valor AB398765UJ1
23 de tipo Numérico de largo 2 y valor 00
<br>

<h1>Instrucciones para ejecutar la Función TLV en Travis CI</h1>
Hacer click en el siguiente link para ver el resultado de la ejecución de la Función TLV   https://travis-ci.org/runningsiete/mi_repositorio/jobs/653483437
<br>
En travis-ci: en el jobs #5.1 debería verse el siguiente resultado de la Función TLV: 
<br>
<img src="https://github.com/runningsiete/mi_repositorio/blob/master/Ejemplo_Resultado_del_Test.png" alt="Test Ejemplo">
