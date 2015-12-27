package rpmheader

import (
	"reflect"
	"github.com/oneumyvakin/rpmfile-go/common"
)

type Rpm_header struct {
	header_start          int64
	header_end            int64
	number_of_entries     int32
	header_structure_size int32
	entries               []common.Entry_header
	store                 []byte

	Headerimmutable []byte   // 63
	Headeri18ntable []string // 100

	Signature  []byte // 267
	Sha1header string
	Name       string
	Version    string // 1001
	Release    string // 1002
	// 1003
	Summary     string // 1004
	Description string // 1005
	Buildtime   int32  // 1006
	Buildhost   string // 1007
	Size        int32  // 1009
	// 1010
	Vendor string // 1011
	// 1012 - 1013
	Copyright string // 1014
	Packager  string // 1015
	Group     string // 1016
	// 1017 - 1020
	Os     string // 1021
	Arch   string // 1022
	Prein  string // 1023
	Postin string // 1024
	Preun  string // 1025

	Filesizes int32 // 1028

	Filemodes   int16    // 1030
	Filerdevs   int16    // 1033
	Filemtimes  int32    // 1034
	Filemd5s    []string // 1035
	Filelinktos []string // 1036
	Fileflags   int32    // 1037

	Fileusername    []string // 1039
	Filegroupname   []string // 1040
	Sourcerpm       string   // 1044
	Fileverifyflags int32    // 1045

	Provides       []string // 1047
	Requireflags   int32    // 1048
	Requirename    []string // 1049
	Requireversion []string // 1050

	Rpmversion string // 1064

	Preinprog         string   // 1085
	Postinprog        string   // 1086
	Preunprog         string   // 1087
	Obsoletename      []string // 1090
	Filedevices       int32    // 1095
	Fileinodes        int32    // 1096
	Filelangs         []string // 1097
	Provideflags      int32    // 1112
	Provideversion    []string // 1113
	Obsoleteflags     int32    // 1114
	Obsoleteversion   []string // 1115
	Dirindexes        int32    // 1116
	Basenames         []string // 1117
	Dirnames          []string // 1118
	Optflags          string   // 1122
	Payloadformat     string   // 1124
	Payloadcompressor string   // 1125
	Payloadflags      string   // 1126
	Rhnplatform       string   // 1131
	Platform          string   // 1132
	Filecontexts      []string // 1147

	/*
		Fscontexts //1148
		Recontexts //1149
		Policies //1150
		Pretrans //1151
		Posttrans //1152
		Pretransprog //1153
		Posttransprog //1154
		Disttag //1155
		Oldsuggestsname //1156
		Oldsuggestsversion //1157
		Oldsuggestsflags //1158
		Oldenhancesname //1159
		Oldenhancesversion //1160
		Oldenhancesflags //1161
		Priority //1162
		Cvsid //1163
		Blinkpkgid //1164
		Blinkhdrid //1165
		Blinknevra //1166
		Flinkpkgid //1167
		Flinkhdrid //1168
		Flinknevra //1169
		Packageorigin //1170
		Triggerprein //1171
		Buildsuggests //1172
		Buildenhances //1173
		Scriptstates //1174
		Scriptmetrics //1175
		Buildcpuclock //1176
		Filedigestalgos //1177
		Variants //1178
		Xmajor //1179
		Xminor //1180
		Repotag //1181
		Keywords //1182
		Buildplatforms //1183
		Packagecolor //1184
		Packageprefcolor //1185
		Xattrsdict //1186
		Filexattrsx //1187
		Depattrsdict //1188
		Conflictattrsx //1189
		Obsoleteattrsx //1190
		Provideattrsx //1191
		Requireattrsx //1192
		Buildprovides //1193
		Buildobsoletes //1194
		Dbinstance //1195
		Nvra //1196
		Filenames //5000
		Fileprovide //5001
		Filerequire //5002
		Fsnames //5003
		Fssizes //5004
		Triggerconds //5005
		Triggertype //5006
		Origfilenames //5007
		Longfilesizes //5008
		Longsize //5009
		Filecaps //5010
		Filedigestalgo //5011
		Bugurl //5012
		Evr //5013
		Nvr //5014
		Nevr //5015
		Nevra //5016
		Headercolor //5017
		Verbose //5018
		Epochnum //5019
		Preinflags //5020
		Postinflags //5021
		Preunflags //5022
		Postunflags //5023
		Pretransflags //5024
		Posttransflags //5025
		Verifyscriptflags //5026
		Triggerscriptflags //5027
		Collections //5029
		Policynames //5030
		Policytypes //5031
		Policytypesindexes //5032
		Policyflags //5033
		Vcs //5034
		Ordername //5035
		Orderversion //5036
		Orderflags //5037
		Mssfmanifest //5038
		Mssfdomain //5039
		Instfilenames //5040
		Requirenevrs //5041
		Providenevrs //5042
		Obsoletenevrs //5043
		Conflictnevrs //5044
		Filenlinks //5045
		Recommendname //5046
		Recommendversion //5047
		Recommendflags //5048
		Suggestname //5049
		Suggestversion //5050
		Suggestflags //5051
		Supplementname //5052
		Supplementversion //5053
		Supplementflags //5054
		Enhancename //5055
		Enhanceversion //5056
		Enhanceflags //5057
		Recommendnevrs //5058
		Suggestnevrs //5059
		Supplementnevrs //5060
		Enhancenevrs //5061
	*/
}

