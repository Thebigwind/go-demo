package main

import (
	"fmt"
	"log"
)

func main() {
	i := []byte{
		0, 0, 0, 1, 9, 240, 0, 0, 1, 1, 159, 84, 106, 66, 127, 0, 0, 3, 0, 0, 3, 0, 0, 3, 0, 0, 3, 0, 0, 3, 0, 0, 3, 0, 0, 3, 0, 13, 211, 201, 87, 239, 30, 72, 80, 57, 71, 2, 178, 42, 159, 249, 170, 212, 137, 34, 128, 5, 212, 202, 23, 19, 163, 64, 113, 219, 124, 138, 224, 71, 82, 235, 114, 241, 177, 16, 97, 150, 49, 237, 3, 87, 170, 231, 80, 120, 134, 15, 146, 49, 93, 167, 3, 73, 78, 133, 236, 27, 208, 101, 191, 63, 173, 157, 237, 118, 35, 195, 218, 67, 88, 61, 21, 209, 190, 232, 112, 65, 39, 73, 205, 253, 254, 106, 123, 237, 227, 2, 171, 54, 223, 108, 150, 89, 149, 164, 44, 26, 46, 56, 247, 49, 242, 211, 127, 2, 11, 156, 93, 223, 128, 40, 79, 205, 115, 63, 218, 46, 188, 24, 130, 81, 60, 201, 34, 58, 127, 195, 212, 190, 106, 204, 166, 112, 6, 145, 102, 155, 88, 78, 86, 173, 169, 83, 53, 237, 234, 34, 188, 52, 142, 84, 139, 102, 176, 232, 24, 51, 38, 151, 111, 191, 247, 70, 153, 102, 234, 169, 56, 171, 204, 217, 51, 165, 212, 18, 64, 165, 169, 145, 196, 219, 192, 7, 248, 251, 155, 169, 77, 83, 180, 215, 149, 28, 228, 134, 86, 64, 64, 124, 60, 253, 198, 25, 232, 231, 44, 65, 117, 3, 235, 82, 161, 227, 106, 251, 184, 149, 236, 158, 194, 42, 217, 36, 59, 34, 211, 147, 107, 59, 56, 108, 79, 139, 152, 244, 64, 227, 131, 127, 37, 85, 204, 72, 139, 54, 47, 192, 129, 250, 162, 173, 193, 74, 149, 100, 234, 63, 142, 165, 224, 99, 106, 114, 139, 152, 189, 0, 114, 133, 173, 39, 17, 190, 111, 110, 132, 164, 121, 219, 206, 45, 242, 135, 92, 74, 75, 175, 169, 179, 71, 232, 94, 40, 129, 31, 110, 132, 203, 11, 193, 255, 212, 155, 31, 56, 203, 167, 102, 71, 64, 194, 99, 166, 19, 47, 140, 151, 62, 113, 195, 51, 135, 70, 46, 183, 141, 234, 157, 249, 231, 80, 19, 146, 97, 88, 144, 63, 220, 162, 85, 189, 192, 150, 49, 79, 173, 247, 179, 172, 93, 147, 148, 156, 205, 188, 7, 161, 242, 162, 132, 77, 65, 118, 108, 195, 177, 76, 210, 39, 192, 119, 241, 37, 51, 16, 121, 100, 240, 247, 42, 92, 64, 86, 128, 252, 5, 14, 75, 10, 255, 10, 242, 125, 188, 154, 43, 29, 202, 202, 211, 117, 137, 83, 8, 174, 203, 131, 244, 241, 128, 183, 84, 161, 186, 65, 111, 165, 8, 79, 158, 163, 67, 227, 110, 123, 241, 222, 177, 156, 55, 187, 249, 89, 119, 135, 11, 181, 0, 247, 141, 246, 41, 168, 213, 154, 213, 124, 250, 95, 74, 2, 35, 131, 243, 77, 45, 125, 176, 58, 120, 212, 50, 153, 195, 237, 126, 161, 195, 23, 167, 216, 210, 3, 21, 76, 179, 200, 90, 156, 37, 178, 27, 137, 116, 201, 193, 209, 2, 120, 138, 176, 241, 200, 207, 174, 163, 135, 79, 192, 71, 206, 217, 106, 155, 156, 36, 120, 5, 79, 59, 34, 245, 104, 193, 100, 63, 242, 92, 195, 75, 67, 6, 205, 43, 66, 35, 26, 205, 16, 128, 24, 11, 22, 22, 153, 137, 6, 123, 199, 208, 162, 25, 49, 104, 154, 25, 123, 18, 87, 187, 53, 163, 179, 222, 119, 193, 113, 149, 153, 143, 247, 112, 229, 13, 77, 162, 86, 160, 146, 152, 123, 137, 163, 87, 217, 39, 129, 66, 225, 28, 19, 191, 174, 252, 40, 54, 234, 37, 191, 70, 250, 74, 159, 173, 224, 152, 94, 171, 75, 183, 216, 235, 191, 149, 25, 39, 60, 64, 0, 0, 3, 0, 0, 3, 0, 0, 3, 0, 0, 3, 0, 0, 3, 0, 0, 35, 97}
	result := UnpackRTP2H264(i)
	fmt.Println(result)
}

