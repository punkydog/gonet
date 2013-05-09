package cfg

import "os"
import "bufio"
import "regexp"
import "fmt"

var _map map[string]string

const CONFIG_PATH = "./config.ini"

//----------------------------------------------- Singleton method for accessing config.ini
func Get() map[string]string {
	if _map == nil {
		_map = _load_config(CONFIG_PATH)
	}

	return _map
}

func _load_config(path string) (ret map[string]string) {
	ret = make(map[string]string)
	f, err := os.Open(path)

	if err != nil {
		fmt.Println("error opening file %v\n", err)
		os.Exit(1)
	}

	re := regexp.MustCompile(`[\t ]*([0-9A-Za-z_]+)[\t ]*=[\t ]*([^\t\n\f\r# ]+)[\t #]*`)

	r := bufio.NewReader(f)

	for {
		line, e := r.ReadString('\n')

		// empty-line & #comment
		if line == "" || []byte(line)[0] == '#' {
			if e == nil {
				continue
			} else {
				break
			}
		}

		// maping
		slice := re.FindStringSubmatch(line)

		if slice != nil {
			ret[slice[1]] = slice[2]
		}

		if e != nil {
			break
		}
	}

	return
}
