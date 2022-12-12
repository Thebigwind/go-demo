package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

const (
	ZdlzRootPubkey = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqRiytRATcOx2rafZTy08\ny5wjfg6G+oaG4f7imGg4vyked3v12j/W1JNN7QGGV8ZR8e3aQ0wS4KKvFVvbFKZY\nkIHwIzsGvZS242cTmViKMyLQB0D0KwxWeveJNzMoQP0d9TzV0cM5+y6CTDVbQmKv\nKfas2w2gI2Q6oQozuVQlSzjpE1ddBuQHZq/qewg+AxMbvhW/I7ge4wMYIFfaPhYC\nsLUAGDJBddY0+9vQsieQXuQBai5UQv1rF40szgSoGLxnNb/qnMdK9Rm39OiH8PWs\nM2zDpB8d04gG8JTvvZwDf0qrlNtma0IjMTfqjQU5PH9wJ3hFU7122TipNcVhr9xf\n8QIDAQAB\n-----END PUBLIC KEY-----\n"
	//ZdlzCfgPubkey  = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA2w2622yzn32zLZgtzVqx\nyS+u5DrVCigUZL/oe+LhdUShDbO/yRYI0O6VIMY7p2Bf2nhovInQP8nfewxrXPfc\nfEiar5xwPjQ9KlTVQh+ITjbv/GQhEEOv9cbbQ9sgHvGcmZF4YUkI9EkMl1A6E72X\nfoUuMw/Nps7khSA8CDenVl3ijsplRpLPaOEbGfxwNrb7SE3K2XnzBr0P63KNENtX\nA6Yp40K+8Q5NQ9i20Hw6BLb7SpusCJiqCDMZfhj2nsPPw2MWmI7YHzDy9wb52MDH\nyv93KbxTDWtnxabOq2OHjMOrbff+Sgm4yxe+T9o7+ie8ONGLAseaAJEHEgp4hH8i\nNwIDAQAB\n-----END PUBLIC KEY-----\n"
	ZdlzCfgPubkey = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAy7JLIGl1JkDTXi47NZpt\n7H5jlWZKJ+SZAgH9ELVVpSYqiqi0WtMgwbUEQeJk1ulP5qEoQHq35szEikrQnork\n2d78AuhYRThyEMvXhqsdc887Gq2k3wY00EXdJENYmX0ITmJbDh1yhhiGowK1Smoe\ni4dxMrOOESJsdaNjkCWMPmXLbJ+WJa4PGa9pUTOqO4qIIT5poI3gokFXX8F8zl/h\nv/PDEwlSd64NrYlKTJR6+aB3KPholU1vY4Zu9/v+MMh5Xgpe9QGz1vMOmkIIC8V9\nnbExdwEZg57Lf/wlhj9iW86utnISzdYPgZoq7zh8BkZdaaBm6eOVxCyIqFrcvYYY\nHwIDAQAB\n-----END PUBLIC KEY-----\n"
	ZdlzIdPubkey  = "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAy7JLIGl1JkDTXi47NZpt\n7H5jlWZKJ+SZAgH9ELVVpSYqiqi0WtMgwbUEQeJk1ulP5qEoQHq35szEikrQnork\n2d78AuhYRThyEMvXhqsdc887Gq2k3wY00EXdJENYmX0ITmJbDh1yhhiGowK1Smoe\ni4dxMrOOESJsdaNjkCWMPmXLbJ+WJa4PGa9pUTOqO4qIIT5poI3gokFXX8F8zl/h\nv/PDEwlSd64NrYlKTJR6+aB3KPholU1vY4Zu9/v+MMh5Xgpe9QGz1vMOmkIIC8V9\nnbExdwEZg57Lf/wlhj9iW86utnISzdYPgZoq7zh8BkZdaaBm6eOVxCyIqFrcvYYY\nHwIDAQAB\n-----END PUBLIC KEY-----\n"
)