func UnpackRTP2H264(rtpPayload []byte) []byte {
	if len(rtpPayload) <= 0 {
		return nil

	}

	var out []byte

	fu_indicator := rtpPayload[0] //获取第一个字节

	fu_header := rtpPayload[1] //获取第二个字节

	nalu_type := fu_indicator & 0x1f //获取FU indicator的类型域

	flag := fu_header & 0xe0 //获取FU header的前三位，判断当前是分包的开始、中间或结束

	nal_fua := ((fu_indicator & 0xe0) | (fu_header & 0x1f)) //FU_A nal

	var FrameType string

	if nal_fua == 0x67 {
		FrameType = "SPS"

	} else if nal_fua == 0x68 {
		FrameType = "PPS"

	} else if nal_fua == 0x65 {
		FrameType = "IDR"

	} else if nal_fua == 0x61 {
		FrameType = "P Frame"

	} else if nal_fua == 0x41 {
		FrameType = "P Frame"

	}

	log.Printf("nalu_type: %x flag: %x FrameType: %s", nalu_type, flag, FrameType)

	if nalu_type == 0x1c { //判断NAL的类型为0x1c=28，说明是FU-A分片

		if flag == 0x80 { //分片NAL单元开始位

			/*

			   o := make([]byte, len(rtpPayload)+5-2) //I帧开头可能为00 00 00 01、00 00 01，组帧时只用00 00 01开头

			   o[0] = 0x00

			   o[1] = 0x00

			   o[2] = 0x00

			   o[3] = 0x01

			   o[4] = nal_fua*/

			o := make([]byte, len(rtpPayload)+4-2) //I帧开头可能为00 00 00 01、00 00 01，组帧时只用00 00 01开头

			o[0] = 0x00

			o[1] = 0x00

			o[2] = 0x01

			o[3] = nal_fua

			copy(o[4:], rtpPayload[2:])

			out = o

		} else { //中间分片包或者最后一个分片包

			o := make([]byte, len(rtpPayload)-2)

			copy(o[0:], rtpPayload[2:])

			out = o

		}

	} else if nalu_type == 0x1 { //单一NAL 单元模式

		o := make([]byte, len(rtpPayload)+4) //将整个rtpPayload一起放进去

		o[0] = 0x00

		o[1] = 0x00

		o[2] = 0x00

		o[3] = 0x01

		copy(o[4:], rtpPayload[0:])

		out = o

	} else {
		log.Printf("Unsport nalu type!")

	}

	return out

}

