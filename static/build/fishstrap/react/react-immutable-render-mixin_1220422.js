define("fishstrap/react/react-immutable-render-mixin.js",function(t,e,r){function n(t,e){if(t===e||i(t,e))return!0;if("object"!=typeof t||null===t||"object"!=typeof e||null===e)return!1;var r=Object.keys(t),n=Object.keys(e);if(r.length!==n.length)return!1;for(var o=Object.prototype.hasOwnProperty.bind(e),u=0;u<r.length;u++)if(!o(r[u])||!i(t[r[u]],e[r[u]]))return!1;return!0}var i=Immutable.is.bind(Immutable),o={shouldComponentUpdate:function(t,e){return!n(this.props,t)||!n(this.state,e)}};r.exports=o});