func main() {
	//fmt.Println(ZdlzRootPubkey)
	//fmt.Println(ZdlzCfgPubkey)
	//fmt.Println(ZdlzIdPubkey)
	//fmt.Println("-----------------")
	cfgSign := "j3nHdiopsMX/ZBjrMC5/eAKVTZP0XXl8xNPP7C1tvrsr3UAQdcuefJ/fIi6KejJR\noe1WtAWbXxuNpPpPgdlL0NfBF+K8Ek4Nw+qQNsfi85DlciMzaHtLOlR/eBeVpU15\n8VUV+6mp+5J7wmvKoyuxONdyFFI6DTGR1u4HiRXP33FiPDLNpYi80A6JohuPTwlE\npVamGpoGVBm1svAZ64NNsF7e5Qz8IK1QFqVkThGJPMI0kcUixgJUy3ISeExawU9W\np4a/0n2q0g8wdZv+gnT6dfFLR+1zWBxi4BNEWSHEJjZ0DRQVhgYa/UqTBtgWI2Af\nCMg7gE6GfJxwK1v/WU01hA=="
	cfgNew := "qCiMpKGlFpcQUAx205Tmt8a6VNKmX6kQWW7IX5euEMTXcNWXRBEGPXzvtm+I2uYu\nrHW226nLQ0HO/D57azq0MEmhtZHL/ve7CcMlYKlF0pbMuBgdCwk6Y7GvKu9PRiCF\nOeCYnvdOBbtSMVb431NCJljqFpBuYQfE6Uy4sWhVnshkwvYjTg7eWRC4kC117boJ\n/zagoQyiZ4u1cidxkHIbXcez7h/gWzsfL8y+MVWIVFB6+MXEK6ZKGgqFsCEkK049\nSX/VqJnjL0jSC32JAVXchcETODTje4zHIRrglrzLnKJpsM2t7LzXsZa5rntRCQYv\n28fD/vVGxk3Vzs8v++KAug=="

	cfgSign = "K9w5xAvEW33q8Ftk7GdlS3hYfL1u+Yu2iO2bR4DidO7DAcv1ifjGyz8/UNV64+G1\n38tjLoWpnPyezNudRHx+BeaIlLRhbikdedmWgZ/4A51KLxoTD3g4dmOijfnEiz+z\nIrDJbwKRuoXejQWS9aIl9ySbWQPdxA4/uIgL2iOtrcImru5/5XKi6lONLK1NySzE\nGqElNrr2CRfO+UBhsceqWR4bRUx4ArQ9bQI2rP4ErqTAGxZu4FChVs7BMGEijljs\n3VesqETmFgu+TP2690YsOx1bGM9BIzfZDB9CuMCn07watqgVR2SFFbZ0MOlmlJDD\nc09bRQdC7/TuFgNxJ/NZsQ=="
	cfgNew = "ti7W67f1R1HkiYOsRV4lzeouA+2EaZRcR6tYIha7zFGxsQyty6iaoYM6uBGGxXbX\nfMd8ROg0Qb+bWJWak3RvTg3mzoCdRIo6ZRE/vg1grdCo8p6Pclt4xktUnjT9d4jL\nQ5Dvl+GiPzJUWPy/FQ+Im0V+B+LUyfKfCX18Gq2DKZJopNrRiskeNykOird+yceU\nc+HZZrN7AfcIduGTnoovnTuaHJxa3O+IN5JP2ufxfFT9SLvcaCQSpD7xFQi1ljwI\nrAn0xwktJsumccra4bEsZlNe+nRiAdS24UWwF9e3YtjGHAyt38LxwLZ3PZJlXb22\nLRXAzph7K6VZooyUAqkPMQ=="
	fmt.Println(cfgNew)
	if err := PackageSignCheck(cfgSign, "cfg"); err != nil {
		fmt.Printf("cfgsign check err:%v", err.Error())
	}
	rootSign := "CYUEEEa2Y4zju6FPrLhapzUvEOMgV7SAMMi5PYO+NcJQ4yZQhpvGQ666DvxCKDtr\n3vS7+oQ/QAzb613age4bo5FhUDadbvAvt64NBkYw4EQKmZr8lDZOC53bha/sgF/7\ntM93U6MhpOia/0ZXYSlUUaj396I9A6w4P2M/3R3rbwoNnW4yjuM6Tr+/0YaPo4iy\nJctF+ns1NMNETJXjkfqmos1zBaXqYRP3ujCZ6iNv3gSaUMXS2bkdlqwCfE9+hlJd\nCD8C7mp0bvq4rpt/zPc7zaGZAStdTnXBLsOD7d/b3MozmbsBCT2f0JQ65UlNKqea\nTLcF6SSaCIT64gemK+LIAA=="
	rootNew := "AhmeotcIiNH3XRPDEZlNAjnElp6ZrWtGXNYvoZ2SPj3uliC08cRAs3ZX/gO+JSz2\nUVlT2vZIqmEZBNFH3Axz5QUnJX4WTkTX/a5LcOSLmBamMqc4uNzGHsDZcPtEojPF\nTy3f1kjJ/Rd6Qkqt9AsftvXnq1gH2g2lTmCI5P6riu3E2RggikcCCAZr7DhJgfFY\n0CKjYD6Xe4b3Gemjrlz1ZOEugl5CyhpzSajulYPwrpwMKeuBzdL08dA1LfI41V0C\nZzs4anPpY9M2WI0IkXk9zc9XOm4g3lqPAWdCbvbJWNFp2x09bqanf7sOo8ttgxQ8\nawxNOIGJeSY292C3S5ko5g=="

	rootSign = "bZI64AhAeEdM05u9WLy/FuTla3EJsMdFQ0D6P0Ba9JflMJ0NvDVJ/tyqa/j9z2/o\nkfRuZisV1X1HEwwA60m1NNyILZ6lkd0cSXgO7eQDJ2ILmYoLN05pzQsoSApkrLl7\n23uTXjYandOQpRQ6yyusprUHLpYcgqKaxh6W4SVysADNqh02FMUYWd7YJMu2rQ8G\nlrvAXJ3MfJ+Nn9sIrclMHru/kUf6OV963QL5MqrySISUiXOx7/gzJzhWAODnEXEV\nNNBHGpmjGG5+7E8tlMSJBCfDxWG6dfRDouB86WXlCTIYR8DSBsJb1EIWYlYZjOOf\nld6TtSp1E6toXQVfWPa8uA=="
	rootNew = "OoUeEyl7S+K1ZEQ3hVjC+YFvPDmoQvxV+AVwbewhw3lqH3nvj1dcqdTJY/4xDuxx\nx8BsAFu/V5EQP9r8L9ZXgWGHzFJq2xeg7fWeSC/AvhJK5OrbA1FRApnL3UK14gME\nNYnRTLINyE8P9i/S2KRY6iAHvxm4VMrOnZaf3UjMbwlE0br3FSO+2TclAHfkf8tg\na6EjQB3NZ9ABlQruNnN8OoTDXVBQp+chq4krngFXz1boEIRbT+mRbfO+UJkFJnGT\nbVpJJMMJmWK4j+Pwu4xI7sxUruV5H1kY++aY4Z/bd5UxnyRnsEcJfW4vzDbTLEUN\nQsD94vbOBtgW1v3N6otrIQ=="
	fmt.Println(rootNew)
	if err := PackageSignCheck(rootSign, "root"); err != nil {
		fmt.Printf("rootsign check err:%v", err.Error())
	}

	cfgRoll := "ti7W67f1R1HkiYOsRV4lzeouA+2EaZRcR6tYIha7zFGxsQyty6iaoYM6uBGGxXbX\nfMd8ROg0Qb+bWJWak3RvTg3mzoCdRIo6ZRE/vg1grdCo8p6Pclt4xktUnjT9d4jL\nQ5Dvl+GiPzJUWPy/FQ+Im0V+B+LUyfKfCX18Gq2DKZJopNrRiskeNykOird+yceU\nc+HZZrN7AfcIduGTnoovnTuaHJxa3O+IN5JP2ufxfFT9SLvcaCQSpD7xFQi1ljwI\nrAn0xwktJsumccra4bEsZlNe+nRiAdS24UWwF9e3YtjGHAyt38LxwLZ3PZJlXb22\nLRXAzph7K6VZooyUAqkPMQ=="
	RootRoll := "OoUeEyl7S+K1ZEQ3hVjC+YFvPDmoQvxV+AVwbewhw3lqH3nvj1dcqdTJY/4xDuxx\nx8BsAFu/V5EQP9r8L9ZXgWGHzFJq2xeg7fWeSC/AvhJK5OrbA1FRApnL3UK14gME\nNYnRTLINyE8P9i/S2KRY6iAHvxm4VMrOnZaf3UjMbwlE0br3FSO+2TclAHfkf8tg\na6EjQB3NZ9ABlQruNnN8OoTDXVBQp+chq4krngFXz1boEIRbT+mRbfO+UJkFJnGT\nbVpJJMMJmWK4j+Pwu4xI7sxUruV5H1kY++aY4Z/bd5UxnyRnsEcJfW4vzDbTLEUN\nQsD94vbOBtgW1v3N6otrIQ=="

	fmt.Printf(cfgRoll, RootRoll)
}

