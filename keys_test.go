package librespot

// import(
// 	"testing"
// 	"math/big"
// 	"bytes"
// 	"fmt"
// )

// func TestGenerateKets(t *testing.T) {
// 	GenerateKeys()
// }

// func TestGeneratePublicKey(t *testing.T) {
// 	private := new(big.Int)
// 	private.SetBytes([]byte{
// 		247, 108, 193, 86, 72, 215, 221, 24, 155, 188, 171, 177, 222, 44, 205, 54, 197, 134, 231, 197, 2, 232, 215, 217, 0, 147, 244, 178, 194, 6, 4, 136, 3, 203, 95, 48, 5, 19, 216, 17, 55, 156, 107, 133, 97, 243, 141, 51, 8, 22, 33, 96, 129, 96, 72, 214, 45, 131, 144, 25, 216, 107, 131, 35, 148, 29, 64, 166, 19, 15, 222, 193, 125, 49, 105, 192, 240, 39, 125, 176, 90, 5, 242, 128, 248, 165, 146, 71, 181, 124, 239, 245, 10, 161, 237,
// 		})

// 	pubSlice := []byte{
// 		196, 29, 15, 134, 227, 241, 34, 138, 35, 50, 183, 115, 248, 124, 11, 21, 118, 167, 59, 55, 177, 21, 178, 220, 167, 219, 174, 145, 188, 33, 35, 187, 94, 19, 251, 229, 112, 139, 137, 208, 46, 52, 96, 35, 127, 82, 29, 183, 87, 210, 172, 87, 136, 166, 130, 211, 71, 17, 82, 211, 202, 55, 117, 30, 137, 96, 154, 49, 43, 188, 20, 20, 78, 107, 45, 93, 161, 206, 234, 167, 152, 41, 1, 195, 51, 118, 239, 201, 204, 228, 15, 208, 145, 82, 213, 11,
// 	}
// 	keys := generateKeysFromPrivate(private)
// 	if !bytes.Equal(keys.publicKey.Bytes(), pubSlice) {
// 		t.Errorf("result does not match, %v %v", keys.publicKey.Bytes(), pubSlice)
// 	}
// }

