define("fishstrap/util/jpegEncoder.js",function(r,a,n){n.exports={JPEGEncoder:function(r){function a(r){for(var a=[16,11,10,16,24,40,51,61,12,12,14,19,26,58,60,55,14,13,16,24,40,57,69,56,14,17,22,29,51,87,80,62,18,22,37,56,68,109,103,77,24,35,55,64,81,104,113,92,49,64,78,87,103,121,120,101,72,92,95,98,112,100,103,99],n=0;64>n;n++){var e=j((a[n]*r+50)/100);1>e?e=1:e>255&&(e=255),T[H[n]]=e}for(var o=[17,18,24,47,99,99,99,99,18,21,26,66,99,99,99,99,24,26,56,99,99,99,99,99,47,66,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99,99],t=0;64>t;t++){var f=j((o[t]*r+50)/100);1>f?f=1:f>255&&(f=255),C[H[t]]=f}for(var i=[1,1.387039845,1.306562965,1.175875602,1,.785694958,.5411961,.275899379],v=0,u=0;8>u;u++)for(var c=0;8>c;c++)M[v]=1/(T[H[v]]*i[u]*i[c]*8),b[v]=1/(C[H[v]]*i[u]*i[c]*8),v++}function n(r,a){for(var n=0,e=0,o=new Array,t=1;16>=t;t++){for(var f=1;f<=r[t];f++)o[a[e]]=[],o[a[e]][0]=n,o[a[e]][1]=t,e++,n++;n*=2}return o}function e(){p=n(K,L),l=n(R,U),D=n(O,Q),E=n(V,W)}function o(){for(var r=1,a=2,n=1;15>=n;n++){for(var e=r;a>e;e++)_[32767+e]=n,x[32767+e]=[],x[32767+e][1]=n,x[32767+e][0]=e;for(var o=-(a-1);-r>=o;o++)_[32767+o]=n,x[32767+o]=[],x[32767+o][1]=n,x[32767+o][0]=a-1+o;r<<=1,a<<=1}}function t(){for(var r=0;256>r;r++)z[r]=19595*r,z[r+256>>0]=38470*r,z[r+512>>0]=7471*r+32768,z[r+768>>0]=-11059*r,z[r+1024>>0]=-21709*r,z[r+1280>>0]=32768*r+8421375,z[r+1536>>0]=-27439*r,z[r+1792>>0]=-5329*r}function f(r){for(var a=r[0],n=r[1]-1;n>=0;)a&1<<n&&(J|=1<<N),n--,N--,0>N&&(255==J?(i(255),i(0)):i(J),N=7,J=0)}function i(r){G.push(q[r])}function v(r){i(r>>8&255),i(255&r)}function u(r,a){var n,e,o,t,f,i,v,u,c,w=0,h=8,g=64;for(c=0;h>c;++c){n=r[w],e=r[w+1],o=r[w+2],t=r[w+3],f=r[w+4],i=r[w+5],v=r[w+6],u=r[w+7];var d=n+u,y=n-u,m=e+v,A=e-v,s=o+i,p=o-i,l=t+f,D=t-f,E=d+l,I=d-l,j=m+s,T=m-s;r[w]=E+j,r[w+4]=E-j;var C=.707106781*(T+I);r[w+2]=I+C,r[w+6]=I-C,E=D+p,j=p+A,T=A+y;var M=.382683433*(E-T),b=.5411961*E+M,x=1.306562965*T+M,_=.707106781*j,F=y+_,G=y-_;r[w+5]=G+b,r[w+3]=G-b,r[w+1]=F+x,r[w+7]=F-x,w+=8}for(w=0,c=0;h>c;++c){n=r[w],e=r[w+8],o=r[w+16],t=r[w+24],f=r[w+32],i=r[w+40],v=r[w+48],u=r[w+56];var J=n+u,N=n-u,P=e+v,S=e-v,k=o+i,q=o-i,z=t+f,H=t-f,K=J+z,L=J-z,O=P+k,Q=P-k;r[w]=K+O,r[w+32]=K-O;var R=.707106781*(Q+L);r[w+16]=L+R,r[w+48]=L-R,K=H+q,O=q+S,Q=S+N;var U=.382683433*(K-Q),V=.5411961*K+U,W=1.306562965*Q+U,X=.707106781*O,Y=N+X,Z=N-X;r[w+40]=Z+V,r[w+24]=Z-V,r[w+8]=Y+W,r[w+56]=Y-W,w++}var $;for(c=0;g>c;++c)$=r[c]*a[c],B[c]=$>0?$+.5|0:$-.5|0;return B}function c(){v(65504),v(16),i(74),i(70),i(73),i(70),i(0),i(1),i(1),i(0),v(1),v(1),i(0),i(0)}function w(r,a){v(65472),v(17),i(8),v(a),v(r),i(3),i(1),i(17),i(0),i(2),i(17),i(1),i(3),i(17),i(1)}function h(){v(65499),v(132),i(0);for(var r=0;64>r;r++)i(T[r]);i(1);for(var a=0;64>a;a++)i(C[a])}function g(){v(65476),v(418),i(0);for(var r=0;16>r;r++)i(K[r+1]);for(var a=0;11>=a;a++)i(L[a]);i(16);for(var n=0;16>n;n++)i(O[n+1]);for(var e=0;161>=e;e++)i(Q[e]);i(1);for(var o=0;16>o;o++)i(R[o+1]);for(var t=0;11>=t;t++)i(U[t]);i(17);for(var f=0;16>f;f++)i(V[f+1]);for(var u=0;161>=u;u++)i(W[u])}function d(){v(65498),v(12),i(3),i(1),i(0),i(2),i(17),i(3),i(17),i(0),i(63),i(0)}function y(r,a,n,e,o){for(var t,i=o[0],v=o[240],c=16,w=63,h=64,g=u(r,a),d=0;h>d;++d)F[H[d]]=g[d];var y=F[0]-n;n=F[0],0==y?f(e[0]):(t=32767+y,f(e[_[t]]),f(x[t]));for(var m=63;m>0&&0==F[m];m--);if(0==m)return f(i),n;for(var A,s=1;m>=s;){for(var p=s;0==F[s]&&m>=s;++s);var l=s-p;if(l>=c){A=l>>4;for(var D=1;A>=D;++D)f(v);l=15&l}t=32767+F[s],f(o[(l<<4)+_[t]]),f(x[t]),s++}return m!=w&&f(i),n}function m(){for(var r=String.fromCharCode,a=0;256>a;a++)q[a]=r(a)}function A(r){if(0>=r&&(r=1),r>100&&(r=100),I!=r){var n=0;n=Math.floor(50>r?5e3/r:200-2*r),a(n),I=r}}function s(){var a=(new Date).getTime();r||(r=50),m(),e(),o(),t(),A(r);(new Date).getTime()-a}var p,l,D,E,I,j=(Math.round,Math.floor),T=new Array(64),C=new Array(64),M=new Array(64),b=new Array(64),x=new Array(65535),_=new Array(65535),B=new Array(64),F=new Array(64),G=[],J=0,N=7,P=new Array(64),S=new Array(64),k=new Array(64),q=new Array(256),z=new Array(2048),H=[0,1,5,6,14,15,27,28,2,4,7,13,16,26,29,42,3,8,12,17,25,30,41,43,9,11,18,24,31,40,44,53,10,19,23,32,39,45,52,54,20,22,33,38,46,51,55,60,21,34,37,47,50,56,59,61,35,36,48,49,57,58,62,63],K=[0,0,1,5,1,1,1,1,1,1,0,0,0,0,0,0,0],L=[0,1,2,3,4,5,6,7,8,9,10,11],O=[0,0,2,1,3,3,2,4,3,5,5,4,4,0,0,1,125],Q=[1,2,3,0,4,17,5,18,33,49,65,6,19,81,97,7,34,113,20,50,129,145,161,8,35,66,177,193,21,82,209,240,36,51,98,114,130,9,10,22,23,24,25,26,37,38,39,40,41,42,52,53,54,55,56,57,58,67,68,69,70,71,72,73,74,83,84,85,86,87,88,89,90,99,100,101,102,103,104,105,106,115,116,117,118,119,120,121,122,131,132,133,134,135,136,137,138,146,147,148,149,150,151,152,153,154,162,163,164,165,166,167,168,169,170,178,179,180,181,182,183,184,185,186,194,195,196,197,198,199,200,201,202,210,211,212,213,214,215,216,217,218,225,226,227,228,229,230,231,232,233,234,241,242,243,244,245,246,247,248,249,250],R=[0,0,3,1,1,1,1,1,1,1,1,1,0,0,0,0,0],U=[0,1,2,3,4,5,6,7,8,9,10,11],V=[0,0,2,1,2,4,4,3,4,7,5,4,4,0,1,2,119],W=[0,1,2,3,17,4,5,33,49,6,18,65,81,7,97,113,19,34,50,129,8,20,66,145,161,177,193,9,35,51,82,240,21,98,114,209,10,22,36,52,225,37,241,23,24,25,26,38,39,40,41,42,53,54,55,56,57,58,67,68,69,70,71,72,73,74,83,84,85,86,87,88,89,90,99,100,101,102,103,104,105,106,115,116,117,118,119,120,121,122,130,131,132,133,134,135,136,137,138,146,147,148,149,150,151,152,153,154,162,163,164,165,166,167,168,169,170,178,179,180,181,182,183,184,185,186,194,195,196,197,198,199,200,201,202,210,211,212,213,214,215,216,217,218,226,227,228,229,230,231,232,233,234,242,243,244,245,246,247,248,249,250];this.encode=function(r,a){var n=(new Date).getTime();a&&A(a),G=new Array,J=0,N=7,v(65496),c(),h(),w(r.width,r.height),g(),d();var e=0,o=0,t=0;J=0,N=7,this.encode.displayName="_encode_";for(var i,u,m,s,I,j,T,C,x,_=r.data,B=r.width,F=r.height,q=4*B,H=0;F>H;){for(i=0;q>i;){for(I=q*H+i,j=I,T=-1,C=0,x=0;64>x;x++)C=x>>3,T=4*(7&x),j=I+C*q+T,H+C>=F&&(j-=q*(H+1+C-F)),i+T>=q&&(j-=i+T-q+4),u=_[j++],m=_[j++],s=_[j++],P[x]=(z[u]+z[m+256>>0]+z[s+512>>0]>>16)-128,S[x]=(z[u+768>>0]+z[m+1024>>0]+z[s+1280>>0]>>16)-128,k[x]=(z[u+1280>>0]+z[m+1536>>0]+z[s+1792>>0]>>16)-128;e=y(P,M,e,p,D),o=y(S,b,o,l,E),t=y(k,b,t,l,E),i+=32}H+=8}if(N>=0){var K=[];K[1]=N+1,K[0]=(1<<N+1)-1,f(K)}v(65497);var L="data:image/jpeg;base64,"+btoa(G.join(""));G=[];(new Date).getTime()-n;return L},s()},getImageDataFromImage:function(r){var a="string"==typeof r?document.getElementById(r):r,n=document.createElement("canvas");n.width=a.width,n.height=a.height;var e=n.getContext("2d");return e.drawImage(a,0,0),e.getImageData(0,0,n.width,n.height)}}});