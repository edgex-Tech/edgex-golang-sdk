package starkcurve

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
)

type StarkCfg struct {
	Alpha          *big.Int     `json:"ALPHA"`
	Beta           *big.Int     `json:"BETA"`
	ConstantPoints [][]*big.Int `json:"CONSTANT_POINTS"`
	EcOrder        *big.Int     `json:"EC_ORDER"`
	FieldGen       *big.Int     `json:"FIELD_GEN"`
	FieldPrime     *big.Int     `json:"FIELD_PRIME"`
	Comment        string       `json:"_comment"`
	License        []string     `json:"_license"`
}

type StarkPoints struct {
	ConstantPoints [][]*big.Int `json:"const_points"`
}

var (
	cfg         StarkCfg
	constPoints StarkPoints
)

func loadCfgFromData() {

	dataBuffer := bytes.NewBuffer([]byte(starkcurveParams))

	// Create JSON decoder
	decoder := json.NewDecoder(dataBuffer)
	err := decoder.Decode(&cfg)
	if err != nil {
		fmt.Println("Decoder cfg failed", err.Error())
	}

	dataBuffer1 := bytes.NewBuffer([]byte(constPointsParams))

	// Create JSON decoder
	decoder1 := json.NewDecoder(dataBuffer1)
	err = decoder1.Decode(&constPoints)
	if err != nil {
		fmt.Println("Decoder constPoints failed", err.Error())
	}
}

func init() {
	loadCfgFromData()
}

func CalcHash(input []*big.Int) []byte {
	shiftPointx := big.NewInt(0).Set(cfg.ConstantPoints[0][0])
	shiftPointy := big.NewInt(0).Set(cfg.ConstantPoints[0][1])
	//fmt.Printf("shiftPointx %x, shiftPointy %x\n", shiftPointx.Bytes(), shiftPointy.Bytes())
	//fmt.Printf("shiftPointx %v, shiftPointy %v\n", shiftPointx.String(), shiftPointy.String())
	curve := NewStarkCurve()
	//     let x = new BN(input[i], 16);
	//     assert(x.gte(zeroBn) && x.lt(prime), 'Invalid input: ' + input[i]);
	//     for (let j = 0; j < 252; j++) {
	//         const pt = constantPoints[2 + i * 252 + j];
	//         assert(!point.getX().eq(pt.getX()));
	//         if (x.and(oneBn).toNumber() !== 0) {
	//             point = point.add(pt);
	//         }
	//         x = x.shrn(1);
	//     }
	// }
	// return point.getX().toString(16);

	for i := 0; i < len(input); i++ {
		x := big.NewInt(0).Set(input[i])
		one := big.NewInt(1)
		zero := big.NewInt(0)
		for j := 0; j < 252; j++ {

			pos := 2 + i*252 + j

			ptx := big.NewInt(0).Set(cfg.ConstantPoints[pos][0])
			pty := big.NewInt(0).Set(cfg.ConstantPoints[pos][1])
			//fmt.Printf("x %x\n", x.Bytes())

			//fmt.Printf("1 ptx %x\n1 pty %x\n", ptx.Bytes(), pty.Bytes())
			//	fmt.Printf("ptx %v, pty %v\n", ptx.String(), pty.String())

			if big.NewInt(0).And(x, one).Cmp(zero) != 0 {
				shiftPointx, shiftPointy = curve.Add(shiftPointx, shiftPointy, ptx, pty)

				//fmt.Printf("2 pointx %x\n2 pointy %x\n", shiftPointx.Bytes(), shiftPointy.Bytes())
				//fmt.Printf("111 shiftPointx %v, shiftPointy %v\n", shiftPointx.String(), shiftPointy.String())

			}
			x = big.NewInt(0).Rsh(x, 1)
		}
	}
	return shiftPointx.Bytes()

}

// Reference https://github.com/apisit/rfc6979
func Sign(privkey []byte, hash []byte) (*big.Int, *big.Int, error) {

	hashInt := big.NewInt(0).SetBytes(hash)
	one := big.NewInt(1)
	maxData := big.NewInt(1).Lsh(one, 251)
	if hashInt.Cmp(maxData) > 0 {
		return nil, nil, errors.New("hash cannot sign ")
	}
	curve := NewStarkCurve()
	privkeyInt := big.NewInt(0).SetBytes(privkey)
	N := curve.N
	r := big.NewInt(0)
	s := big.NewInt(0)
	generateSecret(N, sha256.New, hash, func(k *big.Int) bool {
		// fmt.Println("k ", k)
		inv := new(big.Int).ModInverse(k, N)
		r, _ = curve.ScalarBaseMult(k.Bytes())
		r.Mod(r, N)

		if r.Sign() == 0 {
			return false
		}

		e := hashToInt(hash, curve.BitSize)
		s = new(big.Int).Mul(privkeyInt, r)
		s.Add(s, e)
		s.Mul(s, inv)
		s.Mod(s, N)

		return s.Sign() != 0
	})
	return r, s, nil
}

