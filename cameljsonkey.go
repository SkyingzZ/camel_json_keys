
 
//下划线写法转为驼峰写法,并且将'/'字符转换为''	like "sample_test_name_balabala/dilidili" to "SampleTestNameBalabalaDilidili"
//同时也会去掉左侧的全部数字	like "361_sport" to "Sport"
func camelName(name string) string {
    name = strings.Replace(name, "_", " ", -1)
    name = strings.Replace(name, "/", " ", -1)
    name = strings.Title(name)
    name = strings.Replace(name, " ", "", -1)
    return strings.TrimLeft(name, "0123456789")
}
 
func camelVecKey(my_vec []interface{}) ([]interface{}, error){
    if !strings.HasPrefix(reflect.TypeOf(my_vec).String(), "[]"){
        return nil, errors.New("decode to vector failed")
    }
    rtn_vec := my_vec[0:0]
    for _, node := range my_vec{
        if reflect.TypeOf(node) == nil{
            continue
        }
        if strings.HasPrefix(reflect.TypeOf(node).String(),"map"){
            res_node, err := camelMapKey(node.(map[string]interface{}))
            if err != nil{
                return nil, err
            }
            rtn_vec = append(rtn_vec, res_node)
        }else if strings.HasPrefix(reflect.TypeOf(node).String(),"[]"){
            res_node, err := camelVecKey(node.([]interface{}))
            if err != nil{
                return nil, err
            }
            rtn_vec = append(rtn_vec, res_node)
        }else{
            rtn_vec = append(rtn_vec, node)
        }
    }
    return rtn_vec, nil
}
 
func camelMapKey(my_map map[string]interface{}) (map[string]interface{}, error){
 
    if !strings.HasPrefix(reflect.TypeOf(my_map).String(), "map"){
        return nil, errors.New("decode to map failed")
    }
    rtn_map := make(map[string]interface{})
    for k,v := range my_map{
        if reflect.TypeOf(v) == nil{
            rtn_map[camelName(k)] = nil
            continue
        }
        if strings.HasPrefix(reflect.TypeOf(v).String(),"map"){
            res_map, err := camelMapKey(v.(map[string]interface{}))
            if err != nil{
                return nil, err
            }
            rtn_map[camelName(k)] = res_map
        } else if strings.HasPrefix(reflect.TypeOf(v).String(), "[]"){
           res_node, err := camelVecKey(v.([]interface{}))
           if err != nil{
               return nil, err
           }
           rtn_map[camelName(k)] = res_node
        } else {
            rtn_map[camelName(k)] = v
        }
    }
    return rtn_map, nil
}
 
func CamelJsonKey(json_data []byte, is_vec bool) ([]byte, error){
	var my_map_vec []interface{}
	var my_map map[string]interface{}
	if is_vec {
		err := json.Unmarshal(json_data, &my_map_vec)
		if err != nil{
			return []byte{}, errors.New("Unmarshal json to vector failed")
		}
 
		my_map_vec, err = camelVecKey(my_map_vec)
		if err != nil{
			return []byte{}, err
		}
		my_json,err :=json.Marshal(my_map_vec)
		if err != nil{
			return []byte{}, err
		}
		return my_json, err
 
	} else {
		my_map = make(map[string]interface{})
		err := json.Unmarshal(json_data, &my_map)
		if err != nil{
			return []byte{}, errors.New("Unmarshal json to map failed")
		}
 
		my_map, err = camelMapKey(my_map)
		if err != nil{
			return []byte{}, err
		}
		my_json,err :=json.Marshal(my_map)
		if err != nil{
			return []byte{}, err
		}
		return my_json, err
	}
}