const (
	RootPubType = "root"
	CfgPubType  = "cfg"
	IdPubType   = "id"
	DownPackage = "./qkms-backend-v0.0.1-f90a2382"
)

//预埋公钥的内容到文件的映射
var ZdlzPubTypeMapFile = map[string]string{
	RootPubType: "zdlz-root.pub",
	CfgPubType:  "zdlz-cfg.pub",
	IdPubType:   "zdlz-id.pub",
}

//要替换的公钥的公钥类型到bin文件的映射
var ReplaceKeySignMapFile = map[string]string{
	RootPubType: "replace-key-root.sig.bin",
	CfgPubType:  "replace-key-cfg.sig.bin",
	IdPubType:   "replace-key-id.sig.bin",
}

//软件包的签名公钥类型到bin文件的映射
var PackageSignMapFile = map[string]string{
	RootPubType: "package-root.sig.bin",
	CfgPubType:  "package-cfg.sig.bin",
	IdPubType:   "package-id.sig.bin",
}

//预埋公钥类型到内容的映射
var ZdlzPubTypeMapContent = map[string]string{
	RootPubType: ZdlzRootPubkey,
	CfgPubType:  ZdlzCfgPubkey,
	IdPubType:   ZdlzIdPubkey,
}

// 使用openssl对软件包签名验证
func PackageSignCheck(sign string, signType string) error {
	fmt.Printf("1111111")
	//1.将签名从base64转成需要的bin格式并写入文件
	signFile := "signfile." + signType
	if err := ioutil.WriteFile(signFile, []byte(sign), 0777); err != nil {
		return err
	}
	command := "cat " + signFile + " | base64 -d > " + PackageSignMapFile[signType]
	fmt.Printf("command:%s\n", command)
	cmd := exec.Command("/bin/bash", "-c", command)
	bytes, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Printf("2222222")
	fmt.Println(string(bytes))

	//2.将预埋的公钥写入文件
	if err = ioutil.WriteFile(ZdlzPubTypeMapFile[signType], []byte(ZdlzPubTypeMapContent[signType]), 0777); err != nil {
		return err
	}
	fmt.Printf("3333333")
	//3.验证签名
	// 格式：openssl dgst -sha256 -verify org.pub -signature data.sig.bin data.txt
	command = "openssl dgst -sha256 -verify " + ZdlzPubTypeMapFile[signType] + " -signature " + PackageSignMapFile[signType] + " " + DownPackage
	fmt.Printf("command:%s\n", command)
	cmd = exec.Command("/bin/bash", "-c", command)
	bytes, err = cmd.Output()
	if err != nil {
		return err
	}
	fmt.Printf("4444444")
	fmt.Println(string(bytes))
	return nil
}