/*

2022/10/02 19:39:44 nalu_type: 1c flag: 40 FrameType: P Frame
[219 42 82 28 156 159 200 28 207 110 25 142 135 119 53 172 9 179 252 200 3 235 203 47 36 197 93 192 94 23 92 60 167 77 159 84 215 237 204 122 100 0 168 206 237 75 215 90 162 248 202 93 55 154 227 35 103 182 91 180 60 52 216 121 177 231 112 81 7 243 246 210 231 3 140 200 181 138 131 246 218 11 138 198 165 19 155 145 103 106 155 182 59 58 41 45 68 99 122 206 13 184 35 238 123 214 200 54 17 184 90 43 87 17 9 183 163 176 173 253 173 122 147 65 45 84 229 228 226 153 184 11 14 17 84 163 72 40 75 216 12 60 244 7 204 68 212 144 36 60 237 110 183 102 218 156 222 124 1 246 75 68 248 187 7 22 169 41 253 100 206 220 241 132 148 135 21 48 82 156 233 225 82 134 101 62 134 24 169 235 52 63 203 184 183 71 89 254 84 37 233 167 206 31 219 196 125 199 161 207 117 83 223 144 130 242 166 169 169 86 148 6 233 131 88 40 26 119 128 125 140 33 129 245 241 212 143 104 199 229 74 250 31 178 62 246 123 42 79 129 185 23 17 249 150 169 155 227 188 134 218 249 133 102 239 55 225 233 201 25 79 119 34 133 161 220 196 147 153 150 48 51 242 119 82 250 12 46 34 41 132 8 219 51 124 102 102 227 216 66 27 183 95 194 96 89 177 127 32 190 155 10 188 255 32 225 47 176 170 74 37 221 234 203 89 152 176 200 75 27 124 130 212 25 187 158 66 240 151 110 147 251 64 132 140 151 167 230 252 5 185 69 234 216 125 107 49 123 65 39 127 218 194 62 170 3 136 134 207 215 10 255 7 6 76 66 66 129 97 77 188 123 251 247 158 182 217 155 252 248 111 109 47 194 68 28 52 21 56 61 113 153 102 136 215 193 178 193 57 171 52 168 238 21 15 6 239 112 130 91 162 20 211 30 175 205 121 145 121 2 153 59 157 222 1 68 221 51 148 134 157 164 15 157 1 161 17 172 143 66 244 203 109 131 17 241 5 65 227 241 215 24 99 105 145 216 142 234 38 123 226 180 241 245 252 102 232 214 113 12 252 4 137 65 220 186 221 103 218 137 139 169 211 54 11 117 199 191 218 56 67 247 208 185 37 39 255 27 243 32 125 177 197 146 59 9 191 253 14 175 65 149 105 154 210 147 186 69 43 226 123 116 143 67 192 67 229 184 10 97 46 135 173 162 192 240 114 13 185 0 3 65 27 38 137 70 209 97 180 207 1 31 147 113 6 34 205 254 80 47 217 96 97 87 240 136 114 228 196 164 216 178 13 127 195 207 204 131 181 81 206 1 145 120 222 14 151 232 170 190 54 48 54 110 19 90 176 114 137 199 63 153 69 245 102 177 111 36 120 116 142 240 218 3 242 44 86 53 2 87 203 193 202 18 41 39 170 222 51 65 76 209 158 173 213 141 11 119 157 21 135 177 236 101 238 164 237 209 252 220 101 78 108 97 68 16 25 81 196 116 37 243 110 153 113 209 54 222 176 222 182 227 199 38 241 42 159 23 191 105 106 249 7 245 22 156 41 166 153 241 25 161 165 27 79 198 221 175 194 18 72 186 76 85 203 58 108 57 116 200 215 94 43 198 139 223 184 74 210 144 207 105 165 43 96 79 121 46 93 0 243 48 253 131 87 21 80 233 252 101 168 157 130 119 245 146 136 0 123 159 42 45 163 130 167 244 229 166 125 166 190 84 12 237 143 121 222 37 205 221 74 247 91 77 218 57 202 103 201 159 63 37 202 231 172 54 209 64 67 218 154 105 9 60 207 77 240 28 232 160 183 110 253 190 12 42 237 39 50 2 87 221 33 74 11 77 243 26 164 73 136 113 45 40 4 207 32 143 41 87 38 245 192 28 109 135 253 245 57 140 116 156 13 126 184 155 56 67 203 38 146 190 44 219 58 57 241 115 69 254 127 199 125 104 99 210 36 248 227 208 132 75 54 80 57 223 62 237 32 82 143 22 39 107 110 252 144 68 199 50 166 209 227 195 69 36 27 112 165 80 177 27 74 165 119 182 241 160 234 25 152 39 88 184 124 171 66 125 88 245 97 180 102 111 24 236 241 219 31 182 251 244 254 126 109 130 20 0 223 228 170 59 141 72 99 193 38 234 120 84 224 249 135 252 13 137 252 102 140 225 85 16 26 72 213 100 11 27 188 22 234 138 8 62 175 255 3 116 2 96 147 96 125 167 133 44 255 250 71 23 19 188 216 86 93 61 78 231 255 209 2 246 190 129 3 105 221 5 118 56 104 138 237 64 245 15 20 102 215 100 88 138 53 195 200 240 218 155 203 229 88 56 104 116 112 152 176 169 84 238 192 64]

[219,42,82,28,156,159,200,28,207,110,25,142,135,119,53,172,9,179,252,200,3,235,203,47,36,197,93,192,94,23,92,60,167,77,159,84,215,237,204,122,100,0,168,206,237,75,215,90,162,248,202,93,55,154,227,35,103,182,91,180,60,52,216,121,177,231,112,81,7,243,246,210,231,3,140,200,181,138,131,246,218,11,138,198,165,19,155,145,103,106,155,182,59,58,41,45,68,99,122,206,13,184,35,238,123,214,200,54,17,184,90,43,87,17,9,183,163,176,173,253,173,122,147,65,45,84,229,228,226,153,184,11,14,17,84,163,72,40,75,216,12,60,244,7,204,68,212,144,36,60,237,110,183,102,218,156,222,124,1,246,75,68,248,187,7,22,169,41,253,100,206,220,241,132,148,135,21,48,82,156,233,225,82,134,101,62,134,24,169,235,52,63,203,184,183,71,89,254,84,37,233,167,206,31,219,196,125,199,161,207,117,83,223,144,130,242,166,169,169,86,148,6,233,131,88,40,26,119,128,125,140,33,129,245,241,212,143,104,199,229,74,250,31,178,62,246,123,42,79,129,185,23,17,249,150,169,155,227,188,134,218,249,133,102,239,55,225,233,201,25,79,119,34,133,161,220,196,147,153,150,48,51,242,119,82,250,12,46,34,41,132,8,219,51,124,102,102,227,216,66,27,183,95,194,96,89,177,127,32,190,155,10,188,255,32,225,47,176,170,74,37,221,234,203,89,152,176,200,75,27,124,130,212,25,187,158,66,240,151,110,147,251,64,132,140,151,167,230,252,5,185,69,234,216,125,107,49,123,65,39,127,218,194,62,170,3,136,134,207,215,10,255,7,6,76,66,66,129,97,77,188,123,251,247,158,182,217,155,252,248,111,109,47,194,68,28,52,21,56,61,113,153,102,136,215,193,178,193,57,171,52,168,238,21,15,6,239,112,130,91,162,20,211,30,175,205,121,145,121,2,153,59,157,222,1,68,221,51,148,134,157,164,15,157,1,161,17,172,143,66,244,203,109,131,17,241,5,65,227,241,215,24,99,105,145,216,142,234,38,123,226,180,241,245,252,102,232,214,113,12,252,4,137,65,220,186,221,103,218,137,139,169,211,54,11,117,199,191,218,56,67,247,208,185,37,39,255,27,243,32,125,177,197,146,59,9,191,253,14,175,65,149,105,154,210,147,186,69,43,226,123,116,143,67,192,67,229,184,10,97,46,135,173,162,192,240,114,13,185,0,3,65,27,38,137,70,209,97,180,207,1,31,147,113,6,34,205,254,80,47,217,96,97,87,240,136,114,228,196,164,216,178,13,127,195,207,204,131,181,81,206,1,145,120,222,14,151,232,170,190,54,48,54,110,19,90,176,114,137,199,63,153,69,245,102,177,111,36,120,116,142,240,218,3,242,44,86,53,2,87,203,193,202,18,41,39,170,222,51,65,76,209,158,173,213,141,11,119,157,21,135,177,236,101,238,164,237,209,252,220,101,78,108,97,68,16,25,81,196,116,37,243,110,153,113,209,54,222,176,222,182,227,199,38,241,42,159,23,191,105,106,249,7,245,22,156,41,166,153,241,25,161,165,27,79,198,221,175,194,18,72,186,76,85,203,58,108,57,116,200,215,94,43,198,139,223,184,74,210,144,207,105,165,43,96,79,121,46,93,0,243,48,253,131,87,21,80,233,252,101,168,157,130,119,245,146,136,0,123,159,42,45,163,130,167,244,229,166,125,166,190,84,12,237,143,121,222,37,205,221,74,247,91,77,218,57,202,103,201,159,63,37,202,231,172,54,209,64,67,218,154,105,9,60,207,77,240,28,232,160,183,110,253,190,12,42,237,39,50,2,87,221,33,74,11,77,243,26,164,73,136,113,45,40,4,207,32,143,41,87,38,245,192,28,109,135,253,245,57,140,116,156,13,126,184,155,56,67,203,38,146,190,44,219,58,57,241,115,69,254,127,199,125,104,99,210,36,248,227,208,132,75,54,80,57,223,62,237,32,82,143,22,39,107,110,252,144,68,199,50,166,209,227,195,69,36,27,112,165,80,177,27,74,165,119,182,241,160,234,25,152,39,88,184,124,171,66,125,88,245,97,180,102,111,24,236,241,219,31,182,251,244,254,126,109,130,20,0,223,228,170,59,141,72,99,193,38,234,120,84,224,249,135,252,13,137,252,102,140,225,85,16,26,72,213,100,11,27,188,22,234,138,8,62,175,255,3,116,2,96,147,96,125,167,133,44,255,250,71,23,19,188,216,86,93,61,78,231,255,209,2,246,190,129,3,105,221,5,118,56,104,138,237,64,245,15,20,102,215,100,88,138,53,195,200,240,218,155,203,229,88,56,104,116,112,152,176,169,84,238,192,64]

*/

