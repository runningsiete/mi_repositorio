package main

import (
	"fmt"
	"strings"
	"strconv"
//	"os"
)

func main() {
	fmt.Printf("INICIO FUNCIÓN TLV \n")
	
	var entrada = "A0511AB398765UJ1N230200"
	fmt.Println("Entrada:", entrada )

	var campos  = make(map[string]string)
    var error_TLV = ""

	campos , error_TLV = obtener_campos_TLV( entrada );
	
	for key, element := range campos {
        fmt.Println("Key:", key, "=>", element )
    }
	fmt.Println("Error  :", error_TLV)

	fmt.Println("FIN FUNCIÓN TLV")
}

func validar_entrada(x string) string{
	error := "" 
	if len(x) == 0 {
		error = "Entrada vacía !!" 
		return error
	}
	if len(strings.TrimSpace(x)) == 0 {
		error = "La Entrada tiene solo espacios en blancos !!" 
		return error
	}
	if len(x) < 6 {
		error = "La Entrada no tiene la longuitud mínima de 6 caracteres !!" 
		return error
	}
	return error
}

func obtener_campos_TLV(entrada_TLV string) ( map[string]string , string ) {	
	var mapa = make(map[string]string)
	error := validar_entrada(entrada_TLV)

	if len(error) > 0 {
		return mapa, error;
	}
	mapa, error = procesar_TLV(entrada_TLV)
	return mapa, error;
}

func isNumeric(s string) bool {
    _, err := strconv.ParseInt(s, 0, 64)
    return err == nil
}

func procesar_TLV(entrada_TLV string) ( map[string]string , string ) {	
	var mapa = make(map[string]string)
	var error = ""
	clave := 1
	
	for len(entrada_TLV) >= 6 {

		// --- VALIDAR TIPO DE DATO -------------------------------
		tipoDeDato := entrada_TLV[0:1]
		if tipoDeDato != "A" && tipoDeDato != "N"  {
			error = "El Tipo "+ tipoDeDato +" no es A, tampoco es N"
			return mapa , error;
		}
		entrada_TLV = entrada_TLV[1:len(entrada_TLV)]

		// --- VALIDAR NUMERO DE CAMPO ----------------------------
		nroDeCampo := entrada_TLV[0:2]
		if !isNumeric(nroDeCampo)   {
			error = "El número de campo "+ nroDeCampo + " no es número"
			return mapa , error;
		}
		entrada_TLV = entrada_TLV[2:len(entrada_TLV)]

		// --- VALIDAR EL TIPO DE DATO DEL LARGO ----------------------------------
		largo := entrada_TLV[0:2]
		if !isNumeric(largo)   {
			error = "El largo "+ largo + " no es número"
			return mapa , error;
		}
		entrada_TLV = entrada_TLV[2:len(entrada_TLV)]

		// --- VALIDAR LONGUITUD -------------------------------	
		n, _ := strconv.ParseInt(largo, 10, 0)
		cantidad := int64(len(entrada_TLV)) 
		
		if n > cantidad {
			error = "No hay suficientes caracteres para obtener el valor de largo " + largo 
			return mapa , error;
		}

		// --- OBTENER EL VALOR -------------------------------	
		valor := entrada_TLV[0:n]

		// --- VALIDAR EL VALOR -------------------------------
		if tipoDeDato == "N"  {
			if !isNumeric(valor)   {
				error = "El valor "+ valor + " no es número"
				return mapa , error;
			}
		}

		// --- GUARDAR CAMPO TLV 
		campo :=  "Campo:" + nroDeCampo + " Tipo:" + tipoDeDato + " Largo:" + largo + " Valor:" + valor;		
		mapa[strconv.Itoa(clave)] = campo   

		// SIGUIENTE ENTRADA
		clave = clave + 1
		entrada_TLV = entrada_TLV[n:len(entrada_TLV)]
		// *** fmt.Println("Siguiente Entrada:", entrada_TLV )

	}//END FOR

	if len(entrada_TLV) > 0 && len(entrada_TLV) < 6 {
		error = "La Entrada "+ entrada_TLV +" no tiene la longuitud mínima de 6 caracteres"
	} 	

	// --- ENVIAR RESPUESTA
	return mapa , error;
}