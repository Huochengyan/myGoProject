/**
 * 气象-风向转化工具
 */
var WWindUtil = {
};
/**
 * 风向json
 */
WWindUtil.WDJson = {
    "N": {
        "en": "N",
        "cn": "北",
        "desc": "北"
    },
    "NNE": {
        "en": "NNE",
        "cn": "北东北",
        "desc": "东北偏北"
    },
    "NE": {
        "en": "NE",
        "cn": "东北",
        "desc": "东北"
    },
    "ENE": {
        "en": "ENE",
        "cn": "东东北",
        "desc": "东北偏北"
    },
    "E": {
        "en": "E",
        "cn": "东",
        "desc": "东"
    },
    "ESE": {
        "en": "ESE",
        "cn": "东东南",
        "desc": "东南偏东"
    },
    "SE": {
        "en": "SNE",
        "cn": "东南",
        "desc": "东南"
    },
    "SSE": {
        "en": "SSE",
        "cn": "南东南",
        "desc": "东南偏南"
    },
    "S": {
        "en": "S",
        "cn": "南",
        "desc": "南"
    },
    "SSW": {
        "en": "SSW",
        "cn": "南西南",
        "desc": "西南偏南"
    },
    "SW": {
        "en": "SW",
        "cn": "西南",
        "desc": "西南"
    },
    "WSW": {
        "en": "WSW",
        "cn": "西西南",
        "desc": "西南偏西"
    },
    "W": {
        "en": "W",
        "cn": "西",
        "desc": "西"
    },
    "WNW": {
        "en": "WNW",
        "cn": "西西北",
        "desc": "西北偏西"
    },
    "NW": {
        "en": "NW",
        "cn": "西北",
        "desc": "西北"
    },
    "NNW": {
        "en": "NW",
        "cn": "北西北",
        "desc": "西北偏北"
    },
    "ERROR": {
        "en": "error value",
        "cn": "错误数值",
        "desc": "错误数值"
    }

};

/**
 * 风力等级
 */
WWindUtil.WRatingArr=[
    {
        "min":0.0,
        "max":0.2,
        "level":0,
        "name":"无风",
        "landScene":"静，烟直上"
    },
    {
        "min":0.3,
        "max":1.5,
        "level":1,
        "name":"软风",
        "landScene":"烟示风向"
    },
    {
        "min":1.6,
        "max":3.3,
        "level":2,
        "name":"轻风",
        "landScene":"感觉有风"
    },
    {
        "min":3.4,
        "max":5.4,
        "level":3,
        "name":"微风",
        "landScene":"旌旗展开"
    },
    {
        "min":5.5,
        "max":7.9,
        "level":4,
        "name":"和风",
        "landScene":"吹起尘土"
    },
    {
        "min":8.0,
        "max":10.7,
        "level":5,
        "name":"劲风",
        "landScene":"小树摇摆"
    },
    {
        "min":10.8,
        "max":13.8,
        "level":6,
        "name":"强风",
        "landScene":"电线有声"
    },
    {
        "min":13.9,
        "max":17.1,
        "level":7,
        "name":"疾风",
        "landScene":"步行困难"
    },
    {
        "min":17.2,
        "max":20.7,
        "level":8,
        "name":"大风",
        "landScene":"折毁树枝"
    },
    {
        "min":20.8,
        "max":24.4,
        "level":9,
        "name":"烈风",
        "landScene":"小损房屋"
    },
    {
        "min":24.5,
        "max":28.4,
        "level":10,
        "name":"狂风",
        "landScene":"拔起树木"
    },
    {
        "min":28.5,
        "max":32.6,
        "level":11,
        "name":"暴风",
        "landScene":"损毁重大"
    },
    {
        "min":32.7,
        "max":36.9,
        "level":0,
        "name":"飓风",
        "landScene":"摧毁极大"
    },
    {
        "min":37.0,
        "max":41.4,
        "level":13,
        "name":"",
        "landScene":""
    },
    {
        "min":41.5,
        "max":46.1,
        "level":14,
        "name":"",
        "landScene":""
    },
    {
        "min":46.2,
        "max":50.9,
        "level":15,
        "name":"",
        "landScene":""
    },
    {
        "min":51.0,
        "max":56.0,
        "level":16,
        "name":"",
        "landScene":""
    },
    {
        "min":56.1,
        "max":null,
        "level":17,
        "name":"",
        "landScene":""
    }
]

