define("fishstrap/ui/tree.js",function(n,i,t){var a=n("fishstrap/core/global.js");t.exports={easyTree:function(n){function i(n){var t="";if(n.children&&(t='<span class="expand" style="font-size:16px;">&nbsp;-&nbsp;&nbsp;</span>'),s+='<div class="item btn" data="'+n.data+'" style="margin-top:10px;display:block;text-align:left;">'+t+"<span>"+n.name+"</span></div>",n.children){s+='<div style="padding-left:30px;">';for(var a in n.children)i(n.children[a]);s+="</div>"}}var t={id:"",data:[],button:[],value:null,click:function(){}};for(var d in n)t[d]=n[d];var s="";s+="<div>";for(var d in t.button){var r=t.button[d];r.id=a.uniqueNum(),s+='<button class="btn btn-warn '+r.id+'">'+r.name+"</button>&nbsp;"}s+="</div>",s+="<div>";for(var d in t.data){var e=t.data[d];i(e)}s+="</div>",s=a(s),a("#"+t.id).empty(),a("#"+t.id).append(s),s.find(".item").click(function(){var n=a(this).attr("data");s.find(".btn").removeClass("btn-primary"),a(this).addClass("btn-primary"),t.click(n)}),s.find(".expand").click(function(){var n=-1!=a(this).text().indexOf("+"),i=a(this).parent().next();i.slideToggle(),a(this).html(n?"&nbsp;-&nbsp;&nbsp;":"&nbsp;+&nbsp;&nbsp;")});for(var d in t.button){var r=t.button[d];!function(n){s.find("."+n.id).click(function(){var i=s.find(".btn-primary").attr("data");n.click(i)})}(r)}null!=t.value&&s.find(".item[data="+t.value+"]").click()}}});