func (rpm_header *Rpm_header) Get_tags() map[int]string {
	return map[int]string{
		61:   "Headerimage",
		62:   "Headersignatures",
		63:   "Headerimmutable",
		64:   "Headerregions",
		100:  "Headeri18ntable",
		267:  "Signature",
		269:  "Sha1header",
		1000: "Name",
		1001: "Version",
		1002: "Release",
		1003: "Serial",
		1004: "Summary",
		1005: "Description",
		1006: "Buildtime",
		1007: "Buildhost",
		1008: "Installtime",
		1009: "Size",
		1010: "Distribution",
		1011: "Vendor",
		1012: "Gif",
		1013: "Xpm",
		1014: "Copyright",
		1015: "Packager",
		1016: "Group",
		1017: "Changelog",
		1018: "Source",
		1019: "Patch",
		1020: "Url",
		1021: "Os",
		1022: "Arch",
		1023: "Prein",
		1024: "Postin",
		1025: "Preun",
		1026: "Postun",
		1027: "Filenames",
		1028: "Filesizes",
		1029: "Filestates",
		1030: "Filemodes",
		1031: "Fileuids",
		1032: "Filegids",
		1033: "Filerdevs",
		1034: "Filemtimes",
		1035: "Filemd5s",
		1036: "Filelinktos",
		1037: "Fileflags",
		1038: "Root",
		1039: "Fileusername",
		1040: "Filegroupname",
		1041: "Exclude",
		1042: "Exclusive",
		1043: "Icon",
		1044: "Sourcerpm",
		1045: "Fileverifyflags",
		1046: "Archivesize",
		1047: "Provides",
		1048: "Requireflags",
		1049: "Requirename",
		1050: "Requireversion",
		1051: "Nosource",
		1052: "Nopatch",
		1053: "Conflictflags",
		1054: "Conflictname",
		1055: "Conflictversion",
		1056: "Defaultprefix",
		1057: "Buildroot",
		1058: "Installprefix",
		1059: "Excludearch",
		1060: "Excludeos",
		1061: "Exclusivearch",
		1062: "Exclusiveos",
		1063: "Autoreqprov",
		1064: "Rpmversion",
		1065: "Triggerscripts",
		1066: "Triggername",
		1067: "Triggerversion",
		1068: "Triggerflags",
		1069: "Triggerindex",
		1079: "Verifyscript",
		1080: "Changelogtime",
		1081: "Changelogname",
		1082: "Changelogtext",
		1083: "Brokenmd5",
		1084: "Prereq",
		1085: "Preinprog",
		1086: "Postinprog",
		1087: "Preunprog",
		1088: "Postunprog",
		1089: "Buildarchs",
		1090: "Obsoletename",
		1091: "Verifyscriptprog",
		1092: "Triggerscriptprog",
		1093: "Docdir",
		1094: "Cookie",
		1095: "Filedevices",
		1096: "Fileinodes",
		1097: "Filelangs",
		1098: "Prefixes",
		1099: "Instprefixes",
		1100: "Triggerin",
		1101: "Triggerun",
		1102: "Triggerpostun",
		1103: "Autoreq",
		1104: "Autoprov",
		1105: "Capability",
		1106: "Sourcepackage",
		1107: "Oldorigfilenames",
		1108: "Buildprereq",
		1109: "Buildrequires",
		1110: "Buildconflicts",
		1111: "Buildmacros",
		1112: "Provideflags",
		1113: "Provideversion",
		1114: "Obsoleteflags",
		1115: "Obsoleteversion",
		1116: "Dirindexes",
		1117: "Basenames",
		1118: "Dirnames",
		1119: "Origdirindexes",
		1120: "Origbasenames",
		1121: "Origdirnames",
		1122: "Optflags",
		1123: "Disturl",
		1124: "Payloadformat",
		1125: "Payloadcompressor",
		1126: "Payloadflags",
		1127: "Installcolor",
		1128: "Installtid",
		1129: "Removetid",
		1130: "Shar1rhn",
		1131: "Rhnplatform",
		1132: "Platform",
		1133: "Patchesname",
		1134: "Patchesflags",
		1135: "Patchesversion",
		1136: "Cachectime",
		1137: "Cachepkgpath",
		1138: "Cachepkgsize",
		1139: "Cachepkgmtime",
		1140: "Filecolors",
		1141: "Fileclass",
		1142: "Classdict",
		1143: "Filedependsx",
		1144: "Filedependsn",
		1145: "Depednsdict",
		1146: "Sourcepkgid",
		1147: "Filecontexts",
		1148: "Fscontexts",
		1149: "Recontexts",
		1150: "Policies",
		1151: "Pretrans",
		1152: "Posttrans",
		1153: "Pretransprog",
		1154: "Posttransprog",
		1155: "Disttag",
		1156: "Oldsuggestsname",
		1157: "Oldsuggestsversion",
		1158: "Oldsuggestsflags",
		1159: "Oldenhancesname",
		1160: "Oldenhancesversion",
		1161: "Oldenhancesflags",
		1162: "Priority",
		1163: "Cvsid",
		1164: "Blinkpkgid",
		1165: "Blinkhdrid",
		1166: "Blinknevra",
		1167: "Flinkpkgid",
		1168: "Flinkhdrid",
		1169: "Flinknevra",
		1170: "Packageorigin",
		1171: "Triggerprein",
		1172: "Buildsuggests",
		1173: "Buildenhances",
		1174: "Scriptstates",
		1175: "Scriptmetrics",
		1176: "Buildcpuclock",
		1177: "Filedigestalgos",
		1178: "Variants",
		1179: "Xmajor",
		1180: "Xminor",
		1181: "Repotag",
		1182: "Keywords",
		1183: "Buildplatforms",
		1184: "Packagecolor",
		1185: "Packageprefcolor",
		1186: "Xattrsdict",
		1187: "Filexattrsx",
		1188: "Depattrsdict",
		1189: "Conflictattrsx",
		1190: "Obsoleteattrsx",
		1191: "Provideattrsx",
		1192: "Requireattrsx",
		1193: "Buildprovides",
		1194: "Buildobsoletes",
		1195: "Dbinstance",
		1196: "Nvra",
		5000: "Filenames",
		5001: "Fileprovide",
		5002: "Filerequire",
		5003: "Fsnames",
		5004: "Fssizes",
		5005: "Triggerconds",
		5006: "Triggertype",
		5007: "Origfilenames",
		5008: "Longfilesizes",
		5009: "Longsize",
		5010: "Filecaps",
		5011: "Filedigestalgo",
		5012: "Bugurl",
		5013: "Evr",
		5014: "Nvr",
		5015: "Nevr",
		5016: "Nevra",
		5017: "Headercolor",
		5018: "Verbose",
		5019: "Epochnum",
		5020: "Preinflags",
		5021: "Postinflags",
		5022: "Preunflags",
		5023: "Postunflags",
		5024: "Pretransflags",
		5025: "Posttransflags",
		5026: "Verifyscriptflags",
		5027: "Triggerscriptflags",
		5029: "Collections",
		5030: "Policynames",
		5031: "Policytypes",
		5032: "Policytypesindexes",
		5033: "Policyflags",
		5034: "Vcs",
		5035: "Ordername",
		5036: "Orderversion",
		5037: "Orderflags",
		5038: "Mssfmanifest",
		5039: "Mssfdomain",
		5040: "Instfilenames",
		5041: "Requirenevrs",
		5042: "Providenevrs",
		5043: "Obsoletenevrs",
		5044: "Conflictnevrs",
		5045: "Filenlinks",
		5046: "Recommendname",
		5047: "Recommendversion",
		5048: "Recommendflags",
		5049: "Suggestname",
		5050: "Suggestversion",
		5051: "Suggestflags",
		5052: "Supplementname",
		5053: "Supplementversion",
		5054: "Supplementflags",
		5055: "Enhancename",
		5056: "Enhanceversion",
		5057: "Enhanceflags",
		5058: "Recommendnevrs",
		5059: "Suggestnevrs",
		5060: "Supplementnevrs",
		5061: "Enhancenevrs",
	}
}

