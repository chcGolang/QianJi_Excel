var findCsvTypeUrl = "http://127.0.0.1:8080/qianji/matching_rule/csv_type_info"
var findConsumptionTypeUrl = "http://127.0.0.1:8080/qianji/matching_rule/consumption_type"
var listdataUrl = "http://127.0.0.1:8080/qianji/matching_rule/listdata"
var findRulesInfoUrl = "http://127.0.0.1:8080/qianji/matching_rule/rules_info"
var findFuzzyInfoUrl = "http://127.0.0.1:8080/qianji/matching_rule/fuzzy"
var saveMatchingRule = "http://127.0.0.1:8080/qianji/matching_rule/save_matching_rule"
var findMatchingRuleByIdUrl = "http://127.0.0.1:8080/qianji/matching_rule/find_id"
var updateMatchingRuleUrl = "http://127.0.0.1:8080/qianji/matching_rule/update_matching_rule"
var deleteMatchingRuleUrl = "http://127.0.0.1:8080/qianji/matching_rule/delete_id"
var baseUrl = "http://127.0.0.1:8080/qianji"
var limitPage = 4
// 初始化csv文件类型
function initCsvSelect(dataInfo){
	var selectInfo = "";
	selectInfo = "<option value="+dataInfo.csv_type+">"+dataInfo.csv_name+"</option>"
	$("#csv-typeId").append(selectInfo);
	
}

// 查询csv文件类型
function findCsvType(url){
	$("#csv-typeId").append("<option value=\"0\">全部</option>");
	$.get(url,function(data,status){
		$.each(data.data,function(i,item){
				initCsvSelect(item)
		　　});
	});
}

// 查询消费类型
function findConsumptionType(url){
	$("#consumption-typeId").append("<option value=\"-1\">全部</option>");
	$.get(url,function(data,status){
		$.each(data.data,function(i,item){
				initConsumptionTypeSelect(item)
		　　});
	});
}

// 初始化钱迹消费类型
function initConsumptionTypeSelect(dataInfo){
	var selectInfo = "";
	selectInfo = "<option value="+dataInfo.id+">"+dataInfo.consumption_name+"</option>"
	$("#consumption-typeId").append(selectInfo);
}

// csv标题
function findRulesInfo(url,csvType){
	$("#rules-typeId").empty();
	$("#rules-typeId").append("<option value=\"-1\">全部</option>");
	var findUrl = url+"?csv_type="+csvType
	$.get(findUrl,function(data,status){
		$.each(data.data,function(i,item){
				initRulesSelect(item)
		　　});
	});
}

// 初始化csv表标题
function initRulesSelect(dataInfo){
	var selectInfo = "";
	selectInfo = "<option value="+dataInfo.id+">"+dataInfo.rules_name+"</option>"
	$("#rules-typeId").append(selectInfo)
}


// 根据csv类型刷新csv标题
function csvTitleInit(form){
	form.on("select(csv-typeId)",function(data){
		// console.log(data.elem); //得到select原始DOM对象
		// console.log(data.value); //得到被选中的值
		
		$("#rules-typeId").empty();
		$("#rules-typeId").append("<option value=\"-1\">全部</option>");
		var findUrl = findRulesInfoUrl+"?csv_type="+data.value
		$.get(findUrl,function(data,status){
			$.each(data.data,function(i,item){
					initRulesSelect(item)
			　　});
			form.render('select');
		});
	});
}

// 查询匹配查询方式
function findFuzzyInfo(url){
	$.get(url,function(data,status){
		$.each(data.data,function(i,item){
			initFuzzySelect(item)
		});
	})
}

// 初始化匹配查询方式选择
function initFuzzySelect(dataInfo){
	var selectInfo = "";
	selectInfo = "<option value="+dataInfo.fuzzyid+">"+dataInfo.fuzzy_name+"</option>"
	$("#fuzzyid").append(selectInfo)
}

// 自定义验证规则
function verifyValue(form){
	form.verify({
		csvType:function(value){
			if (value==0){
				 return '请选择csv文件类型';
			}
		},
		consumptionType:function(value){
			if(value==-1){
				 return '请选择钱迹消费类型';
			}
		},
		rulesInfo:function(value){
			console.log(value)
			if(value==-1){
				 return '请选择csv表标题';
			}
		},
		value_data:function(value){
			if (value.length<2){
				 return '匹配值至少得5个字符啊';
			}
		}
	  
	});
}

function openFailView(title,content){
	layer.open({
		  title: title
		  ,content: content
		})
}

function checkData(data){
	if(data.code!=200){
		openFailView("请求失败",data.err_msg);
		return false
	}
	return true
}

function findMatchingRuleById(id){
	layui.use(['form','layer'], function(){
		var form = layui.form;
		var findUrl = findMatchingRuleByIdUrl+"?id="+id
		$.get(findUrl,function(data,status){
			if (data.code != 200){
				openFailView("请求失败",data.err_msg)
			}else{
				$("#consumption-typeId").val(data.data.consumption_type_id)
				$("#csv-typeId").val(data.data.csv_type)
				$("#rules-typeId").val(data.data.rules_type_id)
				$("#fuzzyid").val(data.data.fuzzy)
				$("#value_data").val(data.data.value_data)
				$("#dataId").val(data.data.id)
				form.render()
			}
		});
	});
	
}


function guid() {
    function S4() {
       return (((1+Math.random())*0x10000)|0).toString(16).substring(1);
    }
    return (S4()+S4()+"-"+S4()+"-"+S4()+"-"+S4()+"-"+S4()+S4()+S4());
}