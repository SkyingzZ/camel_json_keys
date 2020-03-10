# camel_json_keys
convers all json keys to camel style, like "the_red_apple":"very_nice" to "TheRedApple":"very_nice"    
把json中的key全部转换为驼峰式（首字母大写形式）    
复杂度  O(n)    

## like:    
### 转换前   
![image](http://anaou.com/photolink/cameljsonkeys/qian.png)   

### 转换后      
![image](http://anaou.com/photolink/cameljsonkeys/hou.png)   

cameljsonkey.go为新实现方法，使用map来实现，效率更高   
cameljsonkey_old.go为旧版实现，效率很低