// func TestAddRemote(t *testing.T) {
// 	remote := []byte{
// 		80, 231, 73, 121, 44, 224, 167, 121, 81, 184, 143, 208, 223, 174, 236, 185, 7, 36, 117, 223, 58, 245, 96, 81, 124, 45, 75, 161, 88, 154, 184, 148, 169, 240, 42, 57, 88, 113, 211, 5, 101, 48, 170, 203, 18, 104, 198, 99, 84, 78, 206, 68, 69, 89, 0, 0, 127, 116, 151, 192, 76, 184, 98, 116, 240, 191, 5, 225, 169, 116, 83, 25, 144, 15, 39, 79, 116, 222, 66, 214, 89, 27, 157, 214, 144, 135, 13, 100, 77, 188, 205, 87, 135, 42, 88, 172}
// 	client := []byte{
// 		0, 4, 0, 0, 0, 154, 82, 13, 80, 5, 240, 1, 2, 192, 2, 128, 128, 128, 128, 128, 33, 240, 1, 0, 146, 3, 103, 82, 101, 82, 96, 32, 17, 240, 213, 13, 116, 134, 254, 34, 194, 214, 192, 123, 82, 245, 150, 111, 96, 229, 176, 139, 255, 207, 118, 237, 164, 151, 167, 196, 117, 223, 192, 174, 74, 22, 41, 160, 136, 160, 194, 74, 242, 176, 151, 239, 93, 185, 8, 144, 76, 73, 206, 89, 98, 87, 188, 228, 241, 163, 85, 153, 215, 52, 59, 184, 192, 134, 61, 164, 177, 57, 81, 208, 170, 206, 84, 95, 119, 9, 253, 218, 7, 133, 13, 169, 43, 247, 217, 114, 232, 230, 53, 147, 152, 203, 158, 160, 1, 1, 226, 3, 16, 161, 17, 134, 95, 119, 52, 99, 38, 34, 100, 149, 63, 201, 89, 227, 139, 130, 5, 2, 8, 1}
// 	server := []byte{
// 		0, 0, 1, 218, 82, 211, 3, 82, 236, 2, 82, 233, 2, 82, 96, 80, 231, 73, 121, 44, 224, 167, 121, 81, 184, 143, 208, 223, 174, 236, 185, 7, 36, 117, 223, 58, 245, 96, 81, 124, 45, 75, 161, 88, 154, 184, 148, 169, 240, 42, 57, 88, 113, 211, 5, 101, 48, 170, 203, 18, 104, 198, 99, 84, 78, 206, 68, 69, 89, 0, 0, 127, 116, 151, 192, 76, 184, 98, 116, 240, 191, 5, 225, 169, 116, 83, 25, 144, 15, 39, 79, 116, 222, 66, 214, 89, 27, 157, 214, 144, 135, 13, 100, 77, 188, 205, 87, 135, 42, 88, 172, 160, 1, 0, 242, 1, 128, 2, 79, 118, 45, 17, 106, 72, 71, 144, 85, 247, 218, 62, 70, 189, 11, 88, 118, 72, 143, 143, 123, 129, 249, 47, 105, 220, 59, 52, 167, 17, 118, 186, 13, 76, 139, 139, 181, 121, 174, 192, 216, 141, 189, 23, 206, 211, 206, 116, 57, 20, 151, 105, 8, 252, 53, 169, 237, 220, 145, 26, 106, 118, 105, 14, 197, 245, 71, 23, 251, 253, 110, 123, 63, 3, 248, 166, 108, 114, 37, 175, 97, 200, 168, 12, 238, 43, 61, 78, 78, 15, 94, 233, 153, 48, 69, 204, 83, 231, 119, 179, 203, 50, 70, 212, 129, 56, 110, 183, 89, 175, 129, 214, 200, 191, 182, 199, 31, 101, 100, 215, 191, 152, 245, 24, 72, 126, 38, 240, 167, 233, 234, 203, 200, 108, 246, 98, 212, 234, 84, 123, 79, 25, 200, 58, 190, 113, 38, 171, 245, 115, 170, 114, 75, 161, 176, 31, 176, 99, 34, 26, 48, 25, 176, 151, 81, 236, 219, 197, 149, 41, 65, 117, 98, 13, 202, 136, 91, 133, 62, 115, 242, 239, 192, 239, 40, 38, 103, 246, 114, 134, 127, 120, 167, 64, 70, 28, 40, 94, 247, 92, 239, 228, 230, 39, 249, 5, 155, 139, 199, 235, 200, 74, 134, 154, 53, 224, 84, 186, 173, 211, 85, 111, 61, 12, 65, 233, 131, 210, 9, 116, 141, 3, 53, 136, 200, 58, 1, 171, 228, 75, 172, 28, 20, 9, 38, 128, 214, 203, 233, 167, 97, 88, 2, 131, 50, 142, 162, 1, 0, 242, 1, 0, 194, 2, 2, 82, 0, 146, 3, 16, 22, 156, 233, 197, 6, 249, 149, 25, 57, 152, 144, 140, 37, 251, 216, 122, 226, 3, 67, 52, 205, 129, 168, 23, 124, 19, 104, 47, 68, 97, 121, 13, 29, 143, 169, 6, 84, 175, 255, 233, 201, 57, 130, 89, 197, 167, 84, 157, 33, 40, 209, 239, 169, 121, 6, 38, 140, 110, 85, 208, 208, 207, 221, 237, 94, 135, 243, 179, 54, 243, 156, 255, 44, 30, 88, 241, 198, 173, 142, 231, 213, 96, 214, 127, 217, 221,}
// 	private := new(big.Int)
// 	private.SetBytes([]byte{
// 		103, 228, 84, 119, 74, 120, 179, 43, 159, 11, 193, 83, 36, 0, 109, 219, 113, 153, 225, 147, 168, 252, 23, 56, 159, 124, 235, 9, 146, 149, 118, 167, 229, 22, 76, 77, 68, 165, 9, 117, 75, 162, 12, 95, 205, 208, 208, 170, 117, 54, 235, 68, 143, 183, 209, 184, 246, 135, 149, 142, 228, 104, 231, 144, 183, 237, 67, 150, 151, 147, 141, 45, 134, 144, 104, 59, 228, 210, 134, 155, 87, 165, 65, 232, 19, 215, 49, 17, 3, 0, 12, 191, 110, 117, 233,
// 		})
// 	fmt.Println("p2",private)

// 	challenge := []byte{
// 		99, 120, 119, 147, 150, 42, 231, 240, 110, 42, 17, 161, 204, 15, 235, 86, 102, 1, 121, 48}
//     sendKey := []byte{
//     	239, 15, 68, 152, 244, 120, 58, 27, 16, 132, 148, 62, 241, 73, 196, 74, 194, 141, 115, 65, 141, 222, 239, 66, 3, 111, 184, 159, 211, 251, 74, 214,
//     }

//     recvKey := []byte{
//     	28, 204, 1, 169, 191, 26, 105, 209, 197, 91, 140, 154, 87, 41, 27, 160, 9, 57, 196, 93, 76, 107, 140, 198, 140, 178, 247, 70, 204, 154, 74, 109,
//     }


