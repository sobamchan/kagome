package data

import(
	"os"
	"time"
)


func dicIpaMorphDicBytes() ([]byte, error) {
	return bindataRead(
		_dicIpaMorphDic,
		"dic/ipa/morph.dic",
	)
}

func dicIpaMorphDic() (*asset, error) {
	bytes, err := dicIpaMorphDicBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "dic/ipa/morph.dic", size: 2352764, mode: os.FileMode(420), modTime: time.Unix(1452316172, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}
