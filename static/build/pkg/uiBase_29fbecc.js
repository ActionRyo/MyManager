define("fishstrap/ui/query.js",function(e,n,r){var o=e("fishstrap/core/global.js"),a=(e("fishstrap/ui/dialog.js"),e("fishstrap/ui/table.js")),t=e("fishstrap/ui/input.js");r.exports={simpleQuery:function(e){var n={id:"",column:[],queryColumn:[],params:{},operate:[],checkAll:!1,url:"",button:[]};for(var r in e)n[r]=e[r];var i=o.uniqueNum(),u=o.uniqueNum(),l=o.uniqueNum(),c=o("#"+n.id),s=o('<form id="'+i+'" class="form-inline m10"></form><div class="m10"><div id="'+l+'"></div><div id="'+u+'"></div></div>');c.empty(),c.append(s),o.console.log(o.JSON.stringify(e.column));for(var f=o.url.toInfo(location.href).search,p=a.staticTable({id:u,params:o.url.toInfo(location.href).search,column:n.column,operate:n.operate,checkAll:n.checkAll,url:n.url,pageIndex:f.pageIndex||0,pageSize:f.pageSize||10,nextPage:function(e){var n=o.url.toInfo(location.href);n.search=o.extend(n.search,e),location.href=o.url.fromInfo(n)}}),m=[],r=0;r!=n.queryColumn.length;++r){for(var d,h=n.queryColumn[r],v=0;v!=n.column.length;++v){var g=n.column[v];if(g.id==h){d=_.clone(g);break}}"enum"==d.type?d.map=o.extend({"":"请选择"},d.map):"check"==d.type&&(d.type="enum",d.map=o.extend({"":"请选择"},d.map)),m.push(d)}t.flowInput({id:i,field:m,value:f,submit:function(e){var n=o.url.toInfo(location.href);n.search=o.extend(n.search,e),n.search.pageIndex=0,location.href=o.url.fromInfo(n)}});for(var r=0;r!=n.button.length;++r){var b=n.button[r];!function(e){var n=o('<button class="btn">'+e.name+"</button>");n.click(e.click),o("#"+l).append(n)}(b)}return p}}});
;define("fishstrap/ui/input.js",function(e,i,t){{var a=e("fishstrap/core/global.js"),d=e("fishstrap/ui/dialog.js"),n=e("fishstrap/ui/editor.js"),r=e("fishstrap/ui/table.js"),l=e("fishstrap/util/upload.js");e("fishstrap/page/subPage.js")}e("fishstrap/util/jqueryDatetimePicker.js"),t.exports={flowInput:function(e){var i={id:"",field:[],value:{},submit:function(){}};for(var t in e)i[t]=e[t];if(0!=i.field.length){for(var d="",t=0;t!=i.field.length;++t){var n=i.field[t];if("text"==n.type)d+="<span>&nbsp;"+n.name+'：</span><input type="text" name="'+n.id+'" class="input-small"/>';else if("time"==n.type)d+="<span>&nbsp;"+n.name+'：</span><input type="text" name="'+n.id+'" class="time input-small"/>';else if("enum"==n.type){var e="";"undefined"!=typeof n.map[""]&&(e+='<option value="">'+n.map[""]+"</option>");for(var r in n.map)""!=r&&(e+='<option value="'+r+'">'+n.map[r]+"</option>");d+="<span>&nbsp;"+n.name+'：</span><select name="'+n.id+'">'+e+"</select>"}}d+='<button type="button" class="btn query">查询</button><button type="reset" class="btn">重置</button>',d=a(d),a("#"+i.id).append(d);for(var t=0;t!=i.field.length;++t){var n=i.field[t];"time"==n.type&&a("#"+i.id).find("input[name="+n.id+"]").datetimepicker({lang:"ch",timepicker:!1,format:"Y-m-d",closeOnDateSelect:!0})}for(var t=0;t!=i.field.length;++t){var n=i.field[t],l=i.value[n.id];l&&("text"==n.type?a("#"+i.id).find("input[name="+n.id+"]").val(l):"time"==n.type?a("#"+i.id).find("input[name="+n.id+"]").val(l):"enum"==n.type&&a("#"+i.id).find("select[name="+n.id+"]").val(l))}a("#"+i.id).find(".query").click(function(){for(var e={},t=0;t!=i.field.length;++t){var d=i.field[t];"text"==d.type?e[d.id]=a.trim(a("#"+i.id).find("input[name="+d.id+"]").val()):"time"==d.type?e[d.id]=a.trim(a("#"+i.id).find("input[name="+d.id+"]").val()):"enum"==d.type&&(e[d.id]=a.trim(a("#"+i.id).find("select[name="+d.id+"]").val()))}i.submit(e)})}},verticalInput:function(e){function i(e){for(var i=0;i!=s.field.length;++i){var t=s.field[i];if("undefined"!=typeof e[t.id])if("read"==t.type)p.find("div[name="+t.id+"]").text(e[t.id]);else if("link"==t.type)p.find("div[name="+t.id+"]").text(e[t.id]),p.find("a[name="+t.id+"]").attr("href",e[t.id]);else if("fullEditor"==t.type)t._editor.setContent(e[t.id]);else if("simpleEditor"==t.type)t._editor.setFormatData(e[t.id]);else if("image"==t.type)p.find("img[name="+t.id+"]").attr("src",e[t.id]);else if("file"==t.type)p.find("div[name="+t.id+"]").text(e[t.id]);else if("compressFile"==t.type){var d=_.map(e[t.id],function(e){return{name:'<a href="'+e+'" target="_blank">'+e+"</a>"}});t.tableOperation.clear(),t.tableOperation.add(d)}else"area"==t.type?p.find("textarea[name="+t.id+"]").val(e[t.id]):"text"==t.type||"password"==t.type||"time"==t.type?p.find("input[name="+t.id+"]").val(e[t.id]):"enum"==t.type?p.find("select[name="+t.id+"]").val(e[t.id]):"check"==t.type?p.find("input[name="+t.id+"]").each(function(){for(var i=0;i!=e[t.id].length;++i){var d=e[t.id][i];if(a(this).val()==d)return void a(this).attr("checked",!0)}a(this).attr("checked",!1)}):"table"==t.type&&t.tableOperation.add(e[t.id])}}function t(){for(var e={},i=0;i!=s.field.length;++i){var t=s.field[i];"read"==t.type?e[t.id]=a.trim(a("#"+s.id).find("div[name="+t.id+"]").text()):"link"==t.type?e[t.id]=p.find("a[name="+t.id+"]").attr("href"):"simpleEditor"==t.type?e[t.id]=t._editor.getFormatData():"fullEditor"==t.type?e[t.id]=t._editor.getContent():"image"==t.type?e[t.id]=a.trim(a("#"+s.id).find("img[name="+t.id+"]").attr("src")):"file"==t.type?e[t.id]=a.trim(a("#"+s.id).find("div[name="+t.id+"]").text()):"compressFile"==t.type?e[t.id]=_.map(t.tableOperation.get(),function(e){return e.name}):"area"==t.type?e[t.id]=a.trim(a("#"+s.id).find("textarea[name="+t.id+"]").val()):"text"==t.type||"password"==t.type||"time"==t.type?e[t.id]=a.trim(a("#"+s.id).find("input[name="+t.id+"]").val()):"enum"==t.type?e[t.id]=a.trim(a("#"+s.id).find("select[name="+t.id+"]").val()):"check"==t.type?(e[t.id]=[],a("#"+s.id).find("input[name="+t.id+"]:checked").each(function(){e[t.id].push(a(this).val())})):"table"==t.type&&(e[t.id]=t.tableOperation.get())}return e}var s={id:"",field:[],value:{},submit:function(){},cancel:function(){}};for(var o in e)s[o]=e[o];var p="",f="";a.console.log(a.JSON.stringify(s.field));for(var o=0;o!=s.field.length;++o){var u=s.field[o];if(f+='<tr><td class="tableleft" style="width:20%;">'+u.name+"</td>",f+="undefined"!=typeof u.targetId?'<td id="'+u.targetId+'">':"<td>","read"==u.type)f+='<div name="'+u.id+'"/>';else if("time"==u.type)f+='<input type="text" name="'+u.id+'" class="time input-small"/>';else if("link"==u.type)f+='<a name="'+u.id+'" target="_blank"><div name="'+u.id+'"/></a>';else if("fullEditor"==u.type)u.editorTargetId=a.uniqueNum(),f+='<div class="alert alert-danger" role="alert">注意！为了保证微信内观看视频的兼容性，强烈建议你只插入腾讯视频网址，不要插入其它视频网站，或直接上传视频。</div><div name="'+u.id+'" id="'+u.editorTargetId+'"/>';else if("simpleEditor"==u.type)u.editorTargetId=a.uniqueNum(),f+='<div name="'+u.id+'" id="'+u.editorTargetId+'"/>';else if("image"==u.type)u.imageTargetId=a.uniqueNum(),u.imageProgressTargetId=a.uniqueNum(),f+='<div><img name="'+u.id+'" src=""/></div><div class="progress"><div class="bar" id="'+u.imageProgressTargetId+'"></div></div><div class="btn" id="'+u.imageTargetId+'"><span>点击这里上传图片</span></div>';else if("file"==u.type)u.fileTargetId=a.uniqueNum(),u.fileProgressTargetId=a.uniqueNum(),f+='<div name="'+u.id+'"></div><div class="progress"><div class="bar" id="'+u.fileProgressTargetId+'"></div></div><div class="btn" id="'+u.fileTargetId+'"><span>点击这里上传文件</span></div>';else if("compressFile"==u.type)u.tableId=a.uniqueNum(),u.fileTargetId=a.uniqueNum(),u.fileProgressTargetId=a.uniqueNum(),f+='<div id="'+u.tableId+'"></div><div class="progress"><div class="bar" id="'+u.fileProgressTargetId+'"></div></div><div class="btn" id="'+u.fileTargetId+'"><span>点击这里上传压缩文件</span></div>';else if("area"==u.type){var m="";0==_.isUndefined(u.disabled)&&u.disabled===!0&&(m='disabled="true"'),f+='<textarea name="'+u.id+'" style="width:90%;height:300px;" '+m+"></textarea>"}else if("text"==u.type)f+='<input type="text" name="'+u.id+'"/>';else if("password"==u.type)f+='<input type="password" name="'+u.id+'"/>';else if("enum"==u.type){var m="";0==_.isUndefined(u.disabled)&&u.disabled===!0&&(m='disabled="true"');var e="";for(var c in u.map)e+='<option value="'+c+'">'+u.map[c]+"</option>";f+="<select "+m+' name="'+u.id+'">'+e+"</select>"}else if("check"==u.type){var m="";0==_.isUndefined(u.disabled)&&u.disabled===!0&&(m='disabled="true"');var e="";for(var c in u.map)e+='<span><input type="checkbox" '+m+' name="'+u.id+'" value="'+c+'">'+u.map[c]+"</span>&nbsp;&nbsp;";f+=e}else if("table"==u.type){u.tableId=a.uniqueNum(),f+="<div>";for(var c=0;c!=u.option.button.length;++c){var v=u.option.button[c];v.buttonId=a.uniqueNum(),f+='<button type="button" class="btn" id="'+v.buttonId+'">'+v.name+"</button>",u.option.button[c]=v}f+="</div>",f+='<div id="'+u.tableId+'"></div>'}f+="</td></tr>"}var g="";(0==_.isUndefined(s.submit)||0==_.isUndefined(s.cancel))&&(g+='<tr><td class="tableleft"></td><td>',0==_.isUndefined(s.submit)&&(g+='<button type="button" class="btn btn-primary submit" >提交</button>'),0==_.isUndefined(s.cancel)&&(g+='<button type="button" class="btn btn-success cancel">返回列表</button>'),g+="</td></tr>"),p+='<table class="table table-bordered table-hover definewidth m10">'+f+g+"</table>",p=a(p),a("#"+s.id).append(p);for(var o=0;o!=s.field.length;++o){var u=s.field[o];!function(e){if("image"==e.type)l.image({url:e.option.url,urlToken:e.option.urlToken,urlType:e.option.urlType,target:e.imageTargetId,field:"data",width:e.option.width,height:e.option.height,quality:.8,onProgress:function(i){a.console.log(i),a("#"+e.imageProgressTargetId).text(i+"%"),a("#"+e.imageProgressTargetId).css("width",i+"%")},onSuccess:function(i){return i=a.JSON.parse(i),0!=i.code?void d.message(i.msg):void p.find("img[name="+e.id+"]").attr("src",i.data)},onFail:function(e){d.message(e)}});else if("file"==e.type)l.file({url:e.option.url,target:e.fileTargetId,field:"data",type:e.option.type,maxSize:e.option.maxSize,onProgress:function(i){a("#"+e.fileProgressTargetId).text(i+"%"),a("#"+e.fileProgressTargetId).css("width",i+"%")},onSuccess:function(i){return i=a.JSON.parse(i),0!=i.code?void d.message(i.msg):void p.find("div[name="+e.id+"]").text(i.data)},onFail:function(e){d.message(e)}});else if("compressFile"==e.type)e.tableOperation=r.staticSimpleTable({id:e.tableId,data:[],column:[{id:"name",type:"html",name:"文件"}],operate:[]}),l.file({url:e.option.url,target:e.fileTargetId,field:"data",type:e.option.type,maxSize:e.option.maxSize,onProgress:function(i){a.console.log(i),a("#"+e.fileProgressTargetId).text(i+"%"),a("#"+e.fileProgressTargetId).css("width",i+"%")},onSuccess:function(i){if(i=a.JSON.parse(i),0!=i.code)return void d.message(i.msg);var t=_.map(i.data,function(e){return{name:'<a href="'+e+'" target="_blank">'+e+"</a>"}});e.tableOperation.clear(),e.tableOperation.add(t)},onFail:function(e){d.message(e)}});else if("simpleEditor"==e.type)e._editor=n.simpleEditor({id:e.editorTargetId,url:e.option.url});else if("fullEditor"==e.type)e._editor=n.fullEditor({id:e.editorTargetId,url:e.option.url});else if("time"==e.type)a("#"+s.id).find("input[name="+e.id+"]").datetimepicker({lang:"ch",timepicker:!1,format:"Y-m-d",closeOnDateSelect:!0});else if("table"==e.type){e.tableOperation=r.staticSimpleTable({id:e.tableId,data:[],column:e.option.column,operate:e.option.operate});for(var i=0;i!=e.option.button.length;++i){var t=e.option.button[i];a("#"+t.buttonId).click(function(){t.click(e.tableOperation)})}}}(u)}return i(s.value),p.find(".submit").click(function(){s.submit(t())}),p.find(".cancel").click(s.cancel),{get:t,set:i}}}});
;define("fishstrap/ui/table.js",function(e,t,a){var n=e("fishstrap/core/global.js"),i=e("fishstrap/ui/dialog.js");a.exports={staticSimpleTable:function(e){function t(){n("#"+p.id).find("tbody").empty()}function a(e,t){e.find("td").each(function(){for(var e=null,a=0;a!=p.column.length;++a)if(p.column[a].id==n(this).attr("class")){e=p.column[a];break}null!=e&&1!=_.isUndefined(t[e.id])&&("hidden"==e.type?n(this).text(t[e.id]):"image"==e.type?n(this).find("img").attr("src",t[e.id]):"textInput"==e.type?n(this).find("input").val(t[e.id]):n(this).text(t[e.id]))})}function i(e){var t={};return e.find("td").each(function(){for(var e=null,a=0;a!=p.column.length;++a)if(p.column[a].id==n(this).attr("class")){e=p.column[a];break}null!=e&&(t[e.id]="hidden"==e.type?n(this).text():"image"==e.type?n(this).find("img").attr("src"):"textInput"==e.type?n(this).find("input").val():n(this).text())}),t}function r(){var e=[];return n("#"+p.id).find("tbody tr").each(function(){for(var t=n(this);0==t.is("tr");)t=t.parent();e.push(i(t))}),e}function o(){for(var e=function(e,t){for(;0==e.is("tr");)e=e.parent();var n=i(e),r={remove:function(){e.remove()},mod:function(t){a(e,t)},moveUp:function(){var t=e.prev();e.insertBefore(t)},moveDown:function(){var t=e.next();t.insertBefore(e)}};t(n,r)},t=0;t!=p.operate.length;++t)!function(t){n("#"+p.id).find(".operate_"+p.operate[t].id).unbind("click").click(function(){e(n(this),p.operate[t].click)})}(t);for(var t=0;t!=p.column.length;++t)!function(t){var a=p.column[t];"textInput"==a.type&&0==_.isUndefined(a.change)&&n("#"+p.id).find("."+a.id+" input").unbind("input").on("input",function(){e(n(this),a.change)})}(t)}function l(e){for(var t="",a=0;a!=p.operate.length;++a)p.operate[a].id=n.uniqueNum(),t+="<a href='javascript: void(0)' class=operate_"+p.operate[a].id+">"+p.operate[a].name+"</a>&nbsp;";for(var i="",a=0;a!=e.length;++a){var r=e[a];if(0==p.operate.length)var o="width:"+1/p.column.length*100+"%;";else var o="width:"+1/(p.column.length+1)*100+"%;";i+="<tr>";for(var l=0;l!=p.column.length;++l){var d=p.column[l];i+="hidden"==d.type?'<td style="display:none;'+o+'" class="'+d.id+'">'+r[d.id]+"</td>":"image"==d.type?'<td style="'+o+'" class="'+d.id+'"><img src="'+r[d.id]+'"/></td>':"textInput"==d.type?'<td style="'+o+'" class="'+d.id+'"><input style="width:100%" type="text" value="'+r[d.id]+'"/></td>':'<td style="'+o+'" class="'+d.id+'">'+r[d.id]+"</td>"}0!=p.operate.length&&(i+='<td style="'+o+'" >'+t+"</td>"),i+="</tr>"}return i}function d(e){var t="";if(t+='<div class="mod-basic">',t+='<table class="mod_table" style="table-layout: auto;">',t+="<thead><tr>",0==p.operate.length)var a="width:"+1/p.column.length*100+"%;";else var a="width:"+1/(p.column.length+1)*100+"%;";for(var i=0;i!=p.column.length;++i){var r=p.column[i];t+="hidden"==r.type?'<th style="display:none;'+a+'"><span class="label">'+r.name+"</span></th>":'<th style="'+a+'"><span class="label">'+r.name+"</span></th>"}0!=p.operate.length&&(t+='<th><span class="label" style="'+a+'">操作</span></th>'),t+="</tr></thead>",t+="<tbody>",t+=l(e),t+="</tbody>",t+="</table>",t+="</div>",t=n(t),n("#"+p.id).empty(),n("#"+p.id).append(t),o()}function s(e){n("#"+p.id).find("tbody").append(l(e)),o()}function c(e){n("#"+p.id).find("tbody").prepend(l(e)),o()}var p={id:"",data:"",column:[],operate:[]};for(var u in e)p[u]=e[u];return d(p.data),{preadd:c,add:s,get:r,clear:t}},staticTable:function(e){function t(e){var t={};return e.find("td").each(function(){for(var e=n(this).attr("class"),a=null,i=0;i!=o.column.length;++i){var r=o.column[i];if(r.id==e){a=r;break}}if(null!=a){var l;l="image"==a.type?n(this).find("img").attr("src"):n(this).text(),t[e]=l}}),t}function a(e,t,a){var i=n.url.toInfo(a),r={};for(var o in s.fields)r[o]=s.fields[o].thText;var l="";l+='<form action="'+i.originpathname+'" method="get" style="display:none">';for(o in i.search){var d=o,c=i.search[o];l+='<input type="text" name="'+d+'" class="input-small" value="'+encodeURIComponent(c)+'"/>'}l+='<input type="text" name="_viewTitle" class="input-small" value="'+encodeURIComponent(t)+'"/>',l+='<input type="text" name="_view" class="input-small" value="'+encodeURIComponent(e)+'"/>',l+='<input type="text" name="_viewFormat" class="input-small" value="'+encodeURIComponent(JSON.stringify(r))+'"/>',l+="</form>",l=n(l),n("body").append(l),l.submit()}function r(e,t,a,r){var o={};for(var l in s.fields)o[l]=s.fields[l].thText;n.get("/export/"+e,{source:d,pageIndex:a,pageSize:r,viewTitle:t,viewFormat:JSON.stringify(o)},function(e){return e=n.JSON.parse(e),0!=e.code?void i.message(e.msg):void i.message("导出数据成功，请稍候留意邮箱！")})}var o={id:"",params:{},column:[],operate:[],checkAll:!1,url:null,pageIndex:0,pageSize:10,nextPage:void 0};for(var l in e)o[l]=e[l];var d="",s={};s.table_id=o.id,s.key_index="_id",s.order_field="_id",s.order_type="asc",s.page_size=o.pageSize,s.index=Math.floor(o.pageIndex/o.pageSize),s.container_class="mod-basic box-table",s.table_class="mod_table",s.fields="",s.summary="",s.ifRealPage=!0,s.data="data",s.checkAll=o.checkAll,s.params={};var c="";for(var l in o.params)""!=n.trim(o.params[l])&&"pageIndex"!=l&&"pageSize"!=l&&(c+=l+"="+encodeURIComponent(n.trim(o.params[l]))+"&");c+="t="+(new Date).getTime(),d=-1==o.url.indexOf("?")?o.url+"?"+c:o.url+"&"+c,s.fields={};for(var l=0;l!=o.column.length;++l){var p=o.column[l];if(!p.hidden){var u;!function(e){"text"==e.type?u={thText:e.name}:"enum"==e.type?u={thText:e.name,format:function(t){return e.map[t]}}:"check"==e.type?u={thText:e.name,format:function(t){for(var a=[],n=0;n!=t.length;++n)a.push(e.map[t[n]]);return a.join(",")}}:"image"==e.type&&(u={thText:e.name,format:function(e){return'<img src="'+e+'" style="width:100%;max-width:128px;">'}})}(p),s.fields[p.id]=u}}for(var f="",l=0;l!=o.operate.length;++l)o.operate[l].id=n.uniqueNum(),f+="<a href='#' class=operate_"+o.operate[l].id+">"+o.operate[l].name+"</a>&nbsp;";return 0!=o.operate.length&&(s.fields._oper={thText:"操作",format:function(){return f}}),n.get(d+"&pageIndex="+o.pageIndex+"&pageSize="+o.pageSize,{},function(e){if(e=n.JSON.parse(e),0!=e.code)return void i.message(e.msg);var t=n("<div/>"),a=e.data[s.data];for(var r in a){var l=a[r];for(var c in l){var p=l[c];"string"==typeof p&&(l[c]=t.text(p).html())}}dt=GRI.initDataTable({resultObj:e,name:s.data,tableId:s.table_id,data:e.data[s.data],summary:s.summary,allFields:s.fields,layout:"auto",checkAll:s.checkAll,enableThClick:!0,stopToday:!1,keyIndex:s.key_index,page:{orderField:s.order_field,orderType:s.order_type,ifRealPage:s.ifRealPage,size:s.page_size,rowCount:e.data.count,index:s.index,url:d},cssSetting:{containerClass:s.container_class,tableClass:s.table_class},callbackJson:o.nextPage,callback:function(){for(var e=0;e!=o.operate.length;++e)!function(e){n(".operate_"+o.operate[e].id).unbind("click").click(function(){for(var t=n(this);0==t.is("tr");)t=t.parent();var a={};t.find("td").each(function(){a[n(this).attr("class")]=n(this).text()}),o.operate[e].click(a)})}(e)}})}),{getCheckData:function(){var e=n("#"+o.id+" .gri_td_checkbox:checked"),a=[];return e.each(function(){var e=n(this).parent().parent();a.push(t(e))}),a},exportDataToTxt:function(e){i.input("请输入需要导出txt的页数（不填代表导出本页数据）",function(t){""==t&&(t=1),t=o.pageSize*t;var n=d+"&pageIndex="+o.pageIndex+"&pageSize="+t;a("txt",e,n)})},exportDataToExcel:function(e){i.input("请输入需要导出excel的页数（不填代表导出本页数据）",function(t){""==t&&(t=1),t=o.pageSize*t;var n=d+"&pageIndex="+o.pageIndex+"&pageSize="+t;a("excel",e,n)})},exportDataToExcelAsync:function(e){i.input("请输入需要导出excel的页数（不填代表导出本页数据）",function(t){""==t&&(t=1),t=o.pageSize*t,r("excel",e,o.pageIndex,t)})}}}}});
;define("fishstrap/ui/editor.js",function(t,e,n){var r=t("fishstrap/core/global.js");n.exports={fullEditor:function(t){var e={id:"",url:"/uedit/control"};for(var n in t)e[n]=t[n];var i=r.uniqueNum();r("#"+e.id).html('<script id="'+i+'" name="content" type="text/plain"></script>');var a=UE.getEditor(i,{serverUrl:e.url,autoHeightEnabled:!0});return{getContent:function(){return a.getContent()},setContent:function(t){a.ready(function(){a.setContent(t)})}}},simpleEditor:function(t){var e={id:"",url:"/uedit/control"};for(var n in t)e[n]=t[n];var i=r.uniqueNum();r("#"+e.id).html('<script id="'+i+'" name="content" type="text/plain"></script>');var a=UE.getEditor(i,{serverUrl:e.url,autoHeightEnabled:!0,toolbars:[["fullscreen","source","simpleupload"]]});return{getContent:function(){return a.getContent()},getFormatData:function(){var t=a.getContent(),e=r(t),n=[];return e.each(function(){var t=r(this);if(1==t.is("p")){var e=t.has("img"),i=t.text();e&&!i?t.find("img").each(function(){n.push({type:1,data:r(this).attr("src")})}):e&&i?(n.push({type:0,data:t.text()}),t.find("img").each(function(){n.push({type:1,data:r(this).attr("src")})})):n.push({type:0,data:t.text()})}else 1==t.is("img")&&n.push({type:1,data:t.attr("src")})}),n},setContent:function(t){a.ready(function(){a.setContent(t)})},setFormatData:function(t){var e="";for(var n in t){var r=t[n];0==r.type?e+="<p>"+r.data+"</p>":1==r.type&&(e+='<p><img src="'+r.data+'"/></p>')}a.ready(function(){a.setContent(e)})}}}}});