// Reference https://github.com/apisit/rfc6979
func Sign2(privkey []byte, hash []byte) (*big.Int, *big.Int, error) {

	hashInt := big.NewInt(0).SetBytes(hash)
	one := big.NewInt(1)
	maxData := big.NewInt(1).Lsh(one, 251)
	if hashInt.Cmp(maxData) > 0 {
		return nil, nil, errors.New("hash cannot sign ")
	}
	curve := NewStarkCurve()
	privkeyInt := big.NewInt(0).SetBytes(privkey)
	N := curve.N
	r := big.NewInt(0)
	s := big.NewInt(0)
	// while True:
	// k = generate_k_rfc6979(msg_hash, priv_key, seed)
	// # Update seed for next iteration in case the value of k is bad.
	// if seed is None:
	// 	seed = 1
	// else:
	// 	seed += 1

	// # Cannot fail because 0 < k < EC_ORDER and EC_ORDER is prime.
	// x = ec_mult(k, EC_GEN, ALPHA, FIELD_PRIME)[0]

	// # DIFF: in classic ECDSA, we take int(x) % n.
	// r = int(x)
	// if not (1 <= r < 2**N_ELEMENT_BITS_ECDSA):
	// 	# Bad value. This fails with negligible probability.
	// 	continue

	// if (msg_hash + r * priv_key) % EC_ORDER == 0:
	// 	# Bad value. This fails with negligible probability.
	// 	continue

	// w = div_mod(k, msg_hash + r * priv_key, EC_ORDER)
	// if not (1 <= w < 2**N_ELEMENT_BITS_ECDSA):
	// 	# Bad value. This fails with negligible probability.
	// 	continue

	// s = inv_mod_curve_size(w)
	// return r, s
	generateSecret(N, sha256.New, hash, func(k *big.Int) bool {
		// fmt.Println("k ", k)
		inv := new(big.Int).ModInverse(k, N)
		r, _ = curve.ScalarBaseMult(k.Bytes())

		if r.Cmp(maxData) > 0 {
			return false
		}
		r.Mod(r, N)

		if r.Sign() == 0 {
			return false
		}

		e := hashToInt(hash, curve.BitSize)
		s = new(big.Int).Mul(privkeyInt, r)
		s.Add(s, e)
		tempData := big.NewInt(0)
		tempData = tempData.Mod(s, cfg.EcOrder)
		if tempData.Sign() == 0 {
			return false
		}
		// tempData = tempData.Set(s)
		// tempData.DivMod(k, tempData, cfg.EcOrder)
		s.Mul(s, inv)
		s.Mod(s, N)

		return s.Sign() != 0
	})
	return r, s, nil
}

// copied from crypto/ecdsa
func hashToInt(hash []byte, orderBits int) *big.Int {

	orderBytes := (orderBits + 7) / 8
	if len(hash) > orderBytes {
		hash = hash[:orderBytes]
	}

	ret := new(big.Int).SetBytes(hash)
	excess := len(hash)*8 - orderBits
	if excess > 0 {
		ret.Rsh(ret, uint(excess))
	}
	return ret
}

func VerifyPubKey(pubkey []byte) bool {
	if len(pubkey) != 64 {
		return false
	}
	pubx := big.NewInt(0).SetBytes(pubkey[0:32])
	puby := big.NewInt(0).SetBytes(pubkey[32:])
	curve := NewStarkCurve()
	return curve.IsOnCurve(pubx, puby)
}
func Verify(hash []byte, pubkeyX, pubkeyY, r, s *big.Int) bool {
	hashInt := big.NewInt(0).SetBytes(hash)

	maxData := big.NewInt(1).Lsh(one, 251)
	if hashInt.Cmp(maxData) > 0 {
		return false
	}
	curve := NewStarkCurve()
	N := curve.N
	if r.Sign() <= 0 || s.Sign() <= 0 {
		return false
	}
	if r.Cmp(N) >= 0 || s.Cmp(N) >= 0 {
		return false
	}
	e := hashToInt(hash, curve.BitSize)
	w := new(big.Int).ModInverse(s, N)
	u1 := e.Mul(e, w)
	u1.Mod(u1, N)
	u2 := w.Mul(r, w)
	u2.Mod(u2, N)

	// Check if implements S1*g + S2*p
	var x, y *big.Int
	x1, y1 := curve.ScalarBaseMult(u1.Bytes())
	x2, y2 := curve.ScalarMult(pubkeyX, pubkeyY, u2.Bytes())
	x, y = curve.Add(x1, y1, x2, y2)

	if x.Sign() == 0 && y.Sign() == 0 {
		return false
	}
	x.Mod(x, N)
	return x.Cmp(r) == 0
}

// def get_msg(
// 	instruction_type: int, vault0: int, vault1: int, amount0: int, amount1: int, token0: int,
// 	token1_or_pub_key: int, nonce: int, expiration_timestamp: int,
// 	hash=pedersen_hash, condition: Optional[int] = None) -> int:
// """
// Creates a message to sign on.
// """
// packed_message = instruction_type
// packed_message = packed_message * 2**31 + vault0
// packed_message = packed_message * 2**31 + vault1
// packed_message = packed_message * 2**63 + amount0
// packed_message = packed_message * 2**63 + amount1
// packed_message = packed_message * 2**31 + nonce
// packed_message = packed_message * 2**22 + expiration_timestamp
// if condition is not None:
// 	# A message representing a conditional transfer. The condition is interpreted by the
// 	# application.
// 	return hash(hash(hash(token0, token1_or_pub_key), condition), packed_message)

// return hash(hash(token0, token1_or_pub_key), packed_message)
func GetMsgHash(instuctType, vault0, vault1, amount0, amount1 *big.Int, token0, token1OrPubkey []byte, nouce, expirationTimestamp uint64, condition []byte) []byte {
	return nil
}
