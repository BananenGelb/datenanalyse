package intlists

// Min berechnet das Minimum einer Liste von Integer-Werten.
// Funktioniert nur für nicht-leere Listen.
func Min(values []int) int {
	min := values[0]
	for _, wert := range values {

		if wert <= min {
			min = wert
		}

	}

	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um das Minimum zu berechnen.
	   Jedes Mal, wenn Sie ein neues Minimum finden, speichern Sie es in der Variablen min.
	*/
	// TODO
	return min
}

// Max berechnet das Maximum einer Liste von Integer-Werten.
// Funktioniert nur für nicht-leere Listen.
func Max(values []int) int {
	max := values[0]
	for _, wert := range values {

		if wert >= max {
			max = wert
		}

	}
	/* Hinweis:
	   Gehen Sie analog zu Min vor.
	*/
	// TODO
	return max
}

// ValueRange erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert eine lückenlose Liste mit allen Zahlen zwischen
// dem Minimum und dem Maximum der Messreihe.
func ValueRange(values []int) []int {
	result := []int{}

	for i := Min(values); i <= Max(values); i++ {
		result = append(result, i)

	}
	/* Hinweis:
	   Verwenden Sie die Funktionen Min und Max, um das Minimum und das Maximum
	   der Messreihe zu berechnen.
	   Fügen Sie dann in einer Schleife alle Zahlen zwischen
	   Minimum und Maximum zu result hinzu.
	*/
	// TODO
	return result

}

// Sum erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert die Summe aller Werte.
func Sum(values []int) int {
	sum := 0

	for i := 0; i < len(values); i++ {

		wert := values[i]

		sum += wert

	}
	// TODO
	return sum
}

// Product erwartet eine Liste mit ganzzahligen Werten.
// Die Funktion liefert das Produkt aller Werte.
func Product(values []int) int {
	product := 1
	for i := 0; i < len(values); i++ {

		wert := values[i]

		product *= wert

	}
	// TODO
	return product
}
