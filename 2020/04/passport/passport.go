package passport

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p Passport) isValid() bool {
	return p.byr != "" && p.iyr != "" && p.eyr != "" && p.hgt != "" && p.hcl != "" && p.ecl != "" && p.pid != ""
}

func (p Passport) addInfo(info []string) {
	switch info[0] {
	case "byr":
		p.byr = info[1]
		break
	case "iyr":
		p.iyr = info[1]
		break
	case "eyr":
		p.eyr = info[1]
		break
	case "hgt":
		p.hgt = info[1]
		break
	case "hcl":
		p.hcl = info[1]
		break
	case "ecl":
		p.ecl = info[1]
		break
	case "pid":
		p.pid = info[1]
		break
	case "cid":
		p.cid = info[1]
		break
	}
}