/*
2022/10/02 19:53:29 nalu_type: 1c flag: 40 FrameType: P Frame
[218 249 125 29 225 4 186 89 234 110 138 187 62 129 31 24 77 200 239 158 114 60 192 90 136 21 153 2 67 248 64 42 134 122 147 135 187 50 17 185 205 88 195 176 3 145 171 211 254 131 113 203 250 73 68 163 208 88 27 221 6 137 119 14 108 51 10 16 5 35 103 80 175 128 55 247 220 51 123 228 101 48 102 84 107 192 202 86 215 180 1 170 107 153 77 101 141 89 245 233 159 142 219 2 69 114 74 86 94 187 80 189 59 86 11 6 3 80 119 10 54 61 94 10 0 0 3 0 0 3 0 0 3 0 13 232]

[218,249,125,29,225,4,186,89,234,110,138,187,62,129,31,24,77,200,239,158,114,60,192,90,136,21,153,2,67,248,64,42,134,122,147,135,187,50,17,185,205,88,195,176,3,145,171,211,254,131,113,203,250,73,68,163,208,88,27,221,6,137,119,14,108,51,10,16,5,35,103,80,175,128,55,247,220,51,123,228,101,48,102,84,107,192,202,86,215,180,1,170,107,153,77,101,141,89,245,233,159,142,219,2,69,114,74,86,94,187,80,189,59,86,11,6,3,80,119,10,54,61,94,10,0,0,3,0,0,3,0,0,3,0,13,232]

*/

