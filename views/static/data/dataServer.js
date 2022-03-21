function methodInfo(methodName,data){
    this.method=methodName
    this.data=data
}
let serverOrigin=location.origin;
let serverUrl=serverOrigin+'/v1';
let code_success=200;
var dataServer= {
    //向服务器法送数据
    postServerData: function (postUrl, postData) {
        $.ajaxSettings.async = false;
        var url = postUrl;
        var serverResultData;
        var jsonData = JSON.stringify(postData)
        $.ajax({
            //请求类型，这里为POST
            type: 'POST', //你要请求的api的URL
            url: url, //是否使用缓存
            cache: false, //数据类型，这里我用的是json
            dataType: "json",
            contentType: "application/json;charset=utf-8",
            //必要的时候需要用JSON.stringify() 将JSON对象转换成字符串
            data: jsonData, //data: {key:value},
            //添加额外的请求头
            headers: {'Access-Control-Allow-Origin': '*', "token": localStorage.getItem("token")}, //请求成功的回调函数
            success: function (data) {
                //函数参数 "data" 为请求成功服务端返回的数据
                // if (data.code==510){
                //     alert("重新登录！");
                //     open("./login.html");
                // }
                serverResultData = data.data;
            }
        });
        $.ajaxSettings.async = true;
        return serverResultData;
    },
    postServerDataSync: function (postUrl, postData) {
        //   $.ajaxSettings.async = false;
        var url = postUrl;
        var serverResultData = false;
        var jsonData = JSON.stringify(postData)
        $.ajax({
            //请求类型，这里为POST
            type: 'POST', //你要请求的api的URL
            url: url, //是否使用缓存
            cache: false, //数据类型，这里我用的是json
            dataType: "json",
            contentType: "application/json;charset=utf-8",
            //必要的时候需要用JSON.stringify() 将JSON对象转换成字符串
            data: jsonData, //data: {key:value},
            //添加额外的请求头
            headers: {'Access-Control-Allow-Origin': '*', "token": localStorage.getItem("token")}, //请求成功的回调函数
            success: function (data) {
                serverResultData = data.data;
                return data.data;
            }
        });
        // $.ajaxSettings.async = true;
        return serverResultData;
    },
    //获取传感器数据 list
    getSensorList: function (postData) {
        var url = serverUrl + "/sensor/list";
        $.ajaxSettings.async = false;
        return sensorServer.postServerData(url, postData);
    },
    // 获取所有用户 getAllUser
    Queryalluser: function (postData) {
        var url = serverUrl + "/user/queryalluser";
        $.ajaxSettings.async = false;
        return sensorServer.postServerData(url, postData);
    },
    // 获取所有用户 getAllUser
    GetUserList: function (postData) {
        var url = serverUrl + "/user/queryalluser";
        $.ajaxSettings.async = false;
        return sensorServer.postServerData(url, postData);
    },
}