package main
import(
	"fmt"
	"math/rand"
)
func fill(dim int) []int{

	slice := []int{}

	for i:=0;i<dim;i++{

		slice = append(slice, rand.Intn(100))
	}
	return slice
}
func selectionSort(s []int) []int{

	for i:=0;i<len(s);i++{

		m_index := i
		m_val := s[i]

		for j:=i+1;j<len(s);j++{

			if s[j] < s[m_index]{

				m_index = j
				m_val = s[j]
			}
		}
		s[m_index] = s[i]
		s[i] = m_val
	}
	return s
}
func insertSort(s []int) []int{

	for i:=1;i<len(s);i++{

		x := s[i]
		j := i-1

		for j>=0 && s[j] > x{

			s[j+1] = s[j]
			j--
		}
		s[j+1] = x
	}
	return s
}
func bubbleSort(s []int) []int{

	ord := true

	for i:=0;i<len(s);i++{

		for j:=1;j<len(s);j++{

			if s[j-1] > s[j]{

				s[j-1], s[j] = s[j], s[j-1]
				ord = false
			}
		}
		if ord{

			break
		}
		ord = true
	}
	return s
}
func mergeSort(s []int) []int{

	if len(s) > 1{

		m := len(s)/2
		b := []int{}
		c := []int{}

		b = append(b, s[:m]...)
		c = append(c, s[m:]...)

		b = mergeSort(b)
		c = mergeSort(c)

		s = merge(b, c)
	}
	return s
}
func merge(b, c []int) []int{

	i, j := 0, 0
	x := []int{}

	for i < len(b) && j < len(c){

		if b[i] <= c[j]{

			x = append(x, b[i])
			i++
		}else{

			x = append(x, c[j])
			j++
		}
	}
	if i < len(b){

		x = append(x, b[i:]...)
	}else{

		x = append(x, c[j:]...)
	}
	return x
}
func quickSort(s []int) []int{

	newSlice := []int{}

	if len(s) > 1{

		index := rand.Intn(len(s)-1)
		x := s[index]
		b := []int{}
		c := []int{}

		for i:=0;i<len(s);i++{

			if index == i{

				i++
			}
			if s[i] <= x{

				b = append(b, s[i])
			}else{

				c = append(c, s[i])
			}
		}
		b = quickSort(b)
		c = quickSort(c)

		newSlice = append(newSlice, b...)
		newSlice = append(newSlice, x)
		newSlice = append(newSlice, c...)

		return newSlice
	}
	return s
}
func main(){

	dim, sc := 10, -1

	for sc != 0{

		slice := fill(dim)
		fmt.Println("Sorting Algorithm....\n1 - Selection\n2 - Insert\n3 - Bubble\n4 - Merge\n5 - Quick\n0 - Esci")
		fmt.Println("Main:", slice)
		fmt.Print("Scelta: ")
		fmt.Scan(&sc)

		if sc == 1{

			fmt.Println("Selection:", selectionSort(slice), "\n")
		}else if sc == 2{

			fmt.Println("Insert:", insertSort(slice), "\n")
		}else if sc == 3{

			fmt.Println("Bubble:", bubbleSort(slice), "\n")
		}else if sc == 4{

			fmt.Println("Merge:", mergeSort(slice), "\n")
		}else if sc == 5{

			fmt.Println("Quick:", quickSort(slice), "\n")
		}else if sc == 0{

			break
		}else{

			fmt.Println("Pick again\n")
		}
	}
}
