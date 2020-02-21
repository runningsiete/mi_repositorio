package main

import "testing"
import "fmt"

func TestObtener_campos_TLV(t *testing.T) {	
	var campos  = make(map[string]string)
	var error_TLV = ""
	
	var entrada = "A0511AB398765UJ1N230200"
	fmt.Println("Entrada:", entrada )

	campos , error_TLV = obtener_campos_TLV( entrada );

	fmt.Println("Ver Resultado:" )
	for key, element := range campos {
        fmt.Println("Key:", key, "=>", element )
	}
	if len(error_TLV) > 0 {
		t.Error("Error  :",  error_TLV )
	}
}