package services

// Retorna a diferença entre dois slices.
func SliceDifference(slice1 []string, slice2 []string) []string {
	var result []string
	
	// Mapa para armazenar os valores do segundo slice
	valuesInSlice2 := make(map[string]bool)
	for _, value := range slice2 {
		valuesInSlice2[value] = true
	}

	// Itera sobre o primeiro slice e verifica se o valor não está no mapa
	for _, value := range slice1 {
		if !valuesInSlice2[value] {
			result = append(result, value)
		}
	}

	return result
}