/*
2022/10/02 19:55:14 nalu_type: 1c flag: 80 FrameType: P Frame
[0 0 1 65 155 92 70 5 23 13 255 254 56 64 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 51 108 177 33 56 158 71 66 192 159 181 69 154 239 37 78 104 223 200 234 81 192 216 175 251 181 187 99 29 49 133 42 172 183 103 252 49 68 244 92 126 196 245 25 21 126 104 74 110 147 236 99 75 205 210 125 186 170 159 80 72 146 168 134 229 212 159 10 252 120 142 180 254 71 101 28 124 179 240 20 198 51 211 29 55 4 148 199 216 126 21 155 72 4 8 67 160 125 5 129 245 132 3 3 104 231 187 168 35 96 61 174 80 182 150 222 215 194 46 188 120 191 100 7 86 27 188 18 118 208 161 141 227 137 138 230 186 134 59 105 77 78 159 222 78 223 62 136 146 52 213 71 19 5 158 219 160 93 217 246 189 65 244 51 36 34 46 136 199 151 157 245 235 231 14 196 84 147 249 229 209 137 255 208 145 24 187 93 195 200 130 96 214 141 195 201 228 10 107 100 19 42 181 243 229 127 74 6 40 249 125 34 166 24 37 213 238 169 122 81 163 180 197 49 107 238 101 92 203 2 105 16 225 114 116 149 212 126 69 147 137 96 148 176 145 120 252 233 125 111 25 242 216 175 1 107 202 224 109 163 161 113 242 2 237 105 225 100 195 173 213 138 183 58 157 0 58 214 174 82 178 10 83 72 110 225 159 78 159 161 110 154 243 197 124 89 68 0 13 114 221 226 66 93 143 188 65 26 9 164 184 4 48 144 244 231 138 88 29 98 193 77 236 235 149 9 103 9 255 152 54 1 83 249 9 83 55 134 144 31 159 244 235 188 112 185 51 184 156 224 76 147 12 21 146 51 234 217 145 15 21 136 14 16 61 141 4 51 221 139 199 51 41 93 114 197 94 214 26 77 44 211 140 205 33 71 194 84 68 164 244 173 36 89 204 181 7 58 93 209 227 98 211 41 51 235 124 134 97 221 73 206 163 247 32 197 217 34 179 132 63 161 250 32 118 105 105 93 43 230 228 213 113 140 18 221 221 60 198 154 134 4 20 51 23 41 108 239 207 183 52 189 120 128 103 219 137 255 182 183 102 231 167 223 217 74 30 73 250 162 165 106 65 0 208 135 118 117 65 160 240 243 33 63 190 243 194 143 174 175 196 98 249 157 121 85 88 153 197 88 119 145 220 88 0 10 2 112 80 2 198 206 197 72 185 102 21 197 30 202 164 150 183 102 138 206 184 165 51 95 125 245 203 200 207 68 166 160 237 133 199 211 144 221 167 158 238 156 19 201 14 125 201 12 250 181 6 70 58 81 76 1 220 184 205 100 178 143 153 230 216 47 7 124 170 76 37 84 161 81 46 186 115 230 207 1 201 87 55 121 185 66 213 199 202 142 151 101 127 130 54 27 35 61 91 249 50 108 22 129 107 166 229 16 11 156 253 205 65 71 15 244 239 58 195 171 218 64 65 132 223 9 191 154 237 84 16 96 46 248 124 162 133 45 241 71 77 12 71 66 131 135 2 20 252 255 240 245 245 181 21 139 150 210 212 107 202 112 221 71 144 90 195 50 149 63 245 105 1 57 21 208 233 28 51 92 43 6 205 71 9 79 7 46 60 160 206 65 60 19 230 185 149 105 193 2 236 93 120 68 204 127 91 250 159 232 159 152 237 231 71 124 120 202 137 129 124 156 241 251 184 0 89 19 247 195 145 193 66 199 50 76 223 2 199 82 109 80 76 16 197 175 203 170 100 126 211 198 73 248 202 81 1 29 16 181 250 197 85 191 147 160 51 55 201 21 39 238 102 170 45 238 4 220 33 108 103 2 157 212 156 81 75 113 91 117 196 123 148 124 250 145 126 59 248 31 186 54 251 149 29 39 146 148 82 72 46 228 157 25 180 78 141 60 239 7 234 103 163 251 139 218 20 224 30 41 203 231 131 149 211 133 56 164 211 74 116 89 62 101 204 3 14 109 28 60 244 156 177 13 98 227 246 147 116 234 247 163 59 196 82 142 97 198 120 21 77 23 73 247 223 170 206 21 189 6 241 230 92 55 180 187 220 44 198 45 87 66 14 229 201 63 23 49 140 5 215 172 90 87 100 203 248 55 251 207 237 211 59 23 151 178 24 238 3 155 193 247 115 28 103 144 90 170 230 116 21 160 114 127 223 213 28 168 42 8 150 72 26 123 110 73 3 22 70 34 161 163 11 124 236 188 81 131 0 59 121 230 118 73 17 215 231 242 149 1 189 114 252 27 223 197 69 215 60 217 152 137 187 11 25 212 65 91 229 44 86 125 242 176 176 185 227 88 76 122 243 234 115 138 124 66 31 48 84 183 143 61 56 129 158 187 144 236 170 147 166 122 29 39 79 198 193 73 251 146 81 111 115 31 71 110 135 171 101 165 53 101 190 53 123 175 102 198 143 178 133 8 162 175 20 170 20 140 223 150 43 247 201 154 204 105 126 180 176 10 254 128 161 193 44 145 104 134 253 158 83 250 239 165 127 229 169 84 8 187 225 145 179 46 104 163 151 63 162 111 114 238 171 123 63 50 167 201 121 73 159 110 2 227 212 145 137 255 151 42 26 0 222 91 236 141 234 174 252 132 104 76 105 5 140 36 63 124 1 63 27 167 61 70 49 200 103 215 57 215 153 95 2 218 60 60 130 27 7 69 6 165 104 175 140 207 12 98 101 183 89 25 148 101 32 177 93 5 151 193 4 61 52 83 44 114 225 6 111 193 57 246 45 167 172 34 255 81 156 132 118 44 45 77 115 160 251 175 171 75 58 145 178 205 192 185 247 36 7 221 13 59 238 73 140 37 92 172 94 208 124 222 45 87 97 77 168 177 161 111 32 67 161 231 112 253 17 245 201 0 63 141 80 37 228 68 236 236 3 27 50 76 10 207 60 160 58 240 112 215 183 81 23 79 170 204 156 167 4 233 30 117 120 36 88 149 153 152 38 174 127 9 62 87 134 222 101 22 125 188 163 33 48 206 247 138 226 216 86 150 190 140 82 159 35 141 171 110 83 158 16 61 66 136 42 107 168 211 67 176 123 61 48 235 126 153 223 23 125 113 196 37 174 13 242 224 77 148 123 58 254 196 181 29 182 51 184 127 190 229 8 187 67 198 78 217 43 146 240 241 214 55 130 63 215 209 104 203 8 253 13 198 82 128 203 242 76 164 74 102 207 76 175 212 206 172 157 186 137 50 56 60 53 165 197 130 114 123 113 101 225 164 46 86 133 154 20 158 250 152 64 144 230]

[0,0,1,65,155,92,70,5,23,13,255,254,56,64,0,0,3,0,0,3,0,0,3,0,0,3,0,0,3,0,0,3,0,51,108,177,33,56,158,71,66,192,159,181,69,154,239,37,78,104,223,200,234,81,192,216,175,251,181,187,99,29,49,133,42,172,183,103,252,49,68,244,92,126,196,245,25,21,126,104,74,110,147,236,99,75,205,210,125,186,170,159,80,72,146,168,134,229,212,159,10,252,120,142,180,254,71,101,28,124,179,240,20,198,51,211,29,55,4,148,199,216,126,21,155,72,4,8,67,160,125,5,129,245,132,3,3,104,231,187,168,35,96,61,174,80,182,150,222,215,194,46,188,120,191,100,7,86,27,188,18,118,208,161,141,227,137,138,230,186,134,59,105,77,78,159,222,78,223,62,136,146,52,213,71,19,5,158,219,160,93,217,246,189,65,244,51,36,34,46,136,199,151,157,245,235,231,14,196,84,147,249,229,209,137,255,208,145,24,187,93,195,200,130,96,214,141,195,201,228,10,107,100,19,42,181,243,229,127,74,6,40,249,125,34,166,24,37,213,238,169,122,81,163,180,197,49,107,238,101,92,203,2,105,16,225,114,116,149,212,126,69,147,137,96,148,176,145,120,252,233,125,111,25,242,216,175,1,107,202,224,109,163,161,113,242,2,237,105,225,100,195,173,213,138,183,58,157,0,58,214,174,82,178,10,83,72,110,225,159,78,159,161,110,154,243,197,124,89,68,0,13,114,221,226,66,93,143,188,65,26,9,164,184,4,48,144,244,231,138,88,29,98,193,77,236,235,149,9,103,9,255,152,54,1,83,249,9,83,55,134,144,31,159,244,235,188,112,185,51,184,156,224,76,147,12,21,146,51,234,217,145,15,21,136,14,16,61,141,4,51,221,139,199,51,41,93,114,197,94,214,26,77,44,211,140,205,33,71,194,84,68,164,244,173,36,89,204,181,7,58,93,209,227,98,211,41,51,235,124,134,97,221,73,206,163,247,32,197,217,34,179,132,63,161,250,32,118,105,105,93,43,230,228,213,113,140,18,221,221,60,198,154,134,4,20,51,23,41,108,239,207,183,52,189,120,128,103,219,137,255,182,183,102,231,167,223,217,74,30,73,250,162,165,106,65,0,208,135,118,117,65,160,240,243,33,63,190,243,194,143,174,175,196,98,249,157,121,85,88,153,197,88,119,145,220,88,0,10,2,112,80,2,198,206,197,72,185,102,21,197,30,202,164,150,183,102,138,206,184,165,51,95,125,245,203,200,207,68,166,160,237,133,199,211,144,221,167,158,238,156,19,201,14,125,201,12,250,181,6,70,58,81,76,1,220,184,205,100,178,143,153,230,216,47,7,124,170,76,37,84,161,81,46,186,115,230,207,1,201,87,55,121,185,66,213,199,202,142,151,101,127,130,54,27,35,61,91,249,50,108,22,129,107,166,229,16,11,156,253,205,65,71,15,244,239,58,195,171,218,64,65,132,223,9,191,154,237,84,16,96,46,248,124,162,133,45,241,71,77,12,71,66,131,135,2,20,252,255,240,245,245,181,21,139,150,210,212,107,202,112,221,71,144,90,195,50,149,63,245,105,1,57,21,208,233,28,51,92,43,6,205,71,9,79,7,46,60,160,206,65,60,19,230,185,149,105,193,2,236,93,120,68,204,127,91,250,159,232,159,152,237,231,71,124,120,202,137,129,124,156,241,251,184,0,89,19,247,195,145,193,66,199,50,76,223,2,199,82,109,80,76,16,197,175,203,170,100,126,211,198,73,248,202,81,1,29,16,181,250,197,85,191,147,160,51,55,201,21,39,238,102,170,45,238,4,220,33,108,103,2,157,212,156,81,75,113,91,117,196,123,148,124,250,145,126,59,248,31,186,54,251,149,29,39,146,148,82,72,46,228,157,25,180,78,141,60,239,7,234,103,163,251,139,218,20,224,30,41,203,231,131,149,211,133,56,164,211,74,116,89,62,101,204,3,14,109,28,60,244,156,177,13,98,227,246,147,116,234,247,163,59,196,82,142,97,198,120,21,77,23,73,247,223,170,206,21,189,6,241,230,92,55,180,187,220,44,198,45,87,66,14,229,201,63,23,49,140,5,215,172,90,87,100,203,248,55,251,207,237,211,59,23,151,178,24,238,3,155,193,247,115,28,103,144,90,170,230,116,21,160,114,127,223,213,28,168,42,8,150,72,26,123,110,73,3,22,70,34,161,163,11,124,236,188,81,131,0,59,121,230,118,73,17,215,231,242,149,1,189,114,252,27,223,197,69,215,60,217,152,137,187,11,25,212,65,91,229,44,86,125,242,176,176,185,227,88,76,122,243,234,115,138,124,66,31,48,84,183,143,61,56,129,158,187,144,236,170,147,166,122,29,39,79,198,193,73,251,146,81,111,115,31,71,110,135,171,101,165,53,101,190,53,123,175,102,198,143,178,133,8,162,175,20,170,20,140,223,150,43,247,201,154,204,105,126,180,176,10,254,128,161,193,44,145,104,134,253,158,83,250,239,165,127,229,169,84,8,187,225,145,179,46,104,163,151,63,162,111,114,238,171,123,63,50,167,201,121,73,159,110,2,227,212,145,137,255,151,42,26,0,222,91,236,141,234,174,252,132,104,76,105,5,140,36,63,124,1,63,27,167,61,70,49,200,103,215,57,215,153,95,2,218,60,60,130,27,7,69,6,165,104,175,140,207,12,98,101,183,89,25,148,101,32,177,93,5,151,193,4,61,52,83,44,114,225,6,111,193,57,246,45,167,172,34,255,81,156,132,118,44,45,77,115,160,251,175,171,75,58,145,178,205,192,185,247,36,7,221,13,59,238,73,140,37,92,172,94,208,124,222,45,87,97,77,168,177,161,111,32,67,161,231,112,253,17,245,201,0,63,141,80,37,228,68,236,236,3,27,50,76,10,207,60,160,58,240,112,215,183,81,23,79,170,204,156,167,4,233,30,117,120,36,88,149,153,152,38,174,127,9,62,87,134,222,101,22,125,188,163,33,48,206,247,138,226,216,86,150,190,140,82,159,35,141,171,110,83,158,16,61,66,136,42,107,168,211,67,176,123,61,48,235,126,153,223,23,125,113,196,37,174,13,242,224,77,148,123,58,254,196,181,29,182,51,184,127,190,229,8,187,67,198,78,217,43,146,240,241,214,55,130,63,215,209,104,203,8,253,13,198,82,128,203,242,76,164,74,102,207,76,175,212,206,172,157,186,137,50,56,60,53,165,197,130,114,123,113,101,225,164,46,86,133,154,20,158,250,152,64,144,230]

*/

