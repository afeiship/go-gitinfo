# weibo_sso:
1. 第一步，请求这个URL，得到JSONP, GET请求
+ https://login.sina.com.cn/sso/prelogin.php?entry=sso&callback=sinaSSOController.preloginCallBack&su=MTgxMDgwNjMwMDU%3D&rsakt=mod&client=ssologin.js(v1.4.15)&_=1535730587143

```js
sinaSSOController.preloginCallBack(
    {
        "retcode": 0,
        "servertime": 1535730692,
        "pcid": "gz-93fd8772ef0c1c8f6fdc3029c3355a6ba9d0",
        "nonce": "8UCHGD",
        "pubkey": "EB2A38568661887FA180BDDB5CABD5F21C7BFD59C090CB2D245A87AC253062882729293E5506350508E7F9AA3BB77F4333231490F915F6D63C55FE2F08A49B353F444AD3993CACC02DB784ABBB8E42A9B1BBFFFB38BE18D78E87A0E41B9B8F73A928EE0CCEE1F6739884B9777E4FE9E88A1BBE495927AC4A799B3181D6442443",
        "rsakv": "1330428213",
        "is_openlock": 0,
        "exectime": 85
    }
)
```

2. login: POST请求
+ https://login.sina.com.cn/sso/login.php?client=ssologin.js(v1.4.15)&_=1535730587298
```conf
# 请求参数如下：

entry: sso
gateway: 1
from: null
savestate: 0
useticket: 0
pagerefer: http://login.sina.com.cn/sso/login.php?client=ssologin.js(v1.4.18)
vsnf: 1
su: encrypt_username
service: sso
servertime: 1535730587
nonce: V18C60
pwencode: rsa2
rsakv: 1330428213
sp: encrypt_password
sr: 1920*1200
encoding: UTF-8
cdult: 3
domain: sina.com.cn
prelt: 56
returntype: TEXT
_: 时间戳
```

## 2. 参数
1. su:
base64(encodeURIComponent(username))

2. servertime
用上面返回的 servertime

3. nonce
用上面返回的 nonce

4. rsakv
用上面返回的 rsakv

5. sp 就是下面的 password
```js
    //算密码值：
	let RSAKey = new sinaSSOEncoder.RSAKey();
	let {nonce, pubkey, servertime, rsakv} = await preLogin()
	RSAKey.setPublic(pubkey, "10001"); //RES E/N[mode]
    passwd = RSAKey.encrypt([servertime, nonce].join("\t") + "\n" + password)
```

## 对应java版
```java
byte[] encryptedContentKey = DigestUtilPlus.RSA256.encryptPublicKey(content, weiBoPreLogin.getPubkey(), "10001");
String encryptPassword = Hex.encodeHexString(encryptedContentKey);

//
public static byte[] encryptPublicKey(String content, String publicKeyModulus, String publicKeyExponent) throws Exception {
    KeyFactory factory = KeyFactory.getInstance("RSA");

    BigInteger modulus = new BigInteger(publicKeyModulus, 16);
    BigInteger publicExponent = new BigInteger(publicKeyExponent, 16);
    RSAPublicKeySpec pubKeySpec = new RSAPublicKeySpec(modulus, publicExponent);
    RSAPublicKey pubKey = (RSAPublicKey) factory.generatePublic(pubKeySpec);

    Cipher cipher = Cipher.getInstance("RSA");
    cipher.init(Cipher.ENCRYPT_MODE, pubKey);
    return cipher.doFinal(content.getBytes());
}
```

6. prelt: 56
这里生成100-1000之间的随机数?

这里的返回值如下：
```json
private String retcode;
private String uid;
private String nick;
private List<String> crossDomainUrlList;
{
    uid:'xxx',
    retcode:0,
    "crossDomainUrlList": [
        "https:\/\/passport.weibo.com\/wbsso\/crossdomain?service=krvideo&savestate=1&ticket=ST-NjA1MjYwMjIxOA%3D%3D-1535710747-gz-AD5403C38C4D32B722FA7D6B4A341727-1&ssosavestate=1567246747",
    ]
}
```

3. 第三步：
把得到有的ticket的url发送GET请求
+ https://passport.weibo.com\/wbsso\/crossdomain?service=krvideo&savestate=1&ticket=ST-NjA1MjYwMjIxOA%3D%3D-1535710747-gz-AD5403C38C4D32B722FA7D6B4A341727-1&ssosavestate=1567246747

拿到最终的 Set-Cookie['SUB']这个COOkie






