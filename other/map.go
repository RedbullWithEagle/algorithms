package other

import "fmt"

/*************************************************
1. value,ok := map[key]  如果元素不存在，ok=false
 ************************************************/
func TestOKMap(){
	mMap := make(map[string]int, 0)
	mMap["liu"] = 123

	//map如果元素不存在，ok=false；反向判断的时候需要注意
	if value, ok := mMap["liu"]; ok {
		fmt.Println(value)
	}

	value, ok := mMap["uuu"]
	fmt.Println(value, ok)    //0,false
}