/** 风向值转文字风向
 * @param {Object} windNumberValue 风向值
 */
WWindUtil.to16WindDirection = function(windNumberValue) {
    if (windNumberValue > 360) {
        return WWindUtil.WDJson.ERROR;
    }
    let a = Math.ceil(windNumberValue / 11.25);
    let ret;
    switch (a) {
        case 32:
        case 1:
            ret = WWindUtil.WDJson.N;
            break;

        case 2:
        case 3:
            ret = WWindUtil.WDJson.NNE;
            break;


        case 4:
        case 5:
            ret = WWindUtil.WDJson.NE;
            break;

        case 6:
        case 7:
            ret = WWindUtil.WDJson.ENE;
            break;

        case 8:
        case 9:
            ret = WWindUtil.WDJson.E;
            break;

        case 10:
        case 11:
            ret = WWindUtil.WDJson.ESE;
            break;

        case 12:
        case 13:
            ret = WWindUtil.WDJson.SE;
            break;

        case 14:
        case 15:
            ret = WWindUtil.WDJson.SSE;
            break;

        case 16:
        case 17:
            ret = WWindUtil.WDJson.S;
            break;

        case 18:
        case 19:
            ret = WWindUtil.WDJson.SSW;
            break;

        case 20:
        case 21:
            ret = WWindUtil.WDJson.SW;
            break;

        case 22:
        case 23:
            ret = WWindUtil.WDJson.WSW;
            break;

        case 24:
        case 25:
            ret = WWindUtil.WDJson.W;
            break;

        case 26:
        case 27:
            ret = WWindUtil.WDJson.WNW;
            break;

        case 28:
        case 29:
            ret = WWindUtil.WDJson.NW;
            break;

        case 30:
        case 31:
            ret = WWindUtil.WDJson.NNW;
            break;

        default:
            ret = WWindUtil.WDJson.ERROR;
            break;

    }
    return ret;
}

/** 风向值转文字风向
 * @param {Object} windNumberValue 风向值
 */
WWindUtil.to8WindDirection = function(windNumberValue) {
    if (windNumberValue > 360) {
        return WWindUtil.WDJson.ERROR;
    }
    let a = Math.ceil(windNumberValue / 22.5);
    let ret;
    switch (a) {
        case 16:
        case 1:
            ret = WWindUtil.WDJson.N;
            break;

        case 2:
        case 3:
            ret = WWindUtil.WDJson.NE;
            break;


        case 4:
        case 5:
            ret = WWindUtil.WDJson.E;
            break;

        case 6:
        case 7:
            ret = WWindUtil.WDJson.SE;
            break;

        case 8:
        case 9:
            ret = WWindUtil.WDJson.S;
            break;

        case 10:
        case 11:
            ret = WWindUtil.WDJson.SW;
            break;

        case 12:
        case 13:
            ret = WWindUtil.WDJson.W;
            break;

        case 14:
        case 15:
            ret = WWindUtil.WDJson.NW;
            break;

        default:
            ret = WWindUtil.WDJson.ERROR;
            break;

    }
    return ret;
}

/** 风速转风力等级
 * @param {Object} value 风速值 (m/s)
 */
WWindUtil.WindtoLevel=function(value){
    let vl=Number(value);
    if(!vl||vl<0){return null;}
    let arr=WWindUtil.WRatingArr;
    if(vl>=arr[arr.length-1].min){
        return arr[arr.length-1];
    }
    for(let i=0;i<arr.length;i++){
        if(vl >= arr[i].min && vl <= arr[i].max){
            return arr[i];
        }
    }
    return null;
}

// 风向值转文字
WWindUtil.windValueFormat = function(value) {
    try{
        let vl=Number(value);
        if (!vl || vl > 360 || vl < 0) {
            return '-';
        }
        return WWindUtil.to16WindDirection(value).desc;
    }catch(e){
        return '-';
    }
}

/** 风速值转文字描述
 * @param {Object} value 风速值
 */
WWindUtil.windFormatLevel = function(value) {
    try{
        let vl=Number(value);
        if (!vl || vl > 100 || vl < 0) {
            return "-";
        }
        return WWindUtil.WindtoLevel(value).level+"级("+WWindUtil.WindtoLevel(value).name+")";
    }catch(e){
        return "-";
    }
}


export {
    WWindUtil
};
