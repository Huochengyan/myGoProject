function methodInfo(methodName,data){
    this.method=methodName
    this.data=data
}
let serverOrigin=window.location.origin;
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
                if(data.code==201)
                {
                    alert("服务器错误！！");
                    return;
                }
                serverResultData = data.data;
            },error:function(XMLHttpRequest, textStatus, errorThrown, data){
                if(XMLHttpRequest.responseJSON.code==2002||XMLHttpRequest.responseJSON.code==2003){
                    alert("超时，请重新登录！");
                    //location.href='login.html'; //tab页内代开
                    //top.location.href ="login.html"; //这种可以新页打开
                    window.parent.location.replace("login.html");//刷新父级页面; 这种可以新页打开
                }
                console.log(errorThrown)    //失败的回调
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

    timestampToTime:function(timestamp) {
        var date = new Date(timestamp * 1000);//时间戳为10位需*1000，时间戳为13位的话不需乘1000
        var Y = date.getFullYear() + '-';
        var M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-';
        var D = date.getDate() + ' ';
        var h = date.getHours() + ':';
        var m = date.getMinutes() + ':';
        var s = date.getSeconds();
        return Y+M+D+h+m+s;
        },


   // 获取所有用户 getAllUser
    login: function (postData) {
        var url = serverUrl + "/user/login";
        $.ajaxSettings.async = false;
        return dataServer.postServerData(url, postData);
    },
    //获取上传上传文件URL
    getUploadFileUrl: function () {
        return url = serverUrl + "/file/upload";
    },
    //获取传感器数据 list
    getSensorList: function (postData) {
        var url = serverUrl + "/sensor/list";
        $.ajaxSettings.async = false;
        return dataServer.postServerData(url, postData);
    },
    // 获取所有用户 getAllUser
    Queryalluser: function (postData) {
        var url = serverUrl + "/user/queryalluser";
        $.ajaxSettings.async = false;
        return dataServer.postServerData(url, postData);
    },
    // 获取所有用户 getAllUser
    getUserList: function (postData) {
        var url = serverUrl + "/user/getUserList";
        $.ajaxSettings.async = false;
        return dataServer.postServerData(url, postData);
    },
    updateUserInfo:function (postData) {
        var url = serverUrl + "/user/update";
        $.ajaxSettings.async = false;
        return dataServer.postServerData(url, postData);
    },
    deleteUserInfo:function (postData) {
        var url = serverUrl + "/user/delete";
        $.ajaxSettings.async = false;
        return dataServer.postServerData(url, postData);
    },
    addUserInfo:function (postData) {
        var url = serverUrl + "/user/add";
        $.ajaxSettings.async = false;
        return dataServer.postServerData(url, postData);
    },
    getPcResourceList:function (postData) {
        var url = serverUrl + "/pcinfo/getPcResourceList";
        $.ajaxSettings.async = false;
        return dataServer.postServerData(url, postData);
    },
    getLoginLogList:function (postData) {
        var url = serverUrl + "/log/getLoginLogList";
        $.ajaxSettings.async = false;
        return dataServer.postServerData(url, postData);
    },

}