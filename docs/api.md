## 接口定义规范

- 接口采用分类分功能的方式进行定义，不同类别下的不同功能分别使用不同的路由进行区分，可类比文件管理器，一级文件夹，二级文件夹的功能

- 链接以https://sy.yonghui.cn开始，根据不同项目，不同分类下的不同功能进行隔离，如:https://sy.yonghui.cn/trace/asset/921108/123

不同的级别定义规则参照以下方式进行定义

### 1、项目

    - 统一一个域名，不同项目的区分，在于服务端的代理区分，目的在于解决端口占用的问题。
    - 对外仅暴露443或80端口，采用nginx代理的方式进行不同项目的区分。
    - 个人环境中可模拟服务器nginx代理进行测试
    
    - 定义项目级别路由时可根据项目名进行定义，如科尔沁牛肉，可定义为:kerchin
    
    注：项目级别的定义不区分前后端，不区分功能，需要进行更加细化的定义需借助其他级别
       项目级别的定义可以代理到某个项目中的某个功能，如sy.yonghui.cn/trace/search代理到：localhost:8889/weixin/search

### 2、模块

    模块的目的在于，区分大的功能，将不同的功能抽象出模块，实现模块化,隔离化，如用户模块、商品模块等
    
    注：模块不区分单双数，如：quality与qualities属统一类别
  
### 3、功能

    功能级别具体到某个功能点，如文件上传、确认收货等
    
### 4、请求方式
    
    使用请求方式来隔离不同的接口借鉴了restful的请求方式，如果根据以上三个级别不能很好的定义接口，可采用请求方式进行更细化的定义
     

根据上述四个级别定义，引申出

* 溯源小程序用户登录接口地址为:https://sy.yonghui.cn/trace/weixin/login

* 湘村黑猪确认收货接口地址为：https://sy.yonghui.cn/admin/blackPig/confirmReceipt



## html 模版定义

   - 接口定义：

    * 请求内容：https://sy.yonghui.cn/trace/html/load
    * 请求参数：router，data
    * 返回内容：html




## 二维码、条码内容规范

- 二维码内容：
    
    https://sy.yonghui.cn/trace/qr?productId=921108&id=123
    接口：https://sy.yonghui.cn/trace/qrcode/scan
        请求内容：url    如： https://sy.yonghui.cn/trace/qr/productId=921108&id=123
        返回内容：router 如： https://sy.yonghui.cn/trace/html/921108/123
    

- 条码内容采用超市生成规则，相应的做映射处理

   接口： https://sy.yonghui.cn/seach/barcode/:功能/:id
