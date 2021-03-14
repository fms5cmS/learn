package some

import (
	"testing"
)

// Go 的代码块和作用域
// https://mp.weixin.qq.com/s?__biz=MzIyNzM0MDk0Mg==&mid=2247484052&idx=1&sn=eab928ea5b67046f4cdc933634db17cc&chksm=e863e575df146c63710dd1ea180c58b3eff451d0247b3e11cad71b1369e03726fecfca6eeb8f&scene=21#wechat_redirect
func TestScopeAndBlock(t *testing.T) {
	if a:=1;false{
	}else if b:=2;false{
	}else{
		println(a,b)
	}
}