func (rpm_header *Rpm_header) Get_field(name string) (field reflect.Value) {
	reflect_of_rpm_header := reflect.ValueOf(rpm_header).Elem()
	field = reflect_of_rpm_header.FieldByName(name)

	return
}

func (r Rpm_header) Get_number_of_entries() int32 {
	return r.number_of_entries
}
func (r Rpm_header) Get_entries() []common.Entry_header {
	return r.entries
}
func (r Rpm_header) Get_header_structure_size() int32 {
	return r.header_structure_size
}
func (r Rpm_header) Get_header_start() int64 {
	return r.header_start
}
func (r Rpm_header) Get_header_end() int64 {
	return r.header_end
}
func (r Rpm_header) Get_store() []byte {
	return r.store
}

func (r *Rpm_header) Set_number_of_entries(number_of_entries int32) {
	r.number_of_entries = number_of_entries
}
func (r *Rpm_header) Set_entries(entries []common.Entry_header) {
	r.entries = entries
}
func (r *Rpm_header) Set_header_structure_size(header_structure_size int32) {
	r.header_structure_size = header_structure_size
}
func (r *Rpm_header) Set_header_start(header_start int64) {
	r.header_start = header_start
}
func (r *Rpm_header) Set_header_end(header_end int64) {
	r.header_end = header_end
}
func (r *Rpm_header) Set_store(store []byte) {
	r.store = store
}