/*
124 69 187 58 221 255 194 214 31 190 210 81 207 5 60 38 213 37 157 50 188 250 108 140 254 33 85 132 154 175 220 93 177 114 141 236 81 50 149 203 231 101 82 42 96 237 176 121 207 54 212 53 230 45 146 237 81 130 45 243 190 174 189 163 129 19 151 3 81 127 53 142 243 247 120 230 198 16 186 93 189 183 141 133 217 111 174 8 38 21 250 76 172 77 173 41 239 129 4 179 188 129 143 2 187 185 202 191 123 178 224 235 106 137 12 17 1 77 5 162 22 67 140 92 58 235 118 9 48 83 78 39 162 6 14 14 35 64 255 25 244 99 250 1 214 99 60 177 47 49 232 121 25 29 141 42 134 23 149 187 190 95 69 50 124 48 26 146 198 101 55 193 191 251 100 202 112 92 138 246 20 14 54 186 165 134 87 253 107 153 144 246 217 172 158 89 144 193 26 56 92 70 118 118 232 188 24 71 86 192 7 149 238 9 159 148 180 51 76 241 140 172 20 120 246 9 224 102 182 11 128 103 44 145 245 253 89 143 166 6 220 120 254 133 32 37 147 212 169 131 204 62 161 154 47 164 39 49 177 88 187 141 64 199 98 60 153 83 175 217 46 216 77 68 56 16 129 187 42 205 11 136 110 77 207 211 163 201 105 26 157 126 164 188 143 82 31 99 81 183 93 135 206 231 77 158 234 158 233 182 237 231 91 184 19 84 250 11 3 128 12 64 183 187 187 66 206 97 22 40 155 12 31 235 142 242 104 121 232 91 103 174 43 75 203 91 1 237 139 133 54 40 235 254 90 0 58 132 67 66 153 14 0 192 80 5 114 13 25 72 63 146 18 89 172 252 105 153 9 108 255 220 116 118 56 65 195 226 216 245 156 3 224 85 249 117 195 129 145 232 130 254 16 10 122 241 199 137 167 30 242 32 49 134 75 132 40 127 249 247 182 73 232 191 19 199 225 98 208 176 93 73 165 192 181 187 218 9 119 206 214 221 210 112 61 124 93 133 128 199 80 89 209 17 49 38 53 51 192 202 56 128 36 101 125 244 77 237 173 107 8 245 172 136 0 84 82 223 2 92 63 32 86 85 86 183 33 243 204 171 207 86 171 226 146 41 51 176 72 225 24 176 179 147 122 134 54 151 13 29 23 185 114 186 59 49 58 61 25 219 134 223 234 40 120 224 172 59 138 213 7 150 10 114 225 229 193 145 206 89 205 43 9 91 98 254 247 136 86 126 64 1 201 145 253 32 73 227 128 143 83 2 57 117 49 230 71 106 68 77 55 170 165 91 224 37 57 149 217 109 73 107 31 103 201 37 56 190 186 129 234 89 55 240 129 37 187 64 32 247 76 247 47 66 67 71 68 211 21 225 210 35 171 226 39 26 76 198 197 250 219 171 125 7 131 5 195 31 6 208 97 243 142 200 142 169 176 169 192 96 2 254 30 203 62 5 235 240 61 1 96 190 213 101 99 71 86 79 72 9 16 32 200 187 34 75 170 110 142 2 203 121 148 29 212 52 82 71 128 240 76 70 121 82 230 234 205 111 192 113 96 191 98 102 83 3 162 9 51 78 118 72 69 90 217 91 112 113 27 139 126 169 224 2 155 165 45 192 176 166 195 110 99 91 35 17 193 211 69 247 100 138 144 170 215 255 44 100 16 142 90 168 7 0 34 158 9 157 20 198 134 114 153 218 153 5 237 56 102 95 158 144 243 255 150 75 47 154 211 101 162 157 66 184 230 109 214 246 172 51 131 122 48 73 90 32 26 253 107 83 105 12 64 122 42 203 112 143 3 148 170 128 0 14 231 189 135 247 67 197 209 243 252 12 203 54 18 187 193 53 169 219 216 234 96 201 169 141 227 9 176 136 252 188 104 11 57 17 193 81 78 220 247 10 249 222 205 165 84 64 44 18 230 159 90 254 245 83 22 17 162 79 153 97 235 50 184 27 181 103 177 193 251 156 62 208 205 124 195 131 158 159 250 208 123 196 63 169 200 184 127 135 149 130 61 185 205 15 113 130 77 198 121 66 94 157 111 232 98 57 146 155 17 26 203 174 104 27 47 82 78 140 70 190 0 142 33 88 238 152 245 34 242 37 218 158 159 128 249 114 76 65 99 113 52 37 214 139 145 153 49 186 210 154 161 109 161 216 201 41 247 39 53 64 111 104 251 129 242 192 208 99 230 230 28 208 109 195 252 114 160 142 92 124 145 253 51 125 237 194 195 168 236 85 254 84 135 106 57 110 220 227 7 77 237 30 101 29 254 154 228 122 187 80 22 1 224 0 0 18 218 165 52 228 128 0 239 213 191 0 123 158 0 0 15 78 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 19 48]

2022/10/03 14:16:21 nalu_type: 1c flag: 40 FrameType: IDR
[187 58 221 255 194 214 31 190 210 81 207 5 60 38 213 37 157 50 188 250 108 140 254 33 85 132 154 175 220 93 177 114 141 236 81 50 149 203 231 101 82 42 96 237 176 121 207 54 212 53 230 45 146 237 81 130 45 243 190 174 189 163 129 19 151 3 81 127 53 142 243 247 120 230 198 16 186 93 189 183 141 133 217 111 174 8 38 21 250 76 172 77 173 41 239 129 4 179 188 129 143 2 187 185 202 191 123 178 224 235 106 137 12 17 1 77 5 162 22 67 140 92 58 235 118 9 48 83 78 39 162 6 14 14 35 64 255 25 244 99 250 1 214 99 60 177 47 49 232 121 25 29 141 42 134 23 149 187 190 95 69 50 124 48 26 146 198 101 55 193 191 251 100 202 112 92 138 246 20 14 54 186 165 134 87 253 107 153 144 246 217 172 158 89 144 193 26 56 92 70 118 118 232 188 24 71 86 192 7 149 238 9 159 148 180 51 76 241 140 172 20 120 246 9 224 102 182 11 128 103 44 145 245 253 89 143 166 6 220 120 254 133 32 37 147 212 169 131 204 62 161 154 47 164 39 49 177 88 187 141 64 199 98 60 153 83 175 217 46 216 77 68 56 16 129 187 42 205 11 136 110 77 207 211 163 201 105 26 157 126 164 188 143 82 31 99 81 183 93 135 206 231 77 158 234 158 233 182 237 231 91 184 19 84 250 11 3 128 12 64 183 187 187 66 206 97 22 40 155 12 31 235 142 242 104 121 232 91 103 174 43 75 203 91 1 237 139 133 54 40 235 254 90 0 58 132 67 66 153 14 0 192 80 5 114 13 25 72 63 146 18 89 172 252 105 153 9 108 255 220 116 118 56 65 195 226 216 245 156 3 224 85 249 117 195 129 145 232 130 254 16 10 122 241 199 137 167 30 242 32 49 134 75 132 40 127 249 247 182 73 232 191 19 199 225 98 208 176 93 73 165 192 181 187 218 9 119 206 214 221 210 112 61 124 93 133 128 199 80 89 209 17 49 38 53 51 192 202 56 128 36 101 125 244 77 237 173 107 8 245 172 136 0 84 82 223 2 92 63 32 86 85 86 183 33 243 204 171 207 86 171 226 146 41 51 176 72 225 24 176 179 147 122 134 54 151 13 29 23 185 114 186 59 49 58 61 25 219 134 223 234 40 120 224 172 59 138 213 7 150 10 114 225 229 193 145 206 89 205 43 9 91 98 254 247 136 86 126 64 1 201 145 253 32 73 227 128 143 83 2 57 117 49 230 71 106 68 77 55 170 165 91 224 37 57 149 217 109 73 107 31 103 201 37 56 190 186 129 234 89 55 240 129 37 187 64 32 247 76 247 47 66 67 71 68 211 21 225 210 35 171 226 39 26 76 198 197 250 219 171 125 7 131 5 195 31 6 208 97 243 142 200 142 169 176 169 192 96 2 254 30 203 62 5 235 240 61 1 96 190 213 101 99 71 86 79 72 9 16 32 200 187 34 75 170 110 142 2 203 121 148 29 212 52 82 71 128 240 76 70 121 82 230 234 205 111 192 113 96 191 98 102 83 3 162 9 51 78 118 72 69 90 217 91 112 113 27 139 126 169 224 2 155 165 45 192 176 166 195 110 99 91 35 17 193 211 69 247 100 138 144 170 215 255 44 100 16 142 90 168 7 0 34 158 9 157 20 198 134 114 153 218 153 5 237 56 102 95 158 144 243 255 150 75 47 154 211 101 162 157 66 184 230 109 214 246 172 51 131 122 48 73 90 32 26 253 107 83 105 12 64 122 42 203 112 143 3 148 170 128 0 14 231 189 135 247 67 197 209 243 252 12 203 54 18 187 193 53 169 219 216 234 96 201 169 141 227 9 176 136 252 188 104 11 57 17 193 81 78 220 247 10 249 222 205 165 84 64 44 18 230 159 90 254 245 83 22 17 162 79 153 97 235 50 184 27 181 103 177 193 251 156 62 208 205 124 195 131 158 159 250 208 123 196 63 169 200 184 127 135 149 130 61 185 205 15 113 130 77 198 121 66 94 157 111 232 98 57 146 155 17 26 203 174 104 27 47 82 78 140 70 190 0 142 33 88 238 152 245 34 242 37 218 158 159 128 249 114 76 65 99 113 52 37 214 139 145 153 49 186 210 154 161 109 161 216 201 41 247 39 53 64 111 104 251 129 242 192 208 99 230 230 28 208 109 195 252 114 160 142 92 124 145 253 51 125 237 194 195 168 236 85 254 84 135 106 57 110 220 227 7 77 237 30 101 29 254 154 228 122 187 80 22 1 224 0 0 18 218 165 52 228 128 0 239 213 191 0 123 158 0 0 15 78 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 3 0 0 19 48]

*/
