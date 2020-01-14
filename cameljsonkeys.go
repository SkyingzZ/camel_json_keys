package main

import (
	"strings"
	"fmt"
	"io/ioutil"
)

//下划线写法转为驼峰写法	like "sample_test_name_balabala" to "SampleTestNameBalabala"
func CamelName(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

func CamelJsonKey(json_data []byte) []byte{
	str := string(json_data)

	var is_quot_first bool = true		//在双引号中为 false  	|	the value in the double quotes is false 
	var first_index int = 0				//左引号的索引			|	the left quote index
	var second_index int = 0			//右引号的索引			|	the right quote index

	var res_str string
	var the_key_index int
	for i, value := range str{
		if !is_quot_first && str[i] == '"'{		//右引号		|	if meet the right quote
			second_index = i
			is_quot_first = !is_quot_first
			
		}else if is_quot_first && str[i] == '"'{	//左引号	|	if meet the left quote
			first_index = i
			the_key_index = len(res_str)
			is_quot_first= !is_quot_first
		}else if is_quot_first && str[i] == ':'{	
			tmp_str := CamelName(str[first_index+1: second_index])
			res_str = res_str[:the_key_index]+"\""+tmp_str+"\""
		}else{}

		res_str += string(value)
	}

	return []byte(res_str)
}


func main(){
	
	
	json_data, err := ioutil.ReadFile("./testjson.json")
	if err != nil{
		return
	}
  fmt.Println(string(json_data))
  fmt.Println()

  fmt.Println(string(CamelJsonKey([]byte(json_data))))
  fmt.Println()

}
