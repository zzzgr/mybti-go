## 北京地铁预约出行
### 简介
微信公众号【北京地铁预约出行】，预约出行免排队。
北京地铁在5号线天通苑站，昌平线沙河站，6号线草房站推出工作日早高峰“预约进站服务”。乘客在限流时间内，按照预约时间到达，即可走预约通道快速进站

### 使用
1. 编译程序
    ```shell
    ## Linux arm
    go env -w CGO_ENABLED=0  GOOS=linux  GOARCH=amd64
    go build
    
    ## Linux amd
    go env -w CGO_ENABLED=0  GOOS=linux  GOARCH=amd64
    go build 
    
    ## windows
    go env -w CGO_ENABLED=0 GOOS=windows  GOARCH=amd64
    go build
    ```
2. 配置文件

    在可执行文件同级目录创建`config`目录，复制`sample/config.yaml`到此目录
    
   注意:
    - pushplus token获取：[pushplus官网](https://www.pushplus.plus/push1.html)
    - name为标识，可识别名字即可
    - accessToken自行抓包获取，理论上有效期7天（程序会检查有效期并推送至微信）
    - line和station可选值：昌平线-沙河站、5号线-天通苑站、6号线-草房站
    - time: 进站时间，参考公众号填写即可

3. 启动
    ```shell
    ./mybti-go 
   ```
   
### 效果
略