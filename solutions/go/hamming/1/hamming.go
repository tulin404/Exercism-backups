package hamming

import (
    "errors"
)

func Distance(a, b string) (int, error) {
    var distance int
	if (len(a) != len(b)) {
        return 0, errors.New("Lenghts are not equal.")
    }
    
    bSlice := []rune(b)
	for i, aChar := range a {
        bChar := bSlice[i]        
        if (aChar == bChar) {
            continue
        } else {
            distance++
        }
    };
    return distance, nil
}