// 	keys := generateKeysFromPrivate(private)
// 	shared := keys.addRemoteKey(remote, client, server)
// 	if !bytes.Equal(shared.challenge, challenge) {
// 		t.Errorf("result does not match, %v %v", shared.challenge, challenge)
// 	}

// 	if !bytes.Equal(shared.sendKey, sendKey) {
// 		t.Errorf("result sendKey  does not match, %v %v", shared.sendKey, sendKey)
// 	}

// 	if !bytes.Equal(shared.recvKey, recvKey) {
// 		t.Errorf("result recvKey does not match, %v %v", shared.recvKey, recvKey)
// 	}
// }

// func TestLoginPacket(t *testing.T) {
// 	loginPacketExpected := []byte{
// 		82, 24, 82, 10, 49, 50, 52, 53, 53, 56, 52, 54, 48, 50, 160, 1, 0, 242, 1, 6, 50, 49, 51, 49, 50, 53, 146, 3, 60, 80, 0, 224, 3, 0, 210, 5, 9, 108, 105, 98, 114, 101, 115, 112, 111, 116, 162, 6, 40, 55, 50, 56, 56, 101, 100, 100, 48, 102, 99, 51, 102, 102, 99, 98, 101, 57, 51, 97, 48, 99, 102, 48, 54, 101, 51, 53, 54, 56, 101, 50, 56, 53, 50, 49, 54, 56, 55, 98, 99, 178, 4, 17, 108, 105, 98, 114, 101, 115, 112, 111, 116, 45, 56, 51, 49, 53, 101, 49, 48, 130, 5, 241, 2, 8, 1, 18, 128, 1, 186, 250, 60, 169, 173, 199, 160, 1, 196, 108, 34, 130, 42, 191, 43, 253, 233, 158, 63, 169, 183, 214, 72, 79, 98, 79, 185, 188, 191, 17, 220, 243, 232, 19, 33, 206, 250, 183, 198, 79, 192, 212, 161, 200, 90, 197, 165, 218, 153, 223, 250, 246, 52, 63, 33, 245, 47, 178, 59, 87, 13, 173, 69, 191, 127, 198, 211, 199, 10, 108, 77, 198, 170, 250, 98, 143, 148, 251, 181, 238, 18, 223, 189, 151, 106, 27, 237, 185, 200, 165, 24, 43, 245, 219, 79, 152, 61, 243, 76, 8, 10, 155, 139, 9, 63, 112, 243, 123, 60, 213, 191, 240, 109, 131, 97, 224, 117, 138, 84, 224, 1, 253, 170, 131, 12, 48, 50, 72, 26, 192, 1, 80, 204, 221, 35, 186, 246, 45, 101, 23, 161, 73, 62, 4, 54, 247, 188, 154, 4, 180, 54, 39, 63, 231, 20, 151, 223, 94, 166, 192, 31, 37, 77, 135, 206, 6, 250, 178, 190, 77, 77, 12, 12, 137, 42, 185, 36, 233, 187, 245, 61, 241, 106, 208, 18, 232, 245, 73, 51, 231, 5, 4, 63, 236, 74, 174, 107, 197, 109, 8, 206, 80, 199, 57, 58, 253, 117, 70, 90, 188, 40, 112, 6, 58, 35, 131, 102, 141, 250, 40, 8, 151, 150, 138, 32, 23, 22, 136, 171, 193, 110, 104, 241, 40, 205, 209, 117, 230, 234, 232, 86, 18, 122, 11, 250, 77, 51, 106, 32, 71, 44, 80, 71, 161, 40, 36, 38, 163, 49, 44, 30, 202, 249, 243, 85, 23, 232, 51, 53, 31, 182, 198, 214, 23, 5, 239, 104, 251, 208, 58, 180, 190, 23, 177, 206, 106, 252, 191, 234, 118, 28, 182, 240, 231, 157, 255, 214, 140, 73, 222, 157, 200, 231, 249, 130, 81, 48, 78, 96, 186, 120, 176, 237, 212, 241, 143, 86, 102, 130, 154, 235, 55, 114, 34, 17, 108, 105, 98, 114, 101, 115, 112, 111, 116, 45, 56, 51, 49, 53, 101, 49, 48, 42, 20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// 	}

// 	loginPacket := loginPacket("./spotify_appkey.key", "1245584602", "213125")
// 	if !bytes.Equal(loginPacketExpected, loginPacket) {
// 		t.Errorf("result does not match, %v %v", loginPacketExpected, loginPacket)
// 	}
// }


