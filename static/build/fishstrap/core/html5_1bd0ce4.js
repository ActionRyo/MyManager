define("fishstrap/core/html5.js",function(e){var r=e("fishstrap/core/global.js"),o={};return o.localStorage=function(e,o,t){function n(){try{return"localStorage"in window&&null!==window.localStorage}catch(e){return!1}}if(0==n())return null;if("undefined"==typeof o){var a=localStorage.getItem(e);if("undefined"==typeof a||null==a)return null;var i=new Date;return a=r.JSON.parse(r.base64.decode(a)),a.expires&&Date.parse(a.expires)<Date.parse(i)?(localStorage.removeItem(e),null):a.data}if(t=t||{},null===o)return o="",localStorage.removeItem(e),null;var l={};if(l.data=o,t.expires&&("number"==typeof t.expires||t.expires.toUTCString)){var f;"number"==typeof t.expires?(f=new Date,f.setTime(f.getTime()+24*t.expires*60*60*1e3)):f=t.expires,l.expires=f.toUTCString()}for(var d=0,u=localStorage.length;u>d;++d){var s=localStorage.key(d),o=localStorage.getItem(s);try{o=r.JSON.parse(r.base64.decode(o))}catch(c){continue}var i=new Date;o.expires&&Date.parse(o.expires)<Date.parse(i)&&localStorage.removeItem(s)}localStorage.removeItem(e),localStorage.setItem(e,r.base64.encode(r.JSON.stringify(l)))},o.history={pushState:function(e,r,o){window.history.pushState?window.history.pushState(e,r,o):location.href=o},replaceState:function(e,r,o){window.history.replaceState?window.history.replaceState(e,r,o):location.href=o}},o.arraybuffer={fromString:function(e){for(var r=new ArrayBuffer(e.length),o=new Uint8Array(r),t=0;t<o.length;t++)o[t]=e.charCodeAt(t);return r}},o.blob={fromArray:function(e){try{var r=new Blob([e],{type:"image/jpeg"})}catch(o){if(alert(o.name),alert(o),window.BlobBuilder=window.BlobBuilder||window.WebKitBlobBuilder||window.MozBlobBuilder||window.MSBlobBuilder,alert(window.BlobBuilder),"TypeError"==o.name&&window.BlobBuilder){alert("TypeError!");var t=new BlobBuilder;t.append(e.buffer);var r=t.getBlob("image/jpeg");alert(r.size)}else if("InvalidStateError"==o.name)var r=new Blob([e.buffer],{type:"image/jpeg"});else var r=null}return r},fromString:function(e){for(var r=new Uint8Array(e.length),o=0,t=e.length;t>o;o++)r[o]=e.charCodeAt(o);return this.fromArray(r)}},o.fileReader={open:function(e){e=e||{};var o={file:null,mode:"binary",onStart:function(){},onProgress:function(){},onSuccess:function(){},onFail:function(){},onStop:function(){}};for(var t in e)o[t]=e[t];if(o.mode=o.mode.toLowerCase(),o.onStart(),"undefined"==typeof window.FileReader||"undefined"==typeof window.Blob)return o.onFail("您的浏览器不支持FileReader"),void o.onStop();var n=new FileReader;n.onabort=function(){o.onFail("读取文件失败"),o.onStop()},n.onerror=function(){o.onFail("读取文件失败"),o.onStop()},n.onprogress=function(e){var r=0;if(e.lengthComputable)var r=Math.ceil(100*(e.loaded/e.total));o.onProgress(r)},n.onload=function(){"binary"==o.mode&&"undefined"==typeof n.readAsBinaryString?(o.onSuccess(r.base64.decode(this.result.substr(this.result.indexOf(",")+1),!1)),o.onStop()):(o.onSuccess(this.result),o.onStop())},"text"==o.mode?n.readAsText(o.file):"dataurl"==o.mode?n.readAsDataURL(o.file):"arraybuffer"==o.mode?n.readAsArrayBuffer(o.file):(o.mode="binary",n.readAsBinaryString?n.readAsBinaryString(o.file):n.readAsDataURL(o.file))}},o});