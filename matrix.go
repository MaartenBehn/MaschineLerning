package main

type matrix struct {
	row    int
	collum int
	data   []float64
}

func NewMatrix(row int, collum int) *matrix {
	return &matrix{
		row:    row,
		collum: collum,
		data:   make([]float64, row*collum),
	}
}

func (m *matrix) Get(row int, collum int) float64 {
	return m.data[row*m.collum+collum]
}

func (m *matrix) Set(row int, collum int, data float64) {
	m.data[row*m.collum+collum] = data
}

func (m *matrix) Mul(m2 *matrix) (result *matrix) {
	if m.row != m2.collum && m.collum != m2.row {
		panic("Diese Mat Mul gehen nicht.")
	}

	var collum int
	if m2.row == m.collum {
		result = NewMatrix(m.row, m2.collum)
		collum = m.collum

	} else if m2.row == m.collum {
		result = NewMatrix(m2.row, m.collum)
		collum = m2.collum
	}

	for i := 0; i < result.row; i++ {
		for j := 0; j < result.collum; j++ {

			var val float64
			for k := 0; k < collum; k++ {
				in1 := m.Get(i, k)
				in2 := m2.Get(k, j)
				val += in1 * in2
			}
			result.Set(i, j, val)
		}
	}

	return result
}

/*
func (m *matrix) AddScala(s float64){
	m.applyToAll(func(v float64, i int) float64 {
		return v + s
	})
}

func (m *matrix) MulScala(s float64){
	m.applyToAll(func(v float64, i int) float64 {
		return v * s
	})
}

func (m *matrix) Min() (val float64, index int){
	val = math.MaxFloat64
	for i := 0; i < len(m.data); i++ {
		if m.data[i] < val{
			val = m.data[i]
			index = i
		}
	}
	return val, index
}

func (m *matrix) Max() (val float64, index int){
	val = 0
	for i := 0; i < len(m.data); i++ {
		if m.data[i] > val{
			val = m.data[i]
			index = i
		}
	}
	return val, index
}

func (m *matrix) Func(f func(float64)float64) (){
	m.applyToAll(func(v float64, i int) float64 {
		return f(v)
	})
}



func main() {

	var mat_row_1, mat_col_1 int

	var mat_row_2, mat_col_2 int

	var mat_1, mat_2, result [10][10]int

	fmt.Print("Enter no of rows of mat_1: ")

	fmt.Scanln(&mat_row_1)

	fmt.Print("Enter no of column of mat_1: ")

	fmt.Scanln(&mat_col_1)

	fmt.Print("Enter no of rows of mat_2: ")

	fmt.Scanln(&mat_row_2)

	fmt.Print("Enter no of column of mat_2: ")

	fmt.Scanln(&mat_col_2)

	fmt.Println("\nEnter matrix_1 elements: ")

	for i := 0; i < mat_row_1; i++ {

		for j := 0; j < mat_col_1; j++ {

			fmt.Scanf("%d ", &mat_1[i][j])

		}

	}

	fmt.Println("\nEnter matrix_2 elements: ")

	for i := 0; i < mat_row_2; i++ {

		for j := 0; j < mat_col_2; j++ {

			fmt.Scanf("%d ", &mat_2[i][j])

		}

	}


	// Multiplication of two matrix

	for i := 0; i < mat_row_1; i++ {

		for j := 0; j < mat_col_2; j++ {

			result[i][j] = 0

			for k := 0; k < mat_col_2; k++ {

				result[i][j] += mat_1[i][k] * mat_2[k][j]

			}

		}

	}

	fmt.Println("\nAfter Multiplication Matrix is: \n")

	for i := 0; i < mat_row_1; i++ {

		for j := 0; j < mat_col_2; j++ {

			fmt.Printf("%d ", result[i][j])

		}
		fmt.Println("\n")

	}


